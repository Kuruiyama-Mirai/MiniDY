package logic

import (
	"context"

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

type PublishVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoListLogic {
	return &PublishVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取投稿列表
func (l *PublishVideoListLogic) PublishVideoList(in *pb.DouyinPublishListRequest) (*pb.DouyinPublishListResponse, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.VideoModel.SelectBuilder().Where(squirrel.Eq{"author_id": in.UserId})

	list, err := l.svcCtx.VideoModel.FindVideoListByAuthorId(l.ctx, whereBuilder)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get user video list err : %v , user :%+v", err, in.UserId)
	}

	userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.DouyinUserRequest{
		UserId: in.UserId,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(dyerr.ErrDBerror, "Failed to get user video list err : %v , user :%+v", err, in.UserId)
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

	var resp []*pb.Video
	if len(list) > 0 {
		for _, videos := range list {
			var pbVideo pb.Video
			pbVideo.Author = &author
			_ = copier.Copy(&pbVideo, videos)

			resp = append(resp, &pbVideo)
		}
	}
	return &pb.DouyinPublishListResponse{
		VideoList: resp,
	}, nil
}
