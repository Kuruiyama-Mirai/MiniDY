package logic

import (
	"context"

	"MiniDY/app/videos/cmd/api/internal/svc"
	"MiniDY/app/videos/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishVideoReq) (resp *types.PublishVideoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
