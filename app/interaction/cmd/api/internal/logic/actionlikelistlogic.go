package logic

import (
	"context"

	"MiniDY/app/interaction/cmd/api/internal/svc"
	"MiniDY/app/interaction/cmd/api/internal/types"
	"MiniDY/app/interaction/cmd/rpc/interactsservice"
	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/common/ctxdata"
	"MiniDY/common/dyerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLikeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLikeListLogic {
	return &ActionLikeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionLikeListLogic) ActionLikeList(req *types.InteractionListReq) (*types.InteractionListResp, error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	favoriteVideoList, err := l.svcCtx.InteractsRpc.VideoInteractionList(l.ctx, &interactsservice.DouyinFavoriteListRequest{
		UserId: userId,
	})
	if err != nil {
		return &types.InteractionListResp{
			StatusCode: 1,
			StatusMsg:  "查询喜欢列表失败" + err.Error(),
			Video_list: nil,
		}, nil
	}

	var resp []*types.Video
	if len(favoriteVideoList.VideoList) > 0 {
		for _, videoId := range favoriteVideoList.VideoList {
			var tempVideo types.Video
			pbVideo, _ := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
			pbUserInfo, _ := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
				UserId: pbVideo.AuthorId,
			})
			var tempUser types.User
			_ = copier.Copy(&tempUser, pbUserInfo.User)

			tempVideo.Author = tempUser

			_ = copier.Copy(&tempVideo, pbVideo)

			resp = append(resp, &tempVideo)
		}
	}
	return &types.InteractionListResp{
		StatusCode: int32(dyerr.OK),
		StatusMsg:  "查询喜欢列表" + dyerr.SUCCESS,
		Video_list: resp,
	}, nil
}
