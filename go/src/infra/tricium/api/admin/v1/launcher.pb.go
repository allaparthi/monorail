// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/tricium/api/admin/v1/launcher.proto

package admin

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// LaunchRequest contains the details needed to launch a workflow for an
// analysis request.
type LaunchRequest struct {
	// The run ID created by a Tricium.Analyze call.
	RunId int64 `protobuf:"varint,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// The name of the project in luci-config.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The full URL of the git repo used for this run.
	GitUrl string `protobuf:"bytes,3,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
	// A git commit-ish, such as a Gerrit revision ref name like
	// "refs/changes/34/1234/1", or any other ref name or commit hash.
	// This is used in the GitFileDetails data type used to pull files.
	GitRef string `protobuf:"bytes,4,opt,name=git_ref,json=gitRef,proto3" json:"git_ref,omitempty"`
	// File metadata from the root of the Git repository.
	Files []*v1.Data_File `protobuf:"bytes,5,rep,name=files,proto3" json:"files,omitempty"`
	// Commit message text from Gerrit if available.
	//
	// If this launch request is not for Gerrit, this could be an empty string.
	CommitMessage        string   `protobuf:"bytes,6,opt,name=commit_message,json=commitMessage,proto3" json:"commit_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchRequest) Reset()         { *m = LaunchRequest{} }
func (m *LaunchRequest) String() string { return proto.CompactTextString(m) }
func (*LaunchRequest) ProtoMessage()    {}
func (*LaunchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aad561e8ba77bd8, []int{0}
}

func (m *LaunchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchRequest.Unmarshal(m, b)
}
func (m *LaunchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchRequest.Marshal(b, m, deterministic)
}
func (m *LaunchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchRequest.Merge(m, src)
}
func (m *LaunchRequest) XXX_Size() int {
	return xxx_messageInfo_LaunchRequest.Size(m)
}
func (m *LaunchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchRequest proto.InternalMessageInfo

func (m *LaunchRequest) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

func (m *LaunchRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *LaunchRequest) GetGitUrl() string {
	if m != nil {
		return m.GitUrl
	}
	return ""
}

func (m *LaunchRequest) GetGitRef() string {
	if m != nil {
		return m.GitRef
	}
	return ""
}

func (m *LaunchRequest) GetFiles() []*v1.Data_File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *LaunchRequest) GetCommitMessage() string {
	if m != nil {
		return m.CommitMessage
	}
	return ""
}

type LaunchResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchResponse) Reset()         { *m = LaunchResponse{} }
func (m *LaunchResponse) String() string { return proto.CompactTextString(m) }
func (*LaunchResponse) ProtoMessage()    {}
func (*LaunchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aad561e8ba77bd8, []int{1}
}

func (m *LaunchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchResponse.Unmarshal(m, b)
}
func (m *LaunchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchResponse.Marshal(b, m, deterministic)
}
func (m *LaunchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchResponse.Merge(m, src)
}
func (m *LaunchResponse) XXX_Size() int {
	return xxx_messageInfo_LaunchResponse.Size(m)
}
func (m *LaunchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LaunchRequest)(nil), "admin.LaunchRequest")
	proto.RegisterType((*LaunchResponse)(nil), "admin.LaunchResponse")
}

func init() {
	proto.RegisterFile("infra/tricium/api/admin/v1/launcher.proto", fileDescriptor_1aad561e8ba77bd8)
}

var fileDescriptor_1aad561e8ba77bd8 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xa9, 0xb5, 0x9d, 0x46, 0x36, 0x24, 0x38, 0x0c, 0x7b, 0xb1, 0x0c, 0x84, 0xfa, 0xd2,
	0xb0, 0x89, 0x1f, 0x40, 0x10, 0x41, 0x98, 0x2f, 0x05, 0x9f, 0x4b, 0x6c, 0x6f, 0xeb, 0x95, 0x36,
	0xa9, 0xf9, 0xb3, 0x2f, 0xe8, 0x17, 0x13, 0x93, 0x55, 0x50, 0xdf, 0x92, 0xf3, 0xbb, 0xf7, 0x72,
	0xce, 0x21, 0x37, 0x28, 0x5b, 0x2d, 0xb8, 0xd5, 0x58, 0xa3, 0x1b, 0xb8, 0x18, 0x91, 0x8b, 0x66,
	0x40, 0xc9, 0xf7, 0x1b, 0xde, 0x0b, 0x27, 0xeb, 0x37, 0xd0, 0xc5, 0xa8, 0x95, 0x55, 0x34, 0xf1,
	0x60, 0x75, 0xf5, 0x7f, 0x63, 0xbf, 0xe1, 0x8d, 0xb0, 0x22, 0xcc, 0xad, 0x3f, 0x23, 0x32, 0xdf,
	0xf9, 0xd5, 0x12, 0x3e, 0x1c, 0x18, 0x4b, 0x97, 0x24, 0xd5, 0x4e, 0x56, 0xd8, 0xb0, 0x28, 0x8b,
	0xf2, 0xb8, 0x4c, 0xb4, 0x93, 0x4f, 0x0d, 0x65, 0x64, 0x36, 0x6a, 0xf5, 0x0e, 0xb5, 0x65, 0x47,
	0x59, 0x94, 0x9f, 0x96, 0xd3, 0x97, 0x5e, 0x92, 0x59, 0x87, 0xb6, 0x72, 0xba, 0x67, 0xb1, 0x27,
	0x69, 0x87, 0xf6, 0x45, 0xf7, 0x13, 0xd0, 0xd0, 0xb2, 0xe3, 0x1f, 0x50, 0x42, 0x4b, 0x73, 0x92,
	0xb4, 0xd8, 0x83, 0x61, 0x49, 0x16, 0xe7, 0x67, 0x5b, 0x5a, 0x1c, 0xfc, 0x15, 0x0f, 0xdf, 0xc6,
	0x1e, 0xb1, 0x87, 0x32, 0x0c, 0xd0, 0x6b, 0xb2, 0xa8, 0xd5, 0x30, 0xa0, 0xad, 0x06, 0x30, 0x46,
	0x74, 0xc0, 0x52, 0x7f, 0x69, 0x1e, 0xd4, 0xe7, 0x20, 0xae, 0xcf, 0xc9, 0x62, 0x0a, 0x61, 0x46,
	0x25, 0x0d, 0x6c, 0xef, 0xc9, 0xc9, 0xee, 0xd0, 0x08, 0xbd, 0x23, 0x69, 0x78, 0xd3, 0x8b, 0xc2,
	0xd7, 0x52, 0xfc, 0x4a, 0xbc, 0x5a, 0xfe, 0x51, 0xc3, 0x89, 0xd7, 0xd4, 0x37, 0x74, 0xfb, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0x1e, 0x41, 0x94, 0x76, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LauncherClient is the client API for Launcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LauncherClient interface {
	// Launch launches a workflow for provided request details.
	Launch(ctx context.Context, in *LaunchRequest, opts ...grpc.CallOption) (*LaunchResponse, error)
}
type launcherPRPCClient struct {
	client *prpc.Client
}

func NewLauncherPRPCClient(client *prpc.Client) LauncherClient {
	return &launcherPRPCClient{client}
}

func (c *launcherPRPCClient) Launch(ctx context.Context, in *LaunchRequest, opts ...grpc.CallOption) (*LaunchResponse, error) {
	out := new(LaunchResponse)
	err := c.client.Call(ctx, "admin.Launcher", "Launch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type launcherClient struct {
	cc grpc.ClientConnInterface
}

func NewLauncherClient(cc grpc.ClientConnInterface) LauncherClient {
	return &launcherClient{cc}
}

func (c *launcherClient) Launch(ctx context.Context, in *LaunchRequest, opts ...grpc.CallOption) (*LaunchResponse, error) {
	out := new(LaunchResponse)
	err := c.cc.Invoke(ctx, "/admin.Launcher/Launch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LauncherServer is the server API for Launcher service.
type LauncherServer interface {
	// Launch launches a workflow for provided request details.
	Launch(context.Context, *LaunchRequest) (*LaunchResponse, error)
}

// UnimplementedLauncherServer can be embedded to have forward compatible implementations.
type UnimplementedLauncherServer struct {
}

func (*UnimplementedLauncherServer) Launch(ctx context.Context, req *LaunchRequest) (*LaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Launch not implemented")
}

func RegisterLauncherServer(s prpc.Registrar, srv LauncherServer) {
	s.RegisterService(&_Launcher_serviceDesc, srv)
}

func _Launcher_Launch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Launch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Launcher/Launch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Launch(ctx, req.(*LaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Launcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Launcher",
	HandlerType: (*LauncherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Launch",
			Handler:    _Launcher_Launch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/tricium/api/admin/v1/launcher.proto",
}
