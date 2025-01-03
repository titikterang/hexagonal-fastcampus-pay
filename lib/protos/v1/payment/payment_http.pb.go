// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.29.1
// source: lib/protos/v1/payment/payment.proto

package payment

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

const OperationPaymentServiceGetPaymentPrecheckInfo = "/fastcampus.payment.v1.PaymentService/GetPaymentPrecheckInfo"
const OperationPaymentServiceGetPaymentStatus = "/fastcampus.payment.v1.PaymentService/GetPaymentStatus"
const OperationPaymentServiceSubmitPayment = "/fastcampus.payment.v1.PaymentService/SubmitPayment"

type PaymentServiceHTTPServer interface {
	// GetPaymentPrecheckInfo /v1/payment/precheck?account_no=1234
	GetPaymentPrecheckInfo(context.Context, *PaymentPrecheckPayload) (*PaymentPrecheckResponse, error)
	// GetPaymentStatus balance return as string
	// /v1/payment/status?invoice_id=1234
	GetPaymentStatus(context.Context, *PaymentStatusPayload) (*PaymentStatusResponse, error)
	// SubmitPayment update user balance
	SubmitPayment(context.Context, *SubmitPaymentPayload) (*SubmitPaymentResponse, error)
}

func RegisterPaymentServiceHTTPServer(s *http.Server, srv PaymentServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/payment/status", _PaymentService_GetPaymentStatus0_HTTP_Handler(srv))
	r.GET("/v1/payment/precheck", _PaymentService_GetPaymentPrecheckInfo0_HTTP_Handler(srv))
	r.POST("/v1/payment/submit", _PaymentService_SubmitPayment0_HTTP_Handler(srv))
}

func _PaymentService_GetPaymentStatus0_HTTP_Handler(srv PaymentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PaymentStatusPayload
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentServiceGetPaymentStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPaymentStatus(ctx, req.(*PaymentStatusPayload))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PaymentStatusResponse)
		return ctx.Result(200, reply)
	}
}

func _PaymentService_GetPaymentPrecheckInfo0_HTTP_Handler(srv PaymentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PaymentPrecheckPayload
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentServiceGetPaymentPrecheckInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPaymentPrecheckInfo(ctx, req.(*PaymentPrecheckPayload))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PaymentPrecheckResponse)
		return ctx.Result(200, reply)
	}
}

func _PaymentService_SubmitPayment0_HTTP_Handler(srv PaymentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SubmitPaymentPayload
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPaymentServiceSubmitPayment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SubmitPayment(ctx, req.(*SubmitPaymentPayload))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SubmitPaymentResponse)
		return ctx.Result(200, reply)
	}
}

type PaymentServiceHTTPClient interface {
	GetPaymentPrecheckInfo(ctx context.Context, req *PaymentPrecheckPayload, opts ...http.CallOption) (rsp *PaymentPrecheckResponse, err error)
	GetPaymentStatus(ctx context.Context, req *PaymentStatusPayload, opts ...http.CallOption) (rsp *PaymentStatusResponse, err error)
	SubmitPayment(ctx context.Context, req *SubmitPaymentPayload, opts ...http.CallOption) (rsp *SubmitPaymentResponse, err error)
}

type PaymentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPaymentServiceHTTPClient(client *http.Client) PaymentServiceHTTPClient {
	return &PaymentServiceHTTPClientImpl{client}
}

func (c *PaymentServiceHTTPClientImpl) GetPaymentPrecheckInfo(ctx context.Context, in *PaymentPrecheckPayload, opts ...http.CallOption) (*PaymentPrecheckResponse, error) {
	var out PaymentPrecheckResponse
	pattern := "/v1/payment/precheck"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPaymentServiceGetPaymentPrecheckInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PaymentServiceHTTPClientImpl) GetPaymentStatus(ctx context.Context, in *PaymentStatusPayload, opts ...http.CallOption) (*PaymentStatusResponse, error) {
	var out PaymentStatusResponse
	pattern := "/v1/payment/status"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPaymentServiceGetPaymentStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PaymentServiceHTTPClientImpl) SubmitPayment(ctx context.Context, in *SubmitPaymentPayload, opts ...http.CallOption) (*SubmitPaymentResponse, error) {
	var out SubmitPaymentResponse
	pattern := "/v1/payment/submit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPaymentServiceSubmitPayment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
