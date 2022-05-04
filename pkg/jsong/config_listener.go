package jsong

import (
	"context"
	"github.com/containous/traefik/v2/pkg/config/dynamic"
	"github.com/containous/traefik/v2/pkg/log"
	"github.com/containous/traefik/v2/pkg/safe"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"strings"
	"sync"
	"time"
)

const PathOption = "72295728"

var (
	Reflections sync.Map
	mux         sync.RWMutex
	config      dynamic.Configuration
)

func init() {
	go func() {
		for range time.Tick(time.Minute) {
			ctx := log.With(context.Background(), log.Str("component", "jsong.timeTick"))
			mux.RLock()
			doParseGRPC(ctx, config)
			mux.RUnlock()
		}
	}()
}

func LoadGRPCReflection(newConfig dynamic.Configuration) {
	ctx := log.With(context.Background(), log.Str("component", "jsong.loadGRPCReflection"))
	mux.Lock()
	config = newConfig
	doParseGRPC(ctx, config)
	mux.Unlock()
}

func doParseGRPC(ctx context.Context, configuration dynamic.Configuration) {
	if configuration.HTTP == nil {
		return
	}
	for name, service := range configuration.HTTP.Services {
		name := name
		service := service
		ctx := log.With(ctx, log.Str("service", name))
		safe.Go(func() {
			resolveService(ctx, name, service)
		})
	}
}

func resolveService(ctx context.Context, name string, service *dynamic.Service) {
	if service.LoadBalancer == nil {
		return
	}
	if len(service.LoadBalancer.Servers) == 0 {
		return
	}

	log.FromContext(ctx).Infof("load grpc reflection from service %s", name)
	server := service.LoadBalancer.Servers[0]
	if !strings.HasPrefix(server.URL, "jsong://") {
		log.FromContext(ctx).Infof("skip %s", server.URL)
		return
	}

	substr := strings.SplitN(server.URL, "://", 2)
	if len(substr) < 2 {
		log.FromContext(ctx).Errorf("parse network address from %s fail, eg: jsong://host:port", server.URL)
		return
	}

	refClient, err := getReflectCC(ctx, substr[1])
	if err != nil {
		log.FromContext(ctx).Errorf("dial to %s fail err: %s", server.URL, err)
		return
	}

	grpcServices, err := refClient.ListServices()
	if err != nil {
		log.FromContext(ctx).Errorf("list %s grpc services fail err: %s", server.URL, err)
		return
	}

	for _, grpcService := range grpcServices {
		serviceDesc, err := refClient.ResolveService(grpcService)
		if err != nil {
			log.FromContext(ctx).Errorf("resolve %s grpc service fail svc:%s err: %s", server.URL, grpcService, err)
			continue
		}

		for _, method := range serviceDesc.GetMethods() {
			options, ok := method.GetOptions().(*descriptor.MethodOptions)
			if !ok {
				continue
			}

			if options == nil {
				continue
			}

			path := parsePath(options)
			if path == "" {
				continue
			}

			if exist, ok := Reflections.Load(path); ok {
				// 方法的Service不一样
				if existMethod := exist.(*desc.MethodDescriptor); existMethod.GetParent().GetName() != method.GetParent().GetName() {
					log.FromContext(ctx).Errorf(
						"found path:%s conflict exist service:%s exist method:%s new service:%s new method:%s",
						path,
						existMethod.GetParent().GetName(),
						existMethod.GetName(),
						method.GetParent().GetName(),
						method.GetName())
					continue
				}
			}

			log.FromContext(ctx).Infof("found path:%s for service:%s method:%s", path, grpcService, method.GetName())
			Reflections.Store(path, method)
		}
	}
}

func parsePath(options *descriptor.MethodOptions) string {
	// 72295728:"\x12//v1/sapiFoundationInfraAPI/uploadFileAssumeRole"
	optionsStr := options.String()

	if !strings.Contains(optionsStr, PathOption) {
		return ""
	}

	var idx int
	if idx = strings.Index(optionsStr, "/"); idx == -1 {
		return ""
	}

	path := optionsStr[idx:]

	if idx = strings.Index(path, ":"); idx != -1 {
		path = path[:idx]
	}

	if idx = strings.Index(path, `"`); idx != -1 {
		path = path[:idx]
	}
	if strings.HasPrefix(path, "//") {
		path = path[1:]
	}

	return path
}
