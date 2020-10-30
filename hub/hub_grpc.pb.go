// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hub

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HubClient is the client API for Hub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubClient interface {
	SelectServer(ctx context.Context, in *MessageJSON, opts ...grpc.CallOption) (*MessageJSON, error)
	RunRoute(ctx context.Context, in *MessageJSON, opts ...grpc.CallOption) (*MessageJSON, error)
}

type hubClient struct {
	cc grpc.ClientConnInterface
}

func NewHubClient(cc grpc.ClientConnInterface) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) SelectServer(ctx context.Context, in *MessageJSON, opts ...grpc.CallOption) (*MessageJSON, error) {
	out := new(MessageJSON)
	err := c.cc.Invoke(ctx, "/hub.Hub/SelectServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) RunRoute(ctx context.Context, in *MessageJSON, opts ...grpc.CallOption) (*MessageJSON, error) {
	out := new(MessageJSON)
	err := c.cc.Invoke(ctx, "/hub.Hub/RunRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubServer is the server API for Hub service.
// All implementations must embed UnimplementedHubServer
// for forward compatibility
type HubServer interface {
	SelectServer(context.Context, *MessageJSON) (*MessageJSON, error)
	RunRoute(context.Context, *MessageJSON) (*MessageJSON, error)
	mustEmbedUnimplementedHubServer()
}

// UnimplementedHubServer must be embedded to have forward compatible implementations.
type UnimplementedHubServer struct {
}

func (UnimplementedHubServer) SelectServer(context.Context, *MessageJSON) (*MessageJSON, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectServer not implemented")
}
func (UnimplementedHubServer) RunRoute(context.Context, *MessageJSON) (*MessageJSON, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunRoute not implemented")
}
func (UnimplementedHubServer) mustEmbedUnimplementedHubServer() {}

// UnsafeHubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServer will
// result in compilation errors.
type UnsafeHubServer interface {
	mustEmbedUnimplementedHubServer()
}

func RegisterHubServer(s grpc.ServiceRegistrar, srv HubServer) {
	s.RegisterService(&_Hub_serviceDesc, srv)
}

func _Hub_SelectServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageJSON)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).SelectServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hub.Hub/SelectServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).SelectServer(ctx, req.(*MessageJSON))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_RunRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageJSON)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).RunRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hub.Hub/RunRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).RunRoute(ctx, req.(*MessageJSON))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hub.Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelectServer",
			Handler:    _Hub_SelectServer_Handler,
		},
		{
			MethodName: "RunRoute",
			Handler:    _Hub_RunRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub/hub.proto",
}
