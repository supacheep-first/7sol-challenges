// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: summary.proto

package summary

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SummaryService_GetSummary_FullMethodName = "/summary.SummaryService/GetSummary"
)

// SummaryServiceClient is the client API for SummaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SummaryServiceClient interface {
	GetSummary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (*SummaryResponse, error)
}

type summaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSummaryServiceClient(cc grpc.ClientConnInterface) SummaryServiceClient {
	return &summaryServiceClient{cc}
}

func (c *summaryServiceClient) GetSummary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (*SummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SummaryResponse)
	err := c.cc.Invoke(ctx, SummaryService_GetSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummaryServiceServer is the server API for SummaryService service.
// All implementations must embed UnimplementedSummaryServiceServer
// for forward compatibility.
type SummaryServiceServer interface {
	GetSummary(context.Context, *SummaryRequest) (*SummaryResponse, error)
	mustEmbedUnimplementedSummaryServiceServer()
}

// UnimplementedSummaryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSummaryServiceServer struct{}

func (UnimplementedSummaryServiceServer) GetSummary(context.Context, *SummaryRequest) (*SummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummary not implemented")
}
func (UnimplementedSummaryServiceServer) mustEmbedUnimplementedSummaryServiceServer() {}
func (UnimplementedSummaryServiceServer) testEmbeddedByValue()                        {}

// UnsafeSummaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SummaryServiceServer will
// result in compilation errors.
type UnsafeSummaryServiceServer interface {
	mustEmbedUnimplementedSummaryServiceServer()
}

func RegisterSummaryServiceServer(s grpc.ServiceRegistrar, srv SummaryServiceServer) {
	// If the following call pancis, it indicates UnimplementedSummaryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SummaryService_ServiceDesc, srv)
}

func _SummaryService_GetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummaryServiceServer).GetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SummaryService_GetSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummaryServiceServer).GetSummary(ctx, req.(*SummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SummaryService_ServiceDesc is the grpc.ServiceDesc for SummaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SummaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "summary.SummaryService",
	HandlerType: (*SummaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSummary",
			Handler:    _SummaryService_GetSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "summary.proto",
}
