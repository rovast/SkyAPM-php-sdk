// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/ApplicationRegisterService.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Application struct {
	ApplicationCode      string   `protobuf:"bytes,1,opt,name=applicationCode,proto3" json:"applicationCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1c2b07c1f73c603, []int{0}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetApplicationCode() string {
	if m != nil {
		return m.ApplicationCode
	}
	return ""
}

type ApplicationMapping struct {
	Application          *KeyWithIntegerValue `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApplicationMapping) Reset()         { *m = ApplicationMapping{} }
func (m *ApplicationMapping) String() string { return proto.CompactTextString(m) }
func (*ApplicationMapping) ProtoMessage()    {}
func (*ApplicationMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1c2b07c1f73c603, []int{1}
}

func (m *ApplicationMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMapping.Unmarshal(m, b)
}
func (m *ApplicationMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMapping.Marshal(b, m, deterministic)
}
func (m *ApplicationMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMapping.Merge(m, src)
}
func (m *ApplicationMapping) XXX_Size() int {
	return xxx_messageInfo_ApplicationMapping.Size(m)
}
func (m *ApplicationMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMapping proto.InternalMessageInfo

func (m *ApplicationMapping) GetApplication() *KeyWithIntegerValue {
	if m != nil {
		return m.Application
	}
	return nil
}

func init() {
	proto.RegisterType((*Application)(nil), "Application")
	proto.RegisterType((*ApplicationMapping)(nil), "ApplicationMapping")
}

func init() {
	proto.RegisterFile("language-agent/ApplicationRegisterService.proto", fileDescriptor_e1c2b07c1f73c603)
}

var fileDescriptor_e1c2b07c1f73c603 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xad, 0x07, 0xc1, 0x54, 0x10, 0xa2, 0xa0, 0xf4, 0x24, 0xc5, 0x43, 0x2f, 0xbe, 0xea,
	0x04, 0xbd, 0x09, 0xea, 0x49, 0xfc, 0xc1, 0xe8, 0xc0, 0x81, 0x78, 0x79, 0xd6, 0x47, 0x16, 0x1a,
	0x93, 0x90, 0x65, 0x2b, 0xfd, 0x97, 0xfc, 0x2b, 0x65, 0x5d, 0xe7, 0x42, 0x75, 0xc7, 0x3c, 0x3e,
	0xdf, 0xf7, 0x4d, 0x3e, 0x61, 0xb9, 0x42, 0x2d, 0x66, 0x28, 0xe8, 0x0c, 0x05, 0x69, 0x9f, 0xdf,
	0x5a, 0xab, 0x64, 0x89, 0x5e, 0x1a, 0x5d, 0x90, 0x90, 0x53, 0x4f, 0x6e, 0x44, 0x6e, 0x2e, 0x4b,
	0x02, 0xeb, 0x8c, 0x37, 0x49, 0xd6, 0x0b, 0x3c, 0x52, 0x33, 0x96, 0x7e, 0xf2, 0xa0, 0x3d, 0x09,
	0x72, 0xaf, 0xa8, 0x66, 0x1d, 0x99, 0x5e, 0xb3, 0x38, 0xd8, 0xc6, 0x33, 0xb6, 0x8f, 0xeb, 0xe3,
	0xbd, 0xf9, 0xa4, 0xe3, 0xe8, 0x24, 0xca, 0x76, 0x8b, 0xfe, 0x38, 0x7d, 0x62, 0x3c, 0x08, 0x3e,
	0xa3, 0xb5, 0x52, 0x0b, 0x7e, 0xc5, 0xe2, 0x00, 0x6c, 0xb3, 0xf1, 0xe0, 0x10, 0xfe, 0xe9, 0x2f,
	0x42, 0x70, 0xf0, 0xce, 0x92, 0xcd, 0x8f, 0xe2, 0x37, 0xec, 0xa8, 0x57, 0xbf, 0x22, 0xf8, 0x1e,
	0x04, 0xb9, 0xe4, 0x00, 0xfe, 0xde, 0x29, 0xdd, 0xba, 0xab, 0xd9, 0xb9, 0x71, 0x02, 0xd0, 0x62,
	0x39, 0x21, 0x98, 0x56, 0x4d, 0x8d, 0xaa, 0x92, 0x7a, 0x31, 0xf9, 0x02, 0x4d, 0xbe, 0x36, 0xae,
	0x82, 0x95, 0x32, 0x68, 0x95, 0x0d, 0xa3, 0xb7, 0xd3, 0x35, 0x98, 0x77, 0xd0, 0xef, 0x47, 0xe4,
	0x4b, 0xaf, 0xf3, 0x8b, 0xef, 0xed, 0x64, 0x54, 0x35, 0xe3, 0x6e, 0xdf, 0xcb, 0x12, 0x1b, 0x2e,
	0xd4, 0x96, 0x46, 0x7d, 0xec, 0xb4, 0x92, 0x2f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x02, 0x5e,
	0x49, 0x1c, 0xc1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationRegisterServiceClient is the client API for ApplicationRegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationRegisterServiceClient interface {
	ApplicationCodeRegister(ctx context.Context, in *Application, opts ...grpc.CallOption) (*ApplicationMapping, error)
}

type applicationRegisterServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationRegisterServiceClient(cc *grpc.ClientConn) ApplicationRegisterServiceClient {
	return &applicationRegisterServiceClient{cc}
}

func (c *applicationRegisterServiceClient) ApplicationCodeRegister(ctx context.Context, in *Application, opts ...grpc.CallOption) (*ApplicationMapping, error) {
	out := new(ApplicationMapping)
	err := c.cc.Invoke(ctx, "/ApplicationRegisterService/applicationCodeRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationRegisterServiceServer is the server API for ApplicationRegisterService service.
type ApplicationRegisterServiceServer interface {
	ApplicationCodeRegister(context.Context, *Application) (*ApplicationMapping, error)
}

// UnimplementedApplicationRegisterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationRegisterServiceServer struct {
}

func (*UnimplementedApplicationRegisterServiceServer) ApplicationCodeRegister(ctx context.Context, req *Application) (*ApplicationMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationCodeRegister not implemented")
}

func RegisterApplicationRegisterServiceServer(s *grpc.Server, srv ApplicationRegisterServiceServer) {
	s.RegisterService(&_ApplicationRegisterService_serviceDesc, srv)
}

func _ApplicationRegisterService_ApplicationCodeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegisterServiceServer).ApplicationCodeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApplicationRegisterService/ApplicationCodeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegisterServiceServer).ApplicationCodeRegister(ctx, req.(*Application))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationRegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ApplicationRegisterService",
	HandlerType: (*ApplicationRegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "applicationCodeRegister",
			Handler:    _ApplicationRegisterService_ApplicationCodeRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/ApplicationRegisterService.proto",
}
