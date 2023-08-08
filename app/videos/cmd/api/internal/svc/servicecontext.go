package svc

import (
	_UserModel "MiniDY/app/usercenter/model"
	"MiniDY/app/videos/cmd/api/internal/config"
	"MiniDY/app/videos/cmd/rpc/videoservice"
	"MiniDY/app/videos/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	VideoRpc        videoservice.Videoservice
	VideoModel      model.VideoModel
	UsercenterModel _UserModel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:          c,
		VideoRpc:        videoservice.NewVideoservice(zrpc.MustNewClient(c.VideoRpcConf)),
		VideoModel:      model.NewVideoModel(sqlconn, c.Cache),
		UsercenterModel: _UserModel.NewUserModel(sqlconn, c.Cache),
	}
}
