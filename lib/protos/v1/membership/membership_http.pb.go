// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.29.1
// source: lib/protos/v1/membership/membership.proto

package membership

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

const OperationMembershipServiceGetUserInfo = "/fastcampus.membership.v1.MembershipService/GetUserInfo"
const OperationMembershipServiceSubmitLogin = "/fastcampus.membership.v1.MembershipService/SubmitLogin"
const OperationMembershipServiceSubmitLogout = "/fastcampus.membership.v1.MembershipService/SubmitLogout"
const OperationMembershipServiceSubmitRegistration = "/fastcampus.membership.v1.MembershipService/SubmitRegistration"

type MembershipServiceHTTPServer interface {
	GetUserInfo(context.Context, *UserInfoPayload) (*UserInfoResponse, error)
	// SubmitLogin login
	SubmitLogin(context.Context, *LoginRequest) (*LoginResponse, error)
	// SubmitLogout login
	SubmitLogout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	// SubmitRegistration Register
	SubmitRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
}

func RegisterMembershipServiceHTTPServer(s *http.Server, srv MembershipServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/membership/info", _MembershipService_GetUserInfo0_HTTP_Handler(srv))
	r.POST("/v1/membership/register", _MembershipService_SubmitRegistration0_HTTP_Handler(srv))
	r.POST("/v1/membership/logout", _MembershipService_SubmitLogout0_HTTP_Handler(srv))
	r.POST("/v1/membership/auth", _MembershipService_SubmitLogin0_HTTP_Handler(srv))
}

func _MembershipService_GetUserInfo0_HTTP_Handler(srv MembershipServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoPayload
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMembershipServiceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*UserInfoPayload))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _MembershipService_SubmitRegistration0_HTTP_Handler(srv MembershipServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegistrationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMembershipServiceSubmitRegistration)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SubmitRegistration(ctx, req.(*RegistrationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegistrationResponse)
		return ctx.Result(200, reply)
	}
}

func _MembershipService_SubmitLogout0_HTTP_Handler(srv MembershipServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMembershipServiceSubmitLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SubmitLogout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutResponse)
		return ctx.Result(200, reply)
	}
}

func _MembershipService_SubmitLogin0_HTTP_Handler(srv MembershipServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMembershipServiceSubmitLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SubmitLogin(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

type MembershipServiceHTTPClient interface {
	GetUserInfo(ctx context.Context, req *UserInfoPayload, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	SubmitLogin(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	SubmitLogout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutResponse, err error)
	SubmitRegistration(ctx context.Context, req *RegistrationRequest, opts ...http.CallOption) (rsp *RegistrationResponse, err error)
}

type MembershipServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMembershipServiceHTTPClient(client *http.Client) MembershipServiceHTTPClient {
	return &MembershipServiceHTTPClientImpl{client}
}

func (c *MembershipServiceHTTPClientImpl) GetUserInfo(ctx context.Context, in *UserInfoPayload, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/v1/membership/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMembershipServiceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MembershipServiceHTTPClientImpl) SubmitLogin(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/v1/membership/auth"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMembershipServiceSubmitLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MembershipServiceHTTPClientImpl) SubmitLogout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutResponse, error) {
	var out LogoutResponse
	pattern := "/v1/membership/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMembershipServiceSubmitLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MembershipServiceHTTPClientImpl) SubmitRegistration(ctx context.Context, in *RegistrationRequest, opts ...http.CallOption) (*RegistrationResponse, error) {
	var out RegistrationResponse
	pattern := "/v1/membership/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMembershipServiceSubmitRegistration))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
