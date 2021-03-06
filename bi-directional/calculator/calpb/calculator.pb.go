// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calpb/calculator.proto

package calpb

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

type Maxrequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Maxrequest) Reset()         { *m = Maxrequest{} }
func (m *Maxrequest) String() string { return proto.CompactTextString(m) }
func (*Maxrequest) ProtoMessage()    {}
func (*Maxrequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48639ba9291a8650, []int{0}
}

func (m *Maxrequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Maxrequest.Unmarshal(m, b)
}
func (m *Maxrequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Maxrequest.Marshal(b, m, deterministic)
}
func (m *Maxrequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Maxrequest.Merge(m, src)
}
func (m *Maxrequest) XXX_Size() int {
	return xxx_messageInfo_Maxrequest.Size(m)
}
func (m *Maxrequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Maxrequest.DiscardUnknown(m)
}

var xxx_messageInfo_Maxrequest proto.InternalMessageInfo

func (m *Maxrequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Maxresponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Maxresponse) Reset()         { *m = Maxresponse{} }
func (m *Maxresponse) String() string { return proto.CompactTextString(m) }
func (*Maxresponse) ProtoMessage()    {}
func (*Maxresponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_48639ba9291a8650, []int{1}
}

func (m *Maxresponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Maxresponse.Unmarshal(m, b)
}
func (m *Maxresponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Maxresponse.Marshal(b, m, deterministic)
}
func (m *Maxresponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Maxresponse.Merge(m, src)
}
func (m *Maxresponse) XXX_Size() int {
	return xxx_messageInfo_Maxresponse.Size(m)
}
func (m *Maxresponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Maxresponse.DiscardUnknown(m)
}

var xxx_messageInfo_Maxresponse proto.InternalMessageInfo

func (m *Maxresponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Maxrequest)(nil), "calpb.maxrequest")
	proto.RegisterType((*Maxresponse)(nil), "calpb.maxresponse")
}

func init() { proto.RegisterFile("calculator/calpb/calculator.proto", fileDescriptor_48639ba9291a8650) }

var fileDescriptor_48639ba9291a8650 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x4f, 0x4e, 0xcc, 0x29, 0x48, 0xd2, 0x47, 0x08, 0xe8,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xb1, 0x82, 0xc5, 0x95, 0x54, 0xb8, 0xb8, 0x72, 0x13, 0x2b,
	0x8a, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0xf2, 0x4a, 0x73, 0x93, 0x52,
	0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x25, 0x55, 0x2e, 0x6e, 0xb0, 0xaa,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0xb2, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x98, 0x32,
	0x08, 0xcf, 0xc8, 0x92, 0x8b, 0xd9, 0x37, 0xb1, 0x42, 0xc8, 0x08, 0x42, 0x09, 0xea, 0x81, 0xad,
	0xd0, 0x43, 0x98, 0x2f, 0x25, 0x84, 0x2c, 0x04, 0x31, 0x4c, 0x89, 0x41, 0x83, 0xd1, 0x80, 0xd1,
	0x89, 0x3d, 0x0a, 0xe2, 0xa0, 0x24, 0x36, 0xb0, 0xf3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0x53, 0x45, 0x2a, 0xc3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MaxClient is the client API for Max service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaxClient interface {
	Max(ctx context.Context, opts ...grpc.CallOption) (Max_MaxClient, error)
}

type maxClient struct {
	cc *grpc.ClientConn
}

func NewMaxClient(cc *grpc.ClientConn) MaxClient {
	return &maxClient{cc}
}

func (c *maxClient) Max(ctx context.Context, opts ...grpc.CallOption) (Max_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Max_serviceDesc.Streams[0], "/calpb.Max/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &maxMaxClient{stream}
	return x, nil
}

type Max_MaxClient interface {
	Send(*Maxrequest) error
	Recv() (*Maxresponse, error)
	grpc.ClientStream
}

type maxMaxClient struct {
	grpc.ClientStream
}

func (x *maxMaxClient) Send(m *Maxrequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maxMaxClient) Recv() (*Maxresponse, error) {
	m := new(Maxresponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxServer is the server API for Max service.
type MaxServer interface {
	Max(Max_MaxServer) error
}

// UnimplementedMaxServer can be embedded to have forward compatible implementations.
type UnimplementedMaxServer struct {
}

func (*UnimplementedMaxServer) Max(srv Max_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}

func RegisterMaxServer(s *grpc.Server, srv MaxServer) {
	s.RegisterService(&_Max_serviceDesc, srv)
}

func _Max_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaxServer).Max(&maxMaxServer{stream})
}

type Max_MaxServer interface {
	Send(*Maxresponse) error
	Recv() (*Maxrequest, error)
	grpc.ServerStream
}

type maxMaxServer struct {
	grpc.ServerStream
}

func (x *maxMaxServer) Send(m *Maxresponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maxMaxServer) Recv() (*Maxrequest, error) {
	m := new(Maxrequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Max_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calpb.Max",
	HandlerType: (*MaxServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Max",
			Handler:       _Max_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calpb/calculator.proto",
}
