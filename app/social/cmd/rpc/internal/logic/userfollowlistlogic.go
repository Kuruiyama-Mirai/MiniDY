package logic

import (
	"context"

	"MiniDY/app/social/cmd/rpc/internal/svc"
	"MiniDY/app/social/cmd/rpc/pb"
	"MiniDY/app/social/model"
	"MiniDY/common/dyerr"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowListLogic {
	return &UserFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户关注列表
func (l *UserFollowListLogic) UserFollowList(in *pb.DouyinRelationFollowListRequest) (*pb.DouyinRelationFollowListResponse, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.FollowModel.SelectBuilder().Where(squirrel.Eq{"user_id": in.UserId})

	followIdList, err := l.svcCtx.FollowModel.FindAllFollow(l.ctx, whereBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get user follow list err : %v , user :%+v", err, in.UserId)
	}
	//把关注的ID返回给api处理
	return &pb.DouyinRelationFollowListResponse{
		UserList: followIdList,
	}, nil
}
