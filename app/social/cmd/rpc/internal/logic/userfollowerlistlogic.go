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

type UserFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowerListLogic {
	return &UserFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户粉丝列表
func (l *UserFollowerListLogic) UserFollowerList(in *pb.DouyinRelationFollowerListRequest) (*pb.DouyinRelationFollowerListResponse, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.FollowModel.SelectBuilder().Where(squirrel.Eq{"to_user_id": in.UserId})

	followerIdList, err := l.svcCtx.FollowModel.FindAllFollower(l.ctx, whereBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get user follower list err : %v , user :%+v", err, in.UserId)
	}
	//把粉丝的ID返回给api处理
	return &pb.DouyinRelationFollowerListResponse{
		UserList: followerIdList,
	}, nil
}
