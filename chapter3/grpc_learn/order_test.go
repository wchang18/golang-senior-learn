package main

import (
	"context"
	"fmt"
	"golang-senior-learn/chapter3/grpc_learn/my/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestGetOrder(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewOrderServiceClient(conn)
	res, err := client.GetOrderList(context.Background(), &pb.GetOrderListReq{Page: 1, PageSize: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

func TestCreateOrder(t *testing.T) {
	conn, err := grpc.NewClient("127.0.0.1:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	client := pb.NewOrderServiceClient(conn)
	stream, err := client.CreateOrder(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for i := 1; i < 10; i++ {
		err = stream.Send(&pb.Order{OrderId: int64(i), OrderName: "order_name" + fmt.Sprint(i), OrderPrice: 1.1 + float32(i)})
		if err != nil {
			t.Fatal(err)
		}
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
