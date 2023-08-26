package svc

import (
	"MiniDY/app/interaction/cmd/rpc/internal/config"
	_comment "MiniDY/app/interaction/model/comment"
	_favorite "MiniDY/app/interaction/model/favorite"
	"MiniDY/app/videos/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	CommentModel  _comment.CommentModel
	FavoriteModel _favorite.FavoriteModel
	VideoModel    model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
		}),
		CommentModel:  _comment.NewCommentModel(sqlconn, c.Cache),
		FavoriteModel: _favorite.NewFavoriteModel(sqlconn, c.Cache),
		VideoModel:    model.NewVideoModel(sqlconn, c.Cache),
	}
}
