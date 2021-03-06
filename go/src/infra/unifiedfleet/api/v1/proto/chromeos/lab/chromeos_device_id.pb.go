// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/unifiedfleet/api/v1/proto/chromeos/lab/chromeos_device_id.proto

package ufspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ChromeOSDeviceID struct {
	// A unique ID for chromeos device, a randomly generated uuid or AssetTag.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChromeOSDeviceID) Reset()         { *m = ChromeOSDeviceID{} }
func (m *ChromeOSDeviceID) String() string { return proto.CompactTextString(m) }
func (*ChromeOSDeviceID) ProtoMessage()    {}
func (*ChromeOSDeviceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_621b60b0dfc256c7, []int{0}
}

func (m *ChromeOSDeviceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChromeOSDeviceID.Unmarshal(m, b)
}
func (m *ChromeOSDeviceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChromeOSDeviceID.Marshal(b, m, deterministic)
}
func (m *ChromeOSDeviceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChromeOSDeviceID.Merge(m, src)
}
func (m *ChromeOSDeviceID) XXX_Size() int {
	return xxx_messageInfo_ChromeOSDeviceID.Size(m)
}
func (m *ChromeOSDeviceID) XXX_DiscardUnknown() {
	xxx_messageInfo_ChromeOSDeviceID.DiscardUnknown(m)
}

var xxx_messageInfo_ChromeOSDeviceID proto.InternalMessageInfo

func (m *ChromeOSDeviceID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ChromeOSDeviceID)(nil), "unifiedfleet.api.v1.proto.chromeos.lab.ChromeOSDeviceID")
}

func init() {
	proto.RegisterFile("infra/unifiedfleet/api/v1/proto/chromeos/lab/chromeos_device_id.proto", fileDescriptor_621b60b0dfc256c7)
}

var fileDescriptor_621b60b0dfc256c7 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcd, 0xcc, 0x4b, 0x2b,
	0x4a, 0xd4, 0x2f, 0xcd, 0xcb, 0x4c, 0xcb, 0x4c, 0x4d, 0x49, 0xcb, 0x49, 0x4d, 0x2d, 0xd1, 0x4f,
	0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xce, 0x28, 0xca,
	0xcf, 0x4d, 0xcd, 0x2f, 0xd6, 0xcf, 0x49, 0x4c, 0x82, 0x73, 0xe2, 0x53, 0x52, 0xcb, 0x32, 0x93,
	0x53, 0xe3, 0x33, 0x53, 0xf4, 0xc0, 0xaa, 0x84, 0xd4, 0x90, 0x0d, 0xd0, 0x4b, 0x2c, 0xc8, 0xd4,
	0x2b, 0x33, 0x84, 0x48, 0xe9, 0xc1, 0xf4, 0xe8, 0xe5, 0x24, 0x26, 0x29, 0x69, 0x70, 0x09, 0x38,
	0x83, 0xf9, 0xfe, 0xc1, 0x2e, 0x60, 0x23, 0x3c, 0x5d, 0x84, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73,
	0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x27, 0x93, 0x28, 0x23, 0x52,
	0x9c, 0x66, 0x5d, 0x9a, 0x56, 0x5c, 0x90, 0x94, 0xc4, 0x06, 0x96, 0x31, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x97, 0x6c, 0xb5, 0x47, 0xd7, 0x00, 0x00, 0x00,
}
