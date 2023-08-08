package logic

import (
	"context"

	"MiniDY/app/social/cmd/api/internal/svc"
	"MiniDY/app/social/cmd/api/internal/types"
	"MiniDY/app/social/cmd/rpc/followerservice"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/common/ctxdata"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendsListLogic {
	return &FriendsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendsListLogic) FriendsList(req *types.FriendsListReq) (resp *types.FriendsListResp, err error) {
	// todo: add your logic here and delete this line
	//从ctx拿到用户Id
	userId := ctxdata.GetUidFromCtx(l.ctx)

	friendsIdList, err := l.svcCtx.FollowRpc.UserFriendList(l.ctx, &followerservice.DouyinRelationFriendListRequest{
		UserId: userId,
	})
	if err != nil {
		return &types.FriendsListResp{
			StatusCode: 1,
			StatusMsg:  "查询好友列表失败" + err.Error(),
			UserList:   nil,
		}, nil
	}

	var friendsList []types.User
	if len(friendsIdList.UserList) > 0 {
		for _, friendsId := range friendsIdList.UserList {
			pbUserInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
				UserId: friendsId,
			})
			if err != nil {
				return nil, err
			}
			var temp types.User
			_ = copier.Copy(&temp, pbUserInfo.User)
			friendsList = append(friendsList, temp)
		}
	}
	return &types.FriendsListResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  "查询好友列表" + dyerr.SUCCESS,
		UserList:   friendsList,
	}, nil

}
