package jsong

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
)

type requestError struct {
	msg string
}

func (e *requestError) Error() string {
	return e.msg
}

func IsRequestError(err error) bool {
	_, ok := err.(*requestError)
	return ok
}

// InvocationEventHandler is a bag of callbacks for handling events that occur in the course
// of invoking an RPC. The handler also provides request data that is sent. The callbacks are
// generally called in the order they are listed below.
type InvocationEventHandler interface {
	// OnResolveMethod is called with a descriptor of the method that is being invoked.
	OnResolveMethod(*desc.MethodDescriptor)
	// OnSendHeaders is called with the request metadata that is being sent.
	OnSendHeaders(metadata.MD)
	// OnReceiveHeaders is called when response headers have been received.
	OnReceiveHeaders(metadata.MD)
	// OnReceiveResponse is called for each response message received.
	OnReceiveResponse(proto.Message)
	// OnReceiveTrailers is called when response trailers and final RPC status have been received.
	OnReceiveTrailers(*status.Status, metadata.MD)
	Reset()
	GetHeaders() metadata.MD
	GetTrailers() metadata.MD
	Err() error
	GetResponse() proto.Message
}

// RequestSupplier is a function that is called to populate messages for a gRPC operation. The
// function should populate the given message or return a non-nil error. If the supplier has no
// more messages, it should return io.EOF. When it returns io.EOF, it should not in any way
// modify the given message argument.
type RequestSupplier func(proto.Message) error

// InvokeRPC uses the given gRPC channel to invoke the given method. The given descriptor source
// is used to determine the type of method and the type of request and response message. The given
// headers are sent as request metadata. Methods on the given event handler are called as the
// invocation proceeds.
//
// The given requestData function supplies the actual data to send. It should return io.EOF when
// there is no more request data. If the method being invoked is a unary or server-streaming RPC
// (e.g. exactly one request message) and there is no request data (e.g. the first invocation of
// the function returns io.EOF), then an empty request message is sent.
//
// If the requestData function and the given event handler coordinate or share any state, they should
// be thread-safe. This is because the requestData function may be called from a different goroutine
// than the one invoking event callbacks. (This only happens for bi-directional streaming RPCs, where
// one goroutine sends request messages and another consumes the response messages).
func InvokeRPC(ctx context.Context, mtd *desc.MethodDescriptor, ch grpcdynamic.Channel,
	headers http.Header, handler InvocationEventHandler, requestData RequestSupplier) error {

	md := MetadataFromHeaders(headers)

	handler.OnResolveMethod(mtd)

	// we also download any applicable extensions so we can provide full support for parsing user-provided data
	var ext dynamic.ExtensionRegistry

	msgFactory := dynamic.NewMessageFactoryWithExtensionRegistry(&ext)
	req := msgFactory.NewMessage(mtd.GetInputType())

	handler.OnSendHeaders(md)
	ctx = metadata.NewOutgoingContext(ctx, md)

	stub := grpcdynamic.NewStubWithMessageFactory(ch, msgFactory)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if mtd.IsClientStreaming() && mtd.IsServerStreaming() {
		return invokeBidi(ctx, stub, mtd, handler, requestData, req)
	} else if mtd.IsClientStreaming() {
		return invokeClientStream(ctx, stub, mtd, handler, requestData, req)
	} else if mtd.IsServerStreaming() {
		return invokeServerStream(ctx, stub, mtd, handler, requestData, req)
	} else {
		return invokeUnary(ctx, stub, mtd, handler, requestData, req)
	}
}

func invokeUnary(ctx context.Context, stub grpcdynamic.Stub, md *desc.MethodDescriptor, handler InvocationEventHandler,
	requestData RequestSupplier, req proto.Message) error {

	err := requestData(req)
	if err != nil && err != io.EOF {
		return &requestError{
			msg: fmt.Sprintf("error getting request data: %v", err),
		}
	}
	if err != io.EOF {
		// verify there is no second message, which is a usage error
		err := requestData(req)
		if err == nil {
			return &requestError{
				msg: fmt.Sprintf("method %q is a unary RPC, but request data contained more than 1 message", md.GetFullyQualifiedName()),
			}
		} else if err != io.EOF {
			return &requestError{
				msg: fmt.Sprintf("error getting request data: %v", err),
			}
		}
	}

	// Now we can actually invoke the RPC!
	var respHeaders metadata.MD
	var respTrailers metadata.MD
	resp, err := stub.InvokeRpc(ctx, md, req, grpc.Trailer(&respTrailers), grpc.Header(&respHeaders))

	stat, ok := status.FromError(err)
	if !ok {
		// Error codes sent from the server will get printed differently below.
		// So just bail for other kinds of errors here.
		return fmt.Errorf("grpc call for %q failed: %v", md.GetFullyQualifiedName(), err)
	}

	handler.OnReceiveHeaders(respHeaders)

	if stat.Code() == codes.OK {
		handler.OnReceiveResponse(resp)
	}

	handler.OnReceiveTrailers(stat, respTrailers)

	return nil
}

func invokeClientStream(ctx context.Context, stub grpcdynamic.Stub, md *desc.MethodDescriptor, handler InvocationEventHandler,
	requestData RequestSupplier, req proto.Message) error {

	// invoke the RPC!
	str, err := stub.InvokeRpcClientStream(ctx, md)

	// Upload each request message in the stream
	var resp proto.Message
	for err == nil {
		err = requestData(req)
		if err == io.EOF {
			resp, err = str.CloseAndReceive()
			break
		}
		if err != nil {
			return fmt.Errorf("error getting request data: %v", err)
		}

		err = str.SendMsg(req)
		if err == io.EOF {
			// We get EOF on send if the server says "go away"
			// We have to use CloseAndReceive to get the actual code
			resp, err = str.CloseAndReceive()
			break
		}

		req.Reset()
	}

	// finally, process response data
	stat, ok := status.FromError(err)
	if !ok {
		// Error codes sent from the server will get printed differently below.
		// So just bail for other kinds of errors here.
		return fmt.Errorf("grpc call for %q failed: %v", md.GetFullyQualifiedName(), err)
	}

	if respHeaders, err := str.Header(); err == nil {
		handler.OnReceiveHeaders(respHeaders)
	}

	if stat.Code() == codes.OK {
		handler.OnReceiveResponse(resp)
	}

	handler.OnReceiveTrailers(stat, str.Trailer())

	return nil
}

func invokeServerStream(ctx context.Context, stub grpcdynamic.Stub, md *desc.MethodDescriptor, handler InvocationEventHandler,
	requestData RequestSupplier, req proto.Message) error {

	err := requestData(req)
	if err != nil && err != io.EOF {
		return fmt.Errorf("error getting request data: %v", err)
	}
	if err != io.EOF {
		// verify there is no second message, which is a usage error
		err := requestData(req)
		if err == nil {
			return fmt.Errorf("method %q is a server-streaming RPC, but request data contained more than 1 message", md.GetFullyQualifiedName())
		} else if err != io.EOF {
			return fmt.Errorf("error getting request data: %v", err)
		}
	}

	// Now we can actually invoke the RPC!
	str, err := stub.InvokeRpcServerStream(ctx, md, req)

	if respHeaders, err := str.Header(); err == nil {
		handler.OnReceiveHeaders(respHeaders)
	}

	// Download each response message
	for err == nil {
		var resp proto.Message
		resp, err = str.RecvMsg()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		handler.OnReceiveResponse(resp)
	}

	stat, ok := status.FromError(err)
	if !ok {
		// Error codes sent from the server will get printed differently below.
		// So just bail for other kinds of errors here.
		return fmt.Errorf("grpc call for %q failed: %v", md.GetFullyQualifiedName(), err)
	}

	handler.OnReceiveTrailers(stat, str.Trailer())

	return nil
}

func invokeBidi(ctx context.Context, stub grpcdynamic.Stub, md *desc.MethodDescriptor, handler InvocationEventHandler,
	requestData RequestSupplier, req proto.Message) error {

	ctx, cancel := context.WithCancel(ctx)

	// invoke the RPC!
	str, err := stub.InvokeRpcBidiStream(ctx, md)

	var wg sync.WaitGroup
	var sendErr atomic.Value

	defer wg.Wait()

	if err == nil {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				cancel()
				err = str.CloseSend()
				//xlog.S(ctx).Debug("--stream主动关闭连接--")
				//xlog.S(ctx).Debug("**stream写消息协程退出**")
			}()

			// Concurrently upload each request message in the stream
			var err error
			for err == nil {
				err = requestData(req)
				if err == io.EOF {
					break
				}
				if err != nil {
					err = fmt.Errorf("error getting request data: %v", err)
					break
				}

				//xlog.S(ctx).Debug("往stream发消息:%v", req)
				err = str.SendMsg(req)

				req.Reset()
			}

			if err != nil {
				//xlog.S(ctx).Errorf("写stream发生错误:%v", err)
				sendErr.Store(err)
			}
		}()
	}

	if respHeaders, err := str.Header(); err == nil {
		handler.OnReceiveHeaders(respHeaders)
	}

	// Download each response message
	for err == nil {
		var resp proto.Message
		resp, err = str.RecvMsg()
		//xlog.S(ctx).Debug("从stream读到消息", resp)

		if err != nil {
			//xlog.S(ctx).Errorf("从stream读消息错误:%v", err)
			if err == io.EOF {
				err = nil
			}
			break
		}
		handler.OnReceiveResponse(resp)
	}

	// 获取写错误
	if se, ok := sendErr.Load().(error); ok && se != io.EOF {
		err = se
	}

	stat, ok := status.FromError(err)
	if !ok {
		// Error codes sent from the server will get printed differently below.
		// So just bail for other kinds of errors here.
		return fmt.Errorf("grpc call for %q failed: %v", md.GetFullyQualifiedName(), err)
	}

	handler.OnReceiveTrailers(stat, str.Trailer())

	return nil
}
