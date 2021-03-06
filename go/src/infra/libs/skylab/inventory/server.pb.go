// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

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

// NEXT TAG: 4
type Server_Status int32

const (
	Server_STATUS_INVALID         Server_Status = 0
	Server_STATUS_PRIMARY         Server_Status = 1
	Server_STATUS_BACKUP          Server_Status = 2
	Server_STATUS_REPAIR_REQUIRED Server_Status = 3
)

var Server_Status_name = map[int32]string{
	0: "STATUS_INVALID",
	1: "STATUS_PRIMARY",
	2: "STATUS_BACKUP",
	3: "STATUS_REPAIR_REQUIRED",
}

var Server_Status_value = map[string]int32{
	"STATUS_INVALID":         0,
	"STATUS_PRIMARY":         1,
	"STATUS_BACKUP":          2,
	"STATUS_REPAIR_REQUIRED": 3,
}

func (x Server_Status) Enum() *Server_Status {
	p := new(Server_Status)
	*p = x
	return p
}

func (x Server_Status) String() string {
	return proto.EnumName(Server_Status_name, int32(x))
}

func (x *Server_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Server_Status_value, data, "Server_Status")
	if err != nil {
		return err
	}
	*x = Server_Status(value)
	return nil
}

func (Server_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0, 0}
}

// NEXT TAG: 14
// Note: Update ROLE_MAP in skylab_inventory/translation_utils.py accordingly
type Server_Role int32

const (
	Server_ROLE_INVALID            Server_Role = 0
	Server_ROLE_AFE                Server_Role = 1
	Server_ROLE_CRASH_SERVER       Server_Role = 2
	Server_ROLE_DATABASE           Server_Role = 3
	Server_ROLE_DATABASE_SLAVE     Server_Role = 4
	Server_ROLE_DEVSERVER          Server_Role = 5
	Server_ROLE_DRONE              Server_Role = 6
	Server_ROLE_GOLO_PROXY         Server_Role = 7
	Server_ROLE_HOST_SCHEDULER     Server_Role = 8
	Server_ROLE_SCHEDULER          Server_Role = 9
	Server_ROLE_SENTINEL           Server_Role = 10
	Server_ROLE_SHARD              Server_Role = 11
	Server_ROLE_SUITE_SCHEDULER    Server_Role = 12
	Server_ROLE_SKYLAB_DRONE       Server_Role = 13
	Server_ROLE_SKYLAB_SUITE_PROXY Server_Role = 14
	Server_ROLE_RPMSERVER          Server_Role = 15
)

var Server_Role_name = map[int32]string{
	0:  "ROLE_INVALID",
	1:  "ROLE_AFE",
	2:  "ROLE_CRASH_SERVER",
	3:  "ROLE_DATABASE",
	4:  "ROLE_DATABASE_SLAVE",
	5:  "ROLE_DEVSERVER",
	6:  "ROLE_DRONE",
	7:  "ROLE_GOLO_PROXY",
	8:  "ROLE_HOST_SCHEDULER",
	9:  "ROLE_SCHEDULER",
	10: "ROLE_SENTINEL",
	11: "ROLE_SHARD",
	12: "ROLE_SUITE_SCHEDULER",
	13: "ROLE_SKYLAB_DRONE",
	14: "ROLE_SKYLAB_SUITE_PROXY",
	15: "ROLE_RPMSERVER",
}

var Server_Role_value = map[string]int32{
	"ROLE_INVALID":            0,
	"ROLE_AFE":                1,
	"ROLE_CRASH_SERVER":       2,
	"ROLE_DATABASE":           3,
	"ROLE_DATABASE_SLAVE":     4,
	"ROLE_DEVSERVER":          5,
	"ROLE_DRONE":              6,
	"ROLE_GOLO_PROXY":         7,
	"ROLE_HOST_SCHEDULER":     8,
	"ROLE_SCHEDULER":          9,
	"ROLE_SENTINEL":           10,
	"ROLE_SHARD":              11,
	"ROLE_SUITE_SCHEDULER":    12,
	"ROLE_SKYLAB_DRONE":       13,
	"ROLE_SKYLAB_SUITE_PROXY": 14,
	"ROLE_RPMSERVER":          15,
}

func (x Server_Role) Enum() *Server_Role {
	p := new(Server_Role)
	*p = x
	return p
}

func (x Server_Role) String() string {
	return proto.EnumName(Server_Role_name, int32(x))
}

func (x *Server_Role) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Server_Role_value, data, "Server_Role")
	if err != nil {
		return err
	}
	*x = Server_Role(value)
	return nil
}

func (Server_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0, 1}
}

// NEXT TAG: 9
type Server struct {
	Hostname    *string        `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	Cname       *string        `protobuf:"bytes,2,opt,name=cname" json:"cname,omitempty"`
	Environment *Environment   `protobuf:"varint,3,req,name=environment,enum=chrome.chromeos_infra.skylab.proto.inventory.Environment" json:"environment,omitempty"`
	Status      *Server_Status `protobuf:"varint,4,req,name=status,enum=chrome.chromeos_infra.skylab.proto.inventory.Server_Status" json:"status,omitempty"`
	Roles       []Server_Role  `protobuf:"varint,5,rep,name=roles,enum=chrome.chromeos_infra.skylab.proto.inventory.Server_Role" json:"roles,omitempty"`
	Attributes  *Attributes    `protobuf:"bytes,6,opt,name=attributes" json:"attributes,omitempty"`
	Notes       *string        `protobuf:"bytes,7,opt,name=notes" json:"notes,omitempty"`
	// List of dut_uids serviced by this server.
	// This can mean different things for different roles.
	//   skylab-drone: These are the DUTs owned by the drone.
	DutUids              []string `protobuf:"bytes,8,rep,name=dut_uids,json=dutUids" json:"dut_uids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Server) GetCname() string {
	if m != nil && m.Cname != nil {
		return *m.Cname
	}
	return ""
}

func (m *Server) GetEnvironment() Environment {
	if m != nil && m.Environment != nil {
		return *m.Environment
	}
	return Environment_ENVIRONMENT_INVALID
}

func (m *Server) GetStatus() Server_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Server_STATUS_INVALID
}

func (m *Server) GetRoles() []Server_Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Server) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Server) GetNotes() string {
	if m != nil && m.Notes != nil {
		return *m.Notes
	}
	return ""
}

func (m *Server) GetDutUids() []string {
	if m != nil {
		return m.DutUids
	}
	return nil
}

// NOTE: Please update SERVER_ATTRIBUTE_TYPE_MAP in
// skylab_inventory/translation_utils.py accordingly.
// NEXT TAG: 6
type Attributes struct {
	Ip                        *string  `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	MaxProcesses              *int32   `protobuf:"varint,2,opt,name=max_processes,json=maxProcesses" json:"max_processes,omitempty"`
	MysqlServerId             *string  `protobuf:"bytes,3,opt,name=mysql_server_id,json=mysqlServerId" json:"mysql_server_id,omitempty"`
	DevserverRestrictedSubnet *string  `protobuf:"bytes,4,opt,name=devserver_restricted_subnet,json=devserverRestrictedSubnet" json:"devserver_restricted_subnet,omitempty"`
	DevserverPort             *int32   `protobuf:"varint,5,opt,name=devserver_port,json=devserverPort" json:"devserver_port,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Attributes) Reset()         { *m = Attributes{} }
func (m *Attributes) String() string { return proto.CompactTextString(m) }
func (*Attributes) ProtoMessage()    {}
func (*Attributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *Attributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attributes.Unmarshal(m, b)
}
func (m *Attributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attributes.Marshal(b, m, deterministic)
}
func (m *Attributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attributes.Merge(m, src)
}
func (m *Attributes) XXX_Size() int {
	return xxx_messageInfo_Attributes.Size(m)
}
func (m *Attributes) XXX_DiscardUnknown() {
	xxx_messageInfo_Attributes.DiscardUnknown(m)
}

var xxx_messageInfo_Attributes proto.InternalMessageInfo

func (m *Attributes) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Attributes) GetMaxProcesses() int32 {
	if m != nil && m.MaxProcesses != nil {
		return *m.MaxProcesses
	}
	return 0
}

func (m *Attributes) GetMysqlServerId() string {
	if m != nil && m.MysqlServerId != nil {
		return *m.MysqlServerId
	}
	return ""
}

func (m *Attributes) GetDevserverRestrictedSubnet() string {
	if m != nil && m.DevserverRestrictedSubnet != nil {
		return *m.DevserverRestrictedSubnet
	}
	return ""
}

func (m *Attributes) GetDevserverPort() int32 {
	if m != nil && m.DevserverPort != nil {
		return *m.DevserverPort
	}
	return 0
}

func init() {
	proto.RegisterEnum("chrome.chromeos_infra.skylab.proto.inventory.Server_Status", Server_Status_name, Server_Status_value)
	proto.RegisterEnum("chrome.chromeos_infra.skylab.proto.inventory.Server_Role", Server_Role_name, Server_Role_value)
	proto.RegisterType((*Server)(nil), "chrome.chromeos_infra.skylab.proto.inventory.Server")
	proto.RegisterType((*Attributes)(nil), "chrome.chromeos_infra.skylab.proto.inventory.Attributes")
}

func init() {
	proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7)
}

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6f, 0xda, 0x3e,
	0x18, 0xc6, 0xbf, 0x84, 0xdf, 0x2f, 0x3f, 0x9a, 0xba, 0xfd, 0xae, 0x69, 0x7b, 0x41, 0x4c, 0x9b,
	0x38, 0x4c, 0x1c, 0x7a, 0xda, 0x34, 0x69, 0x52, 0x28, 0xde, 0x88, 0x4a, 0x81, 0xd9, 0x80, 0xda,
	0xed, 0x10, 0xa5, 0xc4, 0x53, 0xa3, 0x91, 0x98, 0xd9, 0x06, 0xb5, 0xb7, 0xfd, 0x89, 0x3b, 0xec,
	0x0f, 0x9a, 0xb0, 0xb3, 0xc0, 0x8e, 0x9c, 0xa2, 0xf7, 0xf1, 0xfb, 0x7c, 0xfc, 0xbc, 0x8e, 0x13,
	0xa8, 0x4b, 0x26, 0x36, 0x4c, 0x74, 0x57, 0x82, 0x2b, 0x8e, 0xde, 0x2c, 0x1e, 0x05, 0x8f, 0x59,
	0xd7, 0x3c, 0xb8, 0xf4, 0xa3, 0xe4, 0x9b, 0x08, 0xba, 0xf2, 0xfb, 0xf3, 0x32, 0x78, 0x30, 0x3d,
	0xdd, 0x28, 0xd9, 0xb0, 0x44, 0x71, 0xf1, 0x7c, 0x51, 0x5f, 0xf0, 0x38, 0xe6, 0x89, 0xd1, 0xdb,
	0x3f, 0xcb, 0x50, 0xa2, 0x1a, 0x86, 0x2e, 0xa0, 0xf2, 0xc8, 0xa5, 0x4a, 0x82, 0x98, 0x39, 0xb9,
	0x96, 0xd5, 0xa9, 0x92, 0xac, 0x46, 0xa7, 0x50, 0x5c, 0xe8, 0x05, 0xab, 0x95, 0xeb, 0x54, 0x89,
	0x29, 0xd0, 0x57, 0xa8, 0xb1, 0x64, 0x13, 0x09, 0x9e, 0xc4, 0x2c, 0x51, 0x4e, 0xbe, 0x65, 0x75,
	0x9a, 0x57, 0xef, 0xba, 0x87, 0xc4, 0xe9, 0xe2, 0x1d, 0x80, 0xec, 0xd3, 0x10, 0x85, 0x92, 0x54,
	0x81, 0x5a, 0x4b, 0xa7, 0xa0, 0xb9, 0xef, 0x0f, 0xe3, 0x9a, 0xa1, 0xba, 0x54, 0x23, 0x48, 0x8a,
	0x42, 0x63, 0x28, 0x0a, 0xbe, 0x64, 0xd2, 0x29, 0xb6, 0xf2, 0x87, 0x67, 0x4d, 0x99, 0x84, 0x2f,
	0x19, 0x31, 0x1c, 0x74, 0x07, 0x10, 0x28, 0x25, 0xa2, 0x87, 0xb5, 0x62, 0xd2, 0x29, 0xb5, 0x72,
	0x9d, 0xda, 0xd5, 0xdb, 0xc3, 0xa8, 0x6e, 0xe6, 0x27, 0x7b, 0xac, 0xed, 0x91, 0x27, 0x7c, 0x0b,
	0x2d, 0x9b, 0x23, 0xd7, 0x05, 0x3a, 0x87, 0x4a, 0xb8, 0x56, 0xfe, 0x3a, 0x0a, 0xa5, 0x53, 0x69,
	0xe5, 0x3b, 0x55, 0x52, 0x0e, 0xd7, 0x6a, 0x16, 0x85, 0xb2, 0xed, 0x43, 0xc9, 0x4c, 0x8b, 0x10,
	0x34, 0xe9, 0xd4, 0x9d, 0xce, 0xa8, 0xef, 0x8d, 0xe6, 0xee, 0xd0, 0xeb, 0xdb, 0xff, 0xed, 0x69,
	0x13, 0xe2, 0xdd, 0xba, 0xe4, 0xde, 0xce, 0xa1, 0x63, 0x68, 0xa4, 0x5a, 0xcf, 0xbd, 0xbe, 0x99,
	0x4d, 0x6c, 0x0b, 0x5d, 0xc0, 0x8b, 0x54, 0x22, 0x78, 0xe2, 0x7a, 0xc4, 0x27, 0xf8, 0xf3, 0xcc,
	0x23, 0xb8, 0x6f, 0xe7, 0xdb, 0xbf, 0x2d, 0x28, 0x6c, 0x67, 0x47, 0x36, 0xd4, 0xc9, 0x78, 0x88,
	0xf7, 0xe8, 0x75, 0xa8, 0x68, 0xc5, 0xfd, 0x88, 0xed, 0x1c, 0xfa, 0x1f, 0x8e, 0x75, 0x75, 0x4d,
	0x5c, 0x3a, 0xf0, 0x29, 0x26, 0x73, 0x4c, 0x6c, 0x6b, 0xbb, 0x9d, 0x96, 0xfb, 0xee, 0xd4, 0xed,
	0xb9, 0x14, 0xdb, 0x79, 0x74, 0x06, 0x27, 0xff, 0x48, 0x3e, 0x1d, 0xba, 0x73, 0x6c, 0x17, 0xb6,
	0x71, 0xcd, 0x02, 0x9e, 0xa7, 0xfe, 0x22, 0x6a, 0x02, 0x18, 0x8d, 0x8c, 0x47, 0xd8, 0x2e, 0xa1,
	0x13, 0x38, 0xd2, 0xf5, 0xa7, 0xf1, 0x70, 0xec, 0x4f, 0xc8, 0xf8, 0xee, 0xde, 0x2e, 0x67, 0xc4,
	0xc1, 0x98, 0x4e, 0x7d, 0x7a, 0x3d, 0xc0, 0xfd, 0xd9, 0x10, 0x13, 0xbb, 0x92, 0x11, 0x77, 0x5a,
	0x35, 0x4b, 0x44, 0xf1, 0x68, 0xea, 0x8d, 0xf0, 0xd0, 0x86, 0x6c, 0x13, 0x3a, 0x70, 0x49, 0xdf,
	0xae, 0x21, 0x07, 0x4e, 0x4d, 0x3d, 0xf3, 0xa6, 0xfb, 0xe6, 0x7a, 0x36, 0x25, 0xbd, 0xb9, 0x1f,
	0xba, 0xbd, 0x34, 0x55, 0x03, 0x5d, 0xc2, 0xd9, 0xbe, 0x6c, 0x7c, 0x26, 0x5d, 0x33, 0x0b, 0x41,
	0x26, 0xb7, 0xe9, 0x58, 0x47, 0xed, 0x5f, 0x39, 0x80, 0xdd, 0x1d, 0x40, 0x4d, 0xb0, 0xa2, 0x95,
	0x93, 0xd3, 0x2f, 0xdd, 0x8a, 0x56, 0xe8, 0x25, 0x34, 0xe2, 0xe0, 0xc9, 0x5f, 0x09, 0xbe, 0x60,
	0x52, 0x32, 0xa9, 0x3f, 0xc1, 0x22, 0xa9, 0xc7, 0xc1, 0xd3, 0xe4, 0xaf, 0x86, 0x5e, 0xc3, 0x51,
	0xfc, 0x2c, 0x7f, 0x2c, 0x7d, 0xf3, 0x63, 0xf0, 0xa3, 0xd0, 0xc9, 0x6b, 0x42, 0x43, 0xcb, 0xe6,
	0xe2, 0x7a, 0x21, 0xfa, 0x00, 0x97, 0x21, 0xdb, 0xa4, 0x4d, 0x82, 0x49, 0x25, 0xa2, 0x85, 0x62,
	0xa1, 0x2f, 0xd7, 0x0f, 0x09, 0x53, 0x4e, 0x41, 0x7b, 0xce, 0xb3, 0x16, 0x92, 0x75, 0x50, 0xdd,
	0x80, 0x5e, 0x41, 0x73, 0xe7, 0x5f, 0x71, 0xa1, 0x9c, 0xa2, 0x4e, 0xd3, 0xc8, 0xd4, 0x09, 0x17,
	0xaa, 0x57, 0xfb, 0x52, 0xcd, 0xee, 0xf7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xa1, 0x0d,
	0x1c, 0xad, 0x04, 0x00, 0x00,
}
