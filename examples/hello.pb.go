// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	// Types that are valid to be assigned to Arg:
	//	*HelloRequest_Name
	//	*HelloRequest_Age
	//	*HelloRequest_Blip
	Arg  isHelloRequest_Arg `protobuf_oneof:"arg"`
	Ints []int32            `protobuf:"varint,4,rep,packed,name=ints" json:"ints,omitempty"`
	Amap map[int32]int32    `protobuf:"bytes,5,rep,name=amap" json:"amap,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isHelloRequest_Arg interface {
	isHelloRequest_Arg()
}

type HelloRequest_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,oneof"`
}
type HelloRequest_Age struct {
	Age int32 `protobuf:"varint,2,opt,name=age,oneof"`
}
type HelloRequest_Blip struct {
	Blip *HelloReply `protobuf:"bytes,3,opt,name=blip,oneof"`
}

func (*HelloRequest_Name) isHelloRequest_Arg() {}
func (*HelloRequest_Age) isHelloRequest_Arg()  {}
func (*HelloRequest_Blip) isHelloRequest_Arg() {}

func (m *HelloRequest) GetArg() isHelloRequest_Arg {
	if m != nil {
		return m.Arg
	}
	return nil
}

func (m *HelloRequest) GetName() string {
	if x, ok := m.GetArg().(*HelloRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *HelloRequest) GetAge() int32 {
	if x, ok := m.GetArg().(*HelloRequest_Age); ok {
		return x.Age
	}
	return 0
}

func (m *HelloRequest) GetBlip() *HelloReply {
	if x, ok := m.GetArg().(*HelloRequest_Blip); ok {
		return x.Blip
	}
	return nil
}

func (m *HelloRequest) GetInts() []int32 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *HelloRequest) GetAmap() map[int32]int32 {
	if m != nil {
		return m.Amap
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HelloRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HelloRequest_OneofMarshaler, _HelloRequest_OneofUnmarshaler, _HelloRequest_OneofSizer, []interface{}{
		(*HelloRequest_Name)(nil),
		(*HelloRequest_Age)(nil),
		(*HelloRequest_Blip)(nil),
	}
}

func _HelloRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HelloRequest)
	// arg
	switch x := m.Arg.(type) {
	case *HelloRequest_Name:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *HelloRequest_Age:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Age))
	case *HelloRequest_Blip:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Blip); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HelloRequest.Arg has unexpected type %T", x)
	}
	return nil
}

func _HelloRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HelloRequest)
	switch tag {
	case 1: // arg.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Arg = &HelloRequest_Name{x}
		return true, err
	case 2: // arg.age
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Arg = &HelloRequest_Age{int32(x)}
		return true, err
	case 3: // arg.blip
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HelloReply)
		err := b.DecodeMessage(msg)
		m.Arg = &HelloRequest_Blip{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HelloRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HelloRequest)
	// arg
	switch x := m.Arg.(type) {
	case *HelloRequest_Name:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *HelloRequest_Age:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Age))
	case *HelloRequest_Blip:
		s := proto.Size(x.Blip)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Integer int32  `protobuf:"varint,2,opt,name=integer" json:"integer,omitempty"`
	Bloop   bool   `protobuf:"varint,3,opt,name=bloop" json:"bloop,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetInteger() int32 {
	if m != nil {
		return m.Integer
	}
	return 0
}

func (m *HelloReply) GetBloop() bool {
	if m != nil {
		return m.Bloop
	}
	return false
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xec, 0xc6, 0xb6, 0x93, 0x0a, 0x32, 0x08, 0x2e, 0x39, 0x48, 0xe8, 0xc5, 0x9c,
	0x02, 0xb6, 0x07, 0xc5, 0x9b, 0x05, 0xb1, 0xe7, 0xf5, 0x13, 0x6c, 0x61, 0x6c, 0x83, 0x9b, 0x3f,
	0x6e, 0xb6, 0x42, 0x6f, 0xfd, 0x52, 0x7e, 0x3f, 0xd9, 0x4d, 0xa2, 0xf1, 0x96, 0xf9, 0xbd, 0xe1,
	0xe5, 0xcd, 0x5b, 0x88, 0x0f, 0xa4, 0x75, 0x9d, 0x37, 0xa6, 0xb6, 0x75, 0xb2, 0xb0, 0x07, 0x53,
	0xbc, 0xdb, 0x6e, 0x5a, 0x9e, 0x43, 0x58, 0x6c, 0x9d, 0x2a, 0xe9, 0xf3, 0x48, 0xad, 0xc5, 0x04,
	0x78, 0xa5, 0x4a, 0x12, 0x41, 0x1a, 0x64, 0xf3, 0x0d, 0x3f, 0x7f, 0x8b, 0x60, 0x3b, 0x91, 0x9e,
	0xa1, 0x00, 0xa6, 0xf6, 0x24, 0xc2, 0x34, 0xc8, 0x22, 0x2f, 0x85, 0xdb, 0x89, 0x74, 0x08, 0xef,
	0x80, 0xef, 0x74, 0xd1, 0x08, 0x96, 0x06, 0x59, 0xbc, 0x8a, 0xf3, 0xde, 0xb2, 0xd1, 0x27, 0xbf,
	0xc7, 0x9c, 0x85, 0x5b, 0x40, 0x01, 0xbc, 0xa8, 0x6c, 0x2b, 0x78, 0xca, 0x06, 0x0f, 0xe9, 0x09,
	0xde, 0x03, 0x57, 0xa5, 0x6a, 0x44, 0x94, 0xb2, 0x2c, 0x5e, 0xdd, 0xe4, 0xe3, 0x54, 0xf9, 0x73,
	0xa9, 0x9a, 0x97, 0xca, 0x9a, 0xde, 0x4e, 0xfa, 0xd5, 0xe4, 0x01, 0xe6, 0xbf, 0x02, 0x5e, 0x01,
	0xfb, 0xa0, 0x93, 0xcf, 0x1d, 0x49, 0xf7, 0x89, 0xd7, 0x10, 0x7d, 0x29, 0x7d, 0xec, 0x03, 0xcb,
	0x6e, 0x78, 0x0a, 0x1f, 0x83, 0x4d, 0x0c, 0x4c, 0x99, 0x3d, 0xfa, 0xeb, 0x96, 0x07, 0x80, 0xbf,
	0xb8, 0x78, 0x0b, 0xd3, 0x92, 0xda, 0xd6, 0xdd, 0x39, 0xaa, 0x40, 0x0e, 0xd0, 0xe9, 0x45, 0x65,
	0x69, 0x4f, 0x66, 0xdc, 0x83, 0x1c, 0x20, 0x26, 0x10, 0xed, 0x74, 0x5d, 0x77, 0x55, 0xcc, 0xfa,
	0xb8, 0x1d, 0x5a, 0xad, 0x61, 0xfa, 0x6a, 0x88, 0x2c, 0x19, 0xcc, 0x60, 0xf6, 0xa6, 0x4e, 0xfe,
	0xbf, 0x78, 0xf9, 0xef, 0xd6, 0x64, 0xdc, 0xde, 0x72, 0xb2, 0xbb, 0xf0, 0x0f, 0xb5, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0xc5, 0x22, 0x8f, 0xc5, 0x01, 0x00, 0x00,
}
