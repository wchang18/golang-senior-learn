package main

import (
	"golang-senior-learn/chapter3/grpc_learn/my/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, &OrderService{})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
