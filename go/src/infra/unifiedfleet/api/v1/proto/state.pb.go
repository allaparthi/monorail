// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/unifiedfleet/api/v1/proto/state.proto

package ufspb

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

// Next tag: 10
type State int32

const (
	State_STATE_UNSPECIFIED State = 0
	// Equlavant to the concept in ChromeOS lab: needs_deploy
	State_STATE_REGISTERED State = 1
	// Deployed but not placed in prod. It's only a temporarily state for browser machine
	// as there's no service to push a deployed machine to prod automatically yet.
	State_STATE_DEPLOYED_PRE_SERVING State = 9
	// Deployed to the prod infrastructure, but for testing.
	State_STATE_DEPLOYED_TESTING State = 2
	// Deployed to the prod infrastructure, serving.
	State_STATE_SERVING State = 3
	// Deployed to the prod infrastructure, but needs repair.
	State_STATE_NEEDS_REPAIR State = 5
	// Deployed to the prod infrastructure, but get disabled.
	State_STATE_DISABLED State = 6
	// Deployed to the prod infrastructure, but get reserved (e.g. locked).
	State_STATE_RESERVED State = 7
	// Decommissioned from the prod infrastructure, but still leave in UFS record.
	State_STATE_DECOMMISSIONED State = 8
)

var State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "STATE_REGISTERED",
	9: "STATE_DEPLOYED_PRE_SERVING",
	2: "STATE_DEPLOYED_TESTING",
	3: "STATE_SERVING",
	5: "STATE_NEEDS_REPAIR",
	6: "STATE_DISABLED",
	7: "STATE_RESERVED",
	8: "STATE_DECOMMISSIONED",
}

var State_value = map[string]int32{
	"STATE_UNSPECIFIED":          0,
	"STATE_REGISTERED":           1,
	"STATE_DEPLOYED_PRE_SERVING": 9,
	"STATE_DEPLOYED_TESTING":     2,
	"STATE_SERVING":              3,
	"STATE_NEEDS_REPAIR":         5,
	"STATE_DISABLED":             6,
	"STATE_RESERVED":             7,
	"STATE_DECOMMISSIONED":       8,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_348721e056364fee, []int{0}
}

// There's no exposed API for users to directly retrieve a state record.
//
// Ideally, state record can only be modified internally by UFS after some essential
// preconditions are fulfilled.
//
// Users will focus on the tasks triggered by any state change instead of the state
// itself, e.g. once the state of a machine is changed to registered, lab admins will
// know it by founding more machines are listed for waiting for further configurations,
// instead of actively monitoring it by any tooling.
type StateRecord struct {
	// The string resource_name could be an ID of a rack, machine, RPM and switches.
	// It can also be the ID of virtual concepts, e.g. LSE and vlan.
	// The format of the resource name will be “racks/XXX” or “rpms/XXX” to help to
	// distinguish the type of the resource.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	State        State  `protobuf:"varint,2,opt,name=state,proto3,enum=unifiedfleet.api.v1.proto.State" json:"state,omitempty"`
	User         string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Ticket       string `protobuf:"bytes,4,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Description  string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Record the last update timestamp of this machine (In UTC timezone)
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StateRecord) Reset()         { *m = StateRecord{} }
func (m *StateRecord) String() string { return proto.CompactTextString(m) }
func (*StateRecord) ProtoMessage()    {}
func (*StateRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_348721e056364fee, []int{0}
}

func (m *StateRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateRecord.Unmarshal(m, b)
}
func (m *StateRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateRecord.Marshal(b, m, deterministic)
}
func (m *StateRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRecord.Merge(m, src)
}
func (m *StateRecord) XXX_Size() int {
	return xxx_messageInfo_StateRecord.Size(m)
}
func (m *StateRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRecord.DiscardUnknown(m)
}

var xxx_messageInfo_StateRecord proto.InternalMessageInfo

func (m *StateRecord) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *StateRecord) GetState() State {
	if m != nil {
		return m.State
	}
	return State_STATE_UNSPECIFIED
}

func (m *StateRecord) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *StateRecord) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

func (m *StateRecord) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StateRecord) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("unifiedfleet.api.v1.proto.State", State_name, State_value)
	proto.RegisterType((*StateRecord)(nil), "unifiedfleet.api.v1.proto.StateRecord")
}

func init() {
	proto.RegisterFile("infra/unifiedfleet/api/v1/proto/state.proto", fileDescriptor_348721e056364fee)
}

var fileDescriptor_348721e056364fee = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x6b, 0xd4, 0x40,
	0x14, 0x85, 0x4d, 0xdb, 0x5d, 0xed, 0x5d, 0x5b, 0xd2, 0x4b, 0x5d, 0xe2, 0x3e, 0x68, 0x50, 0xc4,
	0x45, 0x21, 0xa1, 0x15, 0x7c, 0xe9, 0xd3, 0xb6, 0xb9, 0x96, 0x81, 0x36, 0xbb, 0xcc, 0x44, 0x41,
	0x5f, 0xc2, 0xec, 0xee, 0xa4, 0x0c, 0x76, 0x37, 0x21, 0x99, 0xf4, 0xdf, 0xfa, 0x33, 0x7c, 0x97,
	0xcc, 0x34, 0x58, 0x04, 0xe9, 0xdb, 0xdc, 0xef, 0x9c, 0x9c, 0x7b, 0x0f, 0x04, 0x3e, 0xea, 0x6d,
	0x51, 0xcb, 0xb8, 0xdd, 0xea, 0x42, 0xab, 0x75, 0x71, 0xab, 0x94, 0x89, 0x65, 0xa5, 0xe3, 0xbb,
	0x93, 0xb8, 0xaa, 0x4b, 0x53, 0xc6, 0x8d, 0x91, 0x46, 0x45, 0xf6, 0x8d, 0x2f, 0x1f, 0xda, 0x22,
	0x59, 0xe9, 0xe8, 0xee, 0xc4, 0x49, 0x93, 0xd7, 0x37, 0x65, 0x79, 0x73, 0xab, 0xdc, 0x47, 0xcb,
	0xb6, 0x88, 0x8d, 0xde, 0xa8, 0xc6, 0xc8, 0x4d, 0xe5, 0x0c, 0x6f, 0x7e, 0x7b, 0x30, 0x12, 0x5d,
	0x16, 0x57, 0xab, 0xb2, 0x5e, 0xe3, 0x5b, 0x38, 0xa8, 0x55, 0x53, 0xb6, 0xf5, 0x4a, 0xe5, 0x5b,
	0xb9, 0x51, 0x81, 0x17, 0x7a, 0xd3, 0x7d, 0xfe, 0xbc, 0x87, 0xa9, 0xdc, 0x28, 0xfc, 0x0c, 0x03,
	0xbb, 0x3f, 0xd8, 0x09, 0xbd, 0xe9, 0xe1, 0x69, 0x18, 0xfd, 0xf7, 0x80, 0xc8, 0x65, 0x3b, 0x3b,
	0x22, 0xec, 0xb5, 0x8d, 0xaa, 0x83, 0x5d, 0x9b, 0x69, 0xdf, 0x38, 0x86, 0xa1, 0xd1, 0xab, 0x9f,
	0xca, 0x04, 0x7b, 0x96, 0xde, 0x4f, 0x18, 0xc2, 0x68, 0xad, 0x9a, 0x55, 0xad, 0x2b, 0xa3, 0xcb,
	0x6d, 0x30, 0xb0, 0xe2, 0x43, 0x84, 0x67, 0x30, 0x6a, 0xab, 0xb5, 0x34, 0x2a, 0xef, 0x4a, 0x05,
	0xc3, 0xd0, 0x9b, 0x8e, 0x4e, 0x27, 0x91, 0x6b, 0x1c, 0xf5, 0x8d, 0xa3, 0xac, 0x6f, 0xcc, 0xc1,
	0xd9, 0x3b, 0xf0, 0xe1, 0x97, 0x07, 0x03, 0x7b, 0x1b, 0xbe, 0x80, 0x23, 0x91, 0xcd, 0x32, 0xca,
	0xbf, 0xa6, 0x62, 0x41, 0x17, 0xec, 0x0b, 0xa3, 0xc4, 0x7f, 0x82, 0xc7, 0xe0, 0x3b, 0xcc, 0xe9,
	0x92, 0x89, 0x8c, 0x38, 0x25, 0xbe, 0x87, 0xaf, 0x60, 0xe2, 0x68, 0x42, 0x8b, 0xab, 0xf9, 0x77,
	0x4a, 0xf2, 0x05, 0xa7, 0x5c, 0x10, 0xff, 0xc6, 0xd2, 0x4b, 0x7f, 0x1f, 0x27, 0x30, 0xfe, 0x47,
	0xcf, 0x48, 0x64, 0x9d, 0xb6, 0x83, 0x47, 0x70, 0xe0, 0xb4, 0xde, 0xbe, 0x8b, 0x63, 0x40, 0x87,
	0x52, 0xa2, 0x44, 0xe4, 0x9c, 0x16, 0x33, 0xc6, 0xfd, 0x01, 0x22, 0x1c, 0xde, 0xc7, 0x30, 0x31,
	0x3b, 0xbf, 0xa2, 0xc4, 0x1f, 0xfe, 0x65, 0x9c, 0xba, 0x00, 0x4a, 0xfc, 0xa7, 0x18, 0xc0, 0x71,
	0xbf, 0xee, 0x62, 0x7e, 0x7d, 0xcd, 0x84, 0x60, 0xf3, 0x94, 0x12, 0xff, 0xd9, 0xf9, 0xfb, 0x1f,
	0xef, 0x1e, 0xf9, 0x85, 0xce, 0xda, 0xa2, 0xa9, 0x96, 0xcb, 0xa1, 0x1d, 0x3e, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xe6, 0xcf, 0x39, 0xc4, 0x72, 0x02, 0x00, 0x00,
}
