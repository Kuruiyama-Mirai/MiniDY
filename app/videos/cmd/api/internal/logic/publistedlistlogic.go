package logic

import (
	"context"

	"github.com/pkg/errors"

	"MiniDY/app/videos/cmd/api/internal/svc"
	"MiniDY/app/videos/cmd/api/internal/types"
	"MiniDY/app/videos/cmd/rpc/videoservice"
	"MiniDY/app/videos/model"
	"MiniDY/common/ctxdata"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublistedListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublistedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublistedListLogic {
	return &PublistedListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublistedListLogic) PublistedList(req *types.PublishedVideoListReq) (*types.PublishedVideoListResp, error) {
	// todo: add your logic here and delete this line
	//从上下文ctx拿到UserID
	userId := ctxdata.GetUidFromCtx(l.ctx)
	if userId != req.AuthorId {
		return &types.PublishedVideoListResp{
			StatusCode: 1,
			StatusMsg:  "当前用户未认证",
		}, nil
	}

	resp, err := l.svcCtx.VideoRpc.PublishVideoList(l.ctx, &videoservice.DouyinPublishListRequest{
		UserId: req.AuthorId,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.NewErrMsg("Failed to get user video list"), "Failed to get user video list err : %v ,req:%+v", err, req)
	}

	var videoList []types.Video
	if len(resp.VideoList) > 0 {
		for _, video := range resp.VideoList {
			var typeVideo types.Video
			_ = copier.Copy(&typeVideo, video)
		}
	}
	return &types.PublishedVideoListResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  dyerr.SUCCESS,
		List:       videoList,
	}, nil
}
