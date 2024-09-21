// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: interactions.proto

package proto

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
	InteractionService_AddInteraction_FullMethodName = "/InteractionService/AddInteraction"
	InteractionService_GetInteraction_FullMethodName = "/InteractionService/GetInteraction"
)

// InteractionServiceClient is the client API for InteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionServiceClient interface {
	AddInteraction(ctx context.Context, in *AddInteractionRequest, opts ...grpc.CallOption) (*InteractionResponse, error)
	GetInteraction(ctx context.Context, in *GetInteractionRequest, opts ...grpc.CallOption) (*GetInteractionResponse, error)
}

type interactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionServiceClient(cc grpc.ClientConnInterface) InteractionServiceClient {
	return &interactionServiceClient{cc}
}

func (c *interactionServiceClient) AddInteraction(ctx context.Context, in *AddInteractionRequest, opts ...grpc.CallOption) (*InteractionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InteractionResponse)
	err := c.cc.Invoke(ctx, InteractionService_AddInteraction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) GetInteraction(ctx context.Context, in *GetInteractionRequest, opts ...grpc.CallOption) (*GetInteractionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInteractionResponse)
	err := c.cc.Invoke(ctx, InteractionService_GetInteraction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServiceServer is the server API for InteractionService service.
// All implementations must embed UnimplementedInteractionServiceServer
// for forward compatibility.
type InteractionServiceServer interface {
	AddInteraction(context.Context, *AddInteractionRequest) (*InteractionResponse, error)
	GetInteraction(context.Context, *GetInteractionRequest) (*GetInteractionResponse, error)
	mustEmbedUnimplementedInteractionServiceServer()
}

// UnimplementedInteractionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInteractionServiceServer struct{}

func (UnimplementedInteractionServiceServer) AddInteraction(context.Context, *AddInteractionRequest) (*InteractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInteraction not implemented")
}
func (UnimplementedInteractionServiceServer) GetInteraction(context.Context, *GetInteractionRequest) (*GetInteractionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInteraction not implemented")
}
func (UnimplementedInteractionServiceServer) mustEmbedUnimplementedInteractionServiceServer() {}
func (UnimplementedInteractionServiceServer) testEmbeddedByValue()                            {}

// UnsafeInteractionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionServiceServer will
// result in compilation errors.
type UnsafeInteractionServiceServer interface {
	mustEmbedUnimplementedInteractionServiceServer()
}

func RegisterInteractionServiceServer(s grpc.ServiceRegistrar, srv InteractionServiceServer) {
	// If the following call pancis, it indicates UnimplementedInteractionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InteractionService_ServiceDesc, srv)
}

func _InteractionService_AddInteraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInteractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).AddInteraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_AddInteraction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).AddInteraction(ctx, req.(*AddInteractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_GetInteraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInteractionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).GetInteraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractionService_GetInteraction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).GetInteraction(ctx, req.(*GetInteractionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionService_ServiceDesc is the grpc.ServiceDesc for InteractionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddInteraction",
			Handler:    _InteractionService_AddInteraction_Handler,
		},
		{
			MethodName: "GetInteraction",
			Handler:    _InteractionService_GetInteraction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interactions.proto",
}
