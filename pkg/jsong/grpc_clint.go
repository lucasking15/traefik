// Package grpcurl provides the core functionality exposed by the grpcurl command, for
// dynamically connecting to a server, using the reflection service to inspect the server,
// and invoking RPCs. The grpcurl command-line tool constructs a DescriptorSource, based
// on the command-line parameters, and supplies an InvocationEventHandler to supply request
// data (which can come from command-line args or the process's stdin) and to log the
// events (to the process's stdout).
package jsong

import (
	"encoding/base64"
	"github.com/containous/traefik/v2/pkg/jsong/copygrpcreflect"
	"github.com/pkg/errors"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"net"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var grpcCCMap sync.Map
var reflectCCMap sync.Map

func getReflectCC(ctx context.Context, addr string) (refCC *copygrpcreflect.Client, err error) {
	if val, ok := reflectCCMap.Load(addr); ok {
		refCC = val.(*copygrpcreflect.Client)
		refCC.CleanCache()
	} else {
		cc, err := GetGrpcCC(ctx, addr)
		if err != nil {
			return nil, err
		}
		refCC = copygrpcreflect.NewClient(ctx, grpc_reflection_v1alpha.NewServerReflectionClient(cc))
		reflectCCMap.Store(addr, refCC)
	}

	return refCC, nil
}

// GetGrpcCC 获取一个连接
func GetGrpcCC(ctx context.Context, addr string) (cc *grpc.ClientConn, err error) {
	if val, ok := grpcCCMap.Load(addr); ok {
		cc = val.(*grpc.ClientConn)
	} else {
		cc, err = dial(ctx, addr)
		if err != nil {
			return nil, err
		}
		grpcCCMap.Store(addr, cc)
	}

	return cc, nil
}

// dial 建立一个连接
func dial(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	network, addr, err := getNetworkAddress(addr)
	if err != nil {
		return nil, errors.Errorf("grpc dial: %v", err)
	}
	return BlockingDial(ctx, network, addr, nil)
}

func getNetworkAddress(address string) (string, string, error) {
	split := strings.SplitN(address, "://", 2)
	if len(split) != 2 {
		return "tcp", address, nil
	}
	switch split[0] {
	case "tcp", "unix":
		return split[0], split[1], nil
	default:
		return "", "", errors.Errorf("invalid network, only tcp or unix allowed: %s", split[0])
	}
}

// MetadataFromHeaders converts a list of header strings (each string in
// "Header-Name: Header-Value" form) into metadata. If a string has a header
// name without a value (e.g. does not contain a colon), the value is assumed
// to be blank. Binary headers (those whose names end in "-bin") should be
// base64-encoded. But if they cannot be base64-decoded, they will be assumed to
// be in raw form and used as is.
func MetadataFromHeaders(headers http.Header) metadata.MD {
	md := make(metadata.MD)
	for k, v := range headers {
		headerName := strings.ToLower(strings.TrimSpace(k))
		if len(v) == 0 {
			v = append(v, "")
		}
		val := strings.TrimSpace(v[0])
		if strings.HasSuffix(headerName, "-bin") {
			if v, err := decode(val); err == nil {
				val = v
			}
		}
		md[headerName] = append(md[headerName], val)
	}
	return md
}

var base64Codecs = []*base64.Encoding{base64.StdEncoding, base64.URLEncoding, base64.RawStdEncoding, base64.RawURLEncoding}

func decode(val string) (string, error) {
	var firstErr error
	var b []byte
	// we are lenient and can accept any of the flavors of base64 encoding
	for _, d := range base64Codecs {
		var err error
		b, err = d.DecodeString(val)
		if err != nil {
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		return string(b), nil
	}
	return "", firstErr
}

// BlockingDial is a helper method to dial the given address, using optional TLS credentials,
// and blocking until the returned connection is ready. If the given credentials are nil, the
// connection will be insecure (plain-text).
func BlockingDial(ctx context.Context, network, address string, creds credentials.TransportCredentials, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	// grpc.Dial doesn't provide any information on permanent connection errors (like
	// TLS handshake failures). So in order to provide good error messages, we need a
	// custom dialer that can provide that info. That means we manage the TLS handshake.
	result := make(chan interface{}, 1)

	writeResult := func(res interface{}) {
		// non-blocking write: we only need the first result
		select {
		case result <- res:
		default:
		}
	}

	dialer := func(ctx context.Context, address string) (net.Conn, error) {
		conn, err := (&net.Dialer{}).DialContext(ctx, network, address)
		if err != nil {
			writeResult(err)
			return nil, err
		}
		if creds != nil {
			conn, _, err = creds.ClientHandshake(ctx, address, conn)
			if err != nil {
				writeResult(err)
				return nil, err
			}
		}
		return conn, nil
	}

	// Even with grpc.FailOnNonTempDialError, this call will usually timeout in
	// the face of TLS handshake errors. So we can't rely on grpc.WithBlock() to
	// know when we're errChan. So we run it in a goroutine and then use result
	// channel to either get the channel or fail-fast.
	go func() {
		opts = append(opts,
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithContextDialer(dialer),
			grpc.WithInsecure(), // we are handling TLS, so tell grpc not to
		)
		conn, err := grpc.DialContext(ctx, address, opts...)
		var res interface{}
		if err != nil {
			res = err
		} else {
			res = conn
		}
		writeResult(res)
	}()

	select {
	case res := <-result:
		if conn, ok := res.(*grpc.ClientConn); ok {
			return conn, nil
		}
		return nil, res.(error)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
