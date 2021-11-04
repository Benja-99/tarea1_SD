package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"bufio"
	"os"
	"strings"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"example.com/m/pozo/github.com/Benja-99/tarea1_SD/pozo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func etapa1_bot() int32 {
	rand.Seed(time.Now().UnixNano())
	resp := int32(rand.Intn(9) + 1)
	return resp
}


func client_bot(conn *grpc.ClientConn, num_jug int) {
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
		var jug int32
		var response_jugada *chat.Message
		for i := 1; i <= 4; i++ {
			log.Printf("RONDA %d", i)
			jug = etapa1_bot()
			response_jugada, err = c.Jugada(context.Background(), &chat.Message{Jugada: jug})
			if err != nil {
				log.Fatalf("Error when calling Peticion: %s", err)
			}
			log.Printf("Response from server: %s", response_jugada.Body)
			if response_jugada.Body == "Cagaste"{
				log.Printf("Response from server: %s", response_jugada.Body)
				break
			}
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
		
		go client_bot(conn, i)
	}

	c := chat.NewChatServiceClient(conn)
	//c_pozo := pozo.NewPozoServiceClient(conn)
	var response *chat.Message
	pri :=	`Iniciar juego?
		[ 1 ] SI
		[ 2 ] NO`
	fmt.Println(pri)
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n') // Leer hasta el separador de salto de línea
	eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
	if eleccion == "1"{
		
		response, err = c.Peticion(context.Background(), &chat.Message{Body: "Quiero jugar"})
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		fmt.Println("Response from server: ", response.Body)
	}

}
