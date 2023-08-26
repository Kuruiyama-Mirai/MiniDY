package svc

import (
	"MiniDY/app/interaction/cmd/api/internal/config"
	"MiniDY/app/interaction/cmd/rpc/interactsservice"
	_comment "MiniDY/app/interaction/model/comment"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/app/videos/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	InteractsRpc  interactsservice.Interactsservice
	VideoModel    model.VideoModel
	CommentModel  _comment.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		InteractsRpc:  interactsservice.NewInteractsservice(zrpc.MustNewClient(c.InteractsRpcConf)),
		VideoModel:    model.NewVideoModel(sqlconn, c.Cache),
		CommentModel:  _comment.NewCommentModel(sqlconn, c.Cache),
	}
}
