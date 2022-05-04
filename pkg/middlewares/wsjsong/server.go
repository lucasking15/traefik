// @Author: wangdehong
// @Description:
// @File: server
// @Date: 2022/4/23 7:14 下午

package wsjsong

import (
	"github.com/containous/traefik/v2/pkg/middlewares/wsjsong/conf"
	"github.com/containous/traefik/v2/pkg/middlewares/wsjsong/protocol/logic"
	"github.com/zhenjl/cityhash"
	"math/rand"
	"time"
)

const (
	minServerHeartbeat = time.Minute * 10
	maxServerHeartbeat = time.Minute * 30
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
	grpcMaxSendMsgSize        = 1 << 24
	grpcMaxCallMsgSize        = 1 << 24
	grpcKeepAliveTime         = time.Second * 10
	grpcKeepAliveTimeout      = time.Second * 3
	grpcBackoffMaxDelay       = time.Second * 3
)

// Server is comet server.
type Server struct {
	c         *conf.Config
	buckets   []*Bucket // subkey bucket
	bucketIdx uint32

	serverID  string

	rpcClient logic.LogicClient
}

// Bucket get the bucket by subkey.
func (s *Server) Bucket(subKey string) *Bucket {
	idx := cityhash.CityHash32([]byte(subKey), uint32(len(subKey))) % s.bucketIdx
	return s.buckets[idx]
}
// RandServerHearbeat rand server heartbeat.
func (s *Server) RandServerHearbeat() time.Duration {
	return (minServerHeartbeat + time.Duration(rand.Int63n(int64(maxServerHeartbeat-minServerHeartbeat))))
}

