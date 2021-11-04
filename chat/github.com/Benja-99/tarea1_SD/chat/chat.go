package chat

import (
	"log"

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

	return &Message{Body: "Jugada registrada"}, nil
}
