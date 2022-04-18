// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package volumegroupsnapshot

import (
	context "context"
	common "github.com/dell/dell-csi-extensions/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VolumeGroupSnapshotClient is the client API for VolumeGroupSnapshot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VolumeGroupSnapshotClient interface {
	// ProbeController is used to probe driver name by making grpc calls
	ProbeController(ctx context.Context, in *common.ProbeControllerRequest, opts ...grpc.CallOption) (*common.ProbeControllerResponse, error)
	// CreateVolumeGroupSnapshot will take in a CreateVolumeGroupSnapshotRequest that will contain:
	// 1. An array of volume IDs to be snapped for the volume snapshot group
	// 2. A name for the volume snapshot group
	// 3. Parameters for the VolumeGroupSnapshot instance
	// It will return a CreateVolumeGroupSnapshotResponse, which contains an array of snapshots, and an id for the group
	CreateVolumeGroupSnapshot(ctx context.Context, in *CreateVolumeGroupSnapshotRequest, opts ...grpc.CallOption) (*CreateVolumeGroupSnapshotResponse, error)
	// ParseVolumeHandle will take driver specific volume handles and returns back array, protocol and volume ID
	ParseVolumeHandle(ctx context.Context, in *VolumeHandleRequest, opts ...grpc.CallOption) (*VolumeHandleResponse, error)
}

type volumeGroupSnapshotClient struct {
	cc grpc.ClientConnInterface
}

func NewVolumeGroupSnapshotClient(cc grpc.ClientConnInterface) VolumeGroupSnapshotClient {
	return &volumeGroupSnapshotClient{cc}
}

func (c *volumeGroupSnapshotClient) ProbeController(ctx context.Context, in *common.ProbeControllerRequest, opts ...grpc.CallOption) (*common.ProbeControllerResponse, error) {
	out := new(common.ProbeControllerResponse)
	err := c.cc.Invoke(ctx, "/volumegroupsnapshot.v1.VolumeGroupSnapshot/ProbeController", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeGroupSnapshotClient) CreateVolumeGroupSnapshot(ctx context.Context, in *CreateVolumeGroupSnapshotRequest, opts ...grpc.CallOption) (*CreateVolumeGroupSnapshotResponse, error) {
	out := new(CreateVolumeGroupSnapshotResponse)
	err := c.cc.Invoke(ctx, "/volumegroupsnapshot.v1.VolumeGroupSnapshot/CreateVolumeGroupSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeGroupSnapshotClient) ParseVolumeHandle(ctx context.Context, in *VolumeHandleRequest, opts ...grpc.CallOption) (*VolumeHandleResponse, error) {
	out := new(VolumeHandleResponse)
	err := c.cc.Invoke(ctx, "/volumegroupsnapshot.v1.VolumeGroupSnapshot/ParseVolumeHandle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumeGroupSnapshotServer is the server API for VolumeGroupSnapshot service.
// All implementations should embed UnimplementedVolumeGroupSnapshotServer
// for forward compatibility
type VolumeGroupSnapshotServer interface {
	// ProbeController is used to probe driver name by making grpc calls
	ProbeController(context.Context, *common.ProbeControllerRequest) (*common.ProbeControllerResponse, error)
	// CreateVolumeGroupSnapshot will take in a CreateVolumeGroupSnapshotRequest that will contain:
	// 1. An array of volume IDs to be snapped for the volume snapshot group
	// 2. A name for the volume snapshot group
	// 3. Parameters for the VolumeGroupSnapshot instance
	// It will return a CreateVolumeGroupSnapshotResponse, which contains an array of snapshots, and an id for the group
	CreateVolumeGroupSnapshot(context.Context, *CreateVolumeGroupSnapshotRequest) (*CreateVolumeGroupSnapshotResponse, error)
	// ParseVolumeHandle will take driver specific volume handles and returns back array, protocol and volume ID
	ParseVolumeHandle(context.Context, *VolumeHandleRequest) (*VolumeHandleResponse, error)
}

// UnimplementedVolumeGroupSnapshotServer should be embedded to have forward compatible implementations.
type UnimplementedVolumeGroupSnapshotServer struct {
}

func (UnimplementedVolumeGroupSnapshotServer) ProbeController(context.Context, *common.ProbeControllerRequest) (*common.ProbeControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProbeController not implemented")
}
func (UnimplementedVolumeGroupSnapshotServer) CreateVolumeGroupSnapshot(context.Context, *CreateVolumeGroupSnapshotRequest) (*CreateVolumeGroupSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolumeGroupSnapshot not implemented")
}
func (UnimplementedVolumeGroupSnapshotServer) ParseVolumeHandle(context.Context, *VolumeHandleRequest) (*VolumeHandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseVolumeHandle not implemented")
}

// UnsafeVolumeGroupSnapshotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VolumeGroupSnapshotServer will
// result in compilation errors.
type UnsafeVolumeGroupSnapshotServer interface {
	mustEmbedUnimplementedVolumeGroupSnapshotServer()
}

func RegisterVolumeGroupSnapshotServer(s grpc.ServiceRegistrar, srv VolumeGroupSnapshotServer) {
	s.RegisterService(&VolumeGroupSnapshot_ServiceDesc, srv)
}

func _VolumeGroupSnapshot_ProbeController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.ProbeControllerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeGroupSnapshotServer).ProbeController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volumegroupsnapshot.v1.VolumeGroupSnapshot/ProbeController",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeGroupSnapshotServer).ProbeController(ctx, req.(*common.ProbeControllerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeGroupSnapshot_CreateVolumeGroupSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeGroupSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeGroupSnapshotServer).CreateVolumeGroupSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volumegroupsnapshot.v1.VolumeGroupSnapshot/CreateVolumeGroupSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeGroupSnapshotServer).CreateVolumeGroupSnapshot(ctx, req.(*CreateVolumeGroupSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeGroupSnapshot_ParseVolumeHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeGroupSnapshotServer).ParseVolumeHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volumegroupsnapshot.v1.VolumeGroupSnapshot/ParseVolumeHandle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeGroupSnapshotServer).ParseVolumeHandle(ctx, req.(*VolumeHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VolumeGroupSnapshot_ServiceDesc is the grpc.ServiceDesc for VolumeGroupSnapshot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VolumeGroupSnapshot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "volumegroupsnapshot.v1.VolumeGroupSnapshot",
	HandlerType: (*VolumeGroupSnapshotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProbeController",
			Handler:    _VolumeGroupSnapshot_ProbeController_Handler,
		},
		{
			MethodName: "CreateVolumeGroupSnapshot",
			Handler:    _VolumeGroupSnapshot_CreateVolumeGroupSnapshot_Handler,
		},
		{
			MethodName: "ParseVolumeHandle",
			Handler:    _VolumeGroupSnapshot_ParseVolumeHandle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "volumeGroupSnapshot.proto",
}
