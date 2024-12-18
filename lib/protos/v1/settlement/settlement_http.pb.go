// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.1
// source: lib/protos/v1/settlement/settlement.proto

package settlement

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

const OperationSettlementServiceGetSettlementReport = "/fastcampus.settlement.v1.SettlementService/GetSettlementReport"

type SettlementServiceHTTPServer interface {
	GetSettlementReport(context.Context, *SettlementReportRequest) (*SettlementReportResponse, error)
}

func RegisterSettlementServiceHTTPServer(s *http.Server, srv SettlementServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/settlement/report", _SettlementService_GetSettlementReport0_HTTP_Handler(srv))
}

func _SettlementService_GetSettlementReport0_HTTP_Handler(srv SettlementServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SettlementReportRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettlementServiceGetSettlementReport)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSettlementReport(ctx, req.(*SettlementReportRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SettlementReportResponse)
		return ctx.Result(200, reply)
	}
}

type SettlementServiceHTTPClient interface {
	GetSettlementReport(ctx context.Context, req *SettlementReportRequest, opts ...http.CallOption) (rsp *SettlementReportResponse, err error)
}

type SettlementServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewSettlementServiceHTTPClient(client *http.Client) SettlementServiceHTTPClient {
	return &SettlementServiceHTTPClientImpl{client}
}

func (c *SettlementServiceHTTPClientImpl) GetSettlementReport(ctx context.Context, in *SettlementReportRequest, opts ...http.CallOption) (*SettlementReportResponse, error) {
	var out SettlementReportResponse
	pattern := "/v1/settlement/report"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettlementServiceGetSettlementReport))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
