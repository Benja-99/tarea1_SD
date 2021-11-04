package main

import (
	"fmt"
	"log"
	"net"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"example.com/m/pozo/github.com/Benja-99/tarea1_SD/pozo"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	s_pozo := pozo.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)
	pozo.RegisterPozoServiceServer(grpcServer, &s_pozo)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
