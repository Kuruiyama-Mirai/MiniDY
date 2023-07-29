package logic

import (
	"context"

	"MiniDY/app/usercenter/cmd/rpc/internal/svc"
	"MiniDY/app/usercenter/cmd/rpc/pb"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/app/usercenter/model"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "GetUserInfo find user db err , id:%d , err:%v", in.UserId, err)
	}
	if user == nil {
		return nil, errors.Wrapf(dyerr.ErrUserNoExistsError, "id:%d", in.UserId)
	}
	var respUser usercenter.User
	_ = copier.Copy(&respUser, user)

	return &pb.DouyinUserResponse{
		User: &respUser,
	}, nil
}
