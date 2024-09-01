package main

import (
	"context"
	"golang-senior-learn/chapter3/grpc_learn/my/pb"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (*HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (res *pb.HelloReply, err error) {
	res = &pb.HelloReply{
		Reply: "hello " + req.Name,
	}
	return
}
