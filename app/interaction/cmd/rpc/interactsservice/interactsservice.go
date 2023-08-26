// Code generated by goctl. DO NOT EDIT.
// Source: interacts.proto

package interactsservice

import (
	"context"

	"MiniDY/app/interaction/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment                      = pb.Comment
	DouyinCommentActionRequest   = pb.DouyinCommentActionRequest
	DouyinCommentActionResponse  = pb.DouyinCommentActionResponse
	DouyinCommentListRequest     = pb.DouyinCommentListRequest
	DouyinCommentListResponse    = pb.DouyinCommentListResponse
	DouyinFavoriteActionRequest  = pb.DouyinFavoriteActionRequest
	DouyinFavoriteActionResponse = pb.DouyinFavoriteActionResponse
	DouyinFavoriteListRequest    = pb.DouyinFavoriteListRequest
	DouyinFavoriteListResponse   = pb.DouyinFavoriteListResponse
	User                         = pb.User

	Interactsservice interface {
		// 点赞
		VideoInteraction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error)
		// 喜欢列表
		VideoInteractionList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error)
		// 评论
		VideoCommentAction(ctx context.Context, in *DouyinCommentActionRequest, opts ...grpc.CallOption) (*DouyinCommentActionResponse, error)
		// 评论列表
		VideoCommentList(ctx context.Context, in *DouyinCommentListRequest, opts ...grpc.CallOption) (*DouyinCommentListResponse, error)
	}

	defaultInteractsservice struct {
		cli zrpc.Client
	}
)

func NewInteractsservice(cli zrpc.Client) Interactsservice {
	return &defaultInteractsservice{
		cli: cli,
	}
}

// 点赞
func (m *defaultInteractsservice) VideoInteraction(ctx context.Context, in *DouyinFavoriteActionRequest, opts ...grpc.CallOption) (*DouyinFavoriteActionResponse, error) {
	client := pb.NewInteractsserviceClient(m.cli.Conn())
	return client.VideoInteraction(ctx, in, opts...)
}

// 喜欢列表
func (m *defaultInteractsservice) VideoInteractionList(ctx context.Context, in *DouyinFavoriteListRequest, opts ...grpc.CallOption) (*DouyinFavoriteListResponse, error) {
	client := pb.NewInteractsserviceClient(m.cli.Conn())
	return client.VideoInteractionList(ctx, in, opts...)
}

// 评论
func (m *defaultInteractsservice) VideoCommentAction(ctx context.Context, in *DouyinCommentActionRequest, opts ...grpc.CallOption) (*DouyinCommentActionResponse, error) {
	client := pb.NewInteractsserviceClient(m.cli.Conn())
	return client.VideoCommentAction(ctx, in, opts...)
}

// 评论列表
func (m *defaultInteractsservice) VideoCommentList(ctx context.Context, in *DouyinCommentListRequest, opts ...grpc.CallOption) (*DouyinCommentListResponse, error) {
	client := pb.NewInteractsserviceClient(m.cli.Conn())
	return client.VideoCommentList(ctx, in, opts...)
}