package logic

import (
	"context"
	"fmt"
	"golang-senior-learn/chapter3/mall/order/model"
	"golang-senior-learn/chapter3/mall/pay/pay/pb"
	"k8s.io/apimachinery/pkg/util/rand"
	"strconv"
	"strings"
	"time"

	"golang-senior-learn/chapter3/mall/order/internal/svc"
	"golang-senior-learn/chapter3/mall/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.Resp, err error) {

	if req.Amount <= 0 {
		err = fmt.Errorf("invalid amount")
		return
	}

	orderId := GetOrderId()

	payRes, err := l.svcCtx.PayRpc.GetPayCode(l.ctx, &pb.PayCodeReq{
		UserId:   1,
		Amount:   req.Amount,
		Platform: pb.PayPlatform(pb.PayPlatform_value[strings.ToUpper(req.PayMethod)]),
		OrderId:  orderId,
	})

	if err != nil {
		return
	}

	fmt.Printf("payRes: %+v", payRes)

	var order model.Order

	order.OrderID = orderId
	order.PayStatus = payRes.Data.Status.String()
	order.UserID = 1
	order.Subject = "用户续费"
	order.MonthCount = req.Month

	err = l.svcCtx.DB.Create(&order).Error
	if err != nil {
		return
	}

	resp = &types.Resp{
		Code: 0,
		Msg:  "success",
		Data: types.Order{
			OrderId:     order.OrderID,
			PayUrl:      payRes.Data.PayUrl,
			PayMethod:   req.PayMethod,
			OrderStatus: order.PayStatus,
			ExpireAt:    payRes.Data.ExpireAt,
		},
	}
	return
}

func GetOrderId() int64 {
	now := time.Now()
	timestamp := now.UnixMilli()
	timestamp = timestamp << 3
	rand.Seed(now.UnixNano())
	randNum := rand.Intn(8999)
	randNum += 1000
	str := fmt.Sprintf("%d%d", timestamp, randNum)
	number, _ := strconv.ParseInt(str, 10, 64)
	return number
}
