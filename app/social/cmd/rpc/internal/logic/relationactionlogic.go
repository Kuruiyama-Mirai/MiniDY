package logic

import (
	"context"
	"database/sql"

	"MiniDY/app/social/cmd/rpc/internal/svc"
	"MiniDY/app/social/cmd/rpc/pb"
	"MiniDY/app/social/model"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/common/dyerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	FOLLOW = 1 //关注
	CANCEL = 2 //取消关注
)

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关系操作
func (l *RelationActionLogic) RelationAction(in *pb.DouyinRelationActionRequest) (*pb.DouyinRelationActionResponse, error) {
	// todo: add your logic here and delete this line

	err := l.chechNum(in.UserId, in.ToUserId)
	if err != nil {
		return nil, err
	}
	switch in.ActionType {
	case FOLLOW:
		if err := l.svcCtx.FollowModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			//1 修改关系数据表
			newFollow := new(model.Follow)
			newFollow.UserId = in.UserId
			newFollow.ToUserId = in.ToUserId
			//反向查找一下这条互关数据存不存在
			isStatus, _ := l.svcCtx.FollowModel.FindOneByUserIdAndFollowId(l.ctx, in.ToUserId, in.UserId)
			if isStatus != nil {
				//如果关注人也关注了自己
				//同时将两个设置为互相关注
				newFollow.Status = 1
				isStatus.Status = 1
			}
			insertResult, err := l.svcCtx.FollowModel.Insert(l.ctx, newFollow)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "new follow Insert err:%v,follow:%+v", err, newFollow)
			}
			lastId, err := insertResult.LastInsertId()
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "new follow insertResult.LastInsertId err:%v,follow:%+v", err, newFollow)
			}
			newFollow.Id = lastId
			//将被关注者的状态也更新一下
			err = l.svcCtx.FollowModel.Update(l.ctx, isStatus)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "new follow update err:%v,follow:%+v", err, isStatus)
			}

			//2修改关注者与被关注者的信息
			//关注+1
			newUser, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
			newUser.FollowCount += 1
			//粉丝+1
			newToUser, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.ToUserId)
			newToUser.IsFollow = sql.NullInt64{Int64: 1, Valid: true}
			newToUser.FollowerCount += 1

			err = l.svcCtx.UserModel.Update(l.ctx, newUser)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "new user update follow err:%v,user:%+v", err, newUser)
			}
			err = l.svcCtx.UserModel.Update(l.ctx, newToUser)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "new user update follower err:%v,follow:%+v", err, newToUser)
			}

			return nil
		}); err != nil {
			return nil, err
		}
	case CANCEL:
		if err := l.svcCtx.FollowModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			//1 修改关系数据表
			sqlResult, err := l.svcCtx.FollowModel.FindOneByUserIdAndFollowId(l.ctx, in.UserId, in.ToUserId)
			if err != nil && err != model.ErrNotFound {
				return errors.Wrapf(dyerr.ErrDBerror, "sql not exist err:%v", err)
			}
			err = l.svcCtx.FollowModel.Delete(l.ctx, sqlResult.Id)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "delete sql err:%v", err)
			}
			//取关后要将被关注者的互关状态取消
			isStatus, err := l.svcCtx.FollowModel.FindOneByUserIdAndFollowId(l.ctx, in.ToUserId, in.UserId)
			if err != nil && isStatus != nil {
				isStatus.Status = 0
			}

			err = l.svcCtx.FollowModel.Update(l.ctx, isStatus)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "update follow status err:%v", err)
			}
			//2修改关注者与被关注者的信息
			//关注-1
			newUser, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
			newUser.FollowCount -= 1
			//粉丝-1
			newToUser, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.ToUserId)
			newToUser.IsFollow = sql.NullInt64{Int64: 0, Valid: false}
			newToUser.FollowerCount -= 1

			err = l.svcCtx.UserModel.Update(l.ctx, newUser)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "new user update follow err:%v,user:%+v", err, newUser)
			}
			err = l.svcCtx.UserModel.Update(l.ctx, newToUser)
			if err != nil {
				return errors.Wrapf(dyerr.ErrDBerror, "new user update follower err:%v,follow:%+v", err, newToUser)
			}

			return nil
		}); err != nil {
			return nil, err
		}

	default:
		return nil, err
	}

	return &pb.DouyinRelationActionResponse{}, nil
}

// 检查关注是否合法
func (l *RelationActionLogic) chechNum(userId, userToId int64) error {
	//判断关注用户是否存在
	user_to_id, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
		UserId: userToId,
	})
	if err != nil && err != model.ErrNotFound {
		return errors.Wrapf(dyerr.ErrDBerror, "GetUserInfo find user db err , id:%d , err:%v", userToId, err)
	}
	if user_to_id == nil {
		return errors.New("关注用户不存在")
	}
	if userId == userToId {
		return errors.New("不能关注自己")
	}
	return nil
}
