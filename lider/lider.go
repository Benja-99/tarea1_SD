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
		response2, err2 = c.EntregarLider(context.Background(), &chat.Message{Body: "Quiero monto"})
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
	conn, err := grpc.Dial("10.6.40.185:9000", grpc.WithInsecure())
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
		response, err = c.IniciarJuego(context.Background(), &chat.Message{Body: "Iniciar", NumJuego: 1})
		if err != nil {
			log.Fatalf("Error when calling Iniciar Juego: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)
		var flag bool = true
		for flag {
			response, err = c.Verificar(context.Background(), &chat.Message{Body: "Verificar"})
			if err != nil {
				log.Fatalf("Error when calling Verificar: %s", err)
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
			for i := 0; i < 4; i++ {
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
					response, err = c.IniciarRonda(context.Background(), &chat.Message{Body: "Iniciar Ronda", NumRonda: int32(i)})
					if err != nil {
						log.Fatalf("Error when calling IniciarRonda: %s", err)
					}
					fmt.Println(response.Body + " - Numero elegido por el LIDER: " + strconv.Itoa(int(response.Jugada)))
					var flag bool = true
					for flag {
						response, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 1, NumRonda: int32(i), Monto: 1})
						if err != nil {
							log.Fatalf("Error when calling VerificarRonda: %s", err)
						}
						flag = response.Aux
						time.Sleep(1 * time.Second)
					}

					response, err = c.TerminarRonda(context.Background(), &chat.Message{Body: "Terminar Ronda", NumRonda: int32(i)})
					if err != nil {
						log.Fatalf("Error when calling TerminarRonda: %s", err)
					}
					//response, err = c.DandoRegistro(context.Background(), &chat.Message{Jugador:)

					fmt.Println("Jugadores: ", response.Jugadores)
					SumJugadoresVivos := 0
					NumeroGanador := 0
					for i := 0; i < 16; i++ {
						if response.Jugadores[i] == 1 {
							SumJugadoresVivos++
							NumeroGanador = i
						}
					}
					if SumJugadoresVivos == 0 {
						log.Printf("Se acabo el juego no quedan jugadores vivos")
						return
					} else if SumJugadoresVivos == 1 {
						log.Printf("El ganador del juego del calamar es %d", NumeroGanador+1)
						return
					}
				} else {
					break
				}
			}
			SumJugadoresVivos := 0
			NumeroGanador := 0
			for i := 0; i < 16; i++ {
				if response.Jugadores[i] == 1 {
					SumJugadoresVivos++
					NumeroGanador = i
				}
			}
			if SumJugadoresVivos == 0 {
				log.Printf("Se acabo el juego no quedan juagdores vivos")
				return
			} else if SumJugadoresVivos == 1 {
				log.Printf("El ganador del juego del calamar es %d", NumeroGanador+1)
				return
			}
		} else {
			return
		}
		pri = `Iniciar juego 2?
			[ 1 ] SI
			[ 2 ] NO`
		fmt.Println(pri)
		reader = bufio.NewReader(os.Stdin)
		entrada, _ = reader.ReadString('\n')          // Leer hasta el separador de salto de línea
		eleccion = strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
		if eleccion == "1" {
			response, err = c.IniciarXJuego(context.Background(), &chat.Message{NumJuego: 2})
			if err != nil {
				log.Fatalf("Error when calling IniciarXJuego: %s", err)
			}
			log.Printf("Respuesta: %s", response.Body)
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
						time.Sleep(3 * time.Second)
					}
					response, err = c.TerminarRonda(context.Background(), &chat.Message{Body: "Terminar Ronda"})
					if err != nil {
						log.Fatalf("Error when calling TerminarRonda: %s", err)
					}
					log.Printf("%s", response.Body)
				}
			}
			SumJugadoresVivos := 0
			NumeroGanador := 0
			for i := 0; i < 16; i++ {
				if response.Jugadores[i] == 1 {
					SumJugadoresVivos++
					NumeroGanador = i
				}
			}
			if SumJugadoresVivos == 0 {
				log.Printf("Se acabo el juego no quedan juagdores vivos")
				os.Exit(1)
			} else if SumJugadoresVivos == 1 {
				log.Printf("El ganador del juego del calamar es %d", NumeroGanador)
				os.Exit(1)
			}
		} else {
			os.Exit(1)
		}
		pri = `Iniciar juego 3?
			[ 1 ] SI
			[ 2 ] NO`
		fmt.Println(pri)
		reader = bufio.NewReader(os.Stdin)
		entrada, _ = reader.ReadString('\n')          // Leer hasta el separador de salto de línea
		eleccion = strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
		if eleccion == "1" {
			response, err = c.IniciarXJuego(context.Background(), &chat.Message{NumJuego: 3})
			if err != nil {
				log.Fatalf("Error when calling IniciarXJuego: %s", err)
			}
			log.Printf("Respuesta: %s", response.Body)
			flagJuego3 := true
			for flagJuego3 {
				fmt.Printf("Iniciar ronda final del juego?")
				pri := `
					[ 1 ] SI
					[ 2 ] NO`
				fmt.Println(pri)
				reader := bufio.NewReader(os.Stdin)
				entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
				eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de línea de la entrada del usuario
				if eleccion == "1" {
					response, err = c.IniciarRonda(context.Background(), &chat.Message{Body: "Iniciar Ronda", NumJuego: 3})
					if err != nil {
						log.Fatalf("Error when calling IniciarRonda: %s", err)
					}
					fmt.Println(response.Body)
					var flag bool = true
					for flag {
						response, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 3, NumRonda: 1})
						if err != nil {
							log.Fatalf("Error when calling VerificarRonda: %s", err)
						}
						flag = response.Aux
						time.Sleep(3 * time.Second)
					}
					response, err = c.TerminarRonda(context.Background(), &chat.Message{Body: "Terminar Ronda"})
					if err != nil {
						log.Fatalf("Error when calling TerminarRonda: %s", err)
					}
					log.Printf("%s", response.Body)

					flagJuego3 = false

					// Terminan los juegos del calamar
				}
			}
		}
	}

}
