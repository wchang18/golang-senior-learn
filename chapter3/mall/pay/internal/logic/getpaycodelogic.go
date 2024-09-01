package logic

import (
	"context"
	"fmt"
	"github.com/thoas/go-funk"
	"golang-senior-learn/chapter3/mall/pay/internal/platform_pay"
	"golang-senior-learn/chapter3/mall/pay/model"

	"golang-senior-learn/chapter3/mall/pay/internal/svc"
	"golang-senior-learn/chapter3/mall/pay/pay/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPayCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayCodeLogic {
	return &GetPayCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPayCodeLogic) GetPayCode(in *pb.PayCodeReq) (*pb.PayCodeRes, error) {

	tradeNo := GetTradeNo()
	platformServer := platform_pay.NewPlatformServer()
	payRes, err := platformServer.GetPayCode(&platform_pay.GetPlatformCodeReq{
		Amount:     in.Amount,
		PlatformId: int32(in.Platform),
		Subject:    in.Subject,
		TradeNo:    tradeNo,
	})
	if err != nil {
		return nil, err
	}

	pay := model.OrderPay{
		OrderID:    in.OrderId,
		PayQrCode:  payRes.PayUrl,
		TradeNo:    tradeNo,
		PlatformID: int32(in.Platform),
		ExpireAt:   payRes.ExpireAt,
		PayStatus:  1,
	}

	err = l.svcCtx.DB.Create(&pay).Error
	if err != nil {
		return nil, err
	}

	fmt.Printf("in:%+v\n", in)
	return &pb.PayCodeRes{
		Data: &pb.PayCode{
			OrderId:  in.OrderId,
			PayUrl:   pay.PayQrCode,
			Status:   pb.PayStatus_PENDING,
			ExpireAt: payRes.ExpireAt,
		},
		Resp: &pb.Resp{
			Code: 1,
			Msg:  "success",
		},
	}, nil
}

func GetTradeNo() string {
	return funk.RandomString(12)
}
