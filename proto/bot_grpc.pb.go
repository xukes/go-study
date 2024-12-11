// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: proto/bot.proto

package proto

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
	HandleService_SendMessage_FullMethodName = "/proto.HandleService/SendMessage"
	HandleService_GetMessage_FullMethodName  = "/proto.HandleService/GetMessage"
)

// HandleServiceClient is the client API for HandleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandleServiceClient interface {
	// SendMessage 发送消息
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	GetMessage(ctx context.Context, in *GetMessageRps, opts ...grpc.CallOption) (*GetMessageResp, error)
}

type handleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHandleServiceClient(cc grpc.ClientConnInterface) HandleServiceClient {
	return &handleServiceClient{cc}
}

func (c *handleServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, HandleService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handleServiceClient) GetMessage(ctx context.Context, in *GetMessageRps, opts ...grpc.CallOption) (*GetMessageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageResp)
	err := c.cc.Invoke(ctx, HandleService_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandleServiceServer is the server API for HandleService service.
// All implementations must embed UnimplementedHandleServiceServer
// for forward compatibility.
type HandleServiceServer interface {
	// SendMessage 发送消息
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	GetMessage(context.Context, *GetMessageRps) (*GetMessageResp, error)
	mustEmbedUnimplementedHandleServiceServer()
}

// UnimplementedHandleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHandleServiceServer struct{}

func (UnimplementedHandleServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedHandleServiceServer) GetMessage(context.Context, *GetMessageRps) (*GetMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedHandleServiceServer) mustEmbedUnimplementedHandleServiceServer() {}
func (UnimplementedHandleServiceServer) testEmbeddedByValue()                       {}

// UnsafeHandleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandleServiceServer will
// result in compilation errors.
type UnsafeHandleServiceServer interface {
	mustEmbedUnimplementedHandleServiceServer()
}

func RegisterHandleServiceServer(s grpc.ServiceRegistrar, srv HandleServiceServer) {
	// If the following call pancis, it indicates UnimplementedHandleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HandleService_ServiceDesc, srv)
}

func _HandleService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HandleService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandleService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRps)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HandleService_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleServiceServer).GetMessage(ctx, req.(*GetMessageRps))
	}
	return interceptor(ctx, in, info, handler)
}

// HandleService_ServiceDesc is the grpc.ServiceDesc for HandleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HandleService",
	HandlerType: (*HandleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _HandleService_SendMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _HandleService_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bot.proto",
}
