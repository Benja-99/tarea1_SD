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
	Monto                int32    `protobuf:"varint,4,opt,name=monto,proto3" json:"monto,omitempty"`
	Aux                  bool     `protobuf:"varint,5,opt,name=aux,proto3" json:"aux,omitempty"`
	Jugador              int32    `protobuf:"varint,6,opt,name=jugador,proto3" json:"jugador,omitempty"`
	Ronda                int32    `protobuf:"varint,7,opt,name=ronda,proto3" json:"ronda,omitempty"`
	NumJuego             int32    `protobuf:"varint,8,opt,name=num_juego,json=numJuego,proto3" json:"num_juego,omitempty"`
	NumRonda             int32    `protobuf:"varint,9,opt,name=num_ronda,json=numRonda,proto3" json:"num_ronda,omitempty"`
	SumJuego             int32    `protobuf:"varint,10,opt,name=sum_juego,json=sumJuego,proto3" json:"sum_juego,omitempty"`
	Jugadores            []int32  `protobuf:"varint,11,rep,packed,name=jugadores,proto3" json:"jugadores,omitempty"`
	Jugadas              []int32  `protobuf:"varint,12,rep,packed,name=jugadas,proto3" json:"jugadas,omitempty"`
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

func (m *Message) GetMonto() int32 {
	if m != nil {
		return m.Monto
	}
	return 0
}

func (m *Message) GetAux() bool {
	if m != nil {
		return m.Aux
	}
	return false
}

func (m *Message) GetJugador() int32 {
	if m != nil {
		return m.Jugador
	}
	return 0
}

func (m *Message) GetRonda() int32 {
	if m != nil {
		return m.Ronda
	}
	return 0
}

func (m *Message) GetNumJuego() int32 {
	if m != nil {
		return m.NumJuego
	}
	return 0
}

func (m *Message) GetNumRonda() int32 {
	if m != nil {
		return m.NumRonda
	}
	return 0
}

func (m *Message) GetSumJuego() int32 {
	if m != nil {
		return m.SumJuego
	}
	return 0
}

func (m *Message) GetJugadores() []int32 {
	if m != nil {
		return m.Jugadores
	}
	return nil
}

func (m *Message) GetJugadas() []int32 {
	if m != nil {
		return m.Jugadas
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "chat.Message")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x49, 0x93, 0x38, 0xf6, 0xa4, 0x2d, 0xb0, 0x20, 0xb4, 0x2a, 0x1c, 0xa2, 0x88, 0x83,
	0x55, 0x68, 0xd2, 0x52, 0x71, 0xe8, 0xb5, 0xb4, 0x12, 0x54, 0x44, 0x8a, 0x1c, 0x84, 0x10, 0x97,
	0x6a, 0x62, 0x0f, 0xce, 0x46, 0xf5, 0x6e, 0xb4, 0xb6, 0x11, 0xe5, 0x41, 0x78, 0x16, 0x1e, 0x0f,
	0x79, 0x37, 0x76, 0x04, 0x17, 0x2f, 0xb7, 0xfd, 0x3d, 0xff, 0x3f, 0x3b, 0xfb, 0x8d, 0x64, 0x80,
	0x78, 0x85, 0xc5, 0x64, 0xa3, 0x55, 0xa1, 0x58, 0xaf, 0x3a, 0x8f, 0x7f, 0xef, 0xc1, 0x60, 0x46,
	0x79, 0x8e, 0x29, 0x31, 0x06, 0xbd, 0xa5, 0x4a, 0xee, 0x79, 0x67, 0xd4, 0x09, 0x83, 0xc8, 0x9c,
	0xd9, 0x11, 0xf8, 0x1b, 0x2a, 0x44, 0x2c, 0x94, 0xe4, 0x7b, 0xa3, 0x4e, 0xd8, 0x8f, 0x1a, 0xcd,
	0x9e, 0x81, 0xb7, 0x2e, 0x53, 0x4c, 0x90, 0x77, 0x4d, 0x65, 0xab, 0xd8, 0x53, 0xe8, 0x67, 0x4a,
	0x16, 0x8a, 0xf7, 0xcc, 0x67, 0x2b, 0xd8, 0x23, 0xe8, 0x62, 0xf9, 0x83, 0xf7, 0x47, 0x9d, 0xd0,
	0x8f, 0xaa, 0x23, 0xe3, 0x30, 0x30, 0x09, 0xa5, 0xb9, 0x67, 0x9c, 0xb5, 0xac, 0x3a, 0x68, 0x25,
	0x13, 0xe4, 0x03, 0xdb, 0xc1, 0x08, 0xf6, 0x1c, 0x02, 0x59, 0x66, 0xb7, 0xeb, 0x92, 0x52, 0xc5,
	0x7d, 0x3b, 0x8c, 0x2c, 0xb3, 0x9b, 0x4a, 0xd7, 0x45, 0x1b, 0x0b, 0x9a, 0x62, 0x54, 0x27, 0xf3,
	0x26, 0x09, 0xb6, 0x98, 0xd7, 0xc9, 0x17, 0x10, 0x6c, 0xef, 0xa5, 0x9c, 0x0f, 0x47, 0xdd, 0xb0,
	0x1f, 0xed, 0x3e, 0x34, 0x43, 0x62, 0xce, 0xf7, 0x4d, 0xad, 0x96, 0x6f, 0x7e, 0x0d, 0x60, 0xf8,
	0x6e, 0x85, 0xc5, 0x82, 0xf4, 0x77, 0x11, 0x13, 0x3b, 0x06, 0x7f, 0x81, 0xf7, 0xef, 0xe9, 0xee,
	0x4e, 0xb1, 0x83, 0x89, 0x21, 0xbd, 0x25, 0x7b, 0xf4, 0xb7, 0x1c, 0x3f, 0xa8, 0xbc, 0xf3, 0x1a,
	0x63, 0x9b, 0x37, 0x04, 0xef, 0xc6, 0x82, 0x6d, 0x73, 0x9e, 0xc1, 0xc3, 0xeb, 0x7c, 0x43, 0x1a,
	0xb5, 0x73, 0xf3, 0xd7, 0x00, 0x73, 0x4a, 0x84, 0x9e, 0x99, 0x1d, 0xb5, 0xb9, 0x27, 0xb0, 0xff,
	0x41, 0x8a, 0x58, 0xa0, 0xb6, 0xe8, 0xda, 0xfc, 0xaf, 0x20, 0xf8, 0x4c, 0x5a, 0x7c, 0x13, 0x31,
	0xea, 0x56, 0xf3, 0x14, 0x0e, 0xb6, 0xcd, 0xbf, 0xb8, 0x75, 0xdf, 0x4d, 0x63, 0xb7, 0xec, 0x70,
	0xc1, 0x27, 0xd2, 0x99, 0x90, 0xae, 0x81, 0x10, 0xbc, 0x59, 0x49, 0xda, 0x01, 0xcc, 0x5b, 0x78,
	0xf2, 0x0f, 0xf9, 0xb9, 0xfa, 0xd9, 0x1e, 0x3b, 0x85, 0xc3, 0x1d, 0x7d, 0xa7, 0xc4, 0x31, 0xf8,
	0xd7, 0xb2, 0xd0, 0x94, 0xba, 0x01, 0xad, 0xbd, 0x1f, 0x45, 0x42, 0xed, 0x81, 0x13, 0x18, 0x2e,
	0x30, 0x46, 0xed, 0xf8, 0xe8, 0x53, 0x38, 0x6c, 0xb6, 0xeb, 0x06, 0xf4, 0x1c, 0x1e, 0x5b, 0x4c,
	0x32, 0x51, 0x11, 0xa5, 0x22, 0x2f, 0xb4, 0x72, 0x79, 0xc6, 0xd5, 0xff, 0x04, 0x2e, 0x5f, 0x7e,
	0x1d, 0xa7, 0xa2, 0x58, 0x95, 0xcb, 0x49, 0xac, 0xb2, 0xe9, 0x25, 0xc9, 0x35, 0x9e, 0x5c, 0x5c,
	0x4c, 0x0b, 0xd4, 0x84, 0x67, 0xb7, 0x8b, 0xab, 0x69, 0xe5, 0x5f, 0x7a, 0xe6, 0x37, 0x78, 0xfe,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xbb, 0xec, 0xf3, 0x14, 0x05, 0x00, 0x00,
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
	EsperarPeticion(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	PedirMonto(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	IniciarJuego(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Verificar(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	IniciarXJuego(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	IniciarRonda(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	TerminarRonda(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Muerto(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EsperarPeticionPozo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	PedirMontoPozo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Entregar(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EntregarLider(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	SacarMuerto(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	VerificarRonda(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	EsperandoRegistro(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	DandoRegistro(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
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

func (c *chatServiceClient) EsperarPeticion(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EsperarPeticion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) PedirMonto(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/PedirMonto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) IniciarJuego(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/IniciarJuego", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Verificar(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Verificar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) IniciarXJuego(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/IniciarXJuego", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) IniciarRonda(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/IniciarRonda", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) TerminarRonda(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/TerminarRonda", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Muerto(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Muerto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EsperarPeticionPozo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EsperarPeticionPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) PedirMontoPozo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/PedirMontoPozo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Entregar(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Entregar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EntregarLider(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EntregarLider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SacarMuerto(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SacarMuerto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) VerificarRonda(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/VerificarRonda", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EsperandoRegistro(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EsperandoRegistro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DandoRegistro(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/DandoRegistro", in, out, opts...)
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
	EsperarPeticion(context.Context, *Message) (*Message, error)
	PedirMonto(context.Context, *Message) (*Message, error)
	IniciarJuego(context.Context, *Message) (*Message, error)
	Verificar(context.Context, *Message) (*Message, error)
	IniciarXJuego(context.Context, *Message) (*Message, error)
	IniciarRonda(context.Context, *Message) (*Message, error)
	TerminarRonda(context.Context, *Message) (*Message, error)
	Muerto(context.Context, *Message) (*Message, error)
	EsperarPeticionPozo(context.Context, *Message) (*Message, error)
	PedirMontoPozo(context.Context, *Message) (*Message, error)
	Entregar(context.Context, *Message) (*Message, error)
	EntregarLider(context.Context, *Message) (*Message, error)
	SacarMuerto(context.Context, *Message) (*Message, error)
	VerificarRonda(context.Context, *Message) (*Message, error)
	EsperandoRegistro(context.Context, *Message) (*Message, error)
	DandoRegistro(context.Context, *Message) (*Message, error)
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
func (*UnimplementedChatServiceServer) EsperarPeticion(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EsperarPeticion not implemented")
}
func (*UnimplementedChatServiceServer) PedirMonto(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirMonto not implemented")
}
func (*UnimplementedChatServiceServer) IniciarJuego(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IniciarJuego not implemented")
}
func (*UnimplementedChatServiceServer) Verificar(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verificar not implemented")
}
func (*UnimplementedChatServiceServer) IniciarXJuego(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IniciarXJuego not implemented")
}
func (*UnimplementedChatServiceServer) IniciarRonda(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IniciarRonda not implemented")
}
func (*UnimplementedChatServiceServer) TerminarRonda(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerminarRonda not implemented")
}
func (*UnimplementedChatServiceServer) Muerto(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Muerto not implemented")
}
func (*UnimplementedChatServiceServer) EsperarPeticionPozo(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EsperarPeticionPozo not implemented")
}
func (*UnimplementedChatServiceServer) PedirMontoPozo(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirMontoPozo not implemented")
}
func (*UnimplementedChatServiceServer) Entregar(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Entregar not implemented")
}
func (*UnimplementedChatServiceServer) EntregarLider(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntregarLider not implemented")
}
func (*UnimplementedChatServiceServer) SacarMuerto(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SacarMuerto not implemented")
}
func (*UnimplementedChatServiceServer) VerificarRonda(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerificarRonda not implemented")
}
func (*UnimplementedChatServiceServer) EsperandoRegistro(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EsperandoRegistro not implemented")
}
func (*UnimplementedChatServiceServer) DandoRegistro(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DandoRegistro not implemented")
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

func _ChatService_EsperarPeticion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EsperarPeticion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EsperarPeticion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EsperarPeticion(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_PedirMonto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).PedirMonto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/PedirMonto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).PedirMonto(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_IniciarJuego_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).IniciarJuego(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/IniciarJuego",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).IniciarJuego(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Verificar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Verificar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Verificar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Verificar(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_IniciarXJuego_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).IniciarXJuego(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/IniciarXJuego",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).IniciarXJuego(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_IniciarRonda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).IniciarRonda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/IniciarRonda",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).IniciarRonda(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_TerminarRonda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).TerminarRonda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/TerminarRonda",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).TerminarRonda(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Muerto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Muerto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Muerto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Muerto(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EsperarPeticionPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EsperarPeticionPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EsperarPeticionPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EsperarPeticionPozo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_PedirMontoPozo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).PedirMontoPozo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/PedirMontoPozo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).PedirMontoPozo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Entregar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Entregar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Entregar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Entregar(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EntregarLider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EntregarLider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EntregarLider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EntregarLider(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SacarMuerto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SacarMuerto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SacarMuerto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SacarMuerto(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_VerificarRonda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).VerificarRonda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/VerificarRonda",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).VerificarRonda(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EsperandoRegistro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EsperandoRegistro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EsperandoRegistro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EsperandoRegistro(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DandoRegistro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DandoRegistro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/DandoRegistro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DandoRegistro(ctx, req.(*Message))
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
		{
			MethodName: "EsperarPeticion",
			Handler:    _ChatService_EsperarPeticion_Handler,
		},
		{
			MethodName: "PedirMonto",
			Handler:    _ChatService_PedirMonto_Handler,
		},
		{
			MethodName: "IniciarJuego",
			Handler:    _ChatService_IniciarJuego_Handler,
		},
		{
			MethodName: "Verificar",
			Handler:    _ChatService_Verificar_Handler,
		},
		{
			MethodName: "IniciarXJuego",
			Handler:    _ChatService_IniciarXJuego_Handler,
		},
		{
			MethodName: "IniciarRonda",
			Handler:    _ChatService_IniciarRonda_Handler,
		},
		{
			MethodName: "TerminarRonda",
			Handler:    _ChatService_TerminarRonda_Handler,
		},
		{
			MethodName: "Muerto",
			Handler:    _ChatService_Muerto_Handler,
		},
		{
			MethodName: "EsperarPeticionPozo",
			Handler:    _ChatService_EsperarPeticionPozo_Handler,
		},
		{
			MethodName: "PedirMontoPozo",
			Handler:    _ChatService_PedirMontoPozo_Handler,
		},
		{
			MethodName: "Entregar",
			Handler:    _ChatService_Entregar_Handler,
		},
		{
			MethodName: "EntregarLider",
			Handler:    _ChatService_EntregarLider_Handler,
		},
		{
			MethodName: "SacarMuerto",
			Handler:    _ChatService_SacarMuerto_Handler,
		},
		{
			MethodName: "VerificarRonda",
			Handler:    _ChatService_VerificarRonda_Handler,
		},
		{
			MethodName: "EsperandoRegistro",
			Handler:    _ChatService_EsperandoRegistro_Handler,
		},
		{
			MethodName: "DandoRegistro",
			Handler:    _ChatService_DandoRegistro_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
