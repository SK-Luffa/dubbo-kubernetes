// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: registry/v1alpha1/docker.proto

package registryv1alpha1

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
	DockerRepoService_CreateDockerRepo_FullMethodName       = "/bufman.dubbo.apache.org.registry.v1alpha1.DockerRepoService/CreateDockerRepo"
	DockerRepoService_GetDockerRepo_FullMethodName          = "/bufman.dubbo.apache.org.registry.v1alpha1.DockerRepoService/GetDockerRepo"
	DockerRepoService_GetDockerRepoByName_FullMethodName    = "/bufman.dubbo.apache.org.registry.v1alpha1.DockerRepoService/GetDockerRepoByName"
	DockerRepoService_ListDockerRepos_FullMethodName        = "/bufman.dubbo.apache.org.registry.v1alpha1.DockerRepoService/ListDockerRepos"
	DockerRepoService_UpdateDockerRepoByName_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.DockerRepoService/UpdateDockerRepoByName"
	DockerRepoService_UpdateDockerRepoByID_FullMethodName   = "/bufman.dubbo.apache.org.registry.v1alpha1.DockerRepoService/UpdateDockerRepoByID"
)

// DockerRepoServiceClient is the client API for DockerRepoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerRepoServiceClient interface {
	// CreateDockerRepo create a docker repo for user
	//
	// This method requires authentication.
	CreateDockerRepo(ctx context.Context, in *CreateDockerRepoRequest, opts ...grpc.CallOption) (*CreateDockerRepoResponse, error)
	// GetDockerRepo get a user's docker repo by id
	//
	// This method requires authentication.
	GetDockerRepo(ctx context.Context, in *GetDockerRepoRequest, opts ...grpc.CallOption) (*GetDockerRepoResponse, error)
	// GetDockerRepoByName get a user's docker repo by name
	//
	// This method requires authentication.
	GetDockerRepoByName(ctx context.Context, in *GetDockerRepoByNameRequest, opts ...grpc.CallOption) (*GetDockerRepoByNameResponse, error)
	// ListDockerRepos lists the user's all docker repo entries
	//
	// This method requires authentication.
	ListDockerRepos(ctx context.Context, in *ListDockerReposRequest, opts ...grpc.CallOption) (*ListDockerReposResponse, error)
	// UpdateDockerRepoByName given a name, to update address、username、password
	//
	// This method requires authentication.
	UpdateDockerRepoByName(ctx context.Context, in *UpdateDockerRepoByNameRequest, opts ...grpc.CallOption) (*UpdateDockerRepoByNameResponse, error)
	// UpdateDockerRepoByName given a id, to update address、username、password
	//
	// This method requires authentication.
	UpdateDockerRepoByID(ctx context.Context, in *UpdateDockerRepoByIDRequest, opts ...grpc.CallOption) (*UpdateDockerRepoByIDResponse, error)
}

type dockerRepoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerRepoServiceClient(cc grpc.ClientConnInterface) DockerRepoServiceClient {
	return &dockerRepoServiceClient{cc}
}

func (c *dockerRepoServiceClient) CreateDockerRepo(ctx context.Context, in *CreateDockerRepoRequest, opts ...grpc.CallOption) (*CreateDockerRepoResponse, error) {
	out := new(CreateDockerRepoResponse)
	err := c.cc.Invoke(ctx, DockerRepoService_CreateDockerRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerRepoServiceClient) GetDockerRepo(ctx context.Context, in *GetDockerRepoRequest, opts ...grpc.CallOption) (*GetDockerRepoResponse, error) {
	out := new(GetDockerRepoResponse)
	err := c.cc.Invoke(ctx, DockerRepoService_GetDockerRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerRepoServiceClient) GetDockerRepoByName(ctx context.Context, in *GetDockerRepoByNameRequest, opts ...grpc.CallOption) (*GetDockerRepoByNameResponse, error) {
	out := new(GetDockerRepoByNameResponse)
	err := c.cc.Invoke(ctx, DockerRepoService_GetDockerRepoByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerRepoServiceClient) ListDockerRepos(ctx context.Context, in *ListDockerReposRequest, opts ...grpc.CallOption) (*ListDockerReposResponse, error) {
	out := new(ListDockerReposResponse)
	err := c.cc.Invoke(ctx, DockerRepoService_ListDockerRepos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerRepoServiceClient) UpdateDockerRepoByName(ctx context.Context, in *UpdateDockerRepoByNameRequest, opts ...grpc.CallOption) (*UpdateDockerRepoByNameResponse, error) {
	out := new(UpdateDockerRepoByNameResponse)
	err := c.cc.Invoke(ctx, DockerRepoService_UpdateDockerRepoByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerRepoServiceClient) UpdateDockerRepoByID(ctx context.Context, in *UpdateDockerRepoByIDRequest, opts ...grpc.CallOption) (*UpdateDockerRepoByIDResponse, error) {
	out := new(UpdateDockerRepoByIDResponse)
	err := c.cc.Invoke(ctx, DockerRepoService_UpdateDockerRepoByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerRepoServiceServer is the server API for DockerRepoService service.
// All implementations must embed UnimplementedDockerRepoServiceServer
// for forward compatibility
type DockerRepoServiceServer interface {
	// CreateDockerRepo create a docker repo for user
	//
	// This method requires authentication.
	CreateDockerRepo(context.Context, *CreateDockerRepoRequest) (*CreateDockerRepoResponse, error)
	// GetDockerRepo get a user's docker repo by id
	//
	// This method requires authentication.
	GetDockerRepo(context.Context, *GetDockerRepoRequest) (*GetDockerRepoResponse, error)
	// GetDockerRepoByName get a user's docker repo by name
	//
	// This method requires authentication.
	GetDockerRepoByName(context.Context, *GetDockerRepoByNameRequest) (*GetDockerRepoByNameResponse, error)
	// ListDockerRepos lists the user's all docker repo entries
	//
	// This method requires authentication.
	ListDockerRepos(context.Context, *ListDockerReposRequest) (*ListDockerReposResponse, error)
	// UpdateDockerRepoByName given a name, to update address、username、password
	//
	// This method requires authentication.
	UpdateDockerRepoByName(context.Context, *UpdateDockerRepoByNameRequest) (*UpdateDockerRepoByNameResponse, error)
	// UpdateDockerRepoByName given a id, to update address、username、password
	//
	// This method requires authentication.
	UpdateDockerRepoByID(context.Context, *UpdateDockerRepoByIDRequest) (*UpdateDockerRepoByIDResponse, error)
	mustEmbedUnimplementedDockerRepoServiceServer()
}

// UnimplementedDockerRepoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDockerRepoServiceServer struct {
}

func (UnimplementedDockerRepoServiceServer) CreateDockerRepo(context.Context, *CreateDockerRepoRequest) (*CreateDockerRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDockerRepo not implemented")
}
func (UnimplementedDockerRepoServiceServer) GetDockerRepo(context.Context, *GetDockerRepoRequest) (*GetDockerRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDockerRepo not implemented")
}
func (UnimplementedDockerRepoServiceServer) GetDockerRepoByName(context.Context, *GetDockerRepoByNameRequest) (*GetDockerRepoByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDockerRepoByName not implemented")
}
func (UnimplementedDockerRepoServiceServer) ListDockerRepos(context.Context, *ListDockerReposRequest) (*ListDockerReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDockerRepos not implemented")
}
func (UnimplementedDockerRepoServiceServer) UpdateDockerRepoByName(context.Context, *UpdateDockerRepoByNameRequest) (*UpdateDockerRepoByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDockerRepoByName not implemented")
}
func (UnimplementedDockerRepoServiceServer) UpdateDockerRepoByID(context.Context, *UpdateDockerRepoByIDRequest) (*UpdateDockerRepoByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDockerRepoByID not implemented")
}
func (UnimplementedDockerRepoServiceServer) mustEmbedUnimplementedDockerRepoServiceServer() {}

// UnsafeDockerRepoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerRepoServiceServer will
// result in compilation errors.
type UnsafeDockerRepoServiceServer interface {
	mustEmbedUnimplementedDockerRepoServiceServer()
}

func RegisterDockerRepoServiceServer(s grpc.ServiceRegistrar, srv DockerRepoServiceServer) {
	s.RegisterService(&DockerRepoService_ServiceDesc, srv)
}

func _DockerRepoService_CreateDockerRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDockerRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerRepoServiceServer).CreateDockerRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerRepoService_CreateDockerRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerRepoServiceServer).CreateDockerRepo(ctx, req.(*CreateDockerRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerRepoService_GetDockerRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDockerRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerRepoServiceServer).GetDockerRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerRepoService_GetDockerRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerRepoServiceServer).GetDockerRepo(ctx, req.(*GetDockerRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerRepoService_GetDockerRepoByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDockerRepoByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerRepoServiceServer).GetDockerRepoByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerRepoService_GetDockerRepoByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerRepoServiceServer).GetDockerRepoByName(ctx, req.(*GetDockerRepoByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerRepoService_ListDockerRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDockerReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerRepoServiceServer).ListDockerRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerRepoService_ListDockerRepos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerRepoServiceServer).ListDockerRepos(ctx, req.(*ListDockerReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerRepoService_UpdateDockerRepoByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDockerRepoByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerRepoServiceServer).UpdateDockerRepoByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerRepoService_UpdateDockerRepoByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerRepoServiceServer).UpdateDockerRepoByName(ctx, req.(*UpdateDockerRepoByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerRepoService_UpdateDockerRepoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDockerRepoByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerRepoServiceServer).UpdateDockerRepoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DockerRepoService_UpdateDockerRepoByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerRepoServiceServer).UpdateDockerRepoByID(ctx, req.(*UpdateDockerRepoByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerRepoService_ServiceDesc is the grpc.ServiceDesc for DockerRepoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerRepoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.DockerRepoService",
	HandlerType: (*DockerRepoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDockerRepo",
			Handler:    _DockerRepoService_CreateDockerRepo_Handler,
		},
		{
			MethodName: "GetDockerRepo",
			Handler:    _DockerRepoService_GetDockerRepo_Handler,
		},
		{
			MethodName: "GetDockerRepoByName",
			Handler:    _DockerRepoService_GetDockerRepoByName_Handler,
		},
		{
			MethodName: "ListDockerRepos",
			Handler:    _DockerRepoService_ListDockerRepos_Handler,
		},
		{
			MethodName: "UpdateDockerRepoByName",
			Handler:    _DockerRepoService_UpdateDockerRepoByName_Handler,
		},
		{
			MethodName: "UpdateDockerRepoByID",
			Handler:    _DockerRepoService_UpdateDockerRepoByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/docker.proto",
}
