// Code generated by protoc-gen-gogo.
// source: tikvpb.proto
// DO NOT EDIT!

/*
	Package tikvpb is a generated protocol buffer package.

	It is generated from these files:
		tikvpb.proto

	It has these top-level messages:
*/
package tikvpb

import (
	"fmt"
	"math"

	proto "github.com/golang/protobuf/proto"

	coprocessor "github.com/pingcap/kvproto/pkg/coprocessor"

	kvrpcpb "github.com/pingcap/kvproto/pkg/kvrpcpb"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TiKV service

type TiKVClient interface {
	// KV commands with mvcc/txn supported.
	KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error)
	KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error)
	KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error)
	KvGC(ctx context.Context, in *kvrpcpb.GCRequest, opts ...grpc.CallOption) (*kvrpcpb.GCResponse, error)
	// RawKV commands.
	RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error)
	RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error)
	RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error)
	// SQL push down commands.
	Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error)
}

type tiKVClient struct {
	cc *grpc.ClientConn
}

func NewTiKVClient(cc *grpc.ClientConn) TiKVClient {
	return &tiKVClient{cc}
}

func (c *tiKVClient) KvGet(ctx context.Context, in *kvrpcpb.GetRequest, opts ...grpc.CallOption) (*kvrpcpb.GetResponse, error) {
	out := new(kvrpcpb.GetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvScan(ctx context.Context, in *kvrpcpb.ScanRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanResponse, error) {
	out := new(kvrpcpb.ScanResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvScan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvPrewrite(ctx context.Context, in *kvrpcpb.PrewriteRequest, opts ...grpc.CallOption) (*kvrpcpb.PrewriteResponse, error) {
	out := new(kvrpcpb.PrewriteResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvPrewrite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvCommit(ctx context.Context, in *kvrpcpb.CommitRequest, opts ...grpc.CallOption) (*kvrpcpb.CommitResponse, error) {
	out := new(kvrpcpb.CommitResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvCommit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvCleanup(ctx context.Context, in *kvrpcpb.CleanupRequest, opts ...grpc.CallOption) (*kvrpcpb.CleanupResponse, error) {
	out := new(kvrpcpb.CleanupResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvCleanup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvBatchGet(ctx context.Context, in *kvrpcpb.BatchGetRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchGetResponse, error) {
	out := new(kvrpcpb.BatchGetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvBatchGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvBatchRollback(ctx context.Context, in *kvrpcpb.BatchRollbackRequest, opts ...grpc.CallOption) (*kvrpcpb.BatchRollbackResponse, error) {
	out := new(kvrpcpb.BatchRollbackResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvBatchRollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvScanLock(ctx context.Context, in *kvrpcpb.ScanLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ScanLockResponse, error) {
	out := new(kvrpcpb.ScanLockResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvScanLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvResolveLock(ctx context.Context, in *kvrpcpb.ResolveLockRequest, opts ...grpc.CallOption) (*kvrpcpb.ResolveLockResponse, error) {
	out := new(kvrpcpb.ResolveLockResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvResolveLock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) KvGC(ctx context.Context, in *kvrpcpb.GCRequest, opts ...grpc.CallOption) (*kvrpcpb.GCResponse, error) {
	out := new(kvrpcpb.GCResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/KvGC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) RawGet(ctx context.Context, in *kvrpcpb.RawGetRequest, opts ...grpc.CallOption) (*kvrpcpb.RawGetResponse, error) {
	out := new(kvrpcpb.RawGetResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/RawGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) RawPut(ctx context.Context, in *kvrpcpb.RawPutRequest, opts ...grpc.CallOption) (*kvrpcpb.RawPutResponse, error) {
	out := new(kvrpcpb.RawPutResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/RawPut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) RawDelete(ctx context.Context, in *kvrpcpb.RawDeleteRequest, opts ...grpc.CallOption) (*kvrpcpb.RawDeleteResponse, error) {
	out := new(kvrpcpb.RawDeleteResponse)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/RawDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiKVClient) Coprocessor(ctx context.Context, in *coprocessor.Request, opts ...grpc.CallOption) (*coprocessor.Response, error) {
	out := new(coprocessor.Response)
	err := grpc.Invoke(ctx, "/tikvpb.TiKV/Coprocessor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TiKV service

type TiKVServer interface {
	// KV commands with mvcc/txn supported.
	KvGet(context.Context, *kvrpcpb.GetRequest) (*kvrpcpb.GetResponse, error)
	KvScan(context.Context, *kvrpcpb.ScanRequest) (*kvrpcpb.ScanResponse, error)
	KvPrewrite(context.Context, *kvrpcpb.PrewriteRequest) (*kvrpcpb.PrewriteResponse, error)
	KvCommit(context.Context, *kvrpcpb.CommitRequest) (*kvrpcpb.CommitResponse, error)
	KvCleanup(context.Context, *kvrpcpb.CleanupRequest) (*kvrpcpb.CleanupResponse, error)
	KvBatchGet(context.Context, *kvrpcpb.BatchGetRequest) (*kvrpcpb.BatchGetResponse, error)
	KvBatchRollback(context.Context, *kvrpcpb.BatchRollbackRequest) (*kvrpcpb.BatchRollbackResponse, error)
	KvScanLock(context.Context, *kvrpcpb.ScanLockRequest) (*kvrpcpb.ScanLockResponse, error)
	KvResolveLock(context.Context, *kvrpcpb.ResolveLockRequest) (*kvrpcpb.ResolveLockResponse, error)
	KvGC(context.Context, *kvrpcpb.GCRequest) (*kvrpcpb.GCResponse, error)
	// RawKV commands.
	RawGet(context.Context, *kvrpcpb.RawGetRequest) (*kvrpcpb.RawGetResponse, error)
	RawPut(context.Context, *kvrpcpb.RawPutRequest) (*kvrpcpb.RawPutResponse, error)
	RawDelete(context.Context, *kvrpcpb.RawDeleteRequest) (*kvrpcpb.RawDeleteResponse, error)
	// SQL push down commands.
	Coprocessor(context.Context, *coprocessor.Request) (*coprocessor.Response, error)
}

func RegisterTiKVServer(s *grpc.Server, srv TiKVServer) {
	s.RegisterService(&_TiKV_serviceDesc, srv)
}

func _TiKV_KvGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvGet(ctx, req.(*kvrpcpb.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvScan(ctx, req.(*kvrpcpb.ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvPrewrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.PrewriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvPrewrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvPrewrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvPrewrite(ctx, req.(*kvrpcpb.PrewriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvCommit(ctx, req.(*kvrpcpb.CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvCleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.CleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvCleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvCleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvCleanup(ctx, req.(*kvrpcpb.CleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvBatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvBatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvBatchGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvBatchGet(ctx, req.(*kvrpcpb.BatchGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvBatchRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.BatchRollbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvBatchRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvBatchRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvBatchRollback(ctx, req.(*kvrpcpb.BatchRollbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvScanLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ScanLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvScanLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvScanLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvScanLock(ctx, req.(*kvrpcpb.ScanLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvResolveLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.ResolveLockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvResolveLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvResolveLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvResolveLock(ctx, req.(*kvrpcpb.ResolveLockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_KvGC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.GCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).KvGC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/KvGC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).KvGC(ctx, req.(*kvrpcpb.GCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_RawGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).RawGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/RawGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).RawGet(ctx, req.(*kvrpcpb.RawGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_RawPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).RawPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/RawPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).RawPut(ctx, req.(*kvrpcpb.RawPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_RawDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(kvrpcpb.RawDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).RawDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/RawDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).RawDelete(ctx, req.(*kvrpcpb.RawDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TiKV_Coprocessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(coprocessor.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiKVServer).Coprocessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tikvpb.TiKV/Coprocessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiKVServer).Coprocessor(ctx, req.(*coprocessor.Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _TiKV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tikvpb.TiKV",
	HandlerType: (*TiKVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KvGet",
			Handler:    _TiKV_KvGet_Handler,
		},
		{
			MethodName: "KvScan",
			Handler:    _TiKV_KvScan_Handler,
		},
		{
			MethodName: "KvPrewrite",
			Handler:    _TiKV_KvPrewrite_Handler,
		},
		{
			MethodName: "KvCommit",
			Handler:    _TiKV_KvCommit_Handler,
		},
		{
			MethodName: "KvCleanup",
			Handler:    _TiKV_KvCleanup_Handler,
		},
		{
			MethodName: "KvBatchGet",
			Handler:    _TiKV_KvBatchGet_Handler,
		},
		{
			MethodName: "KvBatchRollback",
			Handler:    _TiKV_KvBatchRollback_Handler,
		},
		{
			MethodName: "KvScanLock",
			Handler:    _TiKV_KvScanLock_Handler,
		},
		{
			MethodName: "KvResolveLock",
			Handler:    _TiKV_KvResolveLock_Handler,
		},
		{
			MethodName: "KvGC",
			Handler:    _TiKV_KvGC_Handler,
		},
		{
			MethodName: "RawGet",
			Handler:    _TiKV_RawGet_Handler,
		},
		{
			MethodName: "RawPut",
			Handler:    _TiKV_RawPut_Handler,
		},
		{
			MethodName: "RawDelete",
			Handler:    _TiKV_RawDelete_Handler,
		},
		{
			MethodName: "Coprocessor",
			Handler:    _TiKV_Coprocessor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tikvpb.proto",
}

func init() { proto.RegisterFile("tikvpb.proto", fileDescriptorTikvpb) }

var fileDescriptorTikvpb = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xcb, 0x4e, 0xea, 0x40,
	0x18, 0xc7, 0x21, 0xe1, 0xf4, 0x1c, 0xe6, 0x1c, 0x72, 0x74, 0x40, 0x85, 0xaa, 0x5d, 0xf8, 0x00,
	0x18, 0x2f, 0x89, 0x0b, 0x35, 0x31, 0x94, 0x84, 0xc5, 0xb0, 0x20, 0xd5, 0xb8, 0x2f, 0xcd, 0x04,
	0x9b, 0x16, 0xa6, 0xb6, 0xd3, 0xe1, 0x55, 0x78, 0x24, 0x97, 0x3e, 0x82, 0xc1, 0x17, 0x31, 0xed,
	0xf8, 0xcd, 0xb4, 0xb4, 0xee, 0x3a, 0xbf, 0xff, 0x25, 0x73, 0xe9, 0x87, 0xfe, 0x71, 0x3f, 0x10,
	0xd1, 0x7c, 0x18, 0xc5, 0x8c, 0x33, 0x6c, 0xc8, 0x95, 0xb9, 0xef, 0xb1, 0x28, 0x66, 0x1e, 0x4d,
	0x12, 0x16, 0x4b, 0xc9, 0xec, 0x04, 0x22, 0x8e, 0x3c, 0x70, 0x9a, 0xbd, 0x05, 0x5b, 0xb0, 0xfc,
	0xf3, 0x3c, 0xfb, 0x92, 0xf4, 0x72, 0xf3, 0x1b, 0xb5, 0x9e, 0x7c, 0xf2, 0x8c, 0xaf, 0xd1, 0x2f,
	0x22, 0x26, 0x94, 0xe3, 0xee, 0x10, 0x72, 0x13, 0xca, 0x1d, 0xfa, 0x9a, 0xd2, 0x84, 0x9b, 0xbd,
	0x32, 0x4c, 0x22, 0xb6, 0x4a, 0xe8, 0x59, 0x03, 0xdf, 0x20, 0x83, 0x88, 0x47, 0xcf, 0x5d, 0x61,
	0xed, 0xc8, 0x96, 0x90, 0x3b, 0xd8, 0xa1, 0x2a, 0x68, 0x23, 0x44, 0xc4, 0x2c, 0xa6, 0xeb, 0xd8,
	0xe7, 0x14, 0xf7, 0x95, 0x0d, 0x10, 0x14, 0x0c, 0x6a, 0x14, 0x55, 0x72, 0x8f, 0xfe, 0x10, 0x61,
	0xb3, 0xe5, 0xd2, 0xe7, 0xf8, 0x50, 0x19, 0x25, 0x80, 0x82, 0xa3, 0x0a, 0x57, 0xf1, 0x07, 0xd4,
	0x26, 0xc2, 0x0e, 0xa9, 0xbb, 0x4a, 0x23, 0x5c, 0xf0, 0x49, 0x02, 0x05, 0xfd, 0xaa, 0x50, 0x3e,
	0xc5, 0xc8, 0xe5, 0xde, 0x4b, 0x76, 0x73, 0xda, 0x09, 0xa8, 0x7a, 0x0a, 0xad, 0xa8, 0x12, 0x07,
	0xfd, 0xff, 0x2e, 0x71, 0x58, 0x18, 0xce, 0x5d, 0x2f, 0xc0, 0xa7, 0x65, 0x3f, 0x70, 0xa8, 0xb3,
	0x7e, 0x92, 0xcb, 0x1b, 0xcb, 0xae, 0x7c, 0xca, 0xbc, 0xa0, 0xb0, 0x31, 0x40, 0xd5, 0x8d, 0x69,
	0x45, 0x95, 0x4c, 0x51, 0x87, 0x08, 0x87, 0x26, 0x2c, 0x14, 0x34, 0xef, 0x39, 0x56, 0xee, 0x02,
	0x85, 0xaa, 0x93, 0x7a, 0x51, 0xb5, 0x5d, 0xa0, 0x16, 0x11, 0x13, 0x1b, 0x63, 0xfd, 0x2b, 0xd9,
	0x90, 0xed, 0x96, 0x98, 0x8a, 0xdc, 0x22, 0xc3, 0x71, 0xd7, 0xd9, 0xd5, 0xea, 0xd7, 0x95, 0xa0,
	0xfa, 0xba, 0xc0, 0x77, 0xc2, 0xb3, 0x74, 0x27, 0x3c, 0x4b, 0xeb, 0xc3, 0x39, 0x57, 0xe1, 0x31,
	0x6a, 0x3b, 0xee, 0x7a, 0x4c, 0x43, 0xca, 0x29, 0x1e, 0x14, 0x7d, 0x92, 0x41, 0x85, 0x59, 0x27,
	0xa9, 0x96, 0x3b, 0xf4, 0xd7, 0xd6, 0x63, 0x89, 0x7b, 0xc3, 0xe2, 0x90, 0xea, 0x11, 0x29, 0x53,
	0x48, 0x8f, 0xf6, 0xde, 0xb6, 0x56, 0xf3, 0x7d, 0x6b, 0x35, 0x3f, 0xb6, 0x56, 0x73, 0xf3, 0x69,
	0x35, 0xe6, 0x46, 0x3e, 0xb3, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xcc, 0xdb, 0x83,
	0x03, 0x04, 0x00, 0x00,
}