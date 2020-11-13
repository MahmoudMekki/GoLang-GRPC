// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package greetpb

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greetpb.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greetpb.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greetpb.GreetResponse")
}

func init() { proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871) }

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x1d,
	0x2a, 0xa8, 0xe4, 0xc2, 0xc5, 0xe1, 0x0e, 0x62, 0x66, 0xe6, 0xa5, 0x0b, 0xc9, 0x70, 0x71, 0xba,
	0x65, 0x16, 0x15, 0x97, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21,
	0x04, 0x84, 0xa4, 0xb8, 0x38, 0x7c, 0x12, 0xa1, 0x92, 0x4c, 0x60, 0x49, 0x38, 0x5f, 0xc9, 0x96,
	0x8b, 0x07, 0x6c, 0x4a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x2e, 0x17, 0x47, 0x3a,
	0xd4, 0x54, 0xb0, 0x41, 0xdc, 0x46, 0x82, 0x7a, 0x50, 0x1b, 0xf5, 0x60, 0xd6, 0x05, 0xc1, 0x95,
	0x28, 0xa9, 0x73, 0xf1, 0x42, 0xb5, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1,
	0x05, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x40, 0x9d, 0x01, 0xe5, 0x19, 0x79, 0x70, 0xb1, 0x82, 0x15,
	0x0a, 0xd9, 0x23, 0x39, 0x5b, 0x14, 0xd5, 0x68, 0xa8, 0x1b, 0xa4, 0xc4, 0xd0, 0x85, 0x21, 0x66,
	0x2b, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x3a, 0x71, 0x46, 0xc1, 0x82, 0x20, 0x89, 0x0d, 0x1c, 0x24,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x4c, 0x25, 0x99, 0x2f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetClient is the client API for Greet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetClient interface {
	Greeting(ctx context.Context, opts ...grpc.CallOption) (Greet_GreetingClient, error)
}

type greetClient struct {
	cc *grpc.ClientConn
}

func NewGreetClient(cc *grpc.ClientConn) GreetClient {
	return &greetClient{cc}
}

func (c *greetClient) Greeting(ctx context.Context, opts ...grpc.CallOption) (Greet_GreetingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greet_serviceDesc.Streams[0], "/greetpb.Greet/Greeting", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetGreetingClient{stream}
	return x, nil
}

type Greet_GreetingClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetGreetingClient struct {
	grpc.ClientStream
}

func (x *greetGreetingClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetGreetingClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServer is the server API for Greet service.
type GreetServer interface {
	Greeting(Greet_GreetingServer) error
}

// UnimplementedGreetServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServer struct {
}

func (*UnimplementedGreetServer) Greeting(srv Greet_GreetingServer) error {
	return status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}

func RegisterGreetServer(s *grpc.Server, srv GreetServer) {
	s.RegisterService(&_Greet_serviceDesc, srv)
}

func _Greet_Greeting_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServer).Greeting(&greetGreetingServer{stream})
}

type Greet_GreetingServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetGreetingServer struct {
	grpc.ServerStream
}

func (x *greetGreetingServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetGreetingServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greetpb.Greet",
	HandlerType: (*GreetServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greeting",
			Handler:       _Greet_Greeting_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}