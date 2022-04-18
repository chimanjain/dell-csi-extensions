// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package migration

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

// MigrationClient is the client API for Migration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MigrationClient interface {
	VolumeMigrate(ctx context.Context, in *VolumeMigrateRequest, opts ...grpc.CallOption) (*VolumeMigrateResponse, error)
	// GetMigrationCapabilities is used to query CSI drivers for their supported migration capabilities
	GetMigrationCapabilities(ctx context.Context, in *GetMigrationCapabilityRequest, opts ...grpc.CallOption) (*GetMigrationCapabilityResponse, error)
}

type migrationClient struct {
	cc grpc.ClientConnInterface
}

func NewMigrationClient(cc grpc.ClientConnInterface) MigrationClient {
	return &migrationClient{cc}
}

func (c *migrationClient) VolumeMigrate(ctx context.Context, in *VolumeMigrateRequest, opts ...grpc.CallOption) (*VolumeMigrateResponse, error) {
	out := new(VolumeMigrateResponse)
	err := c.cc.Invoke(ctx, "/migration.v1alpha1.Migration/VolumeMigrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationClient) GetMigrationCapabilities(ctx context.Context, in *GetMigrationCapabilityRequest, opts ...grpc.CallOption) (*GetMigrationCapabilityResponse, error) {
	out := new(GetMigrationCapabilityResponse)
	err := c.cc.Invoke(ctx, "/migration.v1alpha1.Migration/GetMigrationCapabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MigrationServer is the server API for Migration service.
// All implementations should embed UnimplementedMigrationServer
// for forward compatibility
type MigrationServer interface {
	VolumeMigrate(context.Context, *VolumeMigrateRequest) (*VolumeMigrateResponse, error)
	// GetMigrationCapabilities is used to query CSI drivers for their supported migration capabilities
	GetMigrationCapabilities(context.Context, *GetMigrationCapabilityRequest) (*GetMigrationCapabilityResponse, error)
}

// UnimplementedMigrationServer should be embedded to have forward compatible implementations.
type UnimplementedMigrationServer struct {
}

func (UnimplementedMigrationServer) VolumeMigrate(context.Context, *VolumeMigrateRequest) (*VolumeMigrateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumeMigrate not implemented")
}
func (UnimplementedMigrationServer) GetMigrationCapabilities(context.Context, *GetMigrationCapabilityRequest) (*GetMigrationCapabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMigrationCapabilities not implemented")
}

// UnsafeMigrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MigrationServer will
// result in compilation errors.
type UnsafeMigrationServer interface {
	mustEmbedUnimplementedMigrationServer()
}

func RegisterMigrationServer(s grpc.ServiceRegistrar, srv MigrationServer) {
	s.RegisterService(&Migration_ServiceDesc, srv)
}

func _Migration_VolumeMigrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeMigrateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).VolumeMigrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.v1alpha1.Migration/VolumeMigrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).VolumeMigrate(ctx, req.(*VolumeMigrateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Migration_GetMigrationCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMigrationCapabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).GetMigrationCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/migration.v1alpha1.Migration/GetMigrationCapabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).GetMigrationCapabilities(ctx, req.(*GetMigrationCapabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Migration_ServiceDesc is the grpc.ServiceDesc for Migration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Migration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "migration.v1alpha1.Migration",
	HandlerType: (*MigrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VolumeMigrate",
			Handler:    _Migration_VolumeMigrate_Handler,
		},
		{
			MethodName: "GetMigrationCapabilities",
			Handler:    _Migration_GetMigrationCapabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "migration.proto",
}
