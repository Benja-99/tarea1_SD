package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	pidiendo int
	monto    int32
	pedido   int
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) Peticion(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %s", in.Body)
	return &Message{Body: "Jugador ingresado"}, nil
}

func (s *Server) Jugada(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Jugada)

	return &Message{Body: "Jugada registrada"}, nil
}

func (s *Server) EsperarPeticion(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)

	s.monto = in.Monto
	if s.pidiendo == 1 {
		s.pidiendo = 0
		s.pedido = 1
		return &Message{Body: "Pidiendo monto"}, nil
	}
	return &Message{Body: "Pidiendo monto"}, nil
}

func (s *Server) PedirMonto(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)

	s.pidiendo = 1
	if s.pedido == 1 {
		s.pedido = 0
		return &Message{Monto: s.monto}, nil
	}
	return &Message{Monto: s.monto}, nil
}
