// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: rng.proto

package dtrngpb

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

// DTRngClient is the client API for DTRng service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DTRngClient interface {
	// getRngs - get rngs
	GetRngs(ctx context.Context, in *RequestRngs, opts ...grpc.CallOption) (*ReplyRngs, error)
}

type dTRngClient struct {
	cc grpc.ClientConnInterface
}

func NewDTRngClient(cc grpc.ClientConnInterface) DTRngClient {
	return &dTRngClient{cc}
}

func (c *dTRngClient) GetRngs(ctx context.Context, in *RequestRngs, opts ...grpc.CallOption) (*ReplyRngs, error) {
	out := new(ReplyRngs)
	err := c.cc.Invoke(ctx, "/dtrngpb.DTRng/getRngs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DTRngServer is the server API for DTRng service.
// All implementations must embed UnimplementedDTRngServer
// for forward compatibility
type DTRngServer interface {
	// getRngs - get rngs
	GetRngs(context.Context, *RequestRngs) (*ReplyRngs, error)
	mustEmbedUnimplementedDTRngServer()
}

// UnimplementedDTRngServer must be embedded to have forward compatible implementations.
type UnimplementedDTRngServer struct {
}

func (UnimplementedDTRngServer) GetRngs(context.Context, *RequestRngs) (*ReplyRngs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRngs not implemented")
}
func (UnimplementedDTRngServer) mustEmbedUnimplementedDTRngServer() {}

// UnsafeDTRngServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DTRngServer will
// result in compilation errors.
type UnsafeDTRngServer interface {
	mustEmbedUnimplementedDTRngServer()
}

func RegisterDTRngServer(s grpc.ServiceRegistrar, srv DTRngServer) {
	s.RegisterService(&DTRng_ServiceDesc, srv)
}

func _DTRng_GetRngs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRngs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DTRngServer).GetRngs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtrngpb.DTRng/getRngs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DTRngServer).GetRngs(ctx, req.(*RequestRngs))
	}
	return interceptor(ctx, in, info, handler)
}

// DTRng_ServiceDesc is the grpc.ServiceDesc for DTRng service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DTRng_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dtrngpb.DTRng",
	HandlerType: (*DTRngServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getRngs",
			Handler:    _DTRng_GetRngs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rng.proto",
}
