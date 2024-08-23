// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0--rc2
// source: lease_rpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CacheNode_OutdateData_FullMethodName = "/CacheNode/OutdateData"
)

// CacheNodeClient is the client API for CacheNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheNodeClient interface {
	// 请求数据，如果缓存中没有数据或租约过期，则从中心节点请求数据
	OutdateData(ctx context.Context, in *OutdateDataRequest, opts ...grpc.CallOption) (*OutdateDataResponse, error)
}

type cacheNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheNodeClient(cc grpc.ClientConnInterface) CacheNodeClient {
	return &cacheNodeClient{cc}
}

func (c *cacheNodeClient) OutdateData(ctx context.Context, in *OutdateDataRequest, opts ...grpc.CallOption) (*OutdateDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OutdateDataResponse)
	err := c.cc.Invoke(ctx, CacheNode_OutdateData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheNodeServer is the server API for CacheNode service.
// All implementations must embed UnimplementedCacheNodeServer
// for forward compatibility.
type CacheNodeServer interface {
	// 请求数据，如果缓存中没有数据或租约过期，则从中心节点请求数据
	OutdateData(context.Context, *OutdateDataRequest) (*OutdateDataResponse, error)
	mustEmbedUnimplementedCacheNodeServer()
}

// UnimplementedCacheNodeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCacheNodeServer struct{}

func (UnimplementedCacheNodeServer) OutdateData(context.Context, *OutdateDataRequest) (*OutdateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OutdateData not implemented")
}
func (UnimplementedCacheNodeServer) mustEmbedUnimplementedCacheNodeServer() {}
func (UnimplementedCacheNodeServer) testEmbeddedByValue()                   {}

// UnsafeCacheNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheNodeServer will
// result in compilation errors.
type UnsafeCacheNodeServer interface {
	mustEmbedUnimplementedCacheNodeServer()
}

func RegisterCacheNodeServer(s grpc.ServiceRegistrar, srv CacheNodeServer) {
	// If the following call pancis, it indicates UnimplementedCacheNodeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CacheNode_ServiceDesc, srv)
}

func _CacheNode_OutdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutdateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheNodeServer).OutdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheNode_OutdateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheNodeServer).OutdateData(ctx, req.(*OutdateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheNode_ServiceDesc is the grpc.ServiceDesc for CacheNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CacheNode",
	HandlerType: (*CacheNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OutdateData",
			Handler:    _CacheNode_OutdateData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lease_rpc.proto",
}

const (
	CenterNode_RequestData_FullMethodName = "/CenterNode/RequestData"
	CenterNode_WriteData_FullMethodName   = "/CenterNode/WriteData"
)

// CenterNodeClient is the client API for CenterNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CenterNodeClient interface {
	// 请求数据，返回数据和租约时间
	RequestData(ctx context.Context, in *RequestDataRequest, opts ...grpc.CallOption) (*RequestDataResponse, error)
	WriteData(ctx context.Context, in *WriteDataRequest, opts ...grpc.CallOption) (*WriteDataResponse, error)
}

type centerNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewCenterNodeClient(cc grpc.ClientConnInterface) CenterNodeClient {
	return &centerNodeClient{cc}
}

func (c *centerNodeClient) RequestData(ctx context.Context, in *RequestDataRequest, opts ...grpc.CallOption) (*RequestDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestDataResponse)
	err := c.cc.Invoke(ctx, CenterNode_RequestData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerNodeClient) WriteData(ctx context.Context, in *WriteDataRequest, opts ...grpc.CallOption) (*WriteDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WriteDataResponse)
	err := c.cc.Invoke(ctx, CenterNode_WriteData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CenterNodeServer is the server API for CenterNode service.
// All implementations must embed UnimplementedCenterNodeServer
// for forward compatibility.
type CenterNodeServer interface {
	// 请求数据，返回数据和租约时间
	RequestData(context.Context, *RequestDataRequest) (*RequestDataResponse, error)
	WriteData(context.Context, *WriteDataRequest) (*WriteDataResponse, error)
	mustEmbedUnimplementedCenterNodeServer()
}

// UnimplementedCenterNodeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCenterNodeServer struct{}

func (UnimplementedCenterNodeServer) RequestData(context.Context, *RequestDataRequest) (*RequestDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestData not implemented")
}
func (UnimplementedCenterNodeServer) WriteData(context.Context, *WriteDataRequest) (*WriteDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteData not implemented")
}
func (UnimplementedCenterNodeServer) mustEmbedUnimplementedCenterNodeServer() {}
func (UnimplementedCenterNodeServer) testEmbeddedByValue()                    {}

// UnsafeCenterNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CenterNodeServer will
// result in compilation errors.
type UnsafeCenterNodeServer interface {
	mustEmbedUnimplementedCenterNodeServer()
}

func RegisterCenterNodeServer(s grpc.ServiceRegistrar, srv CenterNodeServer) {
	// If the following call pancis, it indicates UnimplementedCenterNodeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CenterNode_ServiceDesc, srv)
}

func _CenterNode_RequestData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterNodeServer).RequestData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CenterNode_RequestData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterNodeServer).RequestData(ctx, req.(*RequestDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterNode_WriteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterNodeServer).WriteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CenterNode_WriteData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterNodeServer).WriteData(ctx, req.(*WriteDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CenterNode_ServiceDesc is the grpc.ServiceDesc for CenterNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CenterNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CenterNode",
	HandlerType: (*CenterNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestData",
			Handler:    _CenterNode_RequestData_Handler,
		},
		{
			MethodName: "WriteData",
			Handler:    _CenterNode_WriteData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lease_rpc.proto",
}
