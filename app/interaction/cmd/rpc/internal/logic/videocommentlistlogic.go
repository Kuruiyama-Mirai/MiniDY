package logic

import (
	"context"

	"MiniDY/app/interaction/cmd/rpc/internal/svc"
	"MiniDY/app/interaction/cmd/rpc/pb"
	"MiniDY/app/videos/model"
	"MiniDY/common/dyerr"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type VideoCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoCommentListLogic {
	return &VideoCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论列表
func (l *VideoCommentListLogic) VideoCommentList(in *pb.DouyinCommentListRequest) (*pb.DouyinCommentListResponse, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.CommentModel.SelectBuilder().Where(squirrel.Eq{"video_id": in.VideoId})
	commentIdList, err := l.svcCtx.CommentModel.FindAllCommentsByVideoId(l.ctx, whereBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get comment  list err : %v , video :%+v", err, in.VideoId)
	}
	//把返回的commentlistid返回给api处理
	return &pb.DouyinCommentListResponse{
		CommentList: commentIdList,
	}, nil
}
