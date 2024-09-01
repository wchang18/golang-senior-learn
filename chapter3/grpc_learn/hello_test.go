package main

import (
	"context"
	"golang-senior-learn/chapter3/grpc_learn/my/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestHello(t *testing.T) {

	conn, err := grpc.NewClient("127.0.0.1:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		t.Fatal(err)
	}
	client := pb.NewHelloServiceClient(conn)
	res, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "jack"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res.Reply)
}
