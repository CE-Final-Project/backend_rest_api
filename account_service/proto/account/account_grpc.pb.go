// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: account.proto

package accountService

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error)
	GetAccountById(ctx context.Context, in *GetAccountByIdReq, opts ...grpc.CallOption) (*GetAccountByIdRes, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error) {
	out := new(CreateAccountRes)
	err := c.cc.Invoke(ctx, "/accountService.accountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error) {
	out := new(UpdateAccountRes)
	err := c.cc.Invoke(ctx, "/accountService.accountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountById(ctx context.Context, in *GetAccountByIdReq, opts ...grpc.CallOption) (*GetAccountByIdRes, error) {
	out := new(GetAccountByIdRes)
	err := c.cc.Invoke(ctx, "/accountService.accountService/GetAccountById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations should embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRes, error)
	UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRes, error)
	GetAccountById(context.Context, *GetAccountByIdReq) (*GetAccountByIdRes, error)
}

// UnimplementedAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountById(context.Context, *GetAccountByIdReq) (*GetAccountByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountService.accountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountService.accountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accountService.accountService/GetAccountById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountById(ctx, req.(*GetAccountByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accountService.accountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "GetAccountById",
			Handler:    _AccountService_GetAccountById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
