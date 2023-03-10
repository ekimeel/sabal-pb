// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: point_service.proto

package pb

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

// PointServiceClient is the client API for PointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointServiceClient interface {
	Get(ctx context.Context, in *PointId, opts ...grpc.CallOption) (*Point, error)
	GetAll(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*PointList, error)
	Create(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error)
	Update(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error)
	Delete(ctx context.Context, in *PointId, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetOrCreate(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error)
	GetAllByEquip(ctx context.Context, in *EquipId, opts ...grpc.CallOption) (*PointList, error)
}

type pointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPointServiceClient(cc grpc.ClientConnInterface) PointServiceClient {
	return &pointServiceClient{cc}
}

func (c *pointServiceClient) Get(ctx context.Context, in *PointId, opts ...grpc.CallOption) (*Point, error) {
	out := new(Point)
	err := c.cc.Invoke(ctx, "/pb.PointService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) GetAll(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*PointList, error) {
	out := new(PointList)
	err := c.cc.Invoke(ctx, "/pb.PointService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) Create(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error) {
	out := new(Point)
	err := c.cc.Invoke(ctx, "/pb.PointService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) Update(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error) {
	out := new(Point)
	err := c.cc.Invoke(ctx, "/pb.PointService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) Delete(ctx context.Context, in *PointId, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.PointService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) GetOrCreate(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error) {
	out := new(Point)
	err := c.cc.Invoke(ctx, "/pb.PointService/GetOrCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointServiceClient) GetAllByEquip(ctx context.Context, in *EquipId, opts ...grpc.CallOption) (*PointList, error) {
	out := new(PointList)
	err := c.cc.Invoke(ctx, "/pb.PointService/GetAllByEquip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointServiceServer is the server API for PointService service.
// All implementations must embed UnimplementedPointServiceServer
// for forward compatibility
type PointServiceServer interface {
	Get(context.Context, *PointId) (*Point, error)
	GetAll(context.Context, *ListRequest) (*PointList, error)
	Create(context.Context, *Point) (*Point, error)
	Update(context.Context, *Point) (*Point, error)
	Delete(context.Context, *PointId) (*DeleteResponse, error)
	GetOrCreate(context.Context, *Point) (*Point, error)
	GetAllByEquip(context.Context, *EquipId) (*PointList, error)
	mustEmbedUnimplementedPointServiceServer()
}

// UnimplementedPointServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPointServiceServer struct {
}

func (UnimplementedPointServiceServer) Get(context.Context, *PointId) (*Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPointServiceServer) GetAll(context.Context, *ListRequest) (*PointList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPointServiceServer) Create(context.Context, *Point) (*Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPointServiceServer) Update(context.Context, *Point) (*Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPointServiceServer) Delete(context.Context, *PointId) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPointServiceServer) GetOrCreate(context.Context, *Point) (*Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreate not implemented")
}
func (UnimplementedPointServiceServer) GetAllByEquip(context.Context, *EquipId) (*PointList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByEquip not implemented")
}
func (UnimplementedPointServiceServer) mustEmbedUnimplementedPointServiceServer() {}

// UnsafePointServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointServiceServer will
// result in compilation errors.
type UnsafePointServiceServer interface {
	mustEmbedUnimplementedPointServiceServer()
}

func RegisterPointServiceServer(s grpc.ServiceRegistrar, srv PointServiceServer) {
	s.RegisterService(&PointService_ServiceDesc, srv)
}

func _PointService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PointService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).Get(ctx, req.(*PointId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PointService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).GetAll(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PointService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).Create(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PointService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).Update(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PointService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).Delete(ctx, req.(*PointId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_GetOrCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).GetOrCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PointService/GetOrCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).GetOrCreate(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointService_GetAllByEquip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EquipId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointServiceServer).GetAllByEquip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PointService/GetAllByEquip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointServiceServer).GetAllByEquip(ctx, req.(*EquipId))
	}
	return interceptor(ctx, in, info, handler)
}

// PointService_ServiceDesc is the grpc.ServiceDesc for PointService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PointService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PointService",
	HandlerType: (*PointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PointService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PointService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PointService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PointService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PointService_Delete_Handler,
		},
		{
			MethodName: "GetOrCreate",
			Handler:    _PointService_GetOrCreate_Handler,
		},
		{
			MethodName: "GetAllByEquip",
			Handler:    _PointService_GetAllByEquip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "point_service.proto",
}
