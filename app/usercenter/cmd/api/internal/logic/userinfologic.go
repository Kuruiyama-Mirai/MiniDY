package logic

import (
	"context"

	"MiniDY/app/usercenter/cmd/api/internal/svc"
	"MiniDY/app/usercenter/cmd/api/internal/types"
	"MiniDY/app/usercenter/cmd/rpc/pb"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.GetUserInfoReq) (*types.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	userInfoResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &pb.DouyinUserRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		return nil, err
	}
	var userInfo types.User
	_ = copier.Copy(&userInfo, userInfoResp)

	return &types.GetUserInfoResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  dyerr.SUCCESS,
		UserInfo:   userInfo,
	}, nil

}
