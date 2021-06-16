// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/cros/lab_inventory/api/bigquery/assetRegistration.proto

package apibq

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protos "infra/libs/fleet/protos"
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

type RegisteredAsset struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Asset                *protos.ChopsAsset   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	UpdatedTime          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegisteredAsset) Reset()         { *m = RegisteredAsset{} }
func (m *RegisteredAsset) String() string { return proto.CompactTextString(m) }
func (*RegisteredAsset) ProtoMessage()    {}
func (*RegisteredAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c53bb14c3788c7de, []int{0}
}

func (m *RegisteredAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisteredAsset.Unmarshal(m, b)
}
func (m *RegisteredAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisteredAsset.Marshal(b, m, deterministic)
}
func (m *RegisteredAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredAsset.Merge(m, src)
}
func (m *RegisteredAsset) XXX_Size() int {
	return xxx_messageInfo_RegisteredAsset.Size(m)
}
func (m *RegisteredAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredAsset.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredAsset proto.InternalMessageInfo

func (m *RegisteredAsset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisteredAsset) GetAsset() *protos.ChopsAsset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *RegisteredAsset) GetUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedTime
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisteredAsset)(nil), "apibq.RegisteredAsset")
}

func init() {
	proto.RegisterFile("infra/appengine/cros/lab_inventory/api/bigquery/assetRegistration.proto", fileDescriptor_c53bb14c3788c7de)
}

var fileDescriptor_c53bb14c3788c7de = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xd9, 0x95, 0x13, 0xcc, 0x89, 0xe2, 0x56, 0xcb, 0x36, 0x1e, 0x5a, 0x78, 0x55, 0x02,
	0x5a, 0x5b, 0x88, 0x85, 0xfd, 0x62, 0x7f, 0x24, 0x66, 0x36, 0x0e, 0xec, 0x25, 0xb9, 0x64, 0x56,
	0xf0, 0x11, 0x7c, 0x6b, 0x49, 0xe6, 0xb6, 0x9c, 0x99, 0x6f, 0xfe, 0xff, 0x13, 0x1f, 0xe8, 0xa7,
	0xa4, 0x95, 0x8e, 0x11, 0xbc, 0x43, 0x0f, 0xea, 0x2b, 0x85, 0xac, 0x66, 0x6d, 0x0e, 0xe8, 0x7f,
	0xc0, 0x53, 0x48, 0xbf, 0x4a, 0x47, 0x54, 0x06, 0xdd, 0x69, 0x81, 0x32, 0xe4, 0x0c, 0x34, 0x82,
	0xc3, 0x4c, 0x49, 0x13, 0x06, 0x2f, 0x63, 0x0a, 0x14, 0xba, 0x8d, 0x8e, 0x68, 0x4e, 0xc3, 0xbd,
	0x0b, 0xc1, 0xcd, 0xa0, 0xea, 0xd2, 0x2c, 0x93, 0x22, 0x3c, 0x42, 0x26, 0x7d, 0x8c, 0xcc, 0x0d,
	0x8f, 0x5c, 0x38, 0xa3, 0xc9, 0x6a, 0x9a, 0x01, 0x88, 0xd1, 0xcc, 0xc1, 0x0c, 0x3d, 0xfc, 0x35,
	0xe2, 0x96, 0x3b, 0x20, 0x81, 0x7d, 0x2b, 0x97, 0xee, 0x46, 0xb4, 0x68, 0xfb, 0x66, 0xd7, 0xec,
	0xaf, 0xc6, 0x16, 0x6d, 0xf7, 0x24, 0x36, 0xf5, 0xa5, 0x6f, 0x77, 0xcd, 0x7e, 0xfb, 0x7c, 0x27,
	0x6b, 0x9a, 0x7c, 0xff, 0x0e, 0x31, 0xd7, 0x8f, 0x91, 0xef, 0xdd, 0xab, 0xb8, 0x5e, 0xa2, 0xd5,
	0x04, 0xf6, 0x50, 0x64, 0xfa, 0x8b, 0xca, 0x0f, 0x92, 0x4d, 0xe5, 0x6a, 0x2a, 0x3f, 0x57, 0xd3,
	0x71, 0x7b, 0xe6, 0xcb, 0xc6, 0x5c, 0x56, 0xe0, 0xe5, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x44,
	0x53, 0xb6, 0x2a, 0x01, 0x00, 0x00,
}