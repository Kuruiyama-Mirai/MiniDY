package logic

import (
	"context"

	"MiniDY/app/usercenter/cmd/api/internal/svc"
	"MiniDY/app/usercenter/cmd/api/internal/types"
	"MiniDY/app/usercenter/cmd/rpc/pb"
	"MiniDY/common/dyerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &pb.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &types.LoginResp{
			StatusCode: 1,
			StatusMsg:  "登录失败" + err.Error(),
		}, nil
	}

	return &types.LoginResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  dyerr.SUCCESS,
		UserID:     loginResp.UserId,
		Token:      loginResp.Token,
	}, nil
}
