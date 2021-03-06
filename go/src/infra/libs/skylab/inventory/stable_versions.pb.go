// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stable_versions.proto

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

// NEXT TAG: 5
type StableVersions struct {
	// OS image stable versions.
	AndroidOsVersions []*StableVersion `protobuf:"bytes,1,rep,name=android_os_versions,json=androidOsVersions" json:"android_os_versions,omitempty"`
	ChromeOsVersions  []*StableVersion `protobuf:"bytes,2,rep,name=chrome_os_versions,json=chromeOsVersions" json:"chrome_os_versions,omitempty"`
	// Read-write firmware versions. Only relevant for ChromeOS boards.
	RwFirmwareVersions []*StableVersion `protobuf:"bytes,3,rep,name=rw_firmware_versions,json=rwFirmwareVersions" json:"rw_firmware_versions,omitempty"`
	// Used by FAFT testing to install the RO firmware to test. ChromeOS only.
	FaftFirmwareVersions []*StableVersion `protobuf:"bytes,4,rep,name=faft_firmware_versions,json=faftFirmwareVersions" json:"faft_firmware_versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StableVersions) Reset()         { *m = StableVersions{} }
func (m *StableVersions) String() string { return proto.CompactTextString(m) }
func (*StableVersions) ProtoMessage()    {}
func (*StableVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3aed830256165b3, []int{0}
}

func (m *StableVersions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StableVersions.Unmarshal(m, b)
}
func (m *StableVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StableVersions.Marshal(b, m, deterministic)
}
func (m *StableVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StableVersions.Merge(m, src)
}
func (m *StableVersions) XXX_Size() int {
	return xxx_messageInfo_StableVersions.Size(m)
}
func (m *StableVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_StableVersions.DiscardUnknown(m)
}

var xxx_messageInfo_StableVersions proto.InternalMessageInfo

func (m *StableVersions) GetAndroidOsVersions() []*StableVersion {
	if m != nil {
		return m.AndroidOsVersions
	}
	return nil
}

func (m *StableVersions) GetChromeOsVersions() []*StableVersion {
	if m != nil {
		return m.ChromeOsVersions
	}
	return nil
}

func (m *StableVersions) GetRwFirmwareVersions() []*StableVersion {
	if m != nil {
		return m.RwFirmwareVersions
	}
	return nil
}

func (m *StableVersions) GetFaftFirmwareVersions() []*StableVersion {
	if m != nil {
		return m.FaftFirmwareVersions
	}
	return nil
}

// NEXT TAG: 3
type StableVersion struct {
	Board *string `protobuf:"bytes,1,req,name=board" json:"board,omitempty"`
	// Versions are opaque strings for the inventory. Different boards may use the
	// version strings in different ways to obtain the actual images.
	Version              *string  `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StableVersion) Reset()         { *m = StableVersion{} }
func (m *StableVersion) String() string { return proto.CompactTextString(m) }
func (*StableVersion) ProtoMessage()    {}
func (*StableVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3aed830256165b3, []int{1}
}

func (m *StableVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StableVersion.Unmarshal(m, b)
}
func (m *StableVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StableVersion.Marshal(b, m, deterministic)
}
func (m *StableVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StableVersion.Merge(m, src)
}
func (m *StableVersion) XXX_Size() int {
	return xxx_messageInfo_StableVersion.Size(m)
}
func (m *StableVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_StableVersion.DiscardUnknown(m)
}

var xxx_messageInfo_StableVersion proto.InternalMessageInfo

func (m *StableVersion) GetBoard() string {
	if m != nil && m.Board != nil {
		return *m.Board
	}
	return ""
}

func (m *StableVersion) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*StableVersions)(nil), "chrome.chromeos_infra.skylab.proto.inventory.StableVersions")
	proto.RegisterType((*StableVersion)(nil), "chrome.chromeos_infra.skylab.proto.inventory.StableVersion")
}

func init() {
	proto.RegisterFile("stable_versions.proto", fileDescriptor_c3aed830256165b3)
}

var fileDescriptor_c3aed830256165b3 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0x49, 0x4c,
	0xca, 0x49, 0x8d, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xd2, 0x49, 0xce, 0x28, 0xca, 0xcf, 0x4d, 0xd5, 0x83, 0x50, 0xf9, 0xc5, 0xf1, 0x99,
	0x79, 0x69, 0x45, 0x89, 0x7a, 0xc5, 0xd9, 0x95, 0x39, 0x89, 0x49, 0x10, 0x35, 0x7a, 0x99, 0x79,
	0x65, 0xa9, 0x79, 0x25, 0xf9, 0x45, 0x95, 0x4a, 0xa7, 0x98, 0xb9, 0xf8, 0x82, 0xc1, 0xe6, 0x84,
	0x41, 0x8d, 0x11, 0xca, 0xe6, 0x12, 0x4e, 0xcc, 0x4b, 0x29, 0xca, 0xcf, 0x4c, 0x89, 0xcf, 0x2f,
	0x86, 0x9b, 0x2e, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x64, 0xad, 0x47, 0x8a, 0xf1, 0x7a, 0x28,
	0x46, 0x07, 0x09, 0x42, 0xcd, 0xf5, 0x2f, 0x86, 0x5b, 0x96, 0xc9, 0x25, 0x04, 0x31, 0x09, 0xc5,
	0x2e, 0x26, 0xca, 0xed, 0x12, 0x80, 0x68, 0x42, 0xb2, 0x2a, 0x97, 0x4b, 0xa4, 0xa8, 0x3c, 0x3e,
	0x2d, 0xb3, 0x28, 0xb7, 0x3c, 0xb1, 0x08, 0x11, 0x6c, 0x12, 0xcc, 0x94, 0x5b, 0x26, 0x54, 0x54,
	0xee, 0x06, 0x35, 0x17, 0x6e, 0x5d, 0x21, 0x97, 0x58, 0x5a, 0x62, 0x5a, 0x09, 0x16, 0x0b, 0x59,
	0x28, 0xb7, 0x50, 0x04, 0x64, 0x34, 0xba, 0x95, 0x4a, 0xf6, 0x5c, 0xbc, 0x28, 0xca, 0x84, 0x44,
	0xb8, 0x58, 0x93, 0xf2, 0x13, 0x8b, 0x52, 0x24, 0x18, 0x15, 0x98, 0x34, 0x38, 0x83, 0x20, 0x1c,
	0x21, 0x09, 0x2e, 0x76, 0xa8, 0x5b, 0x24, 0x98, 0xc0, 0xe2, 0x30, 0xae, 0x13, 0x77, 0x14, 0x27,
	0xdc, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xaf, 0x34, 0x2f, 0x60, 0x02, 0x00, 0x00,
}
