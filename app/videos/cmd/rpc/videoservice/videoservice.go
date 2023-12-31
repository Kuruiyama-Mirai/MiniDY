// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoservice

import (
	"context"

	"MiniDY/app/videos/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Author                    = pb.Author
	DouyinFeedRequest         = pb.DouyinFeedRequest
	DouyinFeedResponse        = pb.DouyinFeedResponse
	DouyinPublishListRequest  = pb.DouyinPublishListRequest
	DouyinPublishListResponse = pb.DouyinPublishListResponse
	Video                     = pb.Video

	Videoservice interface {
		// 获取视频流
		GetVideoFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error)
		// 视频投稿
		PublishVideoList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error)
	}

	defaultVideoservice struct {
		cli zrpc.Client
	}
)

func NewVideoservice(cli zrpc.Client) Videoservice {
	return &defaultVideoservice{
		cli: cli,
	}
}

// 获取视频流
func (m *defaultVideoservice) GetVideoFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error) {
	client := pb.NewVideoserviceClient(m.cli.Conn())
	return client.GetVideoFeed(ctx, in, opts...)
}

// 视频投稿
func (m *defaultVideoservice) PublishVideoList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error) {
	client := pb.NewVideoserviceClient(m.cli.Conn())
	return client.PublishVideoList(ctx, in, opts...)
}
