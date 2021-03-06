package main

import (
	"fmt"
	"log"
	"net"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Bienvenidos al Juego del Calamar!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
