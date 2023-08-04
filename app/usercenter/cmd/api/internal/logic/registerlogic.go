package logic

import (
	"context"

	"MiniDY/app/usercenter/cmd/api/internal/svc"
	"MiniDY/app/usercenter/cmd/api/internal/types"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/common/dyerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &types.RegisterResp{
			StatusCode: 1,
			StatusMsg:  "注册失败" + err.Error(),
		}, nil
	}

	return &types.RegisterResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  dyerr.SUCCESS,
		UserID:     registerResp.UserId,
		Token:      registerResp.Token,
	}, nil
}
