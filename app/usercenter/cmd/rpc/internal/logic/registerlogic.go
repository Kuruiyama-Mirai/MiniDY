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
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {
	// todo: add your logic here and delete this line
	//先检查用户是否已经存在
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "username:%s,err:%v", in.Username, err)
	}
	if user != nil {
		return nil, errors.Wrapf(dyerr.ErrUserAlreadyRegisterError, "Register user exists username:%s,err:%v", in.Username, err)
	}

	var userId int64
	//将数据库的事务处理封装成一个函数调用，通过将 fn 函数作为参数传递给 Trans 方法，可以在 fn 函数中执行一系列的数据库操作，这些操作将在一个事务中执行，并且出现任何错误时，事务将自动回滚。
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		//创建新的用户数据库实例对象
		user := new(model.User)
		user.Username = in.Username
		//对密码加密处理
		if len(in.Password) > 0 {
			user.Password = tool.Md5ByString(in.Password)
		}

		insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, user)
		if err != nil {
			return errors.Wrapf(dyerr.ErrDBerror, "Register db user Insert err:%v,user:%+v", err, user)
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(dyerr.ErrDBerror, "Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
		}
		userId = lastId
		return nil
	}); err != nil {
		return nil, err
	}

	//生成Token实例 不要在服务内部再调用rpc
	NGTL := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := NGTL.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(dyerr.ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &pb.DouyinUserRegisterResponse{
		UserId: userId,
		Token:  token.AccessToken,
	}, nil
}
