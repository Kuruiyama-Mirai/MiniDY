package svc

import (
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/app/videos/cmd/rpc/internal/config"
	"MiniDY/app/videos/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	VideoModel    model.VideoModel
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
		}),
		VideoModel:    model.NewVideoModel(sqlconn, c.Cache),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
