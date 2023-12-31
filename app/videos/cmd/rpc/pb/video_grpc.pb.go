// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0--rc3
// source: video.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VideoserviceClient is the client API for Videoservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoserviceClient interface {
	// 获取视频流
	GetVideoFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error)
	// 视频投稿
	// rpc publishVideo(douyin_publish_action_request) returns(douyin_publish_action_response);
	// 获取投稿列表
	PublishVideoList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error)
}

type videoserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoserviceClient(cc grpc.ClientConnInterface) VideoserviceClient {
	return &videoserviceClient{cc}
}

func (c *videoserviceClient) GetVideoFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error) {
	out := new(DouyinFeedResponse)
	err := c.cc.Invoke(ctx, "/pb.videoservice/getVideoFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoserviceClient) PublishVideoList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error) {
	out := new(DouyinPublishListResponse)
	err := c.cc.Invoke(ctx, "/pb.videoservice/publishVideoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoserviceServer is the server API for Videoservice service.
// All implementations must embed UnimplementedVideoserviceServer
// for forward compatibility
type VideoserviceServer interface {
	// 获取视频流
	GetVideoFeed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error)
	// 视频投稿
	// rpc publishVideo(douyin_publish_action_request) returns(douyin_publish_action_response);
	// 获取投稿列表
	PublishVideoList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error)
	mustEmbedUnimplementedVideoserviceServer()
}

// UnimplementedVideoserviceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoserviceServer struct {
}

func (UnimplementedVideoserviceServer) GetVideoFeed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoFeed not implemented")
}
func (UnimplementedVideoserviceServer) PublishVideoList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideoList not implemented")
}
func (UnimplementedVideoserviceServer) mustEmbedUnimplementedVideoserviceServer() {}

// UnsafeVideoserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoserviceServer will
// result in compilation errors.
type UnsafeVideoserviceServer interface {
	mustEmbedUnimplementedVideoserviceServer()
}

func RegisterVideoserviceServer(s grpc.ServiceRegistrar, srv VideoserviceServer) {
	s.RegisterService(&Videoservice_ServiceDesc, srv)
}

func _Videoservice_GetVideoFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoserviceServer).GetVideoFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoservice/getVideoFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoserviceServer).GetVideoFeed(ctx, req.(*DouyinFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Videoservice_PublishVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoserviceServer).PublishVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.videoservice/publishVideoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoserviceServer).PublishVideoList(ctx, req.(*DouyinPublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Videoservice_ServiceDesc is the grpc.ServiceDesc for Videoservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Videoservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.videoservice",
	HandlerType: (*VideoserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getVideoFeed",
			Handler:    _Videoservice_GetVideoFeed_Handler,
		},
		{
			MethodName: "publishVideoList",
			Handler:    _Videoservice_PublishVideoList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
