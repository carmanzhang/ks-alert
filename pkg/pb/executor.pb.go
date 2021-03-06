// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Informer_Signal int32

const (
	Informer_CREATE    Informer_Signal = 0
	Informer_TERMINATE Informer_Signal = 1
	Informer_STOP      Informer_Signal = 2
	Informer_RELOAD    Informer_Signal = 3
	Informer_OTHER     Informer_Signal = 4
)

var Informer_Signal_name = map[int32]string{
	0: "CREATE",
	1: "TERMINATE",
	2: "STOP",
	3: "RELOAD",
	4: "OTHER",
}

var Informer_Signal_value = map[string]int32{
	"CREATE":    0,
	"TERMINATE": 1,
	"STOP":      2,
	"RELOAD":    3,
	"OTHER":     4,
}

func (x Informer_Signal) String() string {
	return proto.EnumName(Informer_Signal_name, int32(x))
}

func (Informer_Signal) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0, 0}
}

type Informer struct {
	AlertConfigId        string          `protobuf:"bytes,1,opt,name=alert_config_id,json=alertConfigId,proto3" json:"alert_config_id,omitempty"`
	Signal               Informer_Signal `protobuf:"varint,2,opt,name=signal,proto3,enum=pb.Informer_Signal" json:"signal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Informer) Reset()         { *m = Informer{} }
func (m *Informer) String() string { return proto.CompactTextString(m) }
func (*Informer) ProtoMessage()    {}
func (*Informer) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0}
}

func (m *Informer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Informer.Unmarshal(m, b)
}
func (m *Informer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Informer.Marshal(b, m, deterministic)
}
func (m *Informer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Informer.Merge(m, src)
}
func (m *Informer) XXX_Size() int {
	return xxx_messageInfo_Informer.Size(m)
}
func (m *Informer) XXX_DiscardUnknown() {
	xxx_messageInfo_Informer.DiscardUnknown(m)
}

var xxx_messageInfo_Informer proto.InternalMessageInfo

func (m *Informer) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *Informer) GetSignal() Informer_Signal {
	if m != nil {
		return m.Signal
	}
	return Informer_CREATE
}

func init() {
	proto.RegisterEnum("pb.Informer_Signal", Informer_Signal_name, Informer_Signal_value)
	proto.RegisterType((*Informer)(nil), "pb.Informer")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f) }

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xad, 0x48, 0x4d,
	0x2e, 0x2d, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2,
	0x4e, 0x2d, 0x2a, 0x82, 0x09, 0x28, 0x2d, 0x65, 0xe4, 0xe2, 0xf0, 0xcc, 0x4b, 0xcb, 0x2f, 0xca,
	0x4d, 0x2d, 0x12, 0x52, 0xe3, 0xe2, 0x4f, 0xcc, 0x49, 0x2d, 0x2a, 0x89, 0x4f, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0x8f, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x05, 0x0b, 0x3b,
	0x83, 0x45, 0x3d, 0x53, 0x84, 0xb4, 0xb9, 0xd8, 0x8a, 0x33, 0xd3, 0xf3, 0x12, 0x73, 0x24, 0x98,
	0x14, 0x18, 0x35, 0xf8, 0x8c, 0x84, 0xf5, 0x0a, 0x92, 0xf4, 0x60, 0xa6, 0xe8, 0x05, 0x83, 0xa5,
	0x82, 0xa0, 0x4a, 0x94, 0x5c, 0xb8, 0xd8, 0x20, 0x22, 0x42, 0x5c, 0x5c, 0x6c, 0xce, 0x41, 0xae,
	0x8e, 0x21, 0xae, 0x02, 0x0c, 0x42, 0xbc, 0x5c, 0x9c, 0x21, 0xae, 0x41, 0xbe, 0x9e, 0x7e, 0x20,
	0x2e, 0xa3, 0x10, 0x07, 0x17, 0x4b, 0x70, 0x88, 0x7f, 0x80, 0x00, 0x13, 0x48, 0x51, 0x90, 0xab,
	0x8f, 0xbf, 0xa3, 0x8b, 0x00, 0xb3, 0x10, 0x27, 0x17, 0xab, 0x7f, 0x88, 0x87, 0x6b, 0x90, 0x00,
	0x8b, 0x91, 0x01, 0x17, 0x87, 0x2b, 0xd4, 0x2b, 0x42, 0x2a, 0x5c, 0xec, 0x10, 0x76, 0xaa, 0x10,
	0x0f, 0xb2, 0xcd, 0x52, 0x9c, 0x20, 0x9e, 0x2b, 0xc8, 0x77, 0x4a, 0x0c, 0x4e, 0x2c, 0x51, 0x4c,
	0x05, 0x49, 0x49, 0x6c, 0x60, 0x6f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x18, 0xc7, 0xbe,
	0xad, 0x09, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecutorClient interface {
	Execute(ctx context.Context, in *Informer, opts ...grpc.CallOption) (*Error, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Execute(ctx context.Context, in *Informer, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/pb.Executor/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
type ExecutorServer interface {
	Execute(context.Context, *Informer) (*Error, error)
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Informer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Executor/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).Execute(ctx, req.(*Informer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Executor_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executor.proto",
}
