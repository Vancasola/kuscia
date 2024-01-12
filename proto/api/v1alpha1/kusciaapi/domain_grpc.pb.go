// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.8
// source: kuscia/proto/api/v1alpha1/kusciaapi/domain.proto

package kusciaapi

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
	DomainService_CreateDomain_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainService/CreateDomain"
	DomainService_QueryDomain_FullMethodName      = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainService/QueryDomain"
	DomainService_UpdateDomain_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainService/UpdateDomain"
	DomainService_DeleteDomain_FullMethodName     = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainService/DeleteDomain"
	DomainService_BatchQueryDomain_FullMethodName = "/kuscia.proto.api.v1alpha1.kusciaapi.DomainService/BatchQueryDomain"
)

// DomainServiceClient is the client API for DomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DomainServiceClient interface {
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error)
	QueryDomain(ctx context.Context, in *QueryDomainRequest, opts ...grpc.CallOption) (*QueryDomainResponse, error)
	UpdateDomain(ctx context.Context, in *UpdateDomainRequest, opts ...grpc.CallOption) (*UpdateDomainResponse, error)
	DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*DeleteDomainResponse, error)
	BatchQueryDomain(ctx context.Context, in *BatchQueryDomainRequest, opts ...grpc.CallOption) (*BatchQueryDomainResponse, error)
}

type domainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDomainServiceClient(cc grpc.ClientConnInterface) DomainServiceClient {
	return &domainServiceClient{cc}
}

func (c *domainServiceClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error) {
	out := new(CreateDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_CreateDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) QueryDomain(ctx context.Context, in *QueryDomainRequest, opts ...grpc.CallOption) (*QueryDomainResponse, error) {
	out := new(QueryDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_QueryDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) UpdateDomain(ctx context.Context, in *UpdateDomainRequest, opts ...grpc.CallOption) (*UpdateDomainResponse, error) {
	out := new(UpdateDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_UpdateDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*DeleteDomainResponse, error) {
	out := new(DeleteDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_DeleteDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainServiceClient) BatchQueryDomain(ctx context.Context, in *BatchQueryDomainRequest, opts ...grpc.CallOption) (*BatchQueryDomainResponse, error) {
	out := new(BatchQueryDomainResponse)
	err := c.cc.Invoke(ctx, DomainService_BatchQueryDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainServiceServer is the server API for DomainService service.
// All implementations must embed UnimplementedDomainServiceServer
// for forward compatibility
type DomainServiceServer interface {
	CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error)
	QueryDomain(context.Context, *QueryDomainRequest) (*QueryDomainResponse, error)
	UpdateDomain(context.Context, *UpdateDomainRequest) (*UpdateDomainResponse, error)
	DeleteDomain(context.Context, *DeleteDomainRequest) (*DeleteDomainResponse, error)
	BatchQueryDomain(context.Context, *BatchQueryDomainRequest) (*BatchQueryDomainResponse, error)
	mustEmbedUnimplementedDomainServiceServer()
}

// UnimplementedDomainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDomainServiceServer struct {
}

func (UnimplementedDomainServiceServer) CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedDomainServiceServer) QueryDomain(context.Context, *QueryDomainRequest) (*QueryDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDomain not implemented")
}
func (UnimplementedDomainServiceServer) UpdateDomain(context.Context, *UpdateDomainRequest) (*UpdateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDomain not implemented")
}
func (UnimplementedDomainServiceServer) DeleteDomain(context.Context, *DeleteDomainRequest) (*DeleteDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomain not implemented")
}
func (UnimplementedDomainServiceServer) BatchQueryDomain(context.Context, *BatchQueryDomainRequest) (*BatchQueryDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchQueryDomain not implemented")
}
func (UnimplementedDomainServiceServer) mustEmbedUnimplementedDomainServiceServer() {}

// UnsafeDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DomainServiceServer will
// result in compilation errors.
type UnsafeDomainServiceServer interface {
	mustEmbedUnimplementedDomainServiceServer()
}

func RegisterDomainServiceServer(s grpc.ServiceRegistrar, srv DomainServiceServer) {
	s.RegisterService(&DomainService_ServiceDesc, srv)
}

func _DomainService_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_CreateDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_QueryDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).QueryDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_QueryDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).QueryDomain(ctx, req.(*QueryDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_UpdateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).UpdateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_UpdateDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).UpdateDomain(ctx, req.(*UpdateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_DeleteDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).DeleteDomain(ctx, req.(*DeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DomainService_BatchQueryDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchQueryDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServiceServer).BatchQueryDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DomainService_BatchQueryDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServiceServer).BatchQueryDomain(ctx, req.(*BatchQueryDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DomainService_ServiceDesc is the grpc.ServiceDesc for DomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kuscia.proto.api.v1alpha1.kusciaapi.DomainService",
	HandlerType: (*DomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDomain",
			Handler:    _DomainService_CreateDomain_Handler,
		},
		{
			MethodName: "QueryDomain",
			Handler:    _DomainService_QueryDomain_Handler,
		},
		{
			MethodName: "UpdateDomain",
			Handler:    _DomainService_UpdateDomain_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _DomainService_DeleteDomain_Handler,
		},
		{
			MethodName: "BatchQueryDomain",
			Handler:    _DomainService_BatchQueryDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kuscia/proto/api/v1alpha1/kusciaapi/domain.proto",
}
