// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0--rc3
// source: follow.proto

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

// FollowerserviceClient is the client API for Followerservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowerserviceClient interface {
	// 关系操作
	RelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error)
	// 用户关注列表
	UserFollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error)
	// 用户粉丝列表
	UserFollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error)
	// 用户好友列表
	UserFriendList(ctx context.Context, in *DouyinRelationFriendListRequest, opts ...grpc.CallOption) (*DouyinRelationFriendListResponse, error)
}

type followerserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowerserviceClient(cc grpc.ClientConnInterface) FollowerserviceClient {
	return &followerserviceClient{cc}
}

func (c *followerserviceClient) RelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error) {
	out := new(DouyinRelationActionResponse)
	err := c.cc.Invoke(ctx, "/pb.followerservice/relationAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerserviceClient) UserFollowList(ctx context.Context, in *DouyinRelationFollowListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowListResponse, error) {
	out := new(DouyinRelationFollowListResponse)
	err := c.cc.Invoke(ctx, "/pb.followerservice/userFollowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerserviceClient) UserFollowerList(ctx context.Context, in *DouyinRelationFollowerListRequest, opts ...grpc.CallOption) (*DouyinRelationFollowerListResponse, error) {
	out := new(DouyinRelationFollowerListResponse)
	err := c.cc.Invoke(ctx, "/pb.followerservice/userFollowerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerserviceClient) UserFriendList(ctx context.Context, in *DouyinRelationFriendListRequest, opts ...grpc.CallOption) (*DouyinRelationFriendListResponse, error) {
	out := new(DouyinRelationFriendListResponse)
	err := c.cc.Invoke(ctx, "/pb.followerservice/userFriendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowerserviceServer is the server API for Followerservice service.
// All implementations must embed UnimplementedFollowerserviceServer
// for forward compatibility
type FollowerserviceServer interface {
	// 关系操作
	RelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error)
	// 用户关注列表
	UserFollowList(context.Context, *DouyinRelationFollowListRequest) (*DouyinRelationFollowListResponse, error)
	// 用户粉丝列表
	UserFollowerList(context.Context, *DouyinRelationFollowerListRequest) (*DouyinRelationFollowerListResponse, error)
	// 用户好友列表
	UserFriendList(context.Context, *DouyinRelationFriendListRequest) (*DouyinRelationFriendListResponse, error)
	mustEmbedUnimplementedFollowerserviceServer()
}

// UnimplementedFollowerserviceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowerserviceServer struct {
}

func (UnimplementedFollowerserviceServer) RelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedFollowerserviceServer) UserFollowList(context.Context, *DouyinRelationFollowListRequest) (*DouyinRelationFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFollowList not implemented")
}
func (UnimplementedFollowerserviceServer) UserFollowerList(context.Context, *DouyinRelationFollowerListRequest) (*DouyinRelationFollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFollowerList not implemented")
}
func (UnimplementedFollowerserviceServer) UserFriendList(context.Context, *DouyinRelationFriendListRequest) (*DouyinRelationFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFriendList not implemented")
}
func (UnimplementedFollowerserviceServer) mustEmbedUnimplementedFollowerserviceServer() {}

// UnsafeFollowerserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowerserviceServer will
// result in compilation errors.
type UnsafeFollowerserviceServer interface {
	mustEmbedUnimplementedFollowerserviceServer()
}

func RegisterFollowerserviceServer(s grpc.ServiceRegistrar, srv FollowerserviceServer) {
	s.RegisterService(&Followerservice_ServiceDesc, srv)
}

func _Followerservice_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerserviceServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.followerservice/relationAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerserviceServer).RelationAction(ctx, req.(*DouyinRelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followerservice_UserFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerserviceServer).UserFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.followerservice/userFollowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerserviceServer).UserFollowList(ctx, req.(*DouyinRelationFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followerservice_UserFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerserviceServer).UserFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.followerservice/userFollowerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerserviceServer).UserFollowerList(ctx, req.(*DouyinRelationFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followerservice_UserFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerserviceServer).UserFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.followerservice/userFriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerserviceServer).UserFriendList(ctx, req.(*DouyinRelationFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Followerservice_ServiceDesc is the grpc.ServiceDesc for Followerservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Followerservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.followerservice",
	HandlerType: (*FollowerserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "relationAction",
			Handler:    _Followerservice_RelationAction_Handler,
		},
		{
			MethodName: "userFollowList",
			Handler:    _Followerservice_UserFollowList_Handler,
		},
		{
			MethodName: "userFollowerList",
			Handler:    _Followerservice_UserFollowerList_Handler,
		},
		{
			MethodName: "userFriendList",
			Handler:    _Followerservice_UserFriendList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow.proto",
}
