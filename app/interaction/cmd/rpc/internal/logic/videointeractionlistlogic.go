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

type VideoInteractionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoInteractionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoInteractionListLogic {
	return &VideoInteractionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 喜欢列表
func (l *VideoInteractionListLogic) VideoInteractionList(in *pb.DouyinFavoriteListRequest) (*pb.DouyinFavoriteListResponse, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.FavoriteModel.SelectBuilder().Where(squirrel.Eq{"user_id": in.UserId})

	favoriteVideoList, err := l.svcCtx.FavoriteModel.FindAllFavoriteByUserId(l.ctx, whereBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get user favorite video list err : %v , user :%+v", err, in.UserId)
	}

	//把视频的ID返回给api处理
	return &pb.DouyinFavoriteListResponse{
		VideoList: favoriteVideoList,
	}, nil
}
