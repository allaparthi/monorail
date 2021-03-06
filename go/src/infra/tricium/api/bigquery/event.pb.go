// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/tricium/api/bigquery/event.proto

package apibq

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	v1 "infra/tricium/api/v1"
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

// Types of event.
type FeedbackEvent_Type int32

const (
	FeedbackEvent_NONE         FeedbackEvent_Type = 0
	FeedbackEvent_NOT_USEFUL   FeedbackEvent_Type = 1
	FeedbackEvent_COMMENT_POST FeedbackEvent_Type = 2
)

var FeedbackEvent_Type_name = map[int32]string{
	0: "NONE",
	1: "NOT_USEFUL",
	2: "COMMENT_POST",
}

var FeedbackEvent_Type_value = map[string]int32{
	"NONE":         0,
	"NOT_USEFUL":   1,
	"COMMENT_POST": 2,
}

func (x FeedbackEvent_Type) String() string {
	return proto.EnumName(FeedbackEvent_Type_name, int32(x))
}

func (FeedbackEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eae2733f3e10ceaa, []int{0, 0}
}

// FeedbackEvent represents one event such as sending comments or a "not
// useful" click.
//
// The purpose of recording these events is to be able to track the feedback on
// comments -- for example, what's the proportion "not useful" clicks to
// comments produced for each analyzer or category?
type FeedbackEvent struct {
	// Type of event.
	Type FeedbackEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=apibq.FeedbackEvent_Type" json:"type,omitempty"`
	// Time when the event occurred.
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// Related comments. For comment post events, this will be the comments
	// sent; for a "not useful" click, this would be the one related comment.
	Comments             []*v1.Data_Comment `protobuf:"bytes,3,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FeedbackEvent) Reset()         { *m = FeedbackEvent{} }
func (m *FeedbackEvent) String() string { return proto.CompactTextString(m) }
func (*FeedbackEvent) ProtoMessage()    {}
func (*FeedbackEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae2733f3e10ceaa, []int{0}
}

func (m *FeedbackEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbackEvent.Unmarshal(m, b)
}
func (m *FeedbackEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbackEvent.Marshal(b, m, deterministic)
}
func (m *FeedbackEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbackEvent.Merge(m, src)
}
func (m *FeedbackEvent) XXX_Size() int {
	return xxx_messageInfo_FeedbackEvent.Size(m)
}
func (m *FeedbackEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbackEvent.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbackEvent proto.InternalMessageInfo

func (m *FeedbackEvent) GetType() FeedbackEvent_Type {
	if m != nil {
		return m.Type
	}
	return FeedbackEvent_NONE
}

func (m *FeedbackEvent) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *FeedbackEvent) GetComments() []*v1.Data_Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func init() {
	proto.RegisterEnum("apibq.FeedbackEvent_Type", FeedbackEvent_Type_name, FeedbackEvent_Type_value)
	proto.RegisterType((*FeedbackEvent)(nil), "apibq.FeedbackEvent")
}

func init() {
	proto.RegisterFile("infra/tricium/api/bigquery/event.proto", fileDescriptor_eae2733f3e10ceaa)
}

var fileDescriptor_eae2733f3e10ceaa = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xce, 0xc1, 0x6a, 0xab, 0x40,
	0x14, 0xc6, 0xf1, 0x6b, 0xe2, 0x2d, 0xe1, 0xa4, 0x0d, 0x32, 0x50, 0xb0, 0x6e, 0x22, 0x59, 0x14,
	0x37, 0x3d, 0x43, 0xec, 0x23, 0xa4, 0x66, 0xd5, 0x68, 0x31, 0x66, 0x1d, 0x46, 0x73, 0x22, 0x43,
	0xab, 0x4e, 0xcc, 0x18, 0xf0, 0x41, 0xfb, 0x3e, 0x45, 0x6d, 0x0a, 0xa5, 0xcb, 0x61, 0x7e, 0x9c,
	0xff, 0x07, 0x8f, 0xb2, 0x3c, 0xd6, 0x82, 0xeb, 0x5a, 0x66, 0xb2, 0x29, 0xb8, 0x50, 0x92, 0xa7,
	0x32, 0x3f, 0x35, 0x54, 0xb7, 0x9c, 0x2e, 0x54, 0x6a, 0x54, 0x75, 0xa5, 0x2b, 0xf6, 0x5f, 0x28,
	0x99, 0x9e, 0x9c, 0xf9, 0x5f, 0x7e, 0x59, 0xf2, 0x83, 0xd0, 0x62, 0x70, 0xce, 0x3c, 0xaf, 0xaa,
	0xfc, 0x83, 0x78, 0xff, 0x4a, 0x9b, 0x23, 0xd7, 0xb2, 0xa0, 0xb3, 0x16, 0x85, 0x1a, 0xc0, 0xe2,
	0xd3, 0x80, 0xbb, 0x35, 0xd1, 0x21, 0x15, 0xd9, 0x7b, 0xd0, 0x05, 0xd8, 0x13, 0x98, 0xba, 0x55,
	0x64, 0x1b, 0xae, 0xe1, 0xcd, 0xfc, 0x07, 0xec, 0x4b, 0xf8, 0xcb, 0x60, 0xd2, 0x2a, 0x8a, 0x7b,
	0xc6, 0x10, 0xcc, 0xee, 0xa6, 0x3d, 0x72, 0x0d, 0x6f, 0xea, 0x3b, 0x38, 0x04, 0xf1, 0x1a, 0xc4,
	0xe4, 0x1a, 0x8c, 0x7b, 0xc7, 0x96, 0x30, 0xc9, 0xaa, 0xa2, 0xa0, 0x52, 0x9f, 0xed, 0xb1, 0x3b,
	0xf6, 0xa6, 0xfe, 0x3d, 0x7e, 0xef, 0xc7, 0x97, 0x6e, 0xf8, 0x6a, 0xf8, 0x8d, 0x7f, 0xd8, 0xc2,
	0x07, 0xb3, 0x0b, 0xb2, 0x09, 0x98, 0x61, 0x14, 0x06, 0xd6, 0x3f, 0x36, 0x03, 0x08, 0xa3, 0x64,
	0xbf, 0xdb, 0x06, 0xeb, 0xdd, 0xab, 0x65, 0x30, 0x0b, 0x6e, 0x57, 0xd1, 0x66, 0x13, 0x84, 0xc9,
	0xfe, 0x2d, 0xda, 0x26, 0xd6, 0x28, 0xbd, 0xe9, 0x07, 0x3c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x77, 0x1a, 0x66, 0xb6, 0x51, 0x01, 0x00, 0x00,
}
