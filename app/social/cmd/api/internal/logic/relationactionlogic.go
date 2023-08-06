package logic

import (
	"context"

	"MiniDY/app/social/cmd/api/internal/svc"
	"MiniDY/app/social/cmd/api/internal/types"
	"MiniDY/app/social/cmd/rpc/followerservice"
	"MiniDY/common/ctxdata"
	"MiniDY/common/dyerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationactionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationactionLogic {
	return &RelationactionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationactionLogic) Relationaction(req *types.ActionReq) (*types.ActionResp, error) {
	// todo: add your logic here and delete this line
	//从ctx拿到用户Id
	userId := ctxdata.GetUidFromCtx(l.ctx)

	_, err := l.svcCtx.FollowRpc.RelationAction(l.ctx, &followerservice.DouyinRelationActionRequest{
		UserId:     userId,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return &types.ActionResp{
			StatusCode: 1,
			StatusMsg:  "操作失败,请稍后再试" + err.Error(),
		}, nil
	}

	return &types.ActionResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  dyerr.SUCCESS,
	}, nil
}
