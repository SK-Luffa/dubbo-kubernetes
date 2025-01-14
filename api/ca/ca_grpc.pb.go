/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: ca.proto

package ca

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

// AuthorityServiceClient is the client API for AuthorityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorityServiceClient interface {
	CreateIdentity(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error)
}

type authorityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorityServiceClient(cc grpc.ClientConnInterface) AuthorityServiceClient {
	return &authorityServiceClient{cc}
}

func (c *authorityServiceClient) CreateIdentity(ctx context.Context, in *IdentityRequest, opts ...grpc.CallOption) (*IdentityResponse, error) {
	out := new(IdentityResponse)
	err := c.cc.Invoke(ctx, "/org.apache.dubbo.auth.v1alpha1.AuthorityService/CreateIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityServiceServer is the server API for AuthorityService service.
// All implementations must embed UnimplementedAuthorityServiceServer
// for forward compatibility
type AuthorityServiceServer interface {
	CreateIdentity(context.Context, *IdentityRequest) (*IdentityResponse, error)
	mustEmbedUnimplementedAuthorityServiceServer()
}

// UnimplementedAuthorityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorityServiceServer struct {
}

func (UnimplementedAuthorityServiceServer) CreateIdentity(context.Context, *IdentityRequest) (*IdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIdentity not implemented")
}
func (UnimplementedAuthorityServiceServer) mustEmbedUnimplementedAuthorityServiceServer() {}

// UnsafeAuthorityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorityServiceServer will
// result in compilation errors.
type UnsafeAuthorityServiceServer interface {
	mustEmbedUnimplementedAuthorityServiceServer()
}

func RegisterAuthorityServiceServer(s grpc.ServiceRegistrar, srv AuthorityServiceServer) {
	s.RegisterService(&AuthorityService_ServiceDesc, srv)
}

func _AuthorityService_CreateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).CreateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.apache.dubbo.auth.v1alpha1.AuthorityService/CreateIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).CreateIdentity(ctx, req.(*IdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorityService_ServiceDesc is the grpc.ServiceDesc for AuthorityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.apache.dubbo.auth.v1alpha1.AuthorityService",
	HandlerType: (*AuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIdentity",
			Handler:    _AuthorityService_CreateIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca.proto",
}
