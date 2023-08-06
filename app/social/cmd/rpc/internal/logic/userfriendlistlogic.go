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

type UserFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFriendListLogic {
	return &UserFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户好友列表
func (l *UserFriendListLogic) UserFriendList(in *pb.DouyinRelationFriendListRequest) (*pb.DouyinRelationFriendListResponse, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.FollowModel.SelectBuilder().Where(squirrel.Eq{"user_id": in.UserId})

	friendsIdList, err := l.svcCtx.FollowModel.FindAllFriends(l.ctx, whereBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get user friends list err : %v , user :%+v", err, in.UserId)
	}
	//把好友的ID返回给api处理
	return &pb.DouyinRelationFriendListResponse{
		UserList: friendsIdList,
	}, nil
}
