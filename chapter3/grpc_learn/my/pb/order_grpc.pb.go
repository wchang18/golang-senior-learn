// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: order.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OrderService_GetOrderList_FullMethodName = "/gary.pb.OrderService/GetOrderList"
	OrderService_CreateOrder_FullMethodName  = "/gary.pb.OrderService/CreateOrder"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetOrderList(ctx context.Context, in *GetOrderListReq, opts ...grpc.CallOption) (*OrderListRes, error)
	CreateOrder(ctx context.Context, opts ...grpc.CallOption) (OrderService_CreateOrderClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrderList(ctx context.Context, in *GetOrderListReq, opts ...grpc.CallOption) (*OrderListRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderListRes)
	err := c.cc.Invoke(ctx, OrderService_GetOrderList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, opts ...grpc.CallOption) (OrderService_CreateOrderClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_CreateOrder_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceCreateOrderClient{ClientStream: stream}
	return x, nil
}

type OrderService_CreateOrderClient interface {
	Send(*Order) error
	CloseAndRecv() (*Res, error)
	grpc.ClientStream
}

type orderServiceCreateOrderClient struct {
	grpc.ClientStream
}

func (x *orderServiceCreateOrderClient) Send(m *Order) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceCreateOrderClient) CloseAndRecv() (*Res, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	GetOrderList(context.Context, *GetOrderListReq) (*OrderListRes, error)
	CreateOrder(OrderService_CreateOrderServer) error
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) GetOrderList(context.Context, *GetOrderListReq) (*OrderListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrder(OrderService_CreateOrderServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderList(ctx, req.(*GetOrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).CreateOrder(&orderServiceCreateOrderServer{ServerStream: stream})
}

type OrderService_CreateOrderServer interface {
	SendAndClose(*Res) error
	Recv() (*Order, error)
	grpc.ServerStream
}

type orderServiceCreateOrderServer struct {
	grpc.ServerStream
}

func (x *orderServiceCreateOrderServer) SendAndClose(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceCreateOrderServer) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gary.pb.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderList",
			Handler:    _OrderService_GetOrderList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateOrder",
			Handler:       _OrderService_CreateOrder_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "order.proto",
}
