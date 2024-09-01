package main

import (
	"context"
	"fmt"
	"golang-senior-learn/chapter3/grpc_learn/my/pb"
	"io"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func (o *OrderService) GetOrderList(ctx context.Context, req *pb.GetOrderListReq) (res *pb.OrderListRes, err error) {
	fmt.Println(req.Page, req.PageSize)
	order := &pb.Order{
		OrderId:     1,
		OrderName:   "order1",
		OrderPrice:  1.1,
		OrderDesc:   nil,
		OrderTag:    map[string]string{"tag1": "tag2"},
		OrderStatus: pb.OrderStatus_Success,
	}
	return &pb.OrderListRes{
		Order: []*pb.Order{order},
		Total: 1,
	}, nil
}

func (o *OrderService) CreateOrder(stream pb.OrderService_CreateOrderServer) error {
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Res{Msg: "success"})
		}
		fmt.Printf("order: %+v\n", order)
	}
}
