package logic

import (
	"context"

	"MiniDY/app/videos/cmd/api/internal/svc"
	"MiniDY/app/videos/cmd/api/internal/types"
	"MiniDY/app/videos/cmd/rpc/videoservice"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type VideofeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideofeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideofeedLogic {
	return &VideofeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideofeedLogic) Videofeed(req *types.VideGetFeedReq) (*types.VideGetFeedResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.VideoRpc.GetVideoFeed(l.ctx, &videoservice.DouyinFeedRequest{
		LatestTime: req.LatestTime,
	})
	if err != nil {
		return nil, errors.Wrapf(dyerr.NewErrMsg("Failed to get video list"), "Failed to get video list err : %v ,req:%+v", err, req)
	}

	var videoList []types.Video

	if len(resp.VideoList) > 0 {
		for _, video := range resp.VideoList {
			var typeVideo types.Video
			_ = copier.Copy(&typeVideo, video)
			videoList = append(videoList, typeVideo)
		}
	}
	return &types.VideGetFeedResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  dyerr.SUCCESS,
		NextTime:   resp.NextTime,
		List:       videoList,
	}, nil
}
