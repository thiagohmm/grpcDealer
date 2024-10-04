// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: listProducts.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ListProductsService_ListProducts_FullMethodName = "/pb.ListProductsService/ListProducts"
)

// ListProductsServiceClient is the client API for ListProductsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListProductsServiceClient interface {
	ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error)
}

type listProductsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListProductsServiceClient(cc grpc.ClientConnInterface) ListProductsServiceClient {
	return &listProductsServiceClient{cc}
}

func (c *listProductsServiceClient) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, ListProductsService_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListProductsServiceServer is the server API for ListProductsService service.
// All implementations must embed UnimplementedListProductsServiceServer
// for forward compatibility.
type ListProductsServiceServer interface {
	ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error)
	mustEmbedUnimplementedListProductsServiceServer()
}

// UnimplementedListProductsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedListProductsServiceServer struct{}

func (UnimplementedListProductsServiceServer) ListProducts(context.Context, *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedListProductsServiceServer) mustEmbedUnimplementedListProductsServiceServer() {}
func (UnimplementedListProductsServiceServer) testEmbeddedByValue()                             {}

// UnsafeListProductsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListProductsServiceServer will
// result in compilation errors.
type UnsafeListProductsServiceServer interface {
	mustEmbedUnimplementedListProductsServiceServer()
}

func RegisterListProductsServiceServer(s grpc.ServiceRegistrar, srv ListProductsServiceServer) {
	// If the following call pancis, it indicates UnimplementedListProductsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ListProductsService_ServiceDesc, srv)
}

func _ListProductsService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListProductsServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ListProductsService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListProductsServiceServer).ListProducts(ctx, req.(*ListProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListProductsService_ServiceDesc is the grpc.ServiceDesc for ListProductsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListProductsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ListProductsService",
	HandlerType: (*ListProductsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _ListProductsService_ListProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "listProducts.proto",
}
