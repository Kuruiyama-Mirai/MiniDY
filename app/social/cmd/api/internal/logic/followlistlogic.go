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

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListReq) (resp *types.FollowListResp, err error) {
	// todo: add your logic here and delete this line
	//从ctx拿到用户Id
	userId := ctxdata.GetUidFromCtx(l.ctx)
	followIdList, err := l.svcCtx.FollowRpc.UserFollowList(l.ctx, &followerservice.DouyinRelationFollowListRequest{
		UserId: userId,
	})
	if err != nil {
		return &types.FollowListResp{
			StatusCode: 1,
			StatusMsg:  "查询关注列表失败" + err.Error(),
			UserList:   nil,
		}, nil
	}
	var followList []types.User
	if len(followIdList.UserList) > 0 {
		for _, followId := range followIdList.UserList {
			pbUserInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
				UserId: followId,
			})
			if err != nil {
				return nil, err
			}
			var temp types.User
			_ = copier.Copy(&temp, pbUserInfo.User)
			followList = append(followList, temp)
		}
	}
	return &types.FollowListResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  "查询关注列表" + dyerr.SUCCESS,
		UserList:   followList,
	}, nil
}
