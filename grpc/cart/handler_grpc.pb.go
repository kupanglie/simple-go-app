// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package handler_proto

import (
	cart "/cart"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserHandlerClient is the client API for UserHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserHandlerClient interface {
	Add(ctx context.Context, in *cart.AddRequest, opts ...grpc.CallOption) (*cart.AddResponse, error)
	Find(ctx context.Context, in *cart.FindRequest, opts ...grpc.CallOption) (*cart.FindResponse, error)
	Delete(ctx context.Context, in *cart.DeleteRequest, opts ...grpc.CallOption) (*cart.DeleteResponse, error)
}

type userHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserHandlerClient(cc grpc.ClientConnInterface) UserHandlerClient {
	return &userHandlerClient{cc}
}

func (c *userHandlerClient) Add(ctx context.Context, in *cart.AddRequest, opts ...grpc.CallOption) (*cart.AddResponse, error) {
	out := new(cart.AddResponse)
	err := c.cc.Invoke(ctx, "/cart.UserHandler/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) Find(ctx context.Context, in *cart.FindRequest, opts ...grpc.CallOption) (*cart.FindResponse, error) {
	out := new(cart.FindResponse)
	err := c.cc.Invoke(ctx, "/cart.UserHandler/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) Delete(ctx context.Context, in *cart.DeleteRequest, opts ...grpc.CallOption) (*cart.DeleteResponse, error) {
	out := new(cart.DeleteResponse)
	err := c.cc.Invoke(ctx, "/cart.UserHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHandlerServer is the server API for UserHandler service.
// All implementations should embed UnimplementedUserHandlerServer
// for forward compatibility
type UserHandlerServer interface {
	Add(context.Context, *cart.AddRequest) (*cart.AddResponse, error)
	Find(context.Context, *cart.FindRequest) (*cart.FindResponse, error)
	Delete(context.Context, *cart.DeleteRequest) (*cart.DeleteResponse, error)
}

// UnimplementedUserHandlerServer should be embedded to have forward compatible implementations.
type UnimplementedUserHandlerServer struct {
}

func (UnimplementedUserHandlerServer) Add(context.Context, *cart.AddRequest) (*cart.AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedUserHandlerServer) Find(context.Context, *cart.FindRequest) (*cart.FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedUserHandlerServer) Delete(context.Context, *cart.DeleteRequest) (*cart.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeUserHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserHandlerServer will
// result in compilation errors.
type UnsafeUserHandlerServer interface {
	mustEmbedUnimplementedUserHandlerServer()
}

func RegisterUserHandlerServer(s grpc.ServiceRegistrar, srv UserHandlerServer) {
	s.RegisterService(&UserHandler_ServiceDesc, srv)
}

func _UserHandler_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cart.AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.UserHandler/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Add(ctx, req.(*cart.AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cart.FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.UserHandler/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Find(ctx, req.(*cart.FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(cart.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.UserHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Delete(ctx, req.(*cart.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserHandler_ServiceDesc is the grpc.ServiceDesc for UserHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.UserHandler",
	HandlerType: (*UserHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UserHandler_Add_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _UserHandler_Find_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "handler.proto",
}
