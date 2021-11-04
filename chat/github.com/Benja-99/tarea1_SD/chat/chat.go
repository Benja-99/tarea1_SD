package chat

import (
	"log"
	"math/rand"
	"time"

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
	rand.Seed(time.Now().UnixNano())
	resp := int32(rand.Intn(4) + 6)
	if in.Jugada >= resp {
		return &Message{Body: "Cagaste"}, nil
	} else {
		return &Message{Body: "Ronda pasada"}, nil
	}

}

func (s *Server) Jugada_2(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Jugada)
	rand.Seed(time.Now().UnixNano())
	resp := int32(rand.Intn(4) + 6)
	if in.Jugada >= resp {
		return &Message{Body: "Cagaste"}, nil
	} else {
		return &Message{Body: "Ronda pasada"}, nil
	}

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
