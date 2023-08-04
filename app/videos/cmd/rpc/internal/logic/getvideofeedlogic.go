package logic

import (
	"context"
	"time"

	"MiniDY/app/usercenter/cmd/rpc/usercenter"
	"MiniDY/app/videos/cmd/rpc/internal/svc"
	"MiniDY/app/videos/cmd/rpc/pb"
	"MiniDY/app/videos/model"
	"MiniDY/common/dyerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoFeedLogic {
	return &GetVideoFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取视频流
func (l *GetVideoFeedLogic) GetVideoFeed(in *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {
	// todo: add your logic here and delete this line
	if in.LatestTime == 0 {
		in.LatestTime = time.Now().Unix()
	}
	whereBuilder := l.svcCtx.VideoModel.SelectBuilder().Where(squirrel.LtOrEq{"create_time": in.LatestTime})
	list, err := l.svcCtx.VideoModel.FindPageListByTimeDESC(l.ctx, whereBuilder)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get video list err : %v , in :%+v", err, in)
	}

	var resp []*pb.Video
	if len(list) > 0 {
		for _, videos := range list {
			var pbVideo pb.Video
			userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
				UserId: videos.AuthorId,
			})
			if err != nil && err != model.ErrNotFound {
				return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get user video list err : %v , user :%+v", err, videos.AuthorId)
			}
			author := pb.Author{
				Id:              userInfo.User.Id,
				Username:        userInfo.User.Username,
				FollowCount:     userInfo.User.FollowCount,
				FollowerCount:   userInfo.User.FollowerCount,
				IsFollow:        userInfo.User.IsFollow,
				Avatar:          userInfo.User.Avatar,
				BackgroundImage: userInfo.User.BackgroundImage,
				Signature:       userInfo.User.Signature,
				TotalFavorited:  userInfo.User.TotalFavorited,
				WorkCount:       userInfo.User.WorkCount,
				FavoriteCount:   userInfo.User.FavoriteCount,
			}
			pbVideo.Author = &author
			_ = copier.Copy(&pbVideo, videos)

			resp = append(resp, &pbVideo)
		}
	}

	//lastVideo, _ := l.svcCtx.VideoModel.FindOne(l.ctx, resp[0].Id)
	//lastTime := lastVideo.UpdateTime.Unix()
	lastTime := time.Now().Unix()
	return &pb.DouyinFeedResponse{
		VideoList: resp,
		NextTime:  lastTime,
	}, nil
}
