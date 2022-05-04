package protocol

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
)

const (
	// MaxBodySize max proto body size
	MaxBodySize = int32(1 << 12)
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

var (
	// ErrProtoPackLen proto packet len error
	ErrProtoPackLen = errors.New("default server codec pack length error")
	// ErrProtoHeaderLen proto header len error
	ErrProtoHeaderLen = errors.New("default server codec header length error")
)

var (
	// ProtoReady proto ready
	ProtoReady = &Proto{Op: OpProtoReady}
	// ProtoFinish proto finish
	ProtoFinish = &Proto{Op: OpProtoFinish}
)

func (p *Proto) ReadWebsocket(ws *websocket.Conn) (err error) {
	_, messageByte, err := ws.ReadMessage()
	if err != nil {
		return err
	}

	err = json.Unmarshal(messageByte, &p)
	if err != nil {
		return err
	}
	
	return nil
}

// WriteWebsocket write a proto to websocket connection.
func (p *Proto) WriteWebsocket(ws *websocket.Conn) (err error) {
	pByte, _ := json.Marshal(&p)
	ws.WriteMessage(1, pByte)

	return nil
}