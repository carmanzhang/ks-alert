// Code generated by protoc-gen-go. DO NOT EDIT.
// source: severity.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Severity
type SeveritySpec struct {
	SeverityId  string `protobuf:"bytes,1,opt,name=severity_id" json:"severity_id,omitempty"`
	ProductId   string `protobuf:"bytes,2,opt,name=product_id" json:"product_id,omitempty"`
	ProductName string `protobuf:"bytes,3,opt,name=product_name" json:"product_name,omitempty"`
}

func (m *SeveritySpec) Reset()                    { *m = SeveritySpec{} }
func (m *SeveritySpec) String() string            { return proto.CompactTextString(m) }
func (*SeveritySpec) ProtoMessage()               {}
func (*SeveritySpec) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *SeveritySpec) GetSeverityId() string {
	if m != nil {
		return m.SeverityId
	}
	return ""
}

func (m *SeveritySpec) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *SeveritySpec) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

type Severity struct {
	SeverityId string `protobuf:"bytes,1,opt,name=severity_id" json:"severity_id,omitempty"`
	SeverityCh string `protobuf:"bytes,2,opt,name=severity_ch" json:"severity_ch,omitempty"`
	SeverityEn string `protobuf:"bytes,3,opt,name=severity_en" json:"severity_en,omitempty"`
	ProductId  string `protobuf:"bytes,4,opt,name=product_id" json:"product_id,omitempty"`
}

func (m *Severity) Reset()                    { *m = Severity{} }
func (m *Severity) String() string            { return proto.CompactTextString(m) }
func (*Severity) ProtoMessage()               {}
func (*Severity) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *Severity) GetSeverityId() string {
	if m != nil {
		return m.SeverityId
	}
	return ""
}

func (m *Severity) GetSeverityCh() string {
	if m != nil {
		return m.SeverityCh
	}
	return ""
}

func (m *Severity) GetSeverityEn() string {
	if m != nil {
		return m.SeverityEn
	}
	return ""
}

func (m *Severity) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

type SeverityResponse struct {
	Severity *Severity `protobuf:"bytes,1,opt,name=Severity" json:"Severity,omitempty"`
	Error    *Error    `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SeverityResponse) Reset()                    { *m = SeverityResponse{} }
func (m *SeverityResponse) String() string            { return proto.CompactTextString(m) }
func (*SeverityResponse) ProtoMessage()               {}
func (*SeverityResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *SeverityResponse) GetSeverity() *Severity {
	if m != nil {
		return m.Severity
	}
	return nil
}

func (m *SeverityResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type SeveritiesResponse struct {
	Severity []*Severity `protobuf:"bytes,1,rep,name=Severity" json:"Severity,omitempty"`
	Error    *Error      `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SeveritiesResponse) Reset()                    { *m = SeveritiesResponse{} }
func (m *SeveritiesResponse) String() string            { return proto.CompactTextString(m) }
func (*SeveritiesResponse) ProtoMessage()               {}
func (*SeveritiesResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *SeveritiesResponse) GetSeverity() []*Severity {
	if m != nil {
		return m.Severity
	}
	return nil
}

func (m *SeveritiesResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*SeveritySpec)(nil), "pb.SeveritySpec")
	proto.RegisterType((*Severity)(nil), "pb.Severity")
	proto.RegisterType((*SeverityResponse)(nil), "pb.SeverityResponse")
	proto.RegisterType((*SeveritiesResponse)(nil), "pb.SeveritiesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeverityHandler service

type SeverityHandlerClient interface {
	// Severity
	CreateSeverity(ctx context.Context, in *Severity, opts ...grpc.CallOption) (*SeverityResponse, error)
	DeleteSeverity(ctx context.Context, in *SeveritySpec, opts ...grpc.CallOption) (*SeverityResponse, error)
	UpdateSeverity(ctx context.Context, in *Severity, opts ...grpc.CallOption) (*SeverityResponse, error)
	GetSeverity(ctx context.Context, in *SeveritySpec, opts ...grpc.CallOption) (*SeveritiesResponse, error)
}

type severityHandlerClient struct {
	cc *grpc.ClientConn
}

func NewSeverityHandlerClient(cc *grpc.ClientConn) SeverityHandlerClient {
	return &severityHandlerClient{cc}
}

func (c *severityHandlerClient) CreateSeverity(ctx context.Context, in *Severity, opts ...grpc.CallOption) (*SeverityResponse, error) {
	out := new(SeverityResponse)
	err := grpc.Invoke(ctx, "/pb.SeverityHandler/CreateSeverity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *severityHandlerClient) DeleteSeverity(ctx context.Context, in *SeveritySpec, opts ...grpc.CallOption) (*SeverityResponse, error) {
	out := new(SeverityResponse)
	err := grpc.Invoke(ctx, "/pb.SeverityHandler/DeleteSeverity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *severityHandlerClient) UpdateSeverity(ctx context.Context, in *Severity, opts ...grpc.CallOption) (*SeverityResponse, error) {
	out := new(SeverityResponse)
	err := grpc.Invoke(ctx, "/pb.SeverityHandler/UpdateSeverity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *severityHandlerClient) GetSeverity(ctx context.Context, in *SeveritySpec, opts ...grpc.CallOption) (*SeveritiesResponse, error) {
	out := new(SeveritiesResponse)
	err := grpc.Invoke(ctx, "/pb.SeverityHandler/GetSeverity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SeverityHandler service

type SeverityHandlerServer interface {
	// Severity
	CreateSeverity(context.Context, *Severity) (*SeverityResponse, error)
	DeleteSeverity(context.Context, *SeveritySpec) (*SeverityResponse, error)
	UpdateSeverity(context.Context, *Severity) (*SeverityResponse, error)
	GetSeverity(context.Context, *SeveritySpec) (*SeveritiesResponse, error)
}

func RegisterSeverityHandlerServer(s *grpc.Server, srv SeverityHandlerServer) {
	s.RegisterService(&_SeverityHandler_serviceDesc, srv)
}

func _SeverityHandler_CreateSeverity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Severity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeverityHandlerServer).CreateSeverity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SeverityHandler/CreateSeverity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeverityHandlerServer).CreateSeverity(ctx, req.(*Severity))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeverityHandler_DeleteSeverity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeveritySpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeverityHandlerServer).DeleteSeverity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SeverityHandler/DeleteSeverity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeverityHandlerServer).DeleteSeverity(ctx, req.(*SeveritySpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeverityHandler_UpdateSeverity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Severity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeverityHandlerServer).UpdateSeverity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SeverityHandler/UpdateSeverity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeverityHandlerServer).UpdateSeverity(ctx, req.(*Severity))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeverityHandler_GetSeverity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeveritySpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeverityHandlerServer).GetSeverity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SeverityHandler/GetSeverity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeverityHandlerServer).GetSeverity(ctx, req.(*SeveritySpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeverityHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SeverityHandler",
	HandlerType: (*SeverityHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSeverity",
			Handler:    _SeverityHandler_CreateSeverity_Handler,
		},
		{
			MethodName: "DeleteSeverity",
			Handler:    _SeverityHandler_DeleteSeverity_Handler,
		},
		{
			MethodName: "UpdateSeverity",
			Handler:    _SeverityHandler_UpdateSeverity_Handler,
		},
		{
			MethodName: "GetSeverity",
			Handler:    _SeverityHandler_GetSeverity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "severity.proto",
}

func init() { proto.RegisterFile("severity.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4b, 0x4b, 0xc3, 0x40,
	0x14, 0x85, 0x9b, 0xb4, 0x8a, 0xbd, 0x13, 0x62, 0x19, 0x8b, 0x84, 0x2e, 0x44, 0x66, 0xe5, 0x2a,
	0x8b, 0x08, 0x05, 0x5d, 0xfa, 0x40, 0x17, 0xea, 0xc2, 0xe2, 0x46, 0x10, 0xc9, 0xe3, 0x82, 0x81,
	0x9a, 0x19, 0x6e, 0x46, 0xc1, 0x7f, 0xee, 0x52, 0x66, 0x9a, 0xa9, 0xa3, 0x18, 0xd1, 0x65, 0xbe,
	0x9c, 0xf3, 0x9d, 0xc9, 0x03, 0xe2, 0x16, 0x5f, 0x91, 0x6a, 0xfd, 0x96, 0x2a, 0x92, 0x5a, 0xf2,
	0x50, 0x15, 0x33, 0x86, 0x44, 0x92, 0x56, 0x40, 0x5c, 0x43, 0xb4, 0xe8, 0x22, 0x0b, 0x85, 0x25,
	0xdf, 0x01, 0xe6, 0x2a, 0x8f, 0x75, 0x95, 0x04, 0xfb, 0xc1, 0xc1, 0x98, 0x73, 0x00, 0x45, 0xb2,
	0x7a, 0x29, 0xb5, 0x61, 0xa1, 0x65, 0x53, 0x88, 0x1c, 0x6b, 0xf2, 0x67, 0x4c, 0x86, 0x86, 0x8a,
	0x07, 0xd8, 0x72, 0xba, 0x9f, 0x55, 0x3e, 0x2c, 0x9f, 0x3a, 0x97, 0x0f, 0xb1, 0x59, 0xa9, 0xbe,
	0x8d, 0x8e, 0xac, 0xfe, 0x0a, 0x26, 0x4e, 0x7f, 0x8b, 0xad, 0x92, 0x4d, 0x8b, 0x7c, 0xef, 0x73,
	0xd2, 0x6e, 0xb0, 0x2c, 0x4a, 0x55, 0x91, 0xae, 0x8f, 0x91, 0xc0, 0x86, 0x7d, 0x60, 0xbb, 0xc5,
	0xb2, 0xb1, 0xb9, 0x79, 0x6e, 0x80, 0xb8, 0x01, 0xde, 0xa5, 0x6a, 0x6c, 0x7b, 0x7c, 0xc3, 0xbf,
	0xfb, 0xb2, 0xf7, 0x00, 0xb6, 0x5d, 0xec, 0x32, 0x6f, 0xaa, 0x25, 0x12, 0x9f, 0x43, 0x7c, 0x4a,
	0x98, 0x6b, 0x5c, 0xf7, 0xbf, 0xd8, 0x66, 0x53, 0xff, 0xca, 0x9d, 0x41, 0x0c, 0xf8, 0x31, 0xc4,
	0x67, 0xb8, 0x44, 0xaf, 0x37, 0xf1, 0x93, 0xe6, 0x5b, 0xf5, 0x76, 0xe7, 0x10, 0xdf, 0xa9, 0xea,
	0xff, 0x9b, 0x47, 0xc0, 0x2e, 0x50, 0xff, 0x32, 0xb8, 0xeb, 0x11, 0xef, 0x95, 0x89, 0xc1, 0xc9,
	0xe8, 0x3e, 0x54, 0x45, 0xb1, 0x69, 0xff, 0xa9, 0xc3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa,
	0xd7, 0xf9, 0xcf, 0x76, 0x02, 0x00, 0x00,
}
