// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Error_ErrorCode int32

const (
	Error_SUCCESS           Error_ErrorCode = 0
	Error_INVALID_PARAM     Error_ErrorCode = 1
	Error_ACCESS_DENIED     Error_ErrorCode = 2
	Error_THRESHOLD_REACHED Error_ErrorCode = 3
)

var Error_ErrorCode_name = map[int32]string{
	0: "SUCCESS",
	1: "INVALID_PARAM",
	2: "ACCESS_DENIED",
	3: "THRESHOLD_REACHED",
}
var Error_ErrorCode_value = map[string]int32{
	"SUCCESS":           0,
	"INVALID_PARAM":     1,
	"ACCESS_DENIED":     2,
	"THRESHOLD_REACHED": 3,
}

func (x Error_ErrorCode) String() string {
	return proto.EnumName(Error_ErrorCode_name, int32(x))
}
func (Error_ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

// Error
type Error struct {
	Code Error_ErrorCode `protobuf:"varint,1,opt,name=code,enum=pb.Error_ErrorCode" json:"code,omitempty"`
	Text string          `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Error) GetCode() Error_ErrorCode {
	if m != nil {
		return m.Code
	}
	return Error_SUCCESS
}

func (m *Error) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "pb.Error")
	proto.RegisterEnum("pb.Error_ErrorCode", Error_ErrorCode_name, Error_ErrorCode_value)
}

func init() { proto.RegisterFile("error.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xea, 0x67, 0xe4, 0x62,
	0x75, 0x05, 0x89, 0x09, 0x29, 0x72, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x19, 0x09, 0xeb, 0x15, 0x24, 0xe9, 0x81, 0x25, 0x20, 0xa4, 0x73, 0x7e, 0x4a, 0xaa, 0x10,
	0x0f, 0x17, 0x4b, 0x49, 0x6a, 0x45, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xa7, 0x52, 0x28, 0x17,
	0x27, 0x42, 0x8a, 0x9b, 0x8b, 0x3d, 0x38, 0xd4, 0xd9, 0xd9, 0x35, 0x38, 0x58, 0x80, 0x41, 0x48,
	0x90, 0x8b, 0xd7, 0xd3, 0x2f, 0xcc, 0xd1, 0xc7, 0xd3, 0x25, 0x3e, 0xc0, 0x31, 0xc8, 0xd1, 0x57,
	0x80, 0x11, 0x24, 0xe4, 0x08, 0x96, 0x8e, 0x77, 0x71, 0xf5, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x12,
	0x12, 0xe5, 0x12, 0x0c, 0xf1, 0x08, 0x72, 0x0d, 0xf6, 0xf0, 0xf7, 0x71, 0x89, 0x0f, 0x72, 0x75,
	0x74, 0xf6, 0x70, 0x75, 0x11, 0x60, 0x76, 0x62, 0x89, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x3b,
	0xd1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x8c, 0xa1, 0x44, 0xb1, 0x00, 0x00, 0x00,
}