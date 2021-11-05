package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
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
	var response *chat.Message
	var err error
	var flag bool = true
	for flag {
		response, err = c.Peticion(context.Background(), &chat.Message{Body: "Quiero jugar", Peticion: int32(num_jug)})
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		if response.Body == "Jugador ingresado" {
			flag = false
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("Primera etapa ", num_jug)
	var jug int32
	var response_jugada *chat.Message
	for i := 0; i < 4; i++ {
		jug = etapa1_bot()
		ronda := i + 1
		var flag_ronda bool = true
		for flag_ronda {
			response_jugada, err = c.Jugada(context.Background(), &chat.Message{Jugada: jug, NumJuego: 1, NumRonda: int32(ronda), Jugador: int32(num_jug)})
			if err != nil {
				log.Fatalf("Error when calling Jugada: %s", err)
			}
			if response.Aux {
				flag_ronda = false
			}
			flag_verificar := true
			for flag_verificar {
				response_verificar, err_verificar := c.VerificarRonda(context.Background(), &chat.Message{NumRonda: int32(ronda)})
				if err != nil {
					log.Fatalf("Error when calling Jugada: %s", err)
				}
				flag_verificar = !response_verificar.Aux
			}

		}
		if response_jugada.Body == "Cagaste" {
			log.Printf("Se murio el jugador: %d", num_jug)
			response, err = c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(response_jugada.NumRonda)})
			break
		}
	}
	fmt.Println("Segunda etapa ", num_jug)
	var jug int32
	var response_jugada *chat.Message
	var flag_juego2 bool = true
	for flag_juego2 {
		jug = etapa2_bot()
		var flag_ronda bool = true
		for flag_ronda {
			response_jugada, err = c.Jugada(context.Background(), &chat.Message{Jugada: jug, NumJuego: 2, NumRonda: 1, Jugador: int32(num_jug)})
			if err != nil {
				log.Fatalf("Error when calling Jugada: %s", err)
			}
			if response_jugada.Aux {
				flag_ronda = false
			}
		}
		if response_jugada.Body == "Cagaste" {
			log.Printf("Se murio el jugador: %d", num_jug)
			response, err = c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(response_jugada.NumRonda)})
			break
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

	for i := 0; i <= 15; i++ {

		go client_bot(conn, i)
	}
	fmt.Scanln()
	/*
		c := chat.NewChatServiceClient(conn)
		//c_pozo := pozo.NewPozoServiceClient(conn)
		var response *chat.Message
		pri := `Iniciar juego?
			[ 1 ] SI
			[ 2 ] NO`
		fmt.Println(pri)
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
		eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
		if eleccion == "1" {

			response, err = c.Peticion(context.Background(), &chat.Message{Body: "Quiero jugar"})
			if err != nil {
				log.Fatalf("Error when calling Peticion: %s", err)
			}
			fmt.Println("Response from server: ", response.Body)
		}*/

}
