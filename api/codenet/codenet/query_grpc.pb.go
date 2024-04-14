// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: codenet/codenet/query.proto

package codenet

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
	Query_Params_FullMethodName                      = "/codenet.codenet.Query/Params"
	Query_GetEncodedDataById_FullMethodName          = "/codenet.codenet.Query/GetEncodedDataById"
	Query_GetProofById_FullMethodName                = "/codenet.codenet.Query/GetProofById"
	Query_GetCreatorById_FullMethodName              = "/codenet.codenet.Query/GetCreatorById"
	Query_GetVerificationStatusById_FullMethodName   = "/codenet.codenet.Query/GetVerificationStatusById"
	Query_GetEncodedDataByCreator_FullMethodName     = "/codenet.codenet.Query/GetEncodedDataByCreator"
	Query_GetEncodedDataByTimestamp_FullMethodName   = "/codenet.codenet.Query/GetEncodedDataByTimestamp"
	Query_GetEncodedDataByChecksum_FullMethodName    = "/codenet.codenet.Query/GetEncodedDataByChecksum"
	Query_GetEncodedDataByBlockNumber_FullMethodName = "/codenet.codenet.Query/GetEncodedDataByBlockNumber"
	Query_GetEncodedDataCount_FullMethodName         = "/codenet.codenet.Query/GetEncodedDataCount"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of GetEncodedDataById items.
	GetEncodedDataById(ctx context.Context, in *QueryGetEncodedDataByIdRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByIdResponse, error)
	// Queries a list of GetProofById items.
	GetProofById(ctx context.Context, in *QueryGetProofByIdRequest, opts ...grpc.CallOption) (*QueryGetProofByIdResponse, error)
	// Queries a list of GetCreatorById items.
	GetCreatorById(ctx context.Context, in *QueryGetCreatorByIdRequest, opts ...grpc.CallOption) (*QueryGetCreatorByIdResponse, error)
	// Queries a list of GetVerificationStatusById items.
	GetVerificationStatusById(ctx context.Context, in *QueryGetVerificationStatusByIdRequest, opts ...grpc.CallOption) (*QueryGetVerificationStatusByIdResponse, error)
	// Queries a list of GetEncodedDataByCreator items.
	GetEncodedDataByCreator(ctx context.Context, in *QueryGetEncodedDataByCreatorRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByCreatorResponse, error)
	// Queries a list of GetEncodedDataByTimestamp items.
	GetEncodedDataByTimestamp(ctx context.Context, in *QueryGetEncodedDataByTimestampRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByTimestampResponse, error)
	// Queries a list of GetEncodedDataByChecksum items.
	GetEncodedDataByChecksum(ctx context.Context, in *QueryGetEncodedDataByChecksumRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByChecksumResponse, error)
	// Queries a list of GetEncodedDataByBlockNumber items.
	GetEncodedDataByBlockNumber(ctx context.Context, in *QueryGetEncodedDataByBlockNumberRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByBlockNumberResponse, error)
	// Queries a list of GetEncodedDataCount items.
	GetEncodedDataCount(ctx context.Context, in *QueryGetEncodedDataCountRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataCountResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetEncodedDataById(ctx context.Context, in *QueryGetEncodedDataByIdRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByIdResponse, error) {
	out := new(QueryGetEncodedDataByIdResponse)
	err := c.cc.Invoke(ctx, Query_GetEncodedDataById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetProofById(ctx context.Context, in *QueryGetProofByIdRequest, opts ...grpc.CallOption) (*QueryGetProofByIdResponse, error) {
	out := new(QueryGetProofByIdResponse)
	err := c.cc.Invoke(ctx, Query_GetProofById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetCreatorById(ctx context.Context, in *QueryGetCreatorByIdRequest, opts ...grpc.CallOption) (*QueryGetCreatorByIdResponse, error) {
	out := new(QueryGetCreatorByIdResponse)
	err := c.cc.Invoke(ctx, Query_GetCreatorById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetVerificationStatusById(ctx context.Context, in *QueryGetVerificationStatusByIdRequest, opts ...grpc.CallOption) (*QueryGetVerificationStatusByIdResponse, error) {
	out := new(QueryGetVerificationStatusByIdResponse)
	err := c.cc.Invoke(ctx, Query_GetVerificationStatusById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetEncodedDataByCreator(ctx context.Context, in *QueryGetEncodedDataByCreatorRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByCreatorResponse, error) {
	out := new(QueryGetEncodedDataByCreatorResponse)
	err := c.cc.Invoke(ctx, Query_GetEncodedDataByCreator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetEncodedDataByTimestamp(ctx context.Context, in *QueryGetEncodedDataByTimestampRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByTimestampResponse, error) {
	out := new(QueryGetEncodedDataByTimestampResponse)
	err := c.cc.Invoke(ctx, Query_GetEncodedDataByTimestamp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetEncodedDataByChecksum(ctx context.Context, in *QueryGetEncodedDataByChecksumRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByChecksumResponse, error) {
	out := new(QueryGetEncodedDataByChecksumResponse)
	err := c.cc.Invoke(ctx, Query_GetEncodedDataByChecksum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetEncodedDataByBlockNumber(ctx context.Context, in *QueryGetEncodedDataByBlockNumberRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataByBlockNumberResponse, error) {
	out := new(QueryGetEncodedDataByBlockNumberResponse)
	err := c.cc.Invoke(ctx, Query_GetEncodedDataByBlockNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetEncodedDataCount(ctx context.Context, in *QueryGetEncodedDataCountRequest, opts ...grpc.CallOption) (*QueryGetEncodedDataCountResponse, error) {
	out := new(QueryGetEncodedDataCountResponse)
	err := c.cc.Invoke(ctx, Query_GetEncodedDataCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of GetEncodedDataById items.
	GetEncodedDataById(context.Context, *QueryGetEncodedDataByIdRequest) (*QueryGetEncodedDataByIdResponse, error)
	// Queries a list of GetProofById items.
	GetProofById(context.Context, *QueryGetProofByIdRequest) (*QueryGetProofByIdResponse, error)
	// Queries a list of GetCreatorById items.
	GetCreatorById(context.Context, *QueryGetCreatorByIdRequest) (*QueryGetCreatorByIdResponse, error)
	// Queries a list of GetVerificationStatusById items.
	GetVerificationStatusById(context.Context, *QueryGetVerificationStatusByIdRequest) (*QueryGetVerificationStatusByIdResponse, error)
	// Queries a list of GetEncodedDataByCreator items.
	GetEncodedDataByCreator(context.Context, *QueryGetEncodedDataByCreatorRequest) (*QueryGetEncodedDataByCreatorResponse, error)
	// Queries a list of GetEncodedDataByTimestamp items.
	GetEncodedDataByTimestamp(context.Context, *QueryGetEncodedDataByTimestampRequest) (*QueryGetEncodedDataByTimestampResponse, error)
	// Queries a list of GetEncodedDataByChecksum items.
	GetEncodedDataByChecksum(context.Context, *QueryGetEncodedDataByChecksumRequest) (*QueryGetEncodedDataByChecksumResponse, error)
	// Queries a list of GetEncodedDataByBlockNumber items.
	GetEncodedDataByBlockNumber(context.Context, *QueryGetEncodedDataByBlockNumberRequest) (*QueryGetEncodedDataByBlockNumberResponse, error)
	// Queries a list of GetEncodedDataCount items.
	GetEncodedDataCount(context.Context, *QueryGetEncodedDataCountRequest) (*QueryGetEncodedDataCountResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetEncodedDataById(context.Context, *QueryGetEncodedDataByIdRequest) (*QueryGetEncodedDataByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncodedDataById not implemented")
}
func (UnimplementedQueryServer) GetProofById(context.Context, *QueryGetProofByIdRequest) (*QueryGetProofByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProofById not implemented")
}
func (UnimplementedQueryServer) GetCreatorById(context.Context, *QueryGetCreatorByIdRequest) (*QueryGetCreatorByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreatorById not implemented")
}
func (UnimplementedQueryServer) GetVerificationStatusById(context.Context, *QueryGetVerificationStatusByIdRequest) (*QueryGetVerificationStatusByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerificationStatusById not implemented")
}
func (UnimplementedQueryServer) GetEncodedDataByCreator(context.Context, *QueryGetEncodedDataByCreatorRequest) (*QueryGetEncodedDataByCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncodedDataByCreator not implemented")
}
func (UnimplementedQueryServer) GetEncodedDataByTimestamp(context.Context, *QueryGetEncodedDataByTimestampRequest) (*QueryGetEncodedDataByTimestampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncodedDataByTimestamp not implemented")
}
func (UnimplementedQueryServer) GetEncodedDataByChecksum(context.Context, *QueryGetEncodedDataByChecksumRequest) (*QueryGetEncodedDataByChecksumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncodedDataByChecksum not implemented")
}
func (UnimplementedQueryServer) GetEncodedDataByBlockNumber(context.Context, *QueryGetEncodedDataByBlockNumberRequest) (*QueryGetEncodedDataByBlockNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncodedDataByBlockNumber not implemented")
}
func (UnimplementedQueryServer) GetEncodedDataCount(context.Context, *QueryGetEncodedDataCountRequest) (*QueryGetEncodedDataCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncodedDataCount not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetEncodedDataById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEncodedDataByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEncodedDataById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetEncodedDataById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEncodedDataById(ctx, req.(*QueryGetEncodedDataByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetProofById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProofByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetProofById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetProofById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetProofById(ctx, req.(*QueryGetProofByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetCreatorById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCreatorByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetCreatorById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetCreatorById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetCreatorById(ctx, req.(*QueryGetCreatorByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetVerificationStatusById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetVerificationStatusByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetVerificationStatusById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetVerificationStatusById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetVerificationStatusById(ctx, req.(*QueryGetVerificationStatusByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetEncodedDataByCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEncodedDataByCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEncodedDataByCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetEncodedDataByCreator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEncodedDataByCreator(ctx, req.(*QueryGetEncodedDataByCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetEncodedDataByTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEncodedDataByTimestampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEncodedDataByTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetEncodedDataByTimestamp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEncodedDataByTimestamp(ctx, req.(*QueryGetEncodedDataByTimestampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetEncodedDataByChecksum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEncodedDataByChecksumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEncodedDataByChecksum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetEncodedDataByChecksum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEncodedDataByChecksum(ctx, req.(*QueryGetEncodedDataByChecksumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetEncodedDataByBlockNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEncodedDataByBlockNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEncodedDataByBlockNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetEncodedDataByBlockNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEncodedDataByBlockNumber(ctx, req.(*QueryGetEncodedDataByBlockNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetEncodedDataCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEncodedDataCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEncodedDataCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetEncodedDataCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEncodedDataCount(ctx, req.(*QueryGetEncodedDataCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "codenet.codenet.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetEncodedDataById",
			Handler:    _Query_GetEncodedDataById_Handler,
		},
		{
			MethodName: "GetProofById",
			Handler:    _Query_GetProofById_Handler,
		},
		{
			MethodName: "GetCreatorById",
			Handler:    _Query_GetCreatorById_Handler,
		},
		{
			MethodName: "GetVerificationStatusById",
			Handler:    _Query_GetVerificationStatusById_Handler,
		},
		{
			MethodName: "GetEncodedDataByCreator",
			Handler:    _Query_GetEncodedDataByCreator_Handler,
		},
		{
			MethodName: "GetEncodedDataByTimestamp",
			Handler:    _Query_GetEncodedDataByTimestamp_Handler,
		},
		{
			MethodName: "GetEncodedDataByChecksum",
			Handler:    _Query_GetEncodedDataByChecksum_Handler,
		},
		{
			MethodName: "GetEncodedDataByBlockNumber",
			Handler:    _Query_GetEncodedDataByBlockNumber_Handler,
		},
		{
			MethodName: "GetEncodedDataCount",
			Handler:    _Query_GetEncodedDataCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "codenet/codenet/query.proto",
}
