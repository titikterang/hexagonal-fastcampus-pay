// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.29.1
// source: lib/protos/v1/money/money.proto

package money

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMoneyServiceGetUserBalance = "/fastcampus.money.v1.MoneyService/GetUserBalance"
const OperationMoneyServiceUpdateUserBalance = "/fastcampus.money.v1.MoneyService/UpdateUserBalance"

type MoneyServiceHTTPServer interface {
	// GetUserBalance balance return as string
	// /v1/money/balance?account_number=1234
	GetUserBalance(context.Context, *UserBalancePayload) (*UserBalanceResponse, error)
	// UpdateUserBalance update user balance
	UpdateUserBalance(context.Context, *UpdateBalancePayload) (*UpdateBalanceResponse, error)
}

func RegisterMoneyServiceHTTPServer(s *http.Server, srv MoneyServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/money/balance", _MoneyService_GetUserBalance0_HTTP_Handler(srv))
	r.POST("/v1/money/private/balance", _MoneyService_UpdateUserBalance0_HTTP_Handler(srv))
}

func _MoneyService_GetUserBalance0_HTTP_Handler(srv MoneyServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserBalancePayload
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMoneyServiceGetUserBalance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserBalance(ctx, req.(*UserBalancePayload))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserBalanceResponse)
		return ctx.Result(200, reply)
	}
}

func _MoneyService_UpdateUserBalance0_HTTP_Handler(srv MoneyServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBalancePayload
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMoneyServiceUpdateUserBalance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserBalance(ctx, req.(*UpdateBalancePayload))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateBalanceResponse)
		return ctx.Result(200, reply)
	}
}

type MoneyServiceHTTPClient interface {
	GetUserBalance(ctx context.Context, req *UserBalancePayload, opts ...http.CallOption) (rsp *UserBalanceResponse, err error)
	UpdateUserBalance(ctx context.Context, req *UpdateBalancePayload, opts ...http.CallOption) (rsp *UpdateBalanceResponse, err error)
}

type MoneyServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMoneyServiceHTTPClient(client *http.Client) MoneyServiceHTTPClient {
	return &MoneyServiceHTTPClientImpl{client}
}

func (c *MoneyServiceHTTPClientImpl) GetUserBalance(ctx context.Context, in *UserBalancePayload, opts ...http.CallOption) (*UserBalanceResponse, error) {
	var out UserBalanceResponse
	pattern := "/v1/money/balance"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMoneyServiceGetUserBalance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MoneyServiceHTTPClientImpl) UpdateUserBalance(ctx context.Context, in *UpdateBalancePayload, opts ...http.CallOption) (*UpdateBalanceResponse, error) {
	var out UpdateBalanceResponse
	pattern := "/v1/money/private/balance"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMoneyServiceUpdateUserBalance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
