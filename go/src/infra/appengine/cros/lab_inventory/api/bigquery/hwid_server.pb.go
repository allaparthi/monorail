// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/cros/lab_inventory/api/bigquery/hwid_server.proto

package apibq

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Next Tag: 4
type HWIDInventory struct {
	Hwid                 string               `protobuf:"bytes,1,opt,name=hwid,proto3" json:"hwid,omitempty"`
	Sku                  string               `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	UpdatedTime          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HWIDInventory) Reset()         { *m = HWIDInventory{} }
func (m *HWIDInventory) String() string { return proto.CompactTextString(m) }
func (*HWIDInventory) ProtoMessage()    {}
func (*HWIDInventory) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a3b8f731ba9c667, []int{0}
}

func (m *HWIDInventory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HWIDInventory.Unmarshal(m, b)
}
func (m *HWIDInventory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HWIDInventory.Marshal(b, m, deterministic)
}
func (m *HWIDInventory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HWIDInventory.Merge(m, src)
}
func (m *HWIDInventory) XXX_Size() int {
	return xxx_messageInfo_HWIDInventory.Size(m)
}
func (m *HWIDInventory) XXX_DiscardUnknown() {
	xxx_messageInfo_HWIDInventory.DiscardUnknown(m)
}

var xxx_messageInfo_HWIDInventory proto.InternalMessageInfo

func (m *HWIDInventory) GetHwid() string {
	if m != nil {
		return m.Hwid
	}
	return ""
}

func (m *HWIDInventory) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *HWIDInventory) GetUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedTime
	}
	return nil
}

func init() {
	proto.RegisterType((*HWIDInventory)(nil), "apibq.HWIDInventory")
}

func init() {
	proto.RegisterFile("infra/appengine/cros/lab_inventory/api/bigquery/hwid_server.proto", fileDescriptor_6a3b8f731ba9c667)
}

var fileDescriptor_6a3b8f731ba9c667 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8d, 0xbf, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0xa9, 0xa7, 0x82, 0x39, 0x05, 0xc9, 0x54, 0x6e, 0xf1, 0x70, 0xba, 0x29, 0x1f, 0xe8,
	0xec, 0x20, 0x38, 0x78, 0x6b, 0x11, 0x1c, 0x4b, 0x62, 0xbf, 0xc6, 0x0f, 0xdb, 0x24, 0xcd, 0x8f,
	0x4a, 0xff, 0x7b, 0x49, 0x6a, 0xb7, 0x97, 0x97, 0x87, 0xe7, 0x61, 0xaf, 0x64, 0x7a, 0x2f, 0x41,
	0x3a, 0x87, 0x46, 0x93, 0x41, 0xf8, 0xf2, 0x36, 0xc0, 0x20, 0x55, 0x4b, 0x66, 0x46, 0x13, 0xad,
	0x5f, 0x40, 0x3a, 0x02, 0x45, 0x7a, 0x4a, 0xe8, 0x17, 0xf8, 0xfe, 0xa5, 0xae, 0x0d, 0xe8, 0x67,
	0xf4, 0xc2, 0x79, 0x1b, 0x2d, 0xbf, 0x92, 0x8e, 0xd4, 0x74, 0x78, 0xd0, 0xd6, 0xea, 0x01, 0xa1,
	0x9c, 0x2a, 0xf5, 0x10, 0x69, 0xc4, 0x10, 0xe5, 0xe8, 0x56, 0xee, 0x31, 0xb2, 0xbb, 0xf7, 0xcf,
	0xf3, 0xdb, 0x79, 0xf3, 0x72, 0xce, 0x2e, 0xb3, 0xad, 0xae, 0x8e, 0xd5, 0xe9, 0xa6, 0x29, 0x9b,
	0xdf, 0xb3, 0x5d, 0xf8, 0x49, 0xf5, 0x45, 0xb9, 0xf2, 0xe4, 0x2f, 0xec, 0x36, 0xb9, 0x4e, 0x46,
	0xec, 0xda, 0x6c, 0xac, 0x77, 0xc7, 0xea, 0xb4, 0x7f, 0x3a, 0x88, 0x35, 0x27, 0xb6, 0x9c, 0xf8,
	0xd8, 0x72, 0xcd, 0xfe, 0x9f, 0xcf, 0x8f, 0xba, 0x2e, 0xc0, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdc, 0x41, 0xb6, 0x17, 0xe9, 0x00, 0x00, 0x00,
}
