package main

import (
	"log"
	"fmt"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func client_bot(conn *grpc.ClientConn){
	c := chat.NewChatServiceClient(conn)
	var response *chat.Message
	var err error
	response, err = c.Peticion(context.Background(), &chat.Message{Body: "Quiero jugar"})
 	if err != nil {
    	log.Fatalf("Error when calling Peticion: %s", err)
  	}
  	log.Printf("Response from server: %s", response.Body)
	if response.Body == "Jugador ingresado"{
		
	}
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	for i := 1; i <= 15; i++ {
		go client_bot(conn)
	} 
	
	fmt.Scanln()
}
