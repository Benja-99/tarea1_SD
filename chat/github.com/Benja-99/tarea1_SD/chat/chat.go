package chat

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/context"
)

type Server struct {
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
