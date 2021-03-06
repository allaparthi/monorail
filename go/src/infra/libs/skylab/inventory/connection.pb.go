// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection.proto

package inventory

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

// NEXT TAG: 6
type ServoHostConnection struct {
	ServoHostId          *string  `protobuf:"bytes,1,req,name=servo_host_id,json=servoHostId" json:"servo_host_id,omitempty"`
	DutId                *string  `protobuf:"bytes,2,req,name=dut_id,json=dutId" json:"dut_id,omitempty"`
	ServoPort            *int32   `protobuf:"varint,3,req,name=servo_port,json=servoPort" json:"servo_port,omitempty"`
	ServoSerial          *string  `protobuf:"bytes,4,opt,name=servo_serial,json=servoSerial" json:"servo_serial,omitempty"`
	ServoType            *string  `protobuf:"bytes,5,opt,name=servo_type,json=servoType" json:"servo_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServoHostConnection) Reset()         { *m = ServoHostConnection{} }
func (m *ServoHostConnection) String() string { return proto.CompactTextString(m) }
func (*ServoHostConnection) ProtoMessage()    {}
func (*ServoHostConnection) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

func (m *ServoHostConnection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServoHostConnection.Unmarshal(m, b)
}
func (m *ServoHostConnection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServoHostConnection.Marshal(b, m, deterministic)
}
func (m *ServoHostConnection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServoHostConnection.Merge(m, src)
}
func (m *ServoHostConnection) XXX_Size() int {
	return xxx_messageInfo_ServoHostConnection.Size(m)
}
func (m *ServoHostConnection) XXX_DiscardUnknown() {
	xxx_messageInfo_ServoHostConnection.DiscardUnknown(m)
}

var xxx_messageInfo_ServoHostConnection proto.InternalMessageInfo

func (m *ServoHostConnection) GetServoHostId() string {
	if m != nil && m.ServoHostId != nil {
		return *m.ServoHostId
	}
	return ""
}

func (m *ServoHostConnection) GetDutId() string {
	if m != nil && m.DutId != nil {
		return *m.DutId
	}
	return ""
}

func (m *ServoHostConnection) GetServoPort() int32 {
	if m != nil && m.ServoPort != nil {
		return *m.ServoPort
	}
	return 0
}

func (m *ServoHostConnection) GetServoSerial() string {
	if m != nil && m.ServoSerial != nil {
		return *m.ServoSerial
	}
	return ""
}

func (m *ServoHostConnection) GetServoType() string {
	if m != nil && m.ServoType != nil {
		return *m.ServoType
	}
	return ""
}

// NEXT TAG: 3
type ChameleonConnection struct {
	Chameleon            *ChameleonDevice `protobuf:"bytes,1,req,name=chameleon" json:"chameleon,omitempty"`
	ControlledDevice     *Device          `protobuf:"bytes,2,req,name=controlled_device,json=controlledDevice" json:"controlled_device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChameleonConnection) Reset()         { *m = ChameleonConnection{} }
func (m *ChameleonConnection) String() string { return proto.CompactTextString(m) }
func (*ChameleonConnection) ProtoMessage()    {}
func (*ChameleonConnection) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

func (m *ChameleonConnection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChameleonConnection.Unmarshal(m, b)
}
func (m *ChameleonConnection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChameleonConnection.Marshal(b, m, deterministic)
}
func (m *ChameleonConnection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChameleonConnection.Merge(m, src)
}
func (m *ChameleonConnection) XXX_Size() int {
	return xxx_messageInfo_ChameleonConnection.Size(m)
}
func (m *ChameleonConnection) XXX_DiscardUnknown() {
	xxx_messageInfo_ChameleonConnection.DiscardUnknown(m)
}

var xxx_messageInfo_ChameleonConnection proto.InternalMessageInfo

func (m *ChameleonConnection) GetChameleon() *ChameleonDevice {
	if m != nil {
		return m.Chameleon
	}
	return nil
}

func (m *ChameleonConnection) GetControlledDevice() *Device {
	if m != nil {
		return m.ControlledDevice
	}
	return nil
}

func init() {
	proto.RegisterType((*ServoHostConnection)(nil), "chrome.chromeos_infra.skylab.proto.inventory.ServoHostConnection")
	proto.RegisterType((*ChameleonConnection)(nil), "chrome.chromeos_infra.skylab.proto.inventory.ChameleonConnection")
}

func init() {
	proto.RegisterFile("connection.proto", fileDescriptor_51baa40a1cc6b48b)
}

var fileDescriptor_51baa40a1cc6b48b = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xb5, 0x42, 0x5f, 0x27, 0xcc, 0x0e, 0xa1, 0x08, 0x42, 0xed, 0x69, 0x07, 0xc9,
	0x61, 0x78, 0xf5, 0xe2, 0x3c, 0xe8, 0x4d, 0x3a, 0x4f, 0x7a, 0x28, 0x35, 0x79, 0xd2, 0x60, 0x97,
	0x57, 0x92, 0xac, 0xd0, 0x8f, 0xe5, 0x87, 0xf1, 0xfb, 0xc8, 0x12, 0xd7, 0x7a, 0xdd, 0x29, 0xf0,
	0x7b, 0xff, 0xfc, 0xf2, 0xf2, 0x87, 0x39, 0x27, 0xa5, 0x90, 0x5b, 0x49, 0x8a, 0x75, 0x9a, 0x2c,
	0xa5, 0xb7, 0xbc, 0xd1, 0xb4, 0x45, 0xe6, 0x0f, 0x32, 0x95, 0x54, 0x9f, 0xba, 0x66, 0xe6, 0x6b,
	0x68, 0xeb, 0x0f, 0x9f, 0x61, 0x52, 0xf5, 0xa8, 0x2c, 0xe9, 0xe1, 0x6a, 0x26, 0xb0, 0x97, 0x1c,
	0x3d, 0x2f, 0xbe, 0x03, 0x58, 0x6c, 0x50, 0xf7, 0xf4, 0x44, 0xc6, 0xae, 0x47, 0x73, 0x5a, 0xc0,
	0xb9, 0xd9, 0xe3, 0xaa, 0x21, 0x63, 0x2b, 0x29, 0xb2, 0x20, 0x0f, 0x97, 0x71, 0x99, 0x98, 0x43,
	0xf6, 0x59, 0xa4, 0x97, 0x70, 0x26, 0x76, 0x6e, 0x18, 0xba, 0x61, 0x24, 0x76, 0x7b, 0x7c, 0x0d,
	0xe0, 0xaf, 0x76, 0xa4, 0x6d, 0x76, 0x92, 0x87, 0xcb, 0xa8, 0x8c, 0x1d, 0x79, 0x21, 0x6d, 0xd3,
	0x1b, 0x98, 0xf9, 0xb1, 0x41, 0x2d, 0xeb, 0x36, 0x3b, 0xcd, 0x83, 0x51, 0xbc, 0x71, 0x68, 0x32,
	0xd8, 0xa1, 0xc3, 0x2c, 0x72, 0x01, 0x6f, 0x78, 0x1d, 0x3a, 0x2c, 0x7e, 0x02, 0x58, 0xac, 0x9b,
	0x7a, 0x8b, 0x2d, 0x92, 0xfa, 0xb7, 0xf3, 0x3b, 0xc4, 0xfc, 0x80, 0xdd, 0xbe, 0xc9, 0xea, 0x9e,
	0x1d, 0xd3, 0x0d, 0x1b, 0xad, 0x8f, 0xae, 0xa3, 0x72, 0xf2, 0xa5, 0x35, 0x5c, 0x70, 0x52, 0x56,
	0x53, 0xdb, 0xa2, 0xa8, 0x7c, 0x87, 0xee, 0xdf, 0xc9, 0xea, 0xee, 0xb8, 0x47, 0xfe, 0xdc, 0xf3,
	0x49, 0xe7, 0xc9, 0x43, 0xf2, 0x16, 0x8f, 0xa9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xb8,
	0xb4, 0x66, 0xe7, 0x01, 0x00, 0x00,
}
