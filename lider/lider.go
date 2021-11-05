package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"google.golang.org/grpc"
)

func esperarPozo(conn *grpc.ClientConn) {
	c := chat.NewChatServiceClient(conn)
	var response *chat.Message
	var err error
	var flag bool = true
	for flag {
		response, err = c.EsperarPeticion(context.Background(), &chat.Message{Body: "Quiero monto"})
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		if response.Body == "Jugador pidiendo monto" {
			flag = false
		}
		time.Sleep(3 * time.Second)
	}
	log.Printf("Response from server: %s", response.Body)
	var response1 *chat.Message
	var err1 error
	response1, err1 = c.PedirMontoPozo(context.Background(), &chat.Message{Body: "Quiero monto"})
	if err1 != nil {
		log.Fatalf("Error when calling Peticion: %s", err)
	}
	log.Printf("Response from server: %s", response1.Body)
	var response2 *chat.Message
	var err2 error
	var flag2 bool = true
	for flag2 {
		response2, err = c.EntregarLider(context.Background(), &chat.Message{Body: "Quiero monto"})
		if err2 != nil {
			log.Fatalf("Error when calling Peticion: %s", err2)
		}
		if response2.Monto != -100 {
			flag = false
		}
		time.Sleep(3 * time.Second)
	}
	log.Printf("Response from server: %d", response.Monto)
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	go esperarPozo(conn)
	c := chat.NewChatServiceClient(conn)
	var response *chat.Message
	pri := `Empezar el juego del calamar?
			[ 1 ] SI
			[ 2 ] NO`
	fmt.Println(pri)
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
	eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
	if eleccion == "1" {
		response, err = c.IniciarJuego(context.Background(), &chat.Message{Body: "Iniciar"})
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)
		var flag bool = true
		for flag {
			response, err = c.Verificar(context.Background(), &chat.Message{Body: "Verificar"})
			if err != nil {
				log.Fatalf("Error when calling Peticion: %s", err)
			}
			if response.Body == "Ahora si" {
				flag = false
			}
			time.Sleep(3 * time.Second)
		}
		log.Printf("Empieza el juego del calamar")
		pri := `Iniciar juego 1?
			[ 1 ] SI
			[ 2 ] NO`
		fmt.Println(pri)
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
		eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
		if eleccion == "1" {
			response, err = c.IniciarXJuego(context.Background(), &chat.Message{NumJuego: 1})
			if err != nil {
				log.Fatalf("Error when calling IniciarXJuego: %s", err)
			}
			for i := 0; i <= 4; i++ {
				ronda := i + 1
				fmt.Printf("Iniciar ronda " + strconv.Itoa(ronda) + "?")
				pri := `
					[ 1 ] SI
					[ 2 ] NO`
				fmt.Println(pri)
				reader := bufio.NewReader(os.Stdin)
				entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
				eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
				if eleccion == "1" {
					response, err = c.IniciarRonda(context.Background(), &chat.Message{Body: "Iniciar Ronda"})
					if err != nil {
						log.Fatalf("Error when calling IniciarRonda: %s", err)
					}
					fmt.Println(response.Body)
					var flag bool = true
					for flag {
						response, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 1, NumRonda: int32(i)})
						if err != nil {
							log.Fatalf("Error when calling VerificarRonda: %s", err)
						}
						flag = response.Aux
					}

					response, err = c.TerminarRonda(context.Background(), &chat.Message{Body: "Terminar Ronda"})
					if err != nil {
						log.Fatalf("Error when calling TerminarRonda: %s", err)
					}
					response, err = c.DandoRegistro(context.Background(), &chat.Message{Jugador: })

				}
			}
			fmt.Println("Jugadores: ", response.Jugadores)
			//Aqui termina el juego 1
		}
		if eleccion == "2" {
			response, err = c.IniciarXJuego(context.Background(), &chat.Message{NumJuego: 2})
			if err != nil {
				log.Fatalf("Error when calling IniciarXJuego: %s", err)
			}
			flagJuego2 := true
			ronda := 1
			for flagJuego2 {
				fmt.Printf("Iniciar ronda " + strconv.Itoa(ronda) + "?")
				pri := `
					[ 1 ] SI
					[ 2 ] NO`
				fmt.Println(pri)
				reader := bufio.NewReader(os.Stdin)
				entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
				eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
				if eleccion == "1" {
					response, err = c.IniciarRonda(context.Background(), &chat.Message{Body: "Iniciar Ronda", NumJuego: 2})
					if err != nil {
						log.Fatalf("Error when calling IniciarRonda: %s", err)
					}
					fmt.Println(response.Body)
					var flag bool = true
					for flag {
						response, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 2, NumRonda: int32(ronda - 1)})
						if err != nil {
							log.Fatalf("Error when calling VerificarRonda: %s", err)
						}
						flag = response.Aux
					}
					response, err = c.TerminarRonda(context.Background(), &chat.Message{Body: "Terminar Ronda"})
					if err != nil {
						log.Fatalf("Error when calling TerminarRonda: %s", err)
					}

				}
			}
		}
	}

}
