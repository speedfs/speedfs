// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trackerpb

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

// TrackerClient is the client API for Tracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	ListGroupStorages(ctx context.Context, in *ListGroupStoragesRequest, opts ...grpc.CallOption) (*ListGroupStoragesResponse, error)
	GetStorage(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*GetStorageResponse, error)
	DeleteStorage(ctx context.Context, in *DeleteStorageRequest, opts ...grpc.CallOption) (*DeleteStorageResponse, error)
}

type trackerClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackerClient(cc grpc.ClientConnInterface) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/trackerpb.Tracker/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, "/trackerpb.Tracker/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/trackerpb.Tracker/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	out := new(DeleteGroupResponse)
	err := c.cc.Invoke(ctx, "/trackerpb.Tracker/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) ListGroupStorages(ctx context.Context, in *ListGroupStoragesRequest, opts ...grpc.CallOption) (*ListGroupStoragesResponse, error) {
	out := new(ListGroupStoragesResponse)
	err := c.cc.Invoke(ctx, "/trackerpb.Tracker/ListGroupStorages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) GetStorage(ctx context.Context, in *GetStorageRequest, opts ...grpc.CallOption) (*GetStorageResponse, error) {
	out := new(GetStorageResponse)
	err := c.cc.Invoke(ctx, "/trackerpb.Tracker/GetStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) DeleteStorage(ctx context.Context, in *DeleteStorageRequest, opts ...grpc.CallOption) (*DeleteStorageResponse, error) {
	out := new(DeleteStorageResponse)
	err := c.cc.Invoke(ctx, "/trackerpb.Tracker/DeleteStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackerServer is the server API for Tracker service.
// All implementations must embed UnimplementedTrackerServer
// for forward compatibility
type TrackerServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)
	ListGroupStorages(context.Context, *ListGroupStoragesRequest) (*ListGroupStoragesResponse, error)
	GetStorage(context.Context, *GetStorageRequest) (*GetStorageResponse, error)
	DeleteStorage(context.Context, *DeleteStorageRequest) (*DeleteStorageResponse, error)
	mustEmbedUnimplementedTrackerServer()
}

// UnimplementedTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedTrackerServer struct {
}

func (UnimplementedTrackerServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTrackerServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedTrackerServer) ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedTrackerServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedTrackerServer) ListGroupStorages(context.Context, *ListGroupStoragesRequest) (*ListGroupStoragesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroupStorages not implemented")
}
func (UnimplementedTrackerServer) GetStorage(context.Context, *GetStorageRequest) (*GetStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorage not implemented")
}
func (UnimplementedTrackerServer) DeleteStorage(context.Context, *DeleteStorageRequest) (*DeleteStorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorage not implemented")
}
func (UnimplementedTrackerServer) mustEmbedUnimplementedTrackerServer() {}

// UnsafeTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackerServer will
// result in compilation errors.
type UnsafeTrackerServer interface {
	mustEmbedUnimplementedTrackerServer()
}

func RegisterTrackerServer(s grpc.ServiceRegistrar, srv TrackerServer) {
	s.RegisterService(&Tracker_ServiceDesc, srv)
}

func _Tracker_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trackerpb.Tracker/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trackerpb.Tracker/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trackerpb.Tracker/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trackerpb.Tracker/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_ListGroupStorages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupStoragesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).ListGroupStorages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trackerpb.Tracker/ListGroupStorages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).ListGroupStorages(ctx, req.(*ListGroupStoragesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_GetStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).GetStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trackerpb.Tracker/GetStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).GetStorage(ctx, req.(*GetStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_DeleteStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).DeleteStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trackerpb.Tracker/DeleteStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).DeleteStorage(ctx, req.(*DeleteStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tracker_ServiceDesc is the grpc.ServiceDesc for Tracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trackerpb.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Tracker_Ping_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _Tracker_GetGroup_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _Tracker_ListGroups_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _Tracker_DeleteGroup_Handler,
		},
		{
			MethodName: "ListGroupStorages",
			Handler:    _Tracker_ListGroupStorages_Handler,
		},
		{
			MethodName: "GetStorage",
			Handler:    _Tracker_GetStorage_Handler,
		},
		{
			MethodName: "DeleteStorage",
			Handler:    _Tracker_DeleteStorage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/trackerpb/tracker.proto",
}