// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/unifiedfleet/api/v1/proto/location.proto

package ufspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Lab refers to the different Labs under chrome org
// More labs to be added later if needed
type Lab int32

const (
	Lab_LAB_UNSPECIFIED         Lab = 0
	Lab_LAB_CHROME_ATLANTA      Lab = 1
	Lab_LAB_CHROMEOS_SANTIEM    Lab = 2
	Lab_LAB_CHROMEOS_DESTINY    Lab = 3
	Lab_LAB_CHROMEOS_PROMETHEUS Lab = 4
	Lab_LAB_CHROMEOS_ATLANTIS   Lab = 5
	Lab_LAB_CHROMEOS_LINDAVISTA Lab = 6
	Lab_LAB_DATACENTER_ATL97    Lab = 7
	Lab_LAB_DATACENTER_IAD97    Lab = 8
	Lab_LAB_DATACENTER_MTV96    Lab = 9
	Lab_LAB_DATACENTER_MTV97    Lab = 10
	Lab_LAB_DATACENTER_FUCHSIA  Lab = 11
)

var Lab_name = map[int32]string{
	0:  "LAB_UNSPECIFIED",
	1:  "LAB_CHROME_ATLANTA",
	2:  "LAB_CHROMEOS_SANTIEM",
	3:  "LAB_CHROMEOS_DESTINY",
	4:  "LAB_CHROMEOS_PROMETHEUS",
	5:  "LAB_CHROMEOS_ATLANTIS",
	6:  "LAB_CHROMEOS_LINDAVISTA",
	7:  "LAB_DATACENTER_ATL97",
	8:  "LAB_DATACENTER_IAD97",
	9:  "LAB_DATACENTER_MTV96",
	10: "LAB_DATACENTER_MTV97",
	11: "LAB_DATACENTER_FUCHSIA",
}

var Lab_value = map[string]int32{
	"LAB_UNSPECIFIED":         0,
	"LAB_CHROME_ATLANTA":      1,
	"LAB_CHROMEOS_SANTIEM":    2,
	"LAB_CHROMEOS_DESTINY":    3,
	"LAB_CHROMEOS_PROMETHEUS": 4,
	"LAB_CHROMEOS_ATLANTIS":   5,
	"LAB_CHROMEOS_LINDAVISTA": 6,
	"LAB_DATACENTER_ATL97":    7,
	"LAB_DATACENTER_IAD97":    8,
	"LAB_DATACENTER_MTV96":    9,
	"LAB_DATACENTER_MTV97":    10,
	"LAB_DATACENTER_FUCHSIA":  11,
}

func (x Lab) String() string {
	return proto.EnumName(Lab_name, int32(x))
}

func (Lab) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ee59bb28e2687ad5, []int{0}
}

// Location of the asset(Rack/Machine) in the lab
// For Browser machine, lab and rack are the only field to fill in.
// The fine-grained location is mainly for OS machine as we care about rack, row, shelf.
type Location struct {
	// Different labs in the chrome org. Required.
	Lab Lab `protobuf:"varint,1,opt,name=lab,proto3,enum=unifiedfleet.api.v1.proto.Lab" json:"lab,omitempty"`
	// Each lab has many aisles.
	// This field refers to the aisle number/name in the lab.
	Aisle string `protobuf:"bytes,2,opt,name=aisle,proto3" json:"aisle,omitempty"`
	// Each aisle has many rows.
	// This field refers to the row number/name in the aisle.
	Row string `protobuf:"bytes,3,opt,name=row,proto3" json:"row,omitempty"`
	// Each row has many racks.
	// This field refers to the rack number/name in the row.
	Rack string `protobuf:"bytes,4,opt,name=rack,proto3" json:"rack,omitempty"`
	// The position of the rack in the row.
	RackNumber string `protobuf:"bytes,5,opt,name=rack_number,json=rackNumber,proto3" json:"rack_number,omitempty"`
	// Each rack has many shelfs.
	// This field refers to the shelf number/name in the rack.
	Shelf string `protobuf:"bytes,6,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// Each shelf has many positions where assets can be placed.
	// This field refers to the position number/name in the shelf
	Position             string   `protobuf:"bytes,7,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee59bb28e2687ad5, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLab() Lab {
	if m != nil {
		return m.Lab
	}
	return Lab_LAB_UNSPECIFIED
}

func (m *Location) GetAisle() string {
	if m != nil {
		return m.Aisle
	}
	return ""
}

func (m *Location) GetRow() string {
	if m != nil {
		return m.Row
	}
	return ""
}

func (m *Location) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Location) GetRackNumber() string {
	if m != nil {
		return m.RackNumber
	}
	return ""
}

func (m *Location) GetShelf() string {
	if m != nil {
		return m.Shelf
	}
	return ""
}

func (m *Location) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func init() {
	proto.RegisterEnum("unifiedfleet.api.v1.proto.Lab", Lab_name, Lab_value)
	proto.RegisterType((*Location)(nil), "unifiedfleet.api.v1.proto.Location")
}

func init() {
	proto.RegisterFile("infra/unifiedfleet/api/v1/proto/location.proto", fileDescriptor_ee59bb28e2687ad5)
}

var fileDescriptor_ee59bb28e2687ad5 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xef, 0x6b, 0xd4, 0x30,
	0x1c, 0xc6, 0xbd, 0x9f, 0xbb, 0x65, 0xa0, 0x21, 0xce, 0xd9, 0x9d, 0xa0, 0x43, 0x18, 0x1b, 0x83,
	0xa5, 0x4e, 0xc1, 0xe3, 0x10, 0x84, 0xdc, 0xb5, 0xe3, 0x0a, 0xbd, 0x6e, 0xb4, 0xbd, 0x81, 0xbe,
	0x39, 0xd2, 0x9a, 0x76, 0x61, 0xed, 0xa5, 0xa4, 0xed, 0xc4, 0xbf, 0xce, 0x3f, 0x4b, 0xf0, 0x95,
	0x24, 0x3d, 0x7f, 0xcc, 0x6d, 0xec, 0x55, 0xfa, 0x3c, 0x9f, 0x7c, 0x9f, 0xe7, 0x5b, 0x08, 0xc0,
	0x7c, 0x95, 0x48, 0x6a, 0xd6, 0x2b, 0x9e, 0x70, 0xf6, 0x25, 0xc9, 0x18, 0xab, 0x4c, 0x5a, 0x70,
	0xf3, 0xfa, 0xc4, 0x2c, 0xa4, 0xa8, 0x84, 0x99, 0x89, 0x98, 0x56, 0x5c, 0xac, 0xb0, 0x96, 0x68,
	0xf7, 0xdf, 0x9b, 0x98, 0x16, 0x1c, 0x5f, 0x9f, 0x34, 0x68, 0x38, 0x4e, 0x05, 0x8e, 0x2f, 0xa5,
	0xc8, 0x79, 0x9d, 0x63, 0x21, 0x53, 0x33, 0xab, 0x63, 0x6e, 0xa6, 0xb2, 0x88, 0xd7, 0x51, 0xa9,
	0x10, 0x69, 0xc6, 0x74, 0xbc, 0x64, 0xa5, 0xa8, 0x65, 0xcc, 0x9a, 0xd1, 0xd7, 0x3f, 0x5a, 0x60,
	0xe0, 0xae, 0x8b, 0xd0, 0x1b, 0xd0, 0xc9, 0x68, 0x64, 0xb4, 0xf6, 0x5a, 0x87, 0x8f, 0xdf, 0xbe,
	0xc4, 0xf7, 0x16, 0x62, 0x97, 0x46, 0xbe, 0xba, 0x8a, 0xb6, 0x41, 0x8f, 0xf2, 0x32, 0x63, 0x46,
	0x7b, 0xaf, 0x75, 0xb8, 0xe9, 0x37, 0x02, 0x41, 0xd0, 0x91, 0xe2, 0xab, 0xd1, 0xd1, 0x9e, 0xfa,
	0x44, 0x1f, 0x41, 0x57, 0xd2, 0xf8, 0xca, 0xe8, 0x2a, 0x6b, 0x72, 0xf4, 0x93, 0x1c, 0x80, 0xfd,
	0x75, 0xfc, 0xb1, 0xce, 0x3f, 0x2e, 0xbf, 0x95, 0x15, 0xcb, 0x31, 0x2d, 0x8a, 0xb2, 0x10, 0x15,
	0x8e, 0x45, 0x6e, 0xfa, 0x34, 0xbe, 0xf2, 0xf5, 0x1c, 0x7a, 0x05, 0xb6, 0xd4, 0xb9, 0x5c, 0xd5,
	0x79, 0xc4, 0xa4, 0xd1, 0xd3, 0xc9, 0x40, 0x59, 0x9e, 0x76, 0xd4, 0x22, 0xe5, 0x25, 0xcb, 0x12,
	0xa3, 0xdf, 0x2c, 0xa2, 0x05, 0x1a, 0x82, 0x41, 0x21, 0x4a, 0xae, 0x7e, 0xce, 0xd8, 0xd0, 0xe0,
	0x8f, 0x3e, 0xfa, 0xde, 0x06, 0x1d, 0x97, 0x46, 0xe8, 0x29, 0x78, 0xe2, 0x92, 0xc9, 0x72, 0xe1,
	0x05, 0xe7, 0xf6, 0xd4, 0x39, 0x75, 0x6c, 0x0b, 0x3e, 0x42, 0x3b, 0x00, 0x29, 0x73, 0x3a, 0xf3,
	0xcf, 0xe6, 0xf6, 0x92, 0x84, 0x2e, 0xf1, 0x42, 0x02, 0x5b, 0xc8, 0x00, 0xdb, 0x7f, 0xfd, 0xb3,
	0x60, 0x19, 0x10, 0x2f, 0x74, 0xec, 0x39, 0x6c, 0xdf, 0x22, 0x96, 0x1d, 0x84, 0x8e, 0xf7, 0x09,
	0x76, 0xd0, 0x0b, 0xf0, 0xfc, 0x06, 0x39, 0x57, 0x67, 0x38, 0xb3, 0x17, 0x01, 0xec, 0xa2, 0x5d,
	0xf0, 0xec, 0x06, 0x6c, 0xaa, 0x9c, 0x00, 0xf6, 0x6e, 0xcd, 0xb9, 0x8e, 0x67, 0x91, 0x0b, 0x27,
	0x08, 0x09, 0xec, 0xff, 0xae, 0xb3, 0x48, 0x48, 0xa6, 0xb6, 0x17, 0xda, 0xbe, 0x9a, 0x1c, 0x8f,
	0xe0, 0xc6, 0x1d, 0xc4, 0x21, 0xd6, 0x78, 0x04, 0x07, 0x77, 0x90, 0x79, 0x78, 0x31, 0x7e, 0x0f,
	0x37, 0xef, 0x21, 0x23, 0x08, 0xd0, 0x10, 0xec, 0xfc, 0x47, 0x4e, 0x17, 0xd3, 0x59, 0xe0, 0x10,
	0xb8, 0x35, 0x39, 0xf8, 0xbc, 0xff, 0xc0, 0x1b, 0xfe, 0x50, 0x27, 0x65, 0x11, 0x45, 0x7d, 0x2d,
	0xde, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x88, 0xad, 0x50, 0x8b, 0xf3, 0x02, 0x00, 0x00,
}