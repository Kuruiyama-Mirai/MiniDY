package svc

import (
	"MiniDY/app/social/cmd/api/internal/config"
	"MiniDY/app/social/cmd/rpc/followerservice"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	FollowRpc     followerservice.Followerservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		FollowRpc:     followerservice.NewFollowerservice(zrpc.MustNewClient(c.FollowRpcConf)),
	}
}
