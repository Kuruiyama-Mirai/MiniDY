package logic

import (
	"context"

	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"
	"MiniDY/app/interaction/cmd/rpc/interactsservice"
	"MiniDY/common/ctxdata"
	"MiniDY/common/dyerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLikeLogic {
	return &ActionLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionLikeLogic) ActionLike(req *types.InteractionReq) (*types.InteractionResp, error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)

	_, err := l.svcCtx.InteractsRpc.VideoInteraction(l.ctx, &interactsservice.DouyinFavoriteActionRequest{
		UserId:     userId,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return &types.InteractionResp{
			StatusCode: 1,
			StatusMsg:  "点赞失败" + err.Error(),
		}, nil
	}
	return &types.InteractionResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  "点赞" + dyerr.SUCCESS,
	}, nil
}
