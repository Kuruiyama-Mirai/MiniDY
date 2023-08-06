package svc

import (
	"MiniDY/app/social/cmd/rpc/internal/config"
	"MiniDY/app/social/model"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	_userModel "MiniDY/app/usercenter/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	FollowModel   model.FollowModel
	UserModel     _userModel.UserModel
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
		}),
		FollowModel:   model.NewFollowModel(sqlconn, c.Cache),
		UserModel:     _userModel.NewUserModel(sqlconn, c.Cache),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
