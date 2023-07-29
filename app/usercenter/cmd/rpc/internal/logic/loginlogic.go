package logic

import (
	"context"

	"MiniDY/app/usercenter/cmd/rpc/internal/svc"
	"MiniDY/app/usercenter/cmd/rpc/pb"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/app/usercenter/model"
	"MiniDY/common/dyerr"
	"MiniDY/common/tool"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "根据用户名查询用户信息失败,name:%s,err:%v", in.Username, err)
	}
	if user == nil {
		return nil, errors.Wrapf(dyerr.ErrUserNoExistsError, "username:%s", in.Username)
	}
	if !(tool.Md5ByString(in.Password) == user.Password) {
		return nil, errors.Wrap(dyerr.ErrUsernamePwdError, "密码不对")
	}
	userId := user.Id
	//生成Token 不要在服务内部再调用rpc
	NGTL := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := NGTL.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(dyerr.ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}
	return &pb.DouyinUserLoginResponse{
		UserId: userId,
		Token:  token.AccessToken,
	}, nil
}
