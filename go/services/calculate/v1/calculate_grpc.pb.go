// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: services/calculate/v1/calculate.proto

package calculatev1

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

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Addition(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error)
	Division(ctx context.Context, in *DivisionRequest, opts ...grpc.CallOption) (*DivisionResponse, error)
	Additions(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AdditionsClient, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Addition(ctx context.Context, in *AdditionRequest, opts ...grpc.CallOption) (*AdditionResponse, error) {
	out := new(AdditionResponse)
	err := c.cc.Invoke(ctx, "/services.calculate.v1.CalculateService/Addition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) Division(ctx context.Context, in *DivisionRequest, opts ...grpc.CallOption) (*DivisionResponse, error) {
	out := new(DivisionResponse)
	err := c.cc.Invoke(ctx, "/services.calculate.v1.CalculateService/Division", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) Additions(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AdditionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculateService_ServiceDesc.Streams[0], "/services.calculate.v1.CalculateService/Additions", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceAdditionsClient{stream}
	return x, nil
}

type CalculateService_AdditionsClient interface {
	Send(*AdditionsRequest) error
	Recv() (*AdditionsResponse, error)
	grpc.ClientStream
}

type calculateServiceAdditionsClient struct {
	grpc.ClientStream
}

func (x *calculateServiceAdditionsClient) Send(m *AdditionsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceAdditionsClient) Recv() (*AdditionsResponse, error) {
	m := new(AdditionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateServiceServer is the server API for CalculateService service.
// All implementations should embed UnimplementedCalculateServiceServer
// for forward compatibility
type CalculateServiceServer interface {
	Addition(context.Context, *AdditionRequest) (*AdditionResponse, error)
	Division(context.Context, *DivisionRequest) (*DivisionResponse, error)
	Additions(CalculateService_AdditionsServer) error
}

// UnimplementedCalculateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (UnimplementedCalculateServiceServer) Addition(context.Context, *AdditionRequest) (*AdditionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Addition not implemented")
}
func (UnimplementedCalculateServiceServer) Division(context.Context, *DivisionRequest) (*DivisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Division not implemented")
}
func (UnimplementedCalculateServiceServer) Additions(CalculateService_AdditionsServer) error {
	return status.Errorf(codes.Unimplemented, "method Additions not implemented")
}

// UnsafeCalculateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculateServiceServer will
// result in compilation errors.
type UnsafeCalculateServiceServer interface {
	mustEmbedUnimplementedCalculateServiceServer()
}

func RegisterCalculateServiceServer(s grpc.ServiceRegistrar, srv CalculateServiceServer) {
	s.RegisterService(&CalculateService_ServiceDesc, srv)
}

func _CalculateService_Addition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Addition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.calculate.v1.CalculateService/Addition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Addition(ctx, req.(*AdditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_Division_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Division(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.calculate.v1.CalculateService/Division",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Division(ctx, req.(*DivisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_Additions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).Additions(&calculateServiceAdditionsServer{stream})
}

type CalculateService_AdditionsServer interface {
	Send(*AdditionsResponse) error
	Recv() (*AdditionsRequest, error)
	grpc.ServerStream
}

type calculateServiceAdditionsServer struct {
	grpc.ServerStream
}

func (x *calculateServiceAdditionsServer) Send(m *AdditionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceAdditionsServer) Recv() (*AdditionsRequest, error) {
	m := new(AdditionsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculateService_ServiceDesc is the grpc.ServiceDesc for CalculateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.calculate.v1.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Addition",
			Handler:    _CalculateService_Addition_Handler,
		},
		{
			MethodName: "Division",
			Handler:    _CalculateService_Division_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Additions",
			Handler:       _CalculateService_Additions_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services/calculate/v1/calculate.proto",
}
