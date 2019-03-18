// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_config.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EmailServiceConfig struct {
	Protocol             *wrappers.StringValue `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	EmailHost            *wrappers.StringValue `protobuf:"bytes,2,opt,name=email_host,json=emailHost,proto3" json:"email_host,omitempty"`
	Port                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	DisplaySender        *wrappers.StringValue `protobuf:"bytes,4,opt,name=display_sender,json=displaySender,proto3" json:"display_sender,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password             *wrappers.StringValue `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	SslEnable            *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=ssl_enable,json=sslEnable,proto3" json:"ssl_enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EmailServiceConfig) Reset()         { *m = EmailServiceConfig{} }
func (m *EmailServiceConfig) String() string { return proto.CompactTextString(m) }
func (*EmailServiceConfig) ProtoMessage()    {}
func (*EmailServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{0}
}

func (m *EmailServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailServiceConfig.Unmarshal(m, b)
}
func (m *EmailServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailServiceConfig.Marshal(b, m, deterministic)
}
func (m *EmailServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailServiceConfig.Merge(m, src)
}
func (m *EmailServiceConfig) XXX_Size() int {
	return xxx_messageInfo_EmailServiceConfig.Size(m)
}
func (m *EmailServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EmailServiceConfig proto.InternalMessageInfo

func (m *EmailServiceConfig) GetProtocol() *wrappers.StringValue {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *EmailServiceConfig) GetEmailHost() *wrappers.StringValue {
	if m != nil {
		return m.EmailHost
	}
	return nil
}

func (m *EmailServiceConfig) GetPort() *wrappers.StringValue {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *EmailServiceConfig) GetDisplaySender() *wrappers.StringValue {
	if m != nil {
		return m.DisplaySender
	}
	return nil
}

func (m *EmailServiceConfig) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *EmailServiceConfig) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *EmailServiceConfig) GetSslEnable() *wrappers.BoolValue {
	if m != nil {
		return m.SslEnable
	}
	return nil
}

type NotificationConfig struct {
	EmailServiceConfig   *EmailServiceConfig `protobuf:"bytes,1,opt,name=email_service_config,json=emailServiceConfig,proto3" json:"email_service_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NotificationConfig) Reset()         { *m = NotificationConfig{} }
func (m *NotificationConfig) String() string { return proto.CompactTextString(m) }
func (*NotificationConfig) ProtoMessage()    {}
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{1}
}

func (m *NotificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationConfig.Unmarshal(m, b)
}
func (m *NotificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationConfig.Marshal(b, m, deterministic)
}
func (m *NotificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationConfig.Merge(m, src)
}
func (m *NotificationConfig) XXX_Size() int {
	return xxx_messageInfo_NotificationConfig.Size(m)
}
func (m *NotificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationConfig proto.InternalMessageInfo

func (m *NotificationConfig) GetEmailServiceConfig() *EmailServiceConfig {
	if m != nil {
		return m.EmailServiceConfig
	}
	return nil
}

type RuntimeItemConfig struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enable               *wrappers.BoolValue   `protobuf:"bytes,2,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RuntimeItemConfig) Reset()         { *m = RuntimeItemConfig{} }
func (m *RuntimeItemConfig) String() string { return proto.CompactTextString(m) }
func (*RuntimeItemConfig) ProtoMessage()    {}
func (*RuntimeItemConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{2}
}

func (m *RuntimeItemConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeItemConfig.Unmarshal(m, b)
}
func (m *RuntimeItemConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeItemConfig.Marshal(b, m, deterministic)
}
func (m *RuntimeItemConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeItemConfig.Merge(m, src)
}
func (m *RuntimeItemConfig) XXX_Size() int {
	return xxx_messageInfo_RuntimeItemConfig.Size(m)
}
func (m *RuntimeItemConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeItemConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeItemConfig proto.InternalMessageInfo

func (m *RuntimeItemConfig) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RuntimeItemConfig) GetEnable() *wrappers.BoolValue {
	if m != nil {
		return m.Enable
	}
	return nil
}

type RuntimeConfig struct {
	ConfigSet            []*RuntimeItemConfig `protobuf:"bytes,1,rep,name=config_set,json=configSet,proto3" json:"config_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RuntimeConfig) Reset()         { *m = RuntimeConfig{} }
func (m *RuntimeConfig) String() string { return proto.CompactTextString(m) }
func (*RuntimeConfig) ProtoMessage()    {}
func (*RuntimeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{3}
}

func (m *RuntimeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeConfig.Unmarshal(m, b)
}
func (m *RuntimeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeConfig.Marshal(b, m, deterministic)
}
func (m *RuntimeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeConfig.Merge(m, src)
}
func (m *RuntimeConfig) XXX_Size() int {
	return xxx_messageInfo_RuntimeConfig.Size(m)
}
func (m *RuntimeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeConfig proto.InternalMessageInfo

func (m *RuntimeConfig) GetConfigSet() []*RuntimeItemConfig {
	if m != nil {
		return m.ConfigSet
	}
	return nil
}

type BasicConfig struct {
	PlatformName         *wrappers.StringValue `protobuf:"bytes,1,opt,name=platform_name,json=platformName,proto3" json:"platform_name,omitempty"`
	PlatformUrl          *wrappers.StringValue `protobuf:"bytes,2,opt,name=platform_url,json=platformUrl,proto3" json:"platform_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BasicConfig) Reset()         { *m = BasicConfig{} }
func (m *BasicConfig) String() string { return proto.CompactTextString(m) }
func (*BasicConfig) ProtoMessage()    {}
func (*BasicConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{4}
}

func (m *BasicConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasicConfig.Unmarshal(m, b)
}
func (m *BasicConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasicConfig.Marshal(b, m, deterministic)
}
func (m *BasicConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicConfig.Merge(m, src)
}
func (m *BasicConfig) XXX_Size() int {
	return xxx_messageInfo_BasicConfig.Size(m)
}
func (m *BasicConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BasicConfig proto.InternalMessageInfo

func (m *BasicConfig) GetPlatformName() *wrappers.StringValue {
	if m != nil {
		return m.PlatformName
	}
	return nil
}

func (m *BasicConfig) GetPlatformUrl() *wrappers.StringValue {
	if m != nil {
		return m.PlatformUrl
	}
	return nil
}

type SetServiceConfigRequest struct {
	NotificationConfig   *NotificationConfig `protobuf:"bytes,1,opt,name=notification_config,json=notificationConfig,proto3" json:"notification_config,omitempty"`
	RuntimeConfig        *RuntimeConfig      `protobuf:"bytes,2,opt,name=runtime_config,json=runtimeConfig,proto3" json:"runtime_config,omitempty"`
	BasicConfig          *BasicConfig        `protobuf:"bytes,3,opt,name=basic_config,json=basicConfig,proto3" json:"basic_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SetServiceConfigRequest) Reset()         { *m = SetServiceConfigRequest{} }
func (m *SetServiceConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetServiceConfigRequest) ProtoMessage()    {}
func (*SetServiceConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{5}
}

func (m *SetServiceConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetServiceConfigRequest.Unmarshal(m, b)
}
func (m *SetServiceConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetServiceConfigRequest.Marshal(b, m, deterministic)
}
func (m *SetServiceConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetServiceConfigRequest.Merge(m, src)
}
func (m *SetServiceConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetServiceConfigRequest.Size(m)
}
func (m *SetServiceConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetServiceConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetServiceConfigRequest proto.InternalMessageInfo

func (m *SetServiceConfigRequest) GetNotificationConfig() *NotificationConfig {
	if m != nil {
		return m.NotificationConfig
	}
	return nil
}

func (m *SetServiceConfigRequest) GetRuntimeConfig() *RuntimeConfig {
	if m != nil {
		return m.RuntimeConfig
	}
	return nil
}

func (m *SetServiceConfigRequest) GetBasicConfig() *BasicConfig {
	if m != nil {
		return m.BasicConfig
	}
	return nil
}

type SetServiceConfigResponse struct {
	IsSucc               *wrappers.BoolValue `protobuf:"bytes,1,opt,name=is_succ,json=isSucc,proto3" json:"is_succ,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SetServiceConfigResponse) Reset()         { *m = SetServiceConfigResponse{} }
func (m *SetServiceConfigResponse) String() string { return proto.CompactTextString(m) }
func (*SetServiceConfigResponse) ProtoMessage()    {}
func (*SetServiceConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{6}
}

func (m *SetServiceConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetServiceConfigResponse.Unmarshal(m, b)
}
func (m *SetServiceConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetServiceConfigResponse.Marshal(b, m, deterministic)
}
func (m *SetServiceConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetServiceConfigResponse.Merge(m, src)
}
func (m *SetServiceConfigResponse) XXX_Size() int {
	return xxx_messageInfo_SetServiceConfigResponse.Size(m)
}
func (m *SetServiceConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetServiceConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetServiceConfigResponse proto.InternalMessageInfo

func (m *SetServiceConfigResponse) GetIsSucc() *wrappers.BoolValue {
	if m != nil {
		return m.IsSucc
	}
	return nil
}

type GetServiceConfigRequest struct {
	ServiceType          []string `protobuf:"bytes,1,rep,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceConfigRequest) Reset()         { *m = GetServiceConfigRequest{} }
func (m *GetServiceConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceConfigRequest) ProtoMessage()    {}
func (*GetServiceConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{7}
}

func (m *GetServiceConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceConfigRequest.Unmarshal(m, b)
}
func (m *GetServiceConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceConfigRequest.Marshal(b, m, deterministic)
}
func (m *GetServiceConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceConfigRequest.Merge(m, src)
}
func (m *GetServiceConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceConfigRequest.Size(m)
}
func (m *GetServiceConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceConfigRequest proto.InternalMessageInfo

func (m *GetServiceConfigRequest) GetServiceType() []string {
	if m != nil {
		return m.ServiceType
	}
	return nil
}

type GetServiceConfigResponse struct {
	NotificationConfig   *NotificationConfig `protobuf:"bytes,1,opt,name=notification_config,json=notificationConfig,proto3" json:"notification_config,omitempty"`
	RuntimeConfig        *RuntimeConfig      `protobuf:"bytes,2,opt,name=runtime_config,json=runtimeConfig,proto3" json:"runtime_config,omitempty"`
	BasicConfig          *BasicConfig        `protobuf:"bytes,3,opt,name=basic_config,json=basicConfig,proto3" json:"basic_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetServiceConfigResponse) Reset()         { *m = GetServiceConfigResponse{} }
func (m *GetServiceConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceConfigResponse) ProtoMessage()    {}
func (*GetServiceConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_452b382cbc0cfd24, []int{8}
}

func (m *GetServiceConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceConfigResponse.Unmarshal(m, b)
}
func (m *GetServiceConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceConfigResponse.Marshal(b, m, deterministic)
}
func (m *GetServiceConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceConfigResponse.Merge(m, src)
}
func (m *GetServiceConfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceConfigResponse.Size(m)
}
func (m *GetServiceConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceConfigResponse proto.InternalMessageInfo

func (m *GetServiceConfigResponse) GetNotificationConfig() *NotificationConfig {
	if m != nil {
		return m.NotificationConfig
	}
	return nil
}

func (m *GetServiceConfigResponse) GetRuntimeConfig() *RuntimeConfig {
	if m != nil {
		return m.RuntimeConfig
	}
	return nil
}

func (m *GetServiceConfigResponse) GetBasicConfig() *BasicConfig {
	if m != nil {
		return m.BasicConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*EmailServiceConfig)(nil), "openpitrix.EmailServiceConfig")
	proto.RegisterType((*NotificationConfig)(nil), "openpitrix.NotificationConfig")
	proto.RegisterType((*RuntimeItemConfig)(nil), "openpitrix.RuntimeItemConfig")
	proto.RegisterType((*RuntimeConfig)(nil), "openpitrix.RuntimeConfig")
	proto.RegisterType((*BasicConfig)(nil), "openpitrix.BasicConfig")
	proto.RegisterType((*SetServiceConfigRequest)(nil), "openpitrix.SetServiceConfigRequest")
	proto.RegisterType((*SetServiceConfigResponse)(nil), "openpitrix.SetServiceConfigResponse")
	proto.RegisterType((*GetServiceConfigRequest)(nil), "openpitrix.GetServiceConfigRequest")
	proto.RegisterType((*GetServiceConfigResponse)(nil), "openpitrix.GetServiceConfigResponse")
}

func init() { proto.RegisterFile("service_config.proto", fileDescriptor_452b382cbc0cfd24) }

var fileDescriptor_452b382cbc0cfd24 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0x41, 0x4f, 0x13, 0x4f,
	0x1c, 0xcd, 0x96, 0x02, 0x7f, 0x7e, 0xa5, 0xe4, 0xef, 0x48, 0xd2, 0xb5, 0x41, 0xac, 0xab, 0x07,
	0x62, 0xa4, 0xc5, 0x72, 0x51, 0x24, 0x41, 0x20, 0xa4, 0x7a, 0x10, 0x4c, 0xab, 0x1e, 0xbc, 0x6c,
	0xa6, 0xcb, 0x74, 0x9d, 0x64, 0x3b, 0x33, 0xce, 0x4c, 0xc1, 0x5e, 0xfd, 0x06, 0xea, 0xdd, 0x2f,
	0xe0, 0xa7, 0x31, 0x7e, 0x05, 0x0f, 0xfa, 0x2d, 0x4c, 0x67, 0x66, 0xa1, 0x65, 0xa1, 0x6e, 0x3c,
	0x7a, 0x6a, 0xa6, 0xfb, 0xde, 0x9b, 0xf7, 0x7b, 0xbf, 0xb7, 0x0b, 0xcb, 0x8a, 0xc8, 0x13, 0x1a,
	0x91, 0x30, 0xe2, 0xac, 0x47, 0xe3, 0xba, 0x90, 0x5c, 0x73, 0x04, 0x5c, 0x10, 0x26, 0xa8, 0x96,
	0xf4, 0x7d, 0x75, 0x25, 0xe6, 0x3c, 0x4e, 0x48, 0x03, 0x0b, 0xda, 0xc0, 0x8c, 0x71, 0x8d, 0x35,
	0xe5, 0x4c, 0x59, 0x64, 0x75, 0xd5, 0x3d, 0x35, 0xa7, 0xee, 0xa0, 0xd7, 0x38, 0x95, 0x58, 0x08,
	0x22, 0xd3, 0xe7, 0xf7, 0xcd, 0x4f, 0xb4, 0x1e, 0x13, 0xb6, 0xae, 0x4e, 0x71, 0x1c, 0x13, 0xd9,
	0xe0, 0xc2, 0x28, 0x5c, 0xa2, 0x76, 0xeb, 0xa2, 0x9a, 0xa6, 0x7d, 0xa2, 0x34, 0xee, 0x0b, 0x0b,
	0x08, 0xbe, 0xce, 0x00, 0x3a, 0xe8, 0x63, 0x9a, 0x74, 0xac, 0xed, 0x7d, 0xe3, 0x1a, 0x3d, 0x84,
	0xff, 0xec, 0x3d, 0x3c, 0xf1, 0xbd, 0x9a, 0xb7, 0x56, 0x6a, 0xae, 0xd4, 0xad, 0x54, 0x3d, 0x95,
	0xaa, 0x77, 0xb4, 0xa4, 0x2c, 0x7e, 0x8d, 0x93, 0x01, 0x69, 0x9f, 0xa1, 0xd1, 0x63, 0x00, 0x32,
	0xd2, 0x0b, 0xdf, 0x72, 0xa5, 0xfd, 0x42, 0x0e, 0xee, 0x82, 0xc1, 0x3f, 0xe5, 0x4a, 0xa3, 0x0d,
	0x28, 0x0a, 0x2e, 0xb5, 0x3f, 0x93, 0x83, 0x66, 0x90, 0x68, 0x1f, 0x96, 0x8e, 0xa9, 0x12, 0x09,
	0x1e, 0x86, 0x8a, 0xb0, 0x63, 0x22, 0xfd, 0x62, 0x0e, 0x6e, 0xd9, 0x71, 0x3a, 0x86, 0x82, 0x9a,
	0x30, 0x6b, 0x3c, 0xf8, 0xb3, 0x39, 0xb8, 0x16, 0x6a, 0x12, 0xc2, 0x4a, 0x9d, 0x72, 0x79, 0xec,
	0xcf, 0xe5, 0x4a, 0xc8, 0xa1, 0xd1, 0x23, 0x00, 0xa5, 0x92, 0x90, 0x30, 0xdc, 0x4d, 0x88, 0x3f,
	0x6f, 0xb8, 0xd5, 0x0c, 0x77, 0x8f, 0xf3, 0xc4, 0xe5, 0xa3, 0x54, 0x72, 0x60, 0xc0, 0x41, 0x0f,
	0xd0, 0x21, 0xd7, 0xb4, 0x47, 0x23, 0xb3, 0x65, 0xb7, 0xac, 0x17, 0xb0, 0x6c, 0x23, 0x9f, 0xac,
	0x9e, 0x5b, 0xdc, 0x6a, 0xfd, 0xbc, 0x7b, 0xf5, 0xec, 0xaa, 0xdb, 0x88, 0x64, 0xfe, 0x0b, 0x86,
	0x70, 0xad, 0x3d, 0x60, 0xa3, 0xae, 0x3c, 0xd3, 0xa4, 0xef, 0xae, 0xd9, 0x80, 0x22, 0xc3, 0x7d,
	0x92, 0xab, 0x0f, 0x06, 0x89, 0x9a, 0x30, 0xe7, 0xa6, 0x2c, 0xfc, 0x71, 0x4a, 0x87, 0x0c, 0x9e,
	0x43, 0xd9, 0x5d, 0xed, 0xae, 0xdd, 0x06, 0xb0, 0xf3, 0x84, 0x8a, 0x68, 0xdf, 0xab, 0xcd, 0xac,
	0x95, 0x9a, 0x37, 0xc7, 0x67, 0xca, 0x38, 0x6d, 0x2f, 0x58, 0x42, 0x87, 0xe8, 0xe0, 0xa3, 0x07,
	0xa5, 0x3d, 0xac, 0x68, 0xe4, 0xd4, 0x76, 0xa1, 0x2c, 0x12, 0xac, 0x7b, 0x5c, 0xf6, 0xc3, 0xdc,
	0xd3, 0x2c, 0xa6, 0x94, 0xc3, 0xd1, 0x54, 0x3b, 0x70, 0x76, 0x0e, 0x07, 0x32, 0xc9, 0xd5, 0xf1,
	0x52, 0xca, 0x78, 0x25, 0x93, 0xe0, 0xa7, 0x07, 0x95, 0x0e, 0xd1, 0x93, 0x6b, 0x20, 0xef, 0x06,
	0x44, 0x69, 0x74, 0x04, 0xd7, 0xd9, 0xd8, 0x86, 0xa7, 0xac, 0x32, 0x5b, 0x84, 0x36, 0x62, 0xd9,
	0x72, 0x3c, 0x81, 0x25, 0x69, 0x03, 0x4a, 0xb5, 0xac, 0xdf, 0x1b, 0x97, 0x44, 0xe8, 0x64, 0xca,
	0x72, 0x62, 0x01, 0x5b, 0xb0, 0xd8, 0x1d, 0x25, 0x98, 0xf2, 0xed, 0xcb, 0x59, 0x19, 0xe7, 0x8f,
	0x25, 0xdc, 0x2e, 0x75, 0xcf, 0x0f, 0xc1, 0x11, 0xf8, 0xd9, 0x49, 0x95, 0xe0, 0x4c, 0x11, 0xb4,
	0x09, 0xf3, 0x54, 0x85, 0x6a, 0x10, 0x45, 0x6e, 0xbc, 0xa9, 0xf5, 0xa0, 0xaa, 0x33, 0x88, 0xa2,
	0x60, 0x1b, 0x2a, 0xad, 0x2b, 0xa2, 0xbb, 0x0d, 0x8b, 0xe9, 0x0b, 0xa0, 0x87, 0x82, 0x98, 0xaa,
	0x2c, 0xb4, 0x4b, 0xee, 0xbf, 0x97, 0x43, 0x41, 0x82, 0x5f, 0x1e, 0xf8, 0xad, 0xab, 0xfc, 0xfc,
	0x5b, 0xd1, 0x37, 0xbf, 0x15, 0xa0, 0x3c, 0xf9, 0x51, 0xff, 0xe2, 0xc1, 0xff, 0x17, 0xb7, 0x81,
	0xee, 0x8c, 0x8b, 0x5d, 0xd1, 0xca, 0xea, 0xdd, 0xe9, 0x20, 0x1b, 0x60, 0xb0, 0xf3, 0x69, 0xb7,
	0x8a, 0x7c, 0x45, 0x74, 0xcd, 0x65, 0x5e, 0xb3, 0xb6, 0xa5, 0x09, 0xe5, 0xc3, 0xf7, 0x1f, 0x9f,
	0x0b, 0x2b, 0x41, 0xa5, 0x71, 0xf2, 0xa0, 0x31, 0xf9, 0xa1, 0x52, 0x0d, 0x45, 0xf4, 0x96, 0x77,
	0xcf, 0x18, 0x6c, 0x4d, 0x35, 0xd8, 0xca, 0x63, 0xb0, 0x35, 0xdd, 0x60, 0xfc, 0x17, 0x06, 0x63,
	0x63, 0x70, 0xaf, 0xf8, 0xa6, 0x20, 0xba, 0xdd, 0x39, 0xd3, 0xcf, 0xcd, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0xd6, 0xaf, 0x9d, 0xeb, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceConfigClient is the client API for ServiceConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceConfigClient interface {
	SetServiceConfig(ctx context.Context, in *SetServiceConfigRequest, opts ...grpc.CallOption) (*SetServiceConfigResponse, error)
	GetServiceConfig(ctx context.Context, in *GetServiceConfigRequest, opts ...grpc.CallOption) (*GetServiceConfigResponse, error)
}

type serviceConfigClient struct {
	cc *grpc.ClientConn
}

func NewServiceConfigClient(cc *grpc.ClientConn) ServiceConfigClient {
	return &serviceConfigClient{cc}
}

func (c *serviceConfigClient) SetServiceConfig(ctx context.Context, in *SetServiceConfigRequest, opts ...grpc.CallOption) (*SetServiceConfigResponse, error) {
	out := new(SetServiceConfigResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.ServiceConfig/SetServiceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceConfigClient) GetServiceConfig(ctx context.Context, in *GetServiceConfigRequest, opts ...grpc.CallOption) (*GetServiceConfigResponse, error) {
	out := new(GetServiceConfigResponse)
	err := c.cc.Invoke(ctx, "/openpitrix.ServiceConfig/GetServiceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceConfigServer is the server API for ServiceConfig service.
type ServiceConfigServer interface {
	SetServiceConfig(context.Context, *SetServiceConfigRequest) (*SetServiceConfigResponse, error)
	GetServiceConfig(context.Context, *GetServiceConfigRequest) (*GetServiceConfigResponse, error)
}

func RegisterServiceConfigServer(s *grpc.Server, srv ServiceConfigServer) {
	s.RegisterService(&_ServiceConfig_serviceDesc, srv)
}

func _ServiceConfig_SetServiceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetServiceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceConfigServer).SetServiceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ServiceConfig/SetServiceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceConfigServer).SetServiceConfig(ctx, req.(*SetServiceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceConfig_GetServiceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceConfigServer).GetServiceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.ServiceConfig/GetServiceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceConfigServer).GetServiceConfig(ctx, req.(*GetServiceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.ServiceConfig",
	HandlerType: (*ServiceConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetServiceConfig",
			Handler:    _ServiceConfig_SetServiceConfig_Handler,
		},
		{
			MethodName: "GetServiceConfig",
			Handler:    _ServiceConfig_GetServiceConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_config.proto",
}
