// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.29.1
// source: lib/protos/v1/money/money.proto

package money

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MoneyService_GetUserBalance_FullMethodName        = "/fastcampus.money.v1.MoneyService/GetUserBalance"
	MoneyService_GetUserBalancePrivate_FullMethodName = "/fastcampus.money.v1.MoneyService/GetUserBalancePrivate"
	MoneyService_UpdateUserBalance_FullMethodName     = "/fastcampus.money.v1.MoneyService/UpdateUserBalance"
)

// MoneyServiceClient is the client API for MoneyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoneyServiceClient interface {
	// balance return as string
	// /v1/money/balance?account_number=1234
	GetUserBalance(ctx context.Context, in *UserBalancePayload, opts ...grpc.CallOption) (*UserBalanceResponse, error)
	// balance will be represented as money
	GetUserBalancePrivate(ctx context.Context, in *UserBalancePayload, opts ...grpc.CallOption) (*UserBalancePrivateResponse, error)
	// update user balance
	UpdateUserBalance(ctx context.Context, in *UpdateBalancePayload, opts ...grpc.CallOption) (*UpdateBalanceResponse, error)
}

type moneyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMoneyServiceClient(cc grpc.ClientConnInterface) MoneyServiceClient {
	return &moneyServiceClient{cc}
}

func (c *moneyServiceClient) GetUserBalance(ctx context.Context, in *UserBalancePayload, opts ...grpc.CallOption) (*UserBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserBalanceResponse)
	err := c.cc.Invoke(ctx, MoneyService_GetUserBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moneyServiceClient) GetUserBalancePrivate(ctx context.Context, in *UserBalancePayload, opts ...grpc.CallOption) (*UserBalancePrivateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserBalancePrivateResponse)
	err := c.cc.Invoke(ctx, MoneyService_GetUserBalancePrivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moneyServiceClient) UpdateUserBalance(ctx context.Context, in *UpdateBalancePayload, opts ...grpc.CallOption) (*UpdateBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBalanceResponse)
	err := c.cc.Invoke(ctx, MoneyService_UpdateUserBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoneyServiceServer is the server API for MoneyService service.
// All implementations must embed UnimplementedMoneyServiceServer
// for forward compatibility
type MoneyServiceServer interface {
	// balance return as string
	// /v1/money/balance?account_number=1234
	GetUserBalance(context.Context, *UserBalancePayload) (*UserBalanceResponse, error)
	// balance will be represented as money
	GetUserBalancePrivate(context.Context, *UserBalancePayload) (*UserBalancePrivateResponse, error)
	// update user balance
	UpdateUserBalance(context.Context, *UpdateBalancePayload) (*UpdateBalanceResponse, error)
	mustEmbedUnimplementedMoneyServiceServer()
}

// UnimplementedMoneyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMoneyServiceServer struct {
}

func (UnimplementedMoneyServiceServer) GetUserBalance(context.Context, *UserBalancePayload) (*UserBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}
func (UnimplementedMoneyServiceServer) GetUserBalancePrivate(context.Context, *UserBalancePayload) (*UserBalancePrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalancePrivate not implemented")
}
func (UnimplementedMoneyServiceServer) UpdateUserBalance(context.Context, *UpdateBalancePayload) (*UpdateBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserBalance not implemented")
}
func (UnimplementedMoneyServiceServer) mustEmbedUnimplementedMoneyServiceServer() {}

// UnsafeMoneyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoneyServiceServer will
// result in compilation errors.
type UnsafeMoneyServiceServer interface {
	mustEmbedUnimplementedMoneyServiceServer()
}

func RegisterMoneyServiceServer(s grpc.ServiceRegistrar, srv MoneyServiceServer) {
	s.RegisterService(&MoneyService_ServiceDesc, srv)
}

func _MoneyService_GetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBalancePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).GetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoneyService_GetUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).GetUserBalance(ctx, req.(*UserBalancePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoneyService_GetUserBalancePrivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBalancePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).GetUserBalancePrivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoneyService_GetUserBalancePrivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).GetUserBalancePrivate(ctx, req.(*UserBalancePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoneyService_UpdateUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalancePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoneyServiceServer).UpdateUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoneyService_UpdateUserBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoneyServiceServer).UpdateUserBalance(ctx, req.(*UpdateBalancePayload))
	}
	return interceptor(ctx, in, info, handler)
}

// MoneyService_ServiceDesc is the grpc.ServiceDesc for MoneyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoneyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fastcampus.money.v1.MoneyService",
	HandlerType: (*MoneyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserBalance",
			Handler:    _MoneyService_GetUserBalance_Handler,
		},
		{
			MethodName: "GetUserBalancePrivate",
			Handler:    _MoneyService_GetUserBalancePrivate_Handler,
		},
		{
			MethodName: "UpdateUserBalance",
			Handler:    _MoneyService_UpdateUserBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lib/protos/v1/money/money.proto",
}
