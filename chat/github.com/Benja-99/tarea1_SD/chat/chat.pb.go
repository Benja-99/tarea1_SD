// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

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

type Message struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Peticion             int32    `protobuf:"varint,2,opt,name=peticion,proto3" json:"peticion,omitempty"`
	Jugada               int32    `protobuf:"varint,3,opt,name=jugada,proto3" json:"jugada,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetPeticion() int32 {
	if m != nil {
		return m.Peticion
	}
	return 0
}

func (m *Message) GetJugada() int32 {
	if m != nil {
		return m.Jugada
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "chat.Message")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x02, 0xb9, 0xd8, 0x7d, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf2, 0x53, 0x2a, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xd4, 0x92, 0xcc, 0xe4, 0xcc,
	0xfc, 0x3c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x38, 0x5f, 0x48, 0x8c, 0x8b, 0x2d, 0xab,
	0x34, 0x3d, 0x31, 0x25, 0x51, 0x82, 0x19, 0x2c, 0x03, 0xe5, 0x19, 0xf5, 0x33, 0x72, 0x71, 0x3b,
	0x67, 0x24, 0x96, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x69, 0x71, 0x71, 0x04, 0x27,
	0x56, 0x7a, 0xa4, 0xe6, 0xe4, 0xe4, 0x0b, 0xf1, 0xea, 0x81, 0x5d, 0x00, 0xb5, 0x52, 0x0a, 0x95,
	0xab, 0xc4, 0x00, 0x52, 0x1b, 0x00, 0x33, 0x9f, 0x90, 0x5a, 0x0d, 0x2e, 0x36, 0x2f, 0xb0, 0x8d,
	0x84, 0x54, 0x3a, 0xa9, 0x44, 0x29, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x3b, 0xa5, 0xe6, 0x65, 0x25, 0xea, 0x5a, 0x5a, 0xea, 0x97, 0x24, 0x16, 0xa5, 0x26, 0x1a,
	0xc6, 0x07, 0xbb, 0xe8, 0x83, 0xd4, 0x27, 0xb1, 0x81, 0xc3, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x7b, 0xa6, 0x11, 0x25, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Peticion(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Jugada(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Peticion(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Peticion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Jugada(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Jugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	Peticion(context.Context, *Message) (*Message, error)
	Jugada(context.Context, *Message) (*Message, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) SayHello(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedChatServiceServer) Peticion(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peticion not implemented")
}
func (*UnimplementedChatServiceServer) Jugada(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Jugada not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Peticion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Peticion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Peticion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Peticion(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Jugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Jugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Jugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Jugada(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatService_SayHello_Handler,
		},
		{
			MethodName: "Peticion",
			Handler:    _ChatService_Peticion_Handler,
		},
		{
			MethodName: "Jugada",
			Handler:    _ChatService_Jugada_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
