// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v3.20.1
// source: lottery/v1/lottery.proto

package v1

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

const OperationLotteryLotteryV1 = "/api.lottery.v1.Lottery/LotteryV1"

type LotteryHTTPServer interface {
	LotteryV1(context.Context, *LotteryReq) (*LotteryRsp, error)
}

func RegisterLotteryHTTPServer(s *http.Server, srv LotteryHTTPServer) {
	r := s.Route("/")
	r.POST("/lottery", _Lottery_LotteryV10_HTTP_Handler(srv))
}

func _Lottery_LotteryV10_HTTP_Handler(srv LotteryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LotteryReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLotteryLotteryV1)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LotteryV1(ctx, req.(*LotteryReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LotteryRsp)
		return ctx.Result(200, reply)
	}
}

type LotteryHTTPClient interface {
	LotteryV1(ctx context.Context, req *LotteryReq, opts ...http.CallOption) (rsp *LotteryRsp, err error)
}

type LotteryHTTPClientImpl struct {
	cc *http.Client
}

func NewLotteryHTTPClient(client *http.Client) LotteryHTTPClient {
	return &LotteryHTTPClientImpl{client}
}

func (c *LotteryHTTPClientImpl) LotteryV1(ctx context.Context, in *LotteryReq, opts ...http.CallOption) (*LotteryRsp, error) {
	var out LotteryRsp
	pattern := "/lottery"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLotteryLotteryV1))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}