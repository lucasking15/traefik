// @Author: wangdehong
// @Description:
// @File: ws
// @Date: 2022/4/7 12:51 下午

package ws

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
	"github.com/traefik/traefik/v2/pkg/tracing"
	"net/http"
)

const (
	typeName = "ws"
)

type ws struct {
	next http.Handler
	name string
}

// New creates a new handler.
func New(ctx context.Context, next http.Handler, config dynamic.Ws, name string) (http.Handler, error) {
	log.FromContext(middlewares.GetLoggerCtx(ctx, name, typeName)).Debug("Creating middleware")
	var result *ws

	result = &ws{
		next: next,
		name: name,
	}
	return result, nil

}

func (a *ws) GetTracingInformation() (string, ext.SpanKindEnum) {
	return a.name, tracing.SpanKindNoneEnum
}

func (a *ws) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	logger := log.FromContext(middlewares.GetLoggerCtx(req.Context(), a.name, typeName))
	logger.Infoln("Hello World")
	fmt.Println("Hello World")
	req.Header.Add("Hello", "World")
	a.next.ServeHTTP(rw, req)
}
