// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package server

import (
	"context"

	"golang-senior-learn/chapter3/mall/pay/internal/logic"
	"golang-senior-learn/chapter3/mall/pay/internal/svc"
	"golang-senior-learn/chapter3/mall/pay/pay/pb"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPayServer
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) GetPayCode(ctx context.Context, in *pb.PayCodeReq) (*pb.PayCodeRes, error) {
	l := logic.NewGetPayCodeLogic(ctx, s.svcCtx)
	return l.GetPayCode(in)
}
