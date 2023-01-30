// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: guestbooks.proto

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

// GuestbookServiceClient is the client API for GuestbookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuestbookServiceClient interface {
	// 新增
	// rpc CreateGuestbook (CreateGuestbookRequest) returns (CreateGuestbookResponse);
	// 查詢全部
	GetGuestbook(ctx context.Context, in *GetGuestbookRequest, opts ...grpc.CallOption) (*GetGuestbookResponse, error)
}

type guestbookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuestbookServiceClient(cc grpc.ClientConnInterface) GuestbookServiceClient {
	return &guestbookServiceClient{cc}
}

func (c *guestbookServiceClient) GetGuestbook(ctx context.Context, in *GetGuestbookRequest, opts ...grpc.CallOption) (*GetGuestbookResponse, error) {
	out := new(GetGuestbookResponse)
	err := c.cc.Invoke(ctx, "/proto.GuestbookService/GetGuestbook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuestbookServiceServer is the server API for GuestbookService service.
// All implementations must embed UnimplementedGuestbookServiceServer
// for forward compatibility
type GuestbookServiceServer interface {
	// 新增
	// rpc CreateGuestbook (CreateGuestbookRequest) returns (CreateGuestbookResponse);
	// 查詢全部
	GetGuestbook(context.Context, *GetGuestbookRequest) (*GetGuestbookResponse, error)
	mustEmbedUnimplementedGuestbookServiceServer()
}

// UnimplementedGuestbookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGuestbookServiceServer struct {
}

func (UnimplementedGuestbookServiceServer) GetGuestbook(context.Context, *GetGuestbookRequest) (*GetGuestbookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuestbook not implemented")
}
func (UnimplementedGuestbookServiceServer) mustEmbedUnimplementedGuestbookServiceServer() {}

// UnsafeGuestbookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuestbookServiceServer will
// result in compilation errors.
type UnsafeGuestbookServiceServer interface {
	mustEmbedUnimplementedGuestbookServiceServer()
}

func RegisterGuestbookServiceServer(s grpc.ServiceRegistrar, srv GuestbookServiceServer) {
	s.RegisterService(&GuestbookService_ServiceDesc, srv)
}

func _GuestbookService_GetGuestbook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuestbookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestbookServiceServer).GetGuestbook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GuestbookService/GetGuestbook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestbookServiceServer).GetGuestbook(ctx, req.(*GetGuestbookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GuestbookService_ServiceDesc is the grpc.ServiceDesc for GuestbookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuestbookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GuestbookService",
	HandlerType: (*GuestbookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGuestbook",
			Handler:    _GuestbookService_GetGuestbook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guestbooks.proto",
}
