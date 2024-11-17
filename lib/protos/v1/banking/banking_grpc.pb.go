// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: lib/protos/v1/banking/banking.proto

package banking

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
	BankingService_SubmitTransfer_FullMethodName = "/fastcampus.banking.v1.BankingService/SubmitTransfer"
	BankingService_SubmitPayment_FullMethodName  = "/fastcampus.banking.v1.BankingService/SubmitPayment"
	BankingService_SubmitDeposit_FullMethodName  = "/fastcampus.banking.v1.BankingService/SubmitDeposit"
)

// BankingServiceClient is the client API for BankingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankingServiceClient interface {
	// Transfer Antar Bank
	SubmitTransfer(ctx context.Context, in *BankTransferRequest, opts ...grpc.CallOption) (*BankTransferResponse, error)
	// Payment Antar Bank
	SubmitPayment(ctx context.Context, in *BankPaymentRequest, opts ...grpc.CallOption) (*BankPaymentResponse, error)
	// Payment Antar Bank
	SubmitDeposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
}

type bankingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankingServiceClient(cc grpc.ClientConnInterface) BankingServiceClient {
	return &bankingServiceClient{cc}
}

func (c *bankingServiceClient) SubmitTransfer(ctx context.Context, in *BankTransferRequest, opts ...grpc.CallOption) (*BankTransferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BankTransferResponse)
	err := c.cc.Invoke(ctx, BankingService_SubmitTransfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingServiceClient) SubmitPayment(ctx context.Context, in *BankPaymentRequest, opts ...grpc.CallOption) (*BankPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BankPaymentResponse)
	err := c.cc.Invoke(ctx, BankingService_SubmitPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingServiceClient) SubmitDeposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, BankingService_SubmitDeposit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankingServiceServer is the server API for BankingService service.
// All implementations must embed UnimplementedBankingServiceServer
// for forward compatibility
type BankingServiceServer interface {
	// Transfer Antar Bank
	SubmitTransfer(context.Context, *BankTransferRequest) (*BankTransferResponse, error)
	// Payment Antar Bank
	SubmitPayment(context.Context, *BankPaymentRequest) (*BankPaymentResponse, error)
	// Payment Antar Bank
	SubmitDeposit(context.Context, *DepositRequest) (*DepositResponse, error)
	mustEmbedUnimplementedBankingServiceServer()
}

// UnimplementedBankingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankingServiceServer struct {
}

func (UnimplementedBankingServiceServer) SubmitTransfer(context.Context, *BankTransferRequest) (*BankTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransfer not implemented")
}
func (UnimplementedBankingServiceServer) SubmitPayment(context.Context, *BankPaymentRequest) (*BankPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPayment not implemented")
}
func (UnimplementedBankingServiceServer) SubmitDeposit(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDeposit not implemented")
}
func (UnimplementedBankingServiceServer) mustEmbedUnimplementedBankingServiceServer() {}

// UnsafeBankingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankingServiceServer will
// result in compilation errors.
type UnsafeBankingServiceServer interface {
	mustEmbedUnimplementedBankingServiceServer()
}

func RegisterBankingServiceServer(s grpc.ServiceRegistrar, srv BankingServiceServer) {
	s.RegisterService(&BankingService_ServiceDesc, srv)
}

func _BankingService_SubmitTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).SubmitTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BankingService_SubmitTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).SubmitTransfer(ctx, req.(*BankTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankingService_SubmitPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).SubmitPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BankingService_SubmitPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).SubmitPayment(ctx, req.(*BankPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankingService_SubmitDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServiceServer).SubmitDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BankingService_SubmitDeposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServiceServer).SubmitDeposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankingService_ServiceDesc is the grpc.ServiceDesc for BankingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fastcampus.banking.v1.BankingService",
	HandlerType: (*BankingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransfer",
			Handler:    _BankingService_SubmitTransfer_Handler,
		},
		{
			MethodName: "SubmitPayment",
			Handler:    _BankingService_SubmitPayment_Handler,
		},
		{
			MethodName: "SubmitDeposit",
			Handler:    _BankingService_SubmitDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lib/protos/v1/banking/banking.proto",
}