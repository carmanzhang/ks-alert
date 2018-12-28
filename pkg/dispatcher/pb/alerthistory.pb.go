// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alerthistory.proto

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

// alert history
type AlertHistory struct {
	AlertHistoryId            string       `protobuf:"bytes,1,opt,name=alert_history_id" json:"alert_history_id,omitempty"`
	AlertConfigId             string       `protobuf:"bytes,2,opt,name=alert_config_id" json:"alert_config_id,omitempty"`
	AlertName                 string       `protobuf:"bytes,3,opt,name=alert_name" json:"alert_name,omitempty"`
	SeverityId                string       `protobuf:"bytes,4,opt,name=severity_id" json:"severity_id,omitempty"`
	SeverityCh                string       `protobuf:"bytes,5,opt,name=severity_ch" json:"severity_ch,omitempty"`
	ResourceGroupId           string       `protobuf:"bytes,6,opt,name=resource_group_id" json:"resource_group_id,omitempty"`
	ResourceGroupName         string       `protobuf:"bytes,7,opt,name=resource_group_name" json:"resource_group_name,omitempty"`
	ResourceTypeId            string       `protobuf:"bytes,8,opt,name=resource_type_id" json:"resource_type_id,omitempty"`
	ResourceType              string       `protobuf:"bytes,9,opt,name=resource_type" json:"resource_type,omitempty"`
	AlertedResource           string       `protobuf:"bytes,10,opt,name=alerted_resource" json:"alerted_resource,omitempty"`
	ReceiverGroupId           string       `protobuf:"bytes,11,opt,name=receiver_group_id" json:"receiver_group_id,omitempty"`
	ReceiverGroupName         string       `protobuf:"bytes,12,opt,name=receiver_group_name" json:"receiver_group_name,omitempty"`
	Receivers                 []*Receiver  `protobuf:"bytes,13,rep,name=receivers" json:"receivers,omitempty"`
	AlertRuleGroupId          string       `protobuf:"bytes,14,opt,name=alert_rule_group_id" json:"alert_rule_group_id,omitempty"`
	TriggerAlertRule          string       `protobuf:"bytes,15,opt,name=trigger_alert_rule" json:"trigger_alert_rule,omitempty"`
	AlertRules                []*AlertRule `protobuf:"bytes,16,rep,name=alert_rules" json:"alert_rules,omitempty"`
	RepeatSend                *RepeatSend  `protobuf:"bytes,17,opt,name=repeat_send" json:"repeat_send,omitempty"`
	RequestNotificationStatus string       `protobuf:"bytes,18,opt,name=request_notification_status" json:"request_notification_status,omitempty"`
	EventTime                 string       `protobuf:"bytes,19,opt,name=event_time" json:"event_time,omitempty"`
}

func (m *AlertHistory) Reset()                    { *m = AlertHistory{} }
func (m *AlertHistory) String() string            { return proto.CompactTextString(m) }
func (*AlertHistory) ProtoMessage()               {}
func (*AlertHistory) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *AlertHistory) GetAlertHistoryId() string {
	if m != nil {
		return m.AlertHistoryId
	}
	return ""
}

func (m *AlertHistory) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *AlertHistory) GetAlertName() string {
	if m != nil {
		return m.AlertName
	}
	return ""
}

func (m *AlertHistory) GetSeverityId() string {
	if m != nil {
		return m.SeverityId
	}
	return ""
}

func (m *AlertHistory) GetSeverityCh() string {
	if m != nil {
		return m.SeverityCh
	}
	return ""
}

func (m *AlertHistory) GetResourceGroupId() string {
	if m != nil {
		return m.ResourceGroupId
	}
	return ""
}

func (m *AlertHistory) GetResourceGroupName() string {
	if m != nil {
		return m.ResourceGroupName
	}
	return ""
}

func (m *AlertHistory) GetResourceTypeId() string {
	if m != nil {
		return m.ResourceTypeId
	}
	return ""
}

func (m *AlertHistory) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *AlertHistory) GetAlertedResource() string {
	if m != nil {
		return m.AlertedResource
	}
	return ""
}

func (m *AlertHistory) GetReceiverGroupId() string {
	if m != nil {
		return m.ReceiverGroupId
	}
	return ""
}

func (m *AlertHistory) GetReceiverGroupName() string {
	if m != nil {
		return m.ReceiverGroupName
	}
	return ""
}

func (m *AlertHistory) GetReceivers() []*Receiver {
	if m != nil {
		return m.Receivers
	}
	return nil
}

func (m *AlertHistory) GetAlertRuleGroupId() string {
	if m != nil {
		return m.AlertRuleGroupId
	}
	return ""
}

func (m *AlertHistory) GetTriggerAlertRule() string {
	if m != nil {
		return m.TriggerAlertRule
	}
	return ""
}

func (m *AlertHistory) GetAlertRules() []*AlertRule {
	if m != nil {
		return m.AlertRules
	}
	return nil
}

func (m *AlertHistory) GetRepeatSend() *RepeatSend {
	if m != nil {
		return m.RepeatSend
	}
	return nil
}

func (m *AlertHistory) GetRequestNotificationStatus() string {
	if m != nil {
		return m.RequestNotificationStatus
	}
	return ""
}

func (m *AlertHistory) GetEventTime() string {
	if m != nil {
		return m.EventTime
	}
	return ""
}

type AlertHistoryRequest struct {
	AlertHistoryId string `protobuf:"bytes,1,opt,name=alert_history_id" json:"alert_history_id,omitempty"`
	AlertConfigId  string `protobuf:"bytes,2,opt,name=alert_config_id" json:"alert_config_id,omitempty"`
	AlertRuleId    string `protobuf:"bytes,3,opt,name=alert_rule_id" json:"alert_rule_id,omitempty"`
	ResourceId     string `protobuf:"bytes,4,opt,name=resource_id" json:"resource_id,omitempty"`
	Product        string `protobuf:"bytes,5,opt,name=product" json:"product,omitempty"`
	Page           int32  `protobuf:"varint,6,opt,name=page" json:"page,omitempty"`
	Limit          int32  `protobuf:"varint,7,opt,name=limit" json:"limit,omitempty"`
	Field          string `protobuf:"bytes,8,opt,name=field" json:"field,omitempty"`
	Fuzz           string `protobuf:"bytes,9,opt,name=fuzz" json:"fuzz,omitempty"`
	StartTime      int64  `protobuf:"varint,10,opt,name=start_time" json:"start_time,omitempty"`
	EndTime        int64  `protobuf:"varint,11,opt,name=end_time" json:"end_time,omitempty"`
}

func (m *AlertHistoryRequest) Reset()                    { *m = AlertHistoryRequest{} }
func (m *AlertHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertHistoryRequest) ProtoMessage()               {}
func (*AlertHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AlertHistoryRequest) GetAlertHistoryId() string {
	if m != nil {
		return m.AlertHistoryId
	}
	return ""
}

func (m *AlertHistoryRequest) GetAlertConfigId() string {
	if m != nil {
		return m.AlertConfigId
	}
	return ""
}

func (m *AlertHistoryRequest) GetAlertRuleId() string {
	if m != nil {
		return m.AlertRuleId
	}
	return ""
}

func (m *AlertHistoryRequest) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *AlertHistoryRequest) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *AlertHistoryRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *AlertHistoryRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AlertHistoryRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *AlertHistoryRequest) GetFuzz() string {
	if m != nil {
		return m.Fuzz
	}
	return ""
}

func (m *AlertHistoryRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AlertHistoryRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type AlertHistoryResponse struct {
	AlertHistory []*AlertHistory `protobuf:"bytes,1,rep,name=alert_history" json:"alert_history,omitempty"`
	Error        []*Error        `protobuf:"bytes,2,rep,name=error" json:"error,omitempty"`
}

func (m *AlertHistoryResponse) Reset()                    { *m = AlertHistoryResponse{} }
func (m *AlertHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertHistoryResponse) ProtoMessage()               {}
func (*AlertHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *AlertHistoryResponse) GetAlertHistory() []*AlertHistory {
	if m != nil {
		return m.AlertHistory
	}
	return nil
}

func (m *AlertHistoryResponse) GetError() []*Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*AlertHistory)(nil), "pb.AlertHistory")
	proto.RegisterType((*AlertHistoryRequest)(nil), "pb.AlertHistoryRequest")
	proto.RegisterType((*AlertHistoryResponse)(nil), "pb.AlertHistoryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AlertHistoryHandler service

type AlertHistoryHandlerClient interface {
	// history
	GetAlertHistory(ctx context.Context, in *AlertHistoryRequest, opts ...grpc.CallOption) (*AlertHistoryResponse, error)
}

type alertHistoryHandlerClient struct {
	cc *grpc.ClientConn
}

func NewAlertHistoryHandlerClient(cc *grpc.ClientConn) AlertHistoryHandlerClient {
	return &alertHistoryHandlerClient{cc}
}

func (c *alertHistoryHandlerClient) GetAlertHistory(ctx context.Context, in *AlertHistoryRequest, opts ...grpc.CallOption) (*AlertHistoryResponse, error) {
	out := new(AlertHistoryResponse)
	err := grpc.Invoke(ctx, "/pb.AlertHistoryHandler/GetAlertHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AlertHistoryHandler service

type AlertHistoryHandlerServer interface {
	// history
	GetAlertHistory(context.Context, *AlertHistoryRequest) (*AlertHistoryResponse, error)
}

func RegisterAlertHistoryHandlerServer(s *grpc.Server, srv AlertHistoryHandlerServer) {
	s.RegisterService(&_AlertHistoryHandler_serviceDesc, srv)
}

func _AlertHistoryHandler_GetAlertHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertHistoryHandlerServer).GetAlertHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlertHistoryHandler/GetAlertHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertHistoryHandlerServer).GetAlertHistory(ctx, req.(*AlertHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlertHistoryHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AlertHistoryHandler",
	HandlerType: (*AlertHistoryHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlertHistory",
			Handler:    _AlertHistoryHandler_GetAlertHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alerthistory.proto",
}

func init() { proto.RegisterFile("alerthistory.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x6e, 0xdb, 0x30,
	0x10, 0xac, 0x5f, 0x49, 0xbc, 0xb2, 0x63, 0x85, 0x6e, 0x10, 0xd6, 0x39, 0xd4, 0x70, 0x0e, 0xf5,
	0xc9, 0x07, 0xf7, 0x0b, 0x5a, 0xa0, 0x6d, 0xce, 0xee, 0xa9, 0x05, 0x0a, 0x41, 0x96, 0xd6, 0x0e,
	0x01, 0x99, 0x64, 0x49, 0xca, 0x40, 0xf2, 0x21, 0xfd, 0xc3, 0xfe, 0x47, 0xc1, 0xa5, 0x64, 0x45,
	0xc9, 0xad, 0x47, 0xce, 0xec, 0x0e, 0x97, 0x3b, 0x43, 0x60, 0x69, 0x81, 0xc6, 0x3d, 0x08, 0xeb,
	0x94, 0x79, 0x5c, 0x69, 0xa3, 0x9c, 0x62, 0x5d, 0xbd, 0x9d, 0x45, 0x68, 0x8c, 0x32, 0x01, 0x98,
	0x5d, 0x1a, 0xcc, 0x50, 0x1c, 0xb1, 0x3e, 0x4f, 0xa8, 0xc9, 0x94, 0x05, 0x56, 0x40, 0x6c, 0x50,
	0x63, 0xea, 0x2c, 0xca, 0x3c, 0x20, 0x8b, 0x3f, 0x7d, 0x18, 0x7d, 0xf2, 0x55, 0xf7, 0x41, 0x9a,
	0x71, 0x88, 0xa9, 0x2b, 0xa9, 0xee, 0x4a, 0x44, 0xce, 0x3b, 0xf3, 0xce, 0x72, 0xc8, 0x6e, 0x20,
	0xe8, 0x25, 0x99, 0x92, 0x3b, 0xb1, 0xf7, 0x44, 0x97, 0x08, 0x06, 0x10, 0x08, 0x99, 0x1e, 0x90,
	0xf7, 0x08, 0x9b, 0x42, 0x64, 0xf1, 0x88, 0x46, 0x38, 0x52, 0xe8, 0xbf, 0x02, 0xb3, 0x07, 0x3e,
	0x20, 0xf0, 0x1d, 0x5c, 0x19, 0xb4, 0xaa, 0x34, 0x19, 0x26, 0x7b, 0xa3, 0x4a, 0xed, 0xeb, 0xcf,
	0x88, 0xba, 0x85, 0xe9, 0x0b, 0x8a, 0x6e, 0x38, 0x27, 0x92, 0x43, 0x7c, 0x22, 0xdd, 0xa3, 0x46,
	0xdf, 0x76, 0x41, 0xcc, 0x35, 0x8c, 0x5b, 0x0c, 0x1f, 0xd6, 0x0d, 0x34, 0x26, 0xe6, 0x49, 0x4d,
	0x73, 0x68, 0x46, 0x08, 0x9b, 0x6b, 0x46, 0x88, 0x9a, 0x11, 0x5a, 0x14, 0x8d, 0x30, 0x22, 0xf2,
	0x3d, 0x0c, 0x6b, 0xd2, 0xf2, 0xf1, 0xbc, 0xb7, 0x8c, 0xd6, 0xa3, 0x95, 0xde, 0xae, 0x36, 0x15,
	0xe8, 0xbb, 0xc3, 0x66, 0xbc, 0x07, 0x8d, 0xf4, 0x25, 0x75, 0xcf, 0x80, 0x39, 0x23, 0xf6, 0x7b,
	0x34, 0x49, 0x53, 0xc4, 0x27, 0xc4, 0x2d, 0x20, 0x6a, 0x30, 0xcb, 0x63, 0xd2, 0x1e, 0x7b, 0x6d,
	0x32, 0x6b, 0x53, 0x16, 0xc8, 0xee, 0x20, 0x0a, 0x76, 0x26, 0xde, 0x4f, 0x7e, 0x35, 0xef, 0x2c,
	0xa3, 0xf5, 0x65, 0xb8, 0xdf, 0xc3, 0xdf, 0x51, 0xe6, 0xec, 0x0e, 0x6e, 0x0d, 0xfe, 0x2e, 0xd1,
	0xba, 0x44, 0x2a, 0x27, 0x76, 0x22, 0x4b, 0x9d, 0x50, 0x32, 0xb1, 0x2e, 0x75, 0xa5, 0xe5, 0xac,
	0x36, 0x10, 0x8f, 0x28, 0x5d, 0xe2, 0xc4, 0x01, 0xf9, 0xd4, 0x63, 0x8b, 0xbf, 0x1d, 0x98, 0x3e,
	0x0f, 0xc6, 0x26, 0xa8, 0xfc, 0x4f, 0x3e, 0xae, 0x61, 0xfc, 0x6c, 0x0b, 0x22, 0x6f, 0x22, 0x72,
	0xb2, 0xe9, 0x14, 0x91, 0x09, 0x9c, 0x6b, 0xa3, 0xf2, 0x32, 0x73, 0x55, 0x3c, 0x46, 0xd0, 0xd7,
	0xe9, 0x1e, 0x29, 0x11, 0x03, 0x36, 0x86, 0x41, 0x21, 0x0e, 0xc2, 0x51, 0x06, 0xe8, 0xb8, 0x13,
	0x58, 0xd4, 0xc6, 0x8f, 0xa0, 0xbf, 0x2b, 0x9f, 0x9e, 0x2a, 0xbf, 0x19, 0x80, 0x75, 0xa9, 0xa9,
	0x5e, 0xe5, 0x9d, 0xee, 0xb1, 0x18, 0x2e, 0x50, 0xe6, 0x01, 0xf1, 0x06, 0xf7, 0x16, 0x3f, 0xe0,
	0x6d, 0xfb, 0x99, 0x56, 0x2b, 0x69, 0x91, 0x7d, 0xa8, 0x87, 0xae, 0xde, 0xc9, 0x3b, 0xe4, 0x41,
	0x7c, 0xf2, 0xa0, 0xf9, 0x30, 0x03, 0xfa, 0x83, 0xbc, 0x4b, 0x05, 0x43, 0x5f, 0xf0, 0xc5, 0x03,
	0xeb, 0x5f, 0xed, 0x0d, 0xde, 0xa7, 0x32, 0x2f, 0xd0, 0xb0, 0xaf, 0x30, 0xf9, 0x86, 0xae, 0xa5,
	0x71, 0xf3, 0x52, 0xb5, 0xda, 0xf6, 0x8c, 0xbf, 0x26, 0xc2, 0x7c, 0x8b, 0x37, 0x9f, 0xfb, 0x3f,
	0xbb, 0x7a, 0xbb, 0x3d, 0xa3, 0x7f, 0xfc, 0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x76,
	0x00, 0xb7, 0x21, 0x04, 0x00, 0x00,
}