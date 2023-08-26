package logic

import (
	"context"

	"MiniDY/app/interaction/cmd/rpc/internal/svc"
	"MiniDY/app/interaction/cmd/rpc/pb"
	_favorite "MiniDY/app/interaction/model/favorite"
	"MiniDY/app/videos/model"
	"MiniDY/common/dyerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	LIKE    = 1
	DISLIKE = 2
)

type VideoInteractionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoInteractionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoInteractionLogic {
	return &VideoInteractionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞
func (l *VideoInteractionLogic) VideoInteraction(in *pb.DouyinFavoriteActionRequest) (*pb.DouyinFavoriteActionResponse, error) {
	// todo: add your logic here and delete this line
	//先判断下video还在不在
	videoId, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get video err : %v , video :%+v", err, in.VideoId)
	}
	if videoId != nil {
		switch in.ActionType {
		case LIKE:
			if err := l.svcCtx.FavoriteModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
				//修改favorite表 先查找是否有点赞记录，有就直接修改，否则插一条新的
				userTovideo, err := l.svcCtx.FavoriteModel.FindOneByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
				if err != nil && err != model.ErrNotFound {
					return errors.Wrapf(dyerr.ErrDBerror, "Failed to get user and video err : %v , user: %+v ,video :%+v", err, in.UserId, in.VideoId)
				}
				if userTovideo != nil {
					userTovideo.Status = 1
					err := l.svcCtx.FavoriteModel.Update(l.ctx, userTovideo)
					if err != nil {
						return err
					}
				} else {
					newFavorite := new(_favorite.Favorite)
					newFavorite.UserId = in.UserId
					newFavorite.VideoId = in.VideoId
					newFavorite.Status = 1

					insertResult, err := l.svcCtx.FavoriteModel.Insert(l.ctx, newFavorite)
					if err != nil {
						return errors.Wrapf(dyerr.ErrDBerror, "new favorite Insert err:%v,favorite:%+v", err, newFavorite)
					}
					lastId, err := insertResult.LastInsertId()
					if err != nil {
						return errors.Wrapf(dyerr.ErrDBerror, "new favorite insertResult.LastInsertId err:%v,favorite:%+v", err, newFavorite)
					}
					newFavorite.Id = lastId
				}

				return nil
			}); err != nil {
				return nil, err
			}
		case DISLIKE:
			if err := l.svcCtx.FavoriteModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
				//修改favorite表 先查找是否有点赞记录，有就直接修改，否则插一条新的
				userTovideo, err := l.svcCtx.FavoriteModel.FindOneByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
				if err != nil && err != model.ErrNotFound {
					return errors.Wrapf(dyerr.ErrDBerror, "Failed to get user and video err : %v , user: %+v ,video :%+v", err, in.UserId, in.VideoId)
				}
				if userTovideo != nil {
					userTovideo.Status = 2
					err := l.svcCtx.FavoriteModel.Update(l.ctx, userTovideo)
					if err != nil {
						return err
					}
				} else {
					newFavorite := new(_favorite.Favorite)
					newFavorite.UserId = in.UserId
					newFavorite.VideoId = in.VideoId
					newFavorite.Status = 2

					insertResult, err := l.svcCtx.FavoriteModel.Insert(l.ctx, newFavorite)
					if err != nil {
						return errors.Wrapf(dyerr.ErrDBerror, "new favorite Insert err:%v,favorite:%+v", err, newFavorite)
					}
					lastId, err := insertResult.LastInsertId()
					if err != nil {
						return errors.Wrapf(dyerr.ErrDBerror, "new favorite insertResult.LastInsertId err:%v,favorite:%+v", err, newFavorite)
					}
					newFavorite.Id = lastId
				}

				return nil
			}); err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}

	return &pb.DouyinFavoriteActionResponse{}, nil
}
