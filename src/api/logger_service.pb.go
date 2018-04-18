// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logger_service.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	logger_service.proto

It has these top-level messages:
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import domain "github.com/yogyrahmawan/logger_service/src/domain"
import domain1 "github.com/yogyrahmawan/logger_service/src/domain"
import domain2 "github.com/yogyrahmawan/logger_service/src/domain"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
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

// Client API for LoggerService service

type LoggerServiceClient interface {
	SendLog(ctx context.Context, opts ...grpc.CallOption) (LoggerService_SendLogClient, error)
	GetLog(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*domain2.LoggerResponsesMessage, error)
}

type loggerServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoggerServiceClient(cc *grpc.ClientConn) LoggerServiceClient {
	return &loggerServiceClient{cc}
}

func (c *loggerServiceClient) SendLog(ctx context.Context, opts ...grpc.CallOption) (LoggerService_SendLogClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_LoggerService_serviceDesc.Streams[0], c.cc, "/api.LoggerService/SendLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &loggerServiceSendLogClient{stream}
	return x, nil
}

type LoggerService_SendLogClient interface {
	Send(*domain.LoggerMessage) error
	CloseAndRecv() (*domain1.LoggerResponse, error)
	grpc.ClientStream
}

type loggerServiceSendLogClient struct {
	grpc.ClientStream
}

func (x *loggerServiceSendLogClient) Send(m *domain.LoggerMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loggerServiceSendLogClient) CloseAndRecv() (*domain1.LoggerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(domain1.LoggerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loggerServiceClient) GetLog(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*domain2.LoggerResponsesMessage, error) {
	out := new(domain2.LoggerResponsesMessage)
	err := grpc.Invoke(ctx, "/api.LoggerService/GetLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoggerService service

type LoggerServiceServer interface {
	SendLog(LoggerService_SendLogServer) error
	GetLog(context.Context, *google_protobuf1.Empty) (*domain2.LoggerResponsesMessage, error)
}

func RegisterLoggerServiceServer(s *grpc.Server, srv LoggerServiceServer) {
	s.RegisterService(&_LoggerService_serviceDesc, srv)
}

func _LoggerService_SendLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoggerServiceServer).SendLog(&loggerServiceSendLogServer{stream})
}

type LoggerService_SendLogServer interface {
	SendAndClose(*domain1.LoggerResponse) error
	Recv() (*domain.LoggerMessage, error)
	grpc.ServerStream
}

type loggerServiceSendLogServer struct {
	grpc.ServerStream
}

func (x *loggerServiceSendLogServer) SendAndClose(m *domain1.LoggerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loggerServiceSendLogServer) Recv() (*domain.LoggerMessage, error) {
	m := new(domain.LoggerMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LoggerService_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServiceServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LoggerService/GetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServiceServer).GetLog(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.LoggerService",
	HandlerType: (*LoggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLog",
			Handler:    _LoggerService_GetLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendLog",
			Handler:       _LoggerService_SendLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "logger_service.proto",
}

func init() { proto.RegisterFile("logger_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xbb, 0x4e, 0xf3, 0x30,
	0x14, 0x56, 0xfa, 0x4b, 0xfd, 0xa5, 0xa8, 0x2c, 0x16, 0xed, 0x90, 0x02, 0x03, 0x0b, 0x54, 0x48,
	0x3e, 0x5c, 0x36, 0x46, 0x24, 0xc4, 0x52, 0x16, 0xba, 0x65, 0xa9, 0xdc, 0xf4, 0xe0, 0x5a, 0x4a,
	0x7c, 0x22, 0xdb, 0x0d, 0x44, 0x51, 0x16, 0x24, 0x36, 0x36, 0x66, 0x9e, 0x89, 0x81, 0x57, 0xe0,
	0x41, 0x50, 0x12, 0x23, 0x51, 0x84, 0xba, 0xd8, 0x3a, 0xdf, 0xcd, 0x9f, 0xed, 0x70, 0x37, 0x25,
	0x29, 0xd1, 0xcc, 0x2d, 0x9a, 0x42, 0x25, 0xc8, 0x73, 0x43, 0x8e, 0xd8, 0x3f, 0x91, 0xab, 0x68,
	0xd0, 0x51, 0x1d, 0x14, 0x0d, 0xbd, 0xd0, 0xa0, 0xcd, 0x49, 0x5b, 0xaf, 0x8c, 0x46, 0xbf, 0x60,
	0xeb, 0xf1, 0xb1, 0x24, 0x92, 0x29, 0x42, 0x3b, 0x2d, 0xd6, 0xf7, 0x80, 0x59, 0xee, 0x4a, 0x4f,
	0xee, 0x79, 0x52, 0xe4, 0x0a, 0x84, 0xd6, 0xe4, 0x84, 0x53, 0xa4, 0xbd, 0xf5, 0xfc, 0xad, 0x17,
	0xee, 0x4c, 0xdb, 0xd4, 0x59, 0x57, 0x8a, 0x5d, 0x86, 0xff, 0x67, 0xa8, 0x97, 0x53, 0x92, 0x6c,
	0xc8, 0x97, 0x94, 0x09, 0xa5, 0x79, 0xa7, 0xb8, 0x45, 0x6b, 0x85, 0xc4, 0x68, 0xb4, 0x09, 0xdf,
	0xf9, 0x36, 0xc7, 0x01, 0x7b, 0x0f, 0xc2, 0xfe, 0x0d, 0xba, 0xc6, 0x3b, 0xe2, 0xdd, 0xb9, 0xfc,
	0xbb, 0x14, 0xbf, 0x6e, 0x4a, 0x45, 0x07, 0x7f, 0x9b, 0xad, 0x0f, 0x3f, 0x7c, 0x09, 0x9e, 0x3e,
	0x3e, 0x5f, 0x7b, 0xcf, 0x01, 0x1b, 0xb4, 0x95, 0x8b, 0x33, 0x48, 0x49, 0xda, 0x78, 0x9f, 0x8d,
	0x7f, 0xce, 0x50, 0xf9, 0x47, 0x9c, 0x6b, 0x91, 0x61, 0x1d, 0x9f, 0xb0, 0xc9, 0x16, 0x1a, 0x52,
	0x2c, 0x30, 0x85, 0xaa, 0xdd, 0xea, 0x78, 0xc2, 0x8e, 0xb6, 0x89, 0x1d, 0x3e, 0x3a, 0xa8, 0x9a,
	0xb5, 0xbe, 0x3a, 0x8d, 0xb9, 0x54, 0x6e, 0xb5, 0x5e, 0xf0, 0x84, 0x32, 0x28, 0x49, 0x96, 0x46,
	0xac, 0x32, 0xf1, 0x20, 0x34, 0x6c, 0x7e, 0x26, 0x58, 0x93, 0x34, 0xa1, 0x8b, 0x7e, 0x7b, 0xe1,
	0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x66, 0x27, 0x6b, 0xed, 0x01, 0x00, 0x00,
}
