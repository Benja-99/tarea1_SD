package main

import (
	"log"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.185:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	log.Printf("Se ejecuta el Data 1: IP 10.0.1.10")
	log.Printf("%s", c)

}
