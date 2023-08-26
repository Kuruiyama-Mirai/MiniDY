package logic

import (
	"context"

	"MiniDY/app/interaction/cmd/rpc/internal/svc"
	"MiniDY/app/interaction/cmd/rpc/pb"
	_comment "MiniDY/app/interaction/model/comment"
	"MiniDY/app/videos/model"
	"MiniDY/common/dyerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	POST_COMMENTS = 1 //评论
	DEL_COMMENTS  = 2 //删除评论
)

type VideoCommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoCommentActionLogic {
	return &VideoCommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论
func (l *VideoCommentActionLogic) VideoCommentAction(in *pb.DouyinCommentActionRequest) (*pb.DouyinCommentActionResponse, error) {
	// todo: add your logic here and delete this line
	//先判断下video还在不在
	videoId, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get video err : %v , video :%+v", err, in.VideoId)
	}
	newComment := new(_comment.Comment)
	if videoId != nil {
		switch in.ActionType {
		case POST_COMMENTS:
			if err := l.svcCtx.FavoriteModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
				newComment.VideoId = in.UserId
				newComment.UserId = in.UserId
				newComment.Content = *in.CommentText

				insertResult, err := l.svcCtx.CommentModel.Insert(l.ctx, newComment)
				if err != nil {
					return errors.Wrapf(dyerr.ErrDBerror, "new comment Insert err:%v,comment:%+v", err, newComment)
				}
				lastId, err := insertResult.LastInsertId()
				if err != nil {
					return errors.Wrapf(dyerr.ErrDBerror, "new comment insertResult.LastInsertId err:%v,comment:%+v", err, newComment)
				}
				newComment.Id = lastId

				return nil
			}); err != nil {
				return nil, err
			}
		case DEL_COMMENTS:
			if err := l.svcCtx.FavoriteModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
				comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, *in.CommentId)
				if err != nil && err != model.ErrNotFound {
					return errors.Wrapf(dyerr.ErrDBerror, "Failed to get user comment err : %v ,  comment :%+v", err, *in.CommentId)
				}
				if comment != nil {
					err = l.svcCtx.CommentModel.Delete(l.ctx, *in.CommentId)
					if err != nil {
						return err
					}
				}
				return nil
			}); err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}

	//把拿到的commentId返回给api处理
	if in.ActionType == POST_COMMENTS {
		return &pb.DouyinCommentActionResponse{
			Comment: newComment.Id,
		}, nil
	} else {
		return &pb.DouyinCommentActionResponse{}, nil
	}

}
