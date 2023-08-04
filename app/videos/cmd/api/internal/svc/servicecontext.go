package svc

import (
	"MiniDY/app/videos/cmd/api/internal/config"
	"MiniDY/app/videos/cmd/rpc/videoservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	VideoRpc videoservice.Videoservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		VideoRpc: videoservice.NewVideoservice(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}
