// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/qscheduler/service/app/config/config.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Config is the configuration data served by luci-config for this app.
type Config struct {
	AccessGroup string `protobuf:"bytes,1,opt,name=access_group,json=accessGroup,proto3" json:"access_group,omitempty"` // Deprecated: Do not use.
	// QuotaScheduler contains QS-specific config information.
	QuotaScheduler *QuotaScheduler `protobuf:"bytes,2,opt,name=quota_scheduler,json=quotaScheduler,proto3" json:"quota_scheduler,omitempty"`
	// Auth describes which access levels are granted to which groups.
	Auth                 *Auth    `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6117cabac368d3, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Config) GetAccessGroup() string {
	if m != nil {
		return m.AccessGroup
	}
	return ""
}

func (m *Config) GetQuotaScheduler() *QuotaScheduler {
	if m != nil {
		return m.QuotaScheduler
	}
	return nil
}

func (m *Config) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type Auth struct {
	// AdminGroup is the luci-auth group controlling access to the administrative
	// endpoints of this server (the QSchedulerAdmin API).
	//
	// Members of this group also receive QSchedulerView access.
	AdminGroup string `protobuf:"bytes,1,opt,name=admin_group,json=adminGroup,proto3" json:"admin_group,omitempty"`
	// SwarmingGroup is the luci-auth group controlling access to the scheduler
	// endpoints of this server (the swarming.ExternalScheduler API).
	SwarmingGroup string `protobuf:"bytes,2,opt,name=swarming_group,json=swarmingGroup,proto3" json:"swarming_group,omitempty"`
	// ViewGroup is the luci-auth group controlloing access to the qscheduler view
	// endpoints of the server (QSchedulerView API).
	ViewGroup            string   `protobuf:"bytes,3,opt,name=view_group,json=viewGroup,proto3" json:"view_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6117cabac368d3, []int{1}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetAdminGroup() string {
	if m != nil {
		return m.AdminGroup
	}
	return ""
}

func (m *Auth) GetSwarmingGroup() string {
	if m != nil {
		return m.SwarmingGroup
	}
	return ""
}

func (m *Auth) GetViewGroup() string {
	if m != nil {
		return m.ViewGroup
	}
	return ""
}

// QuotaScheduler contains configuration information for the QuotaScheduler app.
type QuotaScheduler struct {
	// If specified (non-zero), this internal timeout is applied to all Notify and
	// Assign requests served by the app.
	HandlerTimeout *duration.Duration `protobuf:"bytes,1,opt,name=handler_timeout,json=handlerTimeout,proto3" json:"handler_timeout,omitempty"`
	// If specified, the amount of time that a batch waits to collect requests
	// before executing.
	//
	// If unspecified, defaults to 300ms.
	//
	// A higher value causes batches to be larger and hence more efficient in
	// in terms of datastore operations, but adds overhead.
	BatchConstructionWait *duration.Duration `protobuf:"bytes,2,opt,name=batch_construction_wait,json=batchConstructionWait,proto3" json:"batch_construction_wait,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *QuotaScheduler) Reset()         { *m = QuotaScheduler{} }
func (m *QuotaScheduler) String() string { return proto.CompactTextString(m) }
func (*QuotaScheduler) ProtoMessage()    {}
func (*QuotaScheduler) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad6117cabac368d3, []int{2}
}

func (m *QuotaScheduler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaScheduler.Unmarshal(m, b)
}
func (m *QuotaScheduler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaScheduler.Marshal(b, m, deterministic)
}
func (m *QuotaScheduler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaScheduler.Merge(m, src)
}
func (m *QuotaScheduler) XXX_Size() int {
	return xxx_messageInfo_QuotaScheduler.Size(m)
}
func (m *QuotaScheduler) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaScheduler.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaScheduler proto.InternalMessageInfo

func (m *QuotaScheduler) GetHandlerTimeout() *duration.Duration {
	if m != nil {
		return m.HandlerTimeout
	}
	return nil
}

func (m *QuotaScheduler) GetBatchConstructionWait() *duration.Duration {
	if m != nil {
		return m.BatchConstructionWait
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "qscheduler.config.Config")
	proto.RegisterType((*Auth)(nil), "qscheduler.config.Auth")
	proto.RegisterType((*QuotaScheduler)(nil), "qscheduler.config.QuotaScheduler")
}

func init() {
	proto.RegisterFile("infra/qscheduler/service/app/config/config.proto", fileDescriptor_ad6117cabac368d3)
}

var fileDescriptor_ad6117cabac368d3 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x71, 0x4b, 0xeb, 0x30,
	0x14, 0xc5, 0xe9, 0x36, 0xc6, 0xdb, 0xed, 0x7b, 0x1d, 0x2f, 0x20, 0xab, 0x82, 0x3a, 0x07, 0x83,
	0x81, 0x90, 0xca, 0xfc, 0x04, 0x6e, 0x82, 0xe0, 0x7f, 0xab, 0x82, 0xe0, 0x3f, 0x25, 0x4b, 0xb3,
	0x36, 0xb0, 0x26, 0x5d, 0x9a, 0x6c, 0x1f, 0x49, 0x3f, 0xa6, 0x34, 0xe9, 0x74, 0x43, 0xf1, 0xaf,
	0x90, 0x73, 0x7f, 0xe7, 0xde, 0xc3, 0x81, 0x1b, 0x2e, 0x56, 0x8a, 0x44, 0x9b, 0x8a, 0xe6, 0x2c,
	0x35, 0x6b, 0xa6, 0xa2, 0x8a, 0xa9, 0x2d, 0xa7, 0x2c, 0x22, 0x65, 0x19, 0x51, 0x29, 0x56, 0x3c,
	0x6b, 0x1e, 0x5c, 0x2a, 0xa9, 0x25, 0xfa, 0xff, 0xc5, 0x62, 0x37, 0x38, 0xbb, 0xc8, 0xa4, 0xcc,
	0xd6, 0x2c, 0xb2, 0xc0, 0xd2, 0xac, 0xa2, 0xd4, 0x28, 0xa2, 0xb9, 0x14, 0xce, 0x32, 0x7a, 0xf7,
	0xa0, 0x3b, 0xb7, 0x28, 0x1a, 0xc3, 0x5f, 0x42, 0x29, 0xab, 0xaa, 0x24, 0x53, 0xd2, 0x94, 0xa1,
	0x37, 0xf4, 0x26, 0xbd, 0x59, 0x2b, 0xf4, 0x62, 0xdf, 0xe9, 0x0f, 0xb5, 0x8c, 0x1e, 0xa1, 0xbf,
	0x31, 0x52, 0x93, 0xe4, 0xf3, 0x56, 0xd8, 0x1a, 0x7a, 0x13, 0x7f, 0x7a, 0x85, 0xbf, 0x9d, 0xc7,
	0x8b, 0x9a, 0x7c, 0xda, 0xab, 0x71, 0xb0, 0x39, 0xfa, 0xa3, 0x6b, 0xe8, 0x10, 0xa3, 0xf3, 0xb0,
	0x6d, 0x17, 0x0c, 0x7e, 0x58, 0x70, 0x67, 0x74, 0x1e, 0x5b, 0x68, 0x54, 0x40, 0xa7, 0xfe, 0xa1,
	0x4b, 0xf0, 0x49, 0x5a, 0x70, 0x71, 0x18, 0x33, 0x06, 0x2b, 0xb9, 0x84, 0x63, 0x08, 0xaa, 0x1d,
	0x51, 0x05, 0x17, 0x59, 0xc3, 0xb4, 0x2c, 0xf3, 0x6f, 0xaf, 0x3a, 0xec, 0x1c, 0x60, 0xcb, 0xd9,
	0xae, 0x41, 0xda, 0x16, 0xe9, 0xd5, 0x8a, 0x1d, 0x8f, 0xde, 0x3c, 0x08, 0x8e, 0xe3, 0xa3, 0x19,
	0xf4, 0x73, 0x22, 0xd2, 0x35, 0x53, 0x89, 0xe6, 0x05, 0x93, 0x46, 0xdb, 0xeb, 0xfe, 0xf4, 0x14,
	0xbb, 0x9a, 0xf1, 0xbe, 0x66, 0x7c, 0xdf, 0xd4, 0x1c, 0x07, 0x8d, 0xe3, 0xd9, 0x19, 0xd0, 0x02,
	0x06, 0x4b, 0xa2, 0x69, 0x9e, 0x50, 0x29, 0x2a, 0xad, 0x0c, 0xad, 0xa9, 0x64, 0x47, 0xb8, 0x6e,
	0x6a, 0xfc, 0x65, 0xd7, 0x89, 0x75, 0xce, 0x0f, 0x8c, 0x2f, 0x84, 0xeb, 0xd9, 0x9f, 0xd7, 0xae,
	0x6b, 0x6b, 0xd9, 0xb5, 0x9e, 0xdb, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x24, 0x58, 0x61,
	0x3b, 0x02, 0x00, 0x00,
}
