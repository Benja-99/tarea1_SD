package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"example.com/m/pozo/github.com/Benja-99/tarea1_SD/pozo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func etapa1_bot() int {
	rand.Seed(time.Now().UnixNano())
	resp := rand.Intn(9) + 1
	return resp
}

func client_bot(conn *grpc.ClientConn) {
	c := chat.NewChatServiceClient(conn)
	c_pozo := pozo.NewPozoServiceClient(conn)
	var response *chat.Message
	var err error
	response, err = c.Peticion(context.Background(), &chat.Message{Body: "Quiero jugar"})
	if err != nil {
		log.Fatalf("Error when calling Peticion: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
	if response.Body == "Jugador ingresado" {
		var response *pozo.Monto
		response, err = c_pozo.GetMonto(context.Background(), &pozo.Monto{CantidadTotal: 1})
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		log.Printf("Response from server: %d", response.CantidadTotal)

		fmt.Println("Primera etapa")
		for i := 1; i <= 4; i++ {

		}

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
