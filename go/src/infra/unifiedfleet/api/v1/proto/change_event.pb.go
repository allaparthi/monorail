// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/unifiedfleet/api/v1/proto/change_event.proto

package ufspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// represents the ChangeEvent generated when there is any change for the asset
type ChangeEvent struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The string representation of the changed (multi-component) field paths,
	// See explanation of the field path here:
	// https://github.com/protocolbuffers/protobuf/blob/50e03cdde3ef1fc7e9674db0a98ee1dea93f6fb2/src/google/protobuf/field_mask.proto#L43
	// machine.serial_number, chromeos_machine.model, dut.config.peripherals.wifi.wificell,
	// peripheral_requirement.min
	EventLabel string `protobuf:"bytes,2,opt,name=event_label,json=eventLabel,proto3" json:"event_label,omitempty"`
	// The string representation of the changed item, e.g.
	// machine.serial_number: from "" => A
	// chromeos_machine.model: from modelA => modelB
	// dut.config.peripherals.wifi.wificell: from false => true
	// periphral_requirement.min: from 1 => 3
	OldValue string `protobuf:"bytes,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	NewValue string `protobuf:"bytes,4,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	// Record the last update timestamp of this Event (In UTC timezone)
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	UserEmail            string               `protobuf:"bytes,6,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Comment              string               `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChangeEvent) Reset()         { *m = ChangeEvent{} }
func (m *ChangeEvent) String() string { return proto.CompactTextString(m) }
func (*ChangeEvent) ProtoMessage()    {}
func (*ChangeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_98d19b956948a621, []int{0}
}

func (m *ChangeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeEvent.Unmarshal(m, b)
}
func (m *ChangeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeEvent.Marshal(b, m, deterministic)
}
func (m *ChangeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeEvent.Merge(m, src)
}
func (m *ChangeEvent) XXX_Size() int {
	return xxx_messageInfo_ChangeEvent.Size(m)
}
func (m *ChangeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeEvent proto.InternalMessageInfo

func (m *ChangeEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChangeEvent) GetEventLabel() string {
	if m != nil {
		return m.EventLabel
	}
	return ""
}

func (m *ChangeEvent) GetOldValue() string {
	if m != nil {
		return m.OldValue
	}
	return ""
}

func (m *ChangeEvent) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

func (m *ChangeEvent) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *ChangeEvent) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *ChangeEvent) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func init() {
	proto.RegisterType((*ChangeEvent)(nil), "unifiedfleet.api.v1.proto.ChangeEvent")
}

func init() {
	proto.RegisterFile("infra/unifiedfleet/api/v1/proto/change_event.proto", fileDescriptor_98d19b956948a621)
}

var fileDescriptor_98d19b956948a621 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbb, 0xee, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0x0b, 0x2d, 0x75, 0x24, 0x86, 0x4c, 0xa1, 0x08, 0xb5, 0x42, 0x02, 0xba, 0xd4,
	0x56, 0xcb, 0x02, 0x74, 0x80, 0x16, 0x75, 0x63, 0xaa, 0x10, 0x03, 0x4b, 0xe4, 0x24, 0x27, 0xa9,
	0x25, 0xdf, 0xe4, 0x4b, 0x2a, 0x84, 0x78, 0x2b, 0x1e, 0x88, 0x99, 0xa7, 0x40, 0xb6, 0x53, 0x89,
	0x0d, 0xfd, 0xa7, 0xe4, 0x7c, 0x3f, 0x9f, 0x13, 0x7f, 0xdf, 0x09, 0xda, 0x33, 0xd9, 0x1a, 0x4a,
	0xbc, 0x64, 0x2d, 0x83, 0xa6, 0xe5, 0x00, 0x8e, 0x50, 0xcd, 0x48, 0xbf, 0x23, 0xda, 0x28, 0xa7,
	0x48, 0x7d, 0xa5, 0xb2, 0x83, 0x12, 0x7a, 0x90, 0x0e, 0x47, 0x29, 0x7f, 0xfa, 0xef, 0x69, 0x4c,
	0x35, 0xc3, 0xfd, 0x2e, 0xa1, 0xe5, 0xaa, 0x53, 0xaa, 0xe3, 0x90, 0x7a, 0x2b, 0xdf, 0x12, 0xc7,
	0x04, 0x58, 0x47, 0x85, 0x1e, 0x0e, 0xbc, 0xeb, 0x14, 0xae, 0xaf, 0x46, 0x09, 0xe6, 0x05, 0x56,
	0xa6, 0x23, 0xdc, 0xd7, 0x8c, 0x74, 0x46, 0xd7, 0xc3, 0xf7, 0x86, 0x01, 0xe1, 0x0e, 0x06, 0xac,
	0xf2, 0xa6, 0x86, 0xa1, 0xf5, 0xc3, 0x03, 0x5a, 0x5b, 0x06, 0xbc, 0x29, 0x2b, 0xb8, 0xd2, 0x9e,
	0x29, 0x93, 0x06, 0xbc, 0xf8, 0x35, 0x46, 0xd9, 0xa7, 0x68, 0xe7, 0x1c, 0xdc, 0xe4, 0x39, 0x9a,
	0x4a, 0x2a, 0xa0, 0x18, 0xad, 0x47, 0x9b, 0xc5, 0x25, 0xbe, 0xe7, 0x2b, 0x94, 0x45, 0xab, 0x25,
	0xa7, 0x15, 0xf0, 0x62, 0x1c, 0x11, 0x8a, 0xd2, 0xe7, 0xa0, 0xe4, 0xcf, 0xd0, 0x42, 0xf1, 0xa6,
	0xec, 0x29, 0xf7, 0x50, 0x4c, 0x22, 0x7e, 0xac, 0x78, 0xf3, 0x35, 0xd4, 0x01, 0x4a, 0xb8, 0x0d,
	0x70, 0x9a, 0xa0, 0x84, 0x5b, 0x82, 0x1f, 0x51, 0xe6, 0x75, 0x43, 0x1d, 0x94, 0x21, 0x94, 0xe2,
	0xd1, 0x7a, 0xb4, 0xc9, 0xf6, 0x4b, 0x9c, 0x6e, 0x8d, 0xef, 0x89, 0xe1, 0x2f, 0xf7, 0xc4, 0x4e,
	0x93, 0xdf, 0xc7, 0xc9, 0x05, 0xa5, 0x9e, 0xa0, 0xe6, 0xcf, 0x11, 0xf2, 0x16, 0x4c, 0x09, 0x82,
	0x32, 0x5e, 0xcc, 0xe2, 0xfc, 0x45, 0x50, 0xce, 0x41, 0xc8, 0x0b, 0x34, 0xaf, 0x95, 0x10, 0x20,
	0x5d, 0x31, 0x8f, 0xec, 0x5e, 0xbe, 0x3f, 0xfc, 0x39, 0xbe, 0x45, 0xaf, 0x86, 0xb5, 0x6d, 0xe3,
	0xde, 0xb6, 0xf6, 0xbb, 0x75, 0x20, 0x30, 0xd5, 0xda, 0x6a, 0xe5, 0x70, 0xad, 0x04, 0x49, 0xb1,
	0x3c, 0x89, 0x6e, 0x2d, 0xf9, 0x11, 0x9f, 0x3f, 0x4f, 0xaf, 0xbf, 0xbd, 0xfc, 0xcf, 0x4f, 0x72,
	0xf0, 0xad, 0xd5, 0x55, 0x35, 0x8b, 0xc5, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0x8b,
	0xbe, 0xb5, 0x54, 0x02, 0x00, 0x00,
}
