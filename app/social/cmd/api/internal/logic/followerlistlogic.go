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

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
	// todo: add your logic here and delete this line
	//从ctx拿到用户Id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	followerIdList, err := l.svcCtx.FollowRpc.UserFollowerList(l.ctx, &followerservice.DouyinRelationFollowerListRequest{
		UserId: userId,
	})
	if err != nil {
		return &types.FollowerListResp{
			StatusCode: 1,
			StatusMsg:  "查询粉丝列表失败" + err.Error(),
			UserList:   nil,
		}, nil
	}

	var followerList []types.User
	if len(followerIdList.UserList) > 0 {
		for _, followerId := range followerIdList.UserList {
			pbUserInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
				UserId: followerId,
			})
			if err != nil {
				return nil, err
			}
			var temp types.User
			_ = copier.Copy(&temp, pbUserInfo.User)
			followerList = append(followerList, temp)
		}
	}
	return &types.FollowerListResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  "查询关注列表" + dyerr.SUCCESS,
		UserList:   followerList,
	}, nil
}
