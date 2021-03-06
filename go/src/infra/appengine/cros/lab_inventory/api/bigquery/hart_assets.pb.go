// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/cros/lab_inventory/api/bigquery/hart_assets.proto

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

type HaRTAsset struct {
	AssetTag string `protobuf:"bytes,1,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	// Person|Team the device is assigned to
	Assignee       string `protobuf:"bytes,2,opt,name=assignee,proto3" json:"assignee,omitempty"`
	SerialNumber   string `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Description    string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	GoogleCodeName string `protobuf:"bytes,5,opt,name=google_code_name,json=googleCodeName,proto3" json:"google_code_name,omitempty"`
	// Google Part Number
	Gpn string `protobuf:"bytes,6,opt,name=gpn,proto3" json:"gpn,omitempty"`
	// Description associated to the part number
	GpnDescription       string               `protobuf:"bytes,7,opt,name=gpn_description,json=gpnDescription,proto3" json:"gpn_description,omitempty"`
	FormFactor           string               `protobuf:"bytes,8,opt,name=form_factor,json=formFactor,proto3" json:"form_factor,omitempty"`
	CostCenter           string               `protobuf:"bytes,9,opt,name=cost_center,json=costCenter,proto3" json:"cost_center,omitempty"`
	Project              string               `protobuf:"bytes,10,opt,name=project,proto3" json:"project,omitempty"`
	UpdatedTime          *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HaRTAsset) Reset()         { *m = HaRTAsset{} }
func (m *HaRTAsset) String() string { return proto.CompactTextString(m) }
func (*HaRTAsset) ProtoMessage()    {}
func (*HaRTAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_b37700a90c78be36, []int{0}
}

func (m *HaRTAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HaRTAsset.Unmarshal(m, b)
}
func (m *HaRTAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HaRTAsset.Marshal(b, m, deterministic)
}
func (m *HaRTAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HaRTAsset.Merge(m, src)
}
func (m *HaRTAsset) XXX_Size() int {
	return xxx_messageInfo_HaRTAsset.Size(m)
}
func (m *HaRTAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_HaRTAsset.DiscardUnknown(m)
}

var xxx_messageInfo_HaRTAsset proto.InternalMessageInfo

func (m *HaRTAsset) GetAssetTag() string {
	if m != nil {
		return m.AssetTag
	}
	return ""
}

func (m *HaRTAsset) GetAssignee() string {
	if m != nil {
		return m.Assignee
	}
	return ""
}

func (m *HaRTAsset) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *HaRTAsset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *HaRTAsset) GetGoogleCodeName() string {
	if m != nil {
		return m.GoogleCodeName
	}
	return ""
}

func (m *HaRTAsset) GetGpn() string {
	if m != nil {
		return m.Gpn
	}
	return ""
}

func (m *HaRTAsset) GetGpnDescription() string {
	if m != nil {
		return m.GpnDescription
	}
	return ""
}

func (m *HaRTAsset) GetFormFactor() string {
	if m != nil {
		return m.FormFactor
	}
	return ""
}

func (m *HaRTAsset) GetCostCenter() string {
	if m != nil {
		return m.CostCenter
	}
	return ""
}

func (m *HaRTAsset) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *HaRTAsset) GetUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedTime
	}
	return nil
}

func init() {
	proto.RegisterType((*HaRTAsset)(nil), "apibq.HaRTAsset")
}

func init() {
	proto.RegisterFile("infra/appengine/cros/lab_inventory/api/bigquery/hart_assets.proto", fileDescriptor_b37700a90c78be36)
}

var fileDescriptor_b37700a90c78be36 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbd, 0x6e, 0x22, 0x31,
	0x10, 0x80, 0xc5, 0x71, 0xfc, 0xac, 0x97, 0xbb, 0x43, 0xae, 0x2c, 0xae, 0x00, 0x25, 0x45, 0xa8,
	0xd6, 0x52, 0x52, 0xa7, 0x40, 0x44, 0x51, 0x2a, 0x0a, 0x44, 0x6f, 0xcd, 0xee, 0xce, 0x3a, 0x8e,
	0x58, 0xdb, 0xd8, 0xde, 0x48, 0x3c, 0x43, 0x5e, 0x3a, 0xb2, 0x17, 0x22, 0xba, 0x99, 0x6f, 0xbe,
	0xf1, 0x8c, 0x3c, 0x64, 0xa3, 0x74, 0xe3, 0x80, 0x83, 0xb5, 0xa8, 0xa5, 0xd2, 0xc8, 0x2b, 0x67,
	0x3c, 0x3f, 0x42, 0x29, 0x94, 0xfe, 0x44, 0x1d, 0x8c, 0x3b, 0x73, 0xb0, 0x8a, 0x97, 0x4a, 0x9e,
	0x3a, 0x74, 0x67, 0xfe, 0x0e, 0x2e, 0x08, 0xf0, 0x1e, 0x83, 0x2f, 0xac, 0x33, 0xc1, 0xd0, 0x11,
	0x58, 0x55, 0x9e, 0x16, 0x4b, 0x69, 0x8c, 0x3c, 0x22, 0x4f, 0xb0, 0xec, 0x1a, 0x1e, 0x54, 0x8b,
	0x3e, 0x40, 0x6b, 0x7b, 0xef, 0xee, 0x6b, 0x48, 0xb2, 0x37, 0xd8, 0x1f, 0x36, 0xb1, 0x99, 0xfe,
	0x27, 0x59, 0x7a, 0x45, 0x04, 0x90, 0x6c, 0xb0, 0x1a, 0xac, 0xb3, 0xfd, 0x34, 0x81, 0x03, 0x48,
	0xba, 0x20, 0x31, 0x56, 0x52, 0x23, 0xb2, 0x5f, 0x3f, 0xb5, 0x94, 0xd3, 0x7b, 0xf2, 0xc7, 0xa3,
	0x53, 0x70, 0x14, 0xba, 0x6b, 0x4b, 0x74, 0x6c, 0x98, 0x84, 0x59, 0x0f, 0x77, 0x89, 0xd1, 0x15,
	0xc9, 0x6b, 0xf4, 0x95, 0x53, 0x36, 0x28, 0xa3, 0xd9, 0xef, 0xa4, 0xdc, 0x22, 0xba, 0x26, 0xf3,
	0x7e, 0x61, 0x51, 0x99, 0x1a, 0x85, 0x86, 0x16, 0xd9, 0x28, 0x69, 0x7f, 0x7b, 0xbe, 0x35, 0x35,
	0xee, 0xa0, 0x45, 0x3a, 0x27, 0x43, 0x69, 0x35, 0x1b, 0xa7, 0x62, 0x0c, 0xe9, 0x03, 0xf9, 0x27,
	0xad, 0x16, 0xb7, 0x13, 0x26, 0x97, 0x56, 0xab, 0x5f, 0x6e, 0x86, 0x2c, 0x49, 0xde, 0x18, 0xd7,
	0x8a, 0x06, 0xaa, 0x60, 0x1c, 0x9b, 0x26, 0x89, 0x44, 0xf4, 0x9a, 0x48, 0x14, 0x2a, 0xe3, 0x83,
	0xa8, 0x50, 0x07, 0x74, 0x2c, 0xeb, 0x85, 0x88, 0xb6, 0x89, 0x50, 0x46, 0x26, 0xd6, 0x99, 0x0f,
	0xac, 0x02, 0x23, 0xa9, 0x78, 0x4d, 0xe9, 0x33, 0x99, 0x75, 0xb6, 0x86, 0x80, 0xb5, 0x88, 0x3f,
	0xcd, 0xf2, 0xd5, 0x60, 0x9d, 0x3f, 0x2e, 0x8a, 0x7e, 0xfb, 0xe2, 0x7a, 0x86, 0xe2, 0x70, 0x3d,
	0xc3, 0x3e, 0xbf, 0xf8, 0x91, 0x94, 0xe3, 0x24, 0x3c, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x75, 0x06, 0x64, 0x01, 0x02, 0x00, 0x00,
}
