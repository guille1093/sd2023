// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: pkg/base.proto

package pkg

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

const (
	Base_Guardar_FullMethodName  = "/base.Base/guardar"
	Base_Obtener_FullMethodName  = "/base.Base/obtener"
	Base_Eliminar_FullMethodName = "/base.Base/eliminar"
)

// BaseClient is the client API for Base service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseClient interface {
	Guardar(ctx context.Context, in *ParametroGuardar, opts ...grpc.CallOption) (*ResultadoGuardar, error)
	Obtener(ctx context.Context, in *ParametroObtenerEliminar, opts ...grpc.CallOption) (*ResultadoObtenerEliminar, error)
	Eliminar(ctx context.Context, in *ParametroObtenerEliminar, opts ...grpc.CallOption) (*ResultadoObtenerEliminar, error)
}

type baseClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseClient(cc grpc.ClientConnInterface) BaseClient {
	return &baseClient{cc}
}

func (c *baseClient) Guardar(ctx context.Context, in *ParametroGuardar, opts ...grpc.CallOption) (*ResultadoGuardar, error) {
	out := new(ResultadoGuardar)
	err := c.cc.Invoke(ctx, Base_Guardar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) Obtener(ctx context.Context, in *ParametroObtenerEliminar, opts ...grpc.CallOption) (*ResultadoObtenerEliminar, error) {
	out := new(ResultadoObtenerEliminar)
	err := c.cc.Invoke(ctx, Base_Obtener_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) Eliminar(ctx context.Context, in *ParametroObtenerEliminar, opts ...grpc.CallOption) (*ResultadoObtenerEliminar, error) {
	out := new(ResultadoObtenerEliminar)
	err := c.cc.Invoke(ctx, Base_Eliminar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServer is the server API for Base service.
// All implementations must embed UnimplementedBaseServer
// for forward compatibility
type BaseServer interface {
	Guardar(context.Context, *ParametroGuardar) (*ResultadoGuardar, error)
	Obtener(context.Context, *ParametroObtenerEliminar) (*ResultadoObtenerEliminar, error)
	Eliminar(context.Context, *ParametroObtenerEliminar) (*ResultadoObtenerEliminar, error)
	mustEmbedUnimplementedBaseServer()
}

// UnimplementedBaseServer must be embedded to have forward compatible implementations.
type UnimplementedBaseServer struct {
}

func (UnimplementedBaseServer) Guardar(context.Context, *ParametroGuardar) (*ResultadoGuardar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Guardar not implemented")
}
func (UnimplementedBaseServer) Obtener(context.Context, *ParametroObtenerEliminar) (*ResultadoObtenerEliminar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Obtener not implemented")
}
func (UnimplementedBaseServer) Eliminar(context.Context, *ParametroObtenerEliminar) (*ResultadoObtenerEliminar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Eliminar not implemented")
}
func (UnimplementedBaseServer) mustEmbedUnimplementedBaseServer() {}

// UnsafeBaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServer will
// result in compilation errors.
type UnsafeBaseServer interface {
	mustEmbedUnimplementedBaseServer()
}

func RegisterBaseServer(s grpc.ServiceRegistrar, srv BaseServer) {
	s.RegisterService(&Base_ServiceDesc, srv)
}

func _Base_Guardar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParametroGuardar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).Guardar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_Guardar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).Guardar(ctx, req.(*ParametroGuardar))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_Obtener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParametroObtenerEliminar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).Obtener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_Obtener_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).Obtener(ctx, req.(*ParametroObtenerEliminar))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_Eliminar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParametroObtenerEliminar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).Eliminar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_Eliminar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).Eliminar(ctx, req.(*ParametroObtenerEliminar))
	}
	return interceptor(ctx, in, info, handler)
}

// Base_ServiceDesc is the grpc.ServiceDesc for Base service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Base_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "base.Base",
	HandlerType: (*BaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "guardar",
			Handler:    _Base_Guardar_Handler,
		},
		{
			MethodName: "obtener",
			Handler:    _Base_Obtener_Handler,
		},
		{
			MethodName: "eliminar",
			Handler:    _Base_Eliminar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/base.proto",
}
