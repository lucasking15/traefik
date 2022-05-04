package wsjsong

import (
	"context"
	"fmt"
	"github.com/containous/traefik/v2/pkg/jsong"
	"github.com/containous/traefik/v2/pkg/log"
	"github.com/containous/traefik/v2/pkg/middlewares/wsjsong/protocol"
	"github.com/golang/glog"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jhump/protoreflect/dynamic"
	"io"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024

	maxServeTime = 20 * time.Minute
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	unmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
	marshaler   = &jsonpb.Marshaler{EmitDefaults: true}
)

type Client struct {
	conn       *websocket.Conn
	wsToStream chan []byte
	streamToWs chan []byte
}

type wsjsongMiddleware struct {
	ctx  context.Context
	next http.Handler
}

func NewEntryPointMiddleware(ctx context.Context, next http.Handler) http.Handler {
	return &wsjsongMiddleware{
		ctx:  ctx,
		next: next,
	}
}

func (m *wsjsongMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		glog.Error("websocket握手失败", err)
		rw.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	fmt.Println("-1-1-1-1-1----")

	messageType, body, err := conn.ReadMessage()
	fmt.Println("-0-0-0-0-0---", messageType, string(body), err)

	server := &Server{
		c:        nil,
		serverID: "",
	}

	go serveWebsocket(server, rw, req)
}

func serveWebsocket(s *Server, rw http.ResponseWriter, req *http.Request) {
	s.ServeWebsocket(rw, req)
}

func (s *Server) ServeWebsocket(rw http.ResponseWriter, req *http.Request) {
	var (
		err     error
		rid     string
		accepts []int32
		hb      time.Duration
		b       *Bucket
		ch      = NewChannel(5, 10)
		p       *protocol.Proto
		ws      *websocket.Conn // websocket
		lastHb  = time.Now()
	)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	if ws, err = upgrader.Upgrade(rw, req, nil); err != nil {
		glog.Error("websocket握手失败", err)
		rw.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if p, err = ch.CliProto.Set(); err == nil {
		if ch.Mid, ch.Key, rid, accepts, hb, err = s.authWebsocket(ctx, ws, p, req.Header.Get("Cookie")); err == nil {
			ch.Watch(accepts...)
			b = s.Bucket(ch.Key)
			err = b.Put(rid, ch)
		}
	}
	fmt.Println(hb)

	if err != nil {
		ws.Close()
		return
	}

	go s.dispatchWebsocket(ws, ch)
	serverHeartbeat := s.RandServerHearbeat()
	for {
		if p, err = ch.CliProto.Set(); err != nil {
			break
		}
		if err = p.ReadWebsocket(ws); err != nil {
			break
		}
		if p.Op == protocol.OpHeartbeat {
			p.Op = protocol.OpHeartbeatReply
			p.Body = nil
			// NOTE: send server heartbeat for a long time
			if now := time.Now(); now.Sub(lastHb) > serverHeartbeat {
				if err1 := s.Heartbeat(ctx, ch.Mid, ch.Key); err1 == nil {
					lastHb = now
				}
			}
		} else {
			if err = s.Operate(ctx, p, ch, b); err != nil {
				break
			}
		}
		ch.CliProto.SetAdv()
		ch.Signal()
	}

	b.Del(ch)
	ws.Close()
	ch.Close()
	if err = s.Disconnect(ctx, ch.Mid, ch.Key); err != nil {
		log.Errorf("key: %s mid: %d operator do disconnect error(%v)", ch.Key, ch.Mid, err)
	}

}

// dispatch accepts connections on the listener and serves requests
// for each incoming connection.  dispatch blocks; the caller typically
// invokes it in a go statement.
func (s *Server) dispatchWebsocket(ws *websocket.Conn, ch *Channel) {
	var (
		err    error
		finish bool
	)
	for {
		var p = ch.Ready()
		switch p {
		case protocol.ProtoFinish:
			finish = true
			goto failed
		case protocol.ProtoReady:
			// fetch message from svrbox(client send)
			for {
				if p, err = ch.CliProto.Get(); err != nil {
					break
				}
				if err = p.WriteWebsocket(ws); err != nil {
					goto failed
				}
				p.Body = nil // avoid memory leak
				ch.CliProto.GetAdv()
			}
		default:
			if err = p.WriteWebsocket(ws); err != nil {
				goto failed
			}
		}
	}
failed:
	ws.Close()
	// must ensure all channel message discard, for reader won't blocking Signal
	for !finish {
		finish = (ch.Ready() == protocol.ProtoFinish)
	}
}

// auth for goim handshake with client, use rsa & aes.
func (s *Server) authWebsocket(ctx context.Context, ws *websocket.Conn, p *protocol.Proto, cookie string) (mid int64, key, rid string, accepts []int32, hb time.Duration, err error) {
	for {
		// 消息序列化
		if err = p.ReadWebsocket(ws); err != nil {
			return
		}
		if p.Op == protocol.OpAuth {
			break
		} else {
			log.Errorf("ws request operation(%d) not auth", p.Op)
		}
	}
	if mid, key, rid, accepts, hb, err = s.Connect(ctx, p, cookie); err != nil {
		return
	}
	p.Op = protocol.OpAuthReply
	p.Body = nil
	if err = p.WriteWebsocket(ws); err != nil {
		return
	}
	return
}

func (c *Client) handleStream(ctx context.Context, cancel context.CancelFunc, eventHandler *jsong.StreamEventHandler) {
	defer func() {
		cancel()
		glog.Info("**handleStream quit**")
	}()

	msgChan := eventHandler.GetResponses()

	for {
		select {
		case <-ctx.Done():
			return
		case message := <-msgChan:
			response, ok := message.(*dynamic.Message)
			if !ok {
				glog.Error("grpc request err:grpc response not a *dynamic.Message")
				return
			}

			body, err := response.MarshalJSONPB(marshaler)
			if err != nil {
				glog.Error("marshal response err:%v", err)
			}

			c.streamToWs <- body
		}
	}
}

func (c *Client) handleWs(ctx context.Context, cancel context.CancelFunc) {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
		_ = c.conn.Close()
		cancel()
		glog.Error("--websocket主动关闭连接--")
		glog.Error("**handleWs退出**")
	}()

	//c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(maxServeTime))
	c.conn.SetPongHandler(func(string) error {
		//_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		glog.Info("websocket探活接收到pong")
		return nil
	})

	go func() {
		defer func() {
			glog.Info("**websocket读协程退出**")
			cancel()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					glog.Error("websocket read error:remote closed:%v", err)
				} else {
					glog.Error("websocket read error:%v", err)
				}
				return
			}
			glog.Info("从websocket读到数据", string(msg))
			c.wsToStream <- msg
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-c.streamToWs:
			glog.Info("往websocket发消息", string(msg))
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				glog.Error("写websocket出错:%v", err)
				return
			}
		case <-ticker.C:
			//_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			glog.Info("websocket探活")
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				glog.Error("websocket探活失败:%v", err)
				return
			}
		}
	}
}

func (c *Client) decodeMessage(ctx context.Context) jsong.RequestSupplier {
	return func(msg proto.Message) error {
		var reqBody []byte
		select {
		case reqBody = <-c.wsToStream:
			err := msg.(*dynamic.Message).UnmarshalJSONPB(unmarshaler, reqBody)
			if err != nil {
				return err
			}
			return nil
		case <-ctx.Done():
			glog.Info("返回io.EOF")
			// 返回io.EOF后，双向流关闭
			return io.EOF
		}
	}
}
