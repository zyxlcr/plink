// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: chat.proto3

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProdServiceClient interface {
	//双向流
	GetProdSocket(ctx context.Context, opts ...grpc.CallOption) (ProdService_GetProdSocketClient, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdSocket(ctx context.Context, opts ...grpc.CallOption) (ProdService_GetProdSocketClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProdService_ServiceDesc.Streams[0], "/chat.ProdService/GetProdSocket", opts...)
	if err != nil {
		return nil, err
	}
	x := &prodServiceGetProdSocketClient{stream}
	return x, nil
}

type ProdService_GetProdSocketClient interface {
	Send(*ProductRequest) error
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type prodServiceGetProdSocketClient struct {
	grpc.ClientStream
}

func (x *prodServiceGetProdSocketClient) Send(m *ProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *prodServiceGetProdSocketClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProdServiceServer is the server API for ProdService service.
// All implementations must embed UnimplementedProdServiceServer
// for forward compatibility
type ProdServiceServer interface {
	//双向流
	GetProdSocket(ProdService_GetProdSocketServer) error
	mustEmbedUnimplementedProdServiceServer()
}

// UnimplementedProdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (UnimplementedProdServiceServer) GetProdSocket(ProdService_GetProdSocketServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProdSocket not implemented")
}
func (UnimplementedProdServiceServer) mustEmbedUnimplementedProdServiceServer() {}

// UnsafeProdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProdServiceServer will
// result in compilation errors.
type UnsafeProdServiceServer interface {
	mustEmbedUnimplementedProdServiceServer()
}

func RegisterProdServiceServer(s grpc.ServiceRegistrar, srv ProdServiceServer) {
	s.RegisterService(&ProdService_ServiceDesc, srv)
}

func _ProdService_GetProdSocket_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProdServiceServer).GetProdSocket(&prodServiceGetProdSocketServer{stream})
}

type ProdService_GetProdSocketServer interface {
	Send(*ProductResponse) error
	Recv() (*ProductRequest, error)
	grpc.ServerStream
}

type prodServiceGetProdSocketServer struct {
	grpc.ServerStream
}

func (x *prodServiceGetProdSocketServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *prodServiceGetProdSocketServer) Recv() (*ProductRequest, error) {
	m := new(ProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProdService_ServiceDesc is the grpc.ServiceDesc for ProdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProdSocket",
			Handler:       _ProdService_GetProdSocket_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto3",
}
