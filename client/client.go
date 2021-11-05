package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func etapa2_bot() int32 {
	rand.Seed(time.Now().UnixNano())
	resp := int32(rand.Intn(4) + 1)
	return resp
}

func client_bot(conn *grpc.ClientConn, num_jug int) {
	c := chat.NewChatServiceClient(conn)
	var response *chat.Message
	var err error
	var flag bool = true
	for flag {

		response, err = c.Peticion(context.Background(), &chat.Message{Body: "Quiero jugar", Peticion: int32(num_jug)})
		//fmt.Println("ERROR")
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		if response.Body == "Jugador ingresado" {
			flag = false
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println("El jugador " + strconv.Itoa(num_jug) + " esta en la primera etapa ")
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

		}
		flagVerificarRonda := true
		for flagVerificarRonda {
			response, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 3, NumRonda: int32(ronda), Jugador: int32(num_jug)})
			if err != nil {
				log.Fatalf("Error when calling Jugada: %s", err)
			}
			flagVerificarRonda = response.Aux
		}
		if response_jugada.Body == "Cagaste" {
			log.Printf("Se murio el jugador: %d", num_jug)
			response4, err4 := c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(response_jugada.NumRonda)})
			if err4 != nil {
				log.Fatalf("Error when calling Jugada: %s", err4)
			}
			log.Printf("%s", response4.Body)
			os.Exit(1)
		}
	}
	fmt.Println("El jugador " + strconv.Itoa(num_jug) + " esta en la segunda etapa ")
	flag_juego2 := true
	ronda := 0
	for flag_juego2 {
		ronda = ronda + 1
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
			time.Sleep(3 * time.Second)
		}
		flagRonda := true
		for flagRonda {
			response_jugada, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 2, NumRonda: int32(ronda), Jugador: int32(num_jug)})
			if err != nil {
				log.Fatalf("Error when calling Jugada: %s", err)
			}
			if response_jugada.Aux {
				flag_ronda = false
			}
			time.Sleep(3 * time.Second)
		}
		if response_jugada.Body == "Cagaste" {
			log.Printf("Se murio el jugador: %d", num_jug)
			response4, err4 := c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(ronda)})
			if err4 != nil {
				log.Fatalf("Error when calling Jugada: %s", err4)
			}
			log.Printf("%s", response4.Body)
			os.Exit(1)
		}
	}
	fmt.Println("El jugador " + strconv.Itoa(num_jug) + " esta en la tercera etapa ")
	flag_juego3 := true
	for flag_juego3 {
		ronda = ronda + 1
		jug = etapa1_bot()
		flag_ronda := true
		for flag_ronda {
			response_jugada, err = c.Jugada(context.Background(), &chat.Message{Jugada: jug, NumJuego: 2, NumRonda: 1, Jugador: int32(num_jug)})
			if err != nil {
				log.Fatalf("Error when calling Jugada: %s", err)
			}
			if response_jugada.Aux {
				flag_ronda = false
			}
			time.Sleep(3 * time.Second)
		}
		flagVerificarRonda := true
		for flagVerificarRonda {
			response_jugada, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 3, NumRonda: int32(ronda), Jugador: int32(num_jug)})
			if err != nil {
				log.Fatalf("Error when calling Jugada: %s", err)
			}
			flagVerificarRonda = response_jugada.Aux
		}
		if response_jugada.Body == "Cagaste" {
			log.Printf("Se murio el jugador: %d", num_jug)
			response4, err4 := c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(ronda)})
			if err4 != nil {
				log.Fatalf("Error when calling Muerto: %s", err4)
			}
			log.Printf("%s", response4.Body)
			os.Exit(1)
		} else {
			// se premian a los ganadores
			log.Printf("El jugador %d es ganador de los juegos del calamar!", num_jug)
		}
	}
}

func client_real(conn *grpc.ClientConn, num_jug int) {
	c := chat.NewChatServiceClient(conn)
	var response *chat.Message
	var err error
	var flag bool = true
	pri := `Iniciar juego?
			[ 1 ] SI
			[ 2 ] NO`
	fmt.Println(pri)
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	eleccion := strings.TrimRight(entrada, "\r\n")
	if eleccion == "1" {
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
		fmt.Println("El jugador " + strconv.Itoa(num_jug) + " esta en la primera etapa ")
		var response_jugada *chat.Message
		fmt.Println("Inicio Juego 1")
		for i := 0; i < 4; i++ {
			ronda := i + 1
			fmt.Println("Ronda " + strconv.Itoa(ronda) + " - Elija un numero del 1 al 10")
			reader := bufio.NewReader(os.Stdin)
			entrada, _ := reader.ReadString('\n')
			eleccion := strings.TrimRight(entrada, "\r\n")
			eleccionInt, _ := strconv.Atoi(eleccion)
			var flag_ronda bool = true
			for flag_ronda {
				response_jugada, err = c.Jugada(context.Background(), &chat.Message{Jugada: int32(eleccionInt), NumJuego: 1, NumRonda: int32(ronda), Jugador: int32(num_jug)})
				if err != nil {
					log.Fatalf("Error when calling Jugada: %s", err)
				}

				if response_jugada.Aux {
					flag_ronda = false
				}
				log.Printf("Pegado en este while")
				time.Sleep(3 * time.Second)
			}
			if response_jugada.Body == "Cagaste" {
				log.Printf("Se murio el jugador: %d", num_jug)
				response4, err4 := c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(response_jugada.NumRonda)})
				if err4 != nil {
					log.Fatalf("Error when calling Jugada: %s", err4)
				}
				log.Printf("%s", response4.Body)
				os.Exit(1)
			}
		}
		//Se termina el juego 1
		fmt.Println("El jugador " + strconv.Itoa(num_jug) + " esta en la segunda etapa ")
		fmt.Println("Inicio Juego 2")
		flag_juego2 := true
		IndiceRonda := 1
		for flag_juego2 {
			fmt.Println("Ronda " + strconv.Itoa(IndiceRonda) + " - Elija un numero del 1 al 4")
			reader := bufio.NewReader(os.Stdin)
			entrada, _ := reader.ReadString('\n')
			eleccion := strings.TrimRight(entrada, "\r\n")
			eleccionInt, _ := strconv.Atoi(eleccion)
			var flag_ronda bool = true
			for flag_ronda {
				response_jugada, err = c.Jugada(context.Background(), &chat.Message{Jugada: int32(eleccionInt), NumJuego: 2, NumRonda: int32(IndiceRonda), Jugador: int32(num_jug)})
				if err != nil {
					log.Fatalf("Error when calling Jugada: %s", err)
				}
				if response_jugada.Aux {
					flag_ronda = false
				}
			}
			if response_jugada.Body == "Cagaste" {
				log.Printf("Se murio el jugador: %d", num_jug)
				response4, err4 := c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(response_jugada.NumRonda)})
				if err4 != nil {
					log.Fatalf("Error when calling Jugada: %s", err4)
				}
				log.Printf("%s", response4.Body)
				os.Exit(1)
			}
			IndiceRonda++
		}
		//Se termina el juego 2
		fmt.Println("El jugador " + strconv.Itoa(num_jug) + " esta en la tercera etapa ")
		fmt.Println("Inicio Juego 2")
		flag_juego3 := true
		IndiceRonda = 1
		for flag_juego3 {
			fmt.Println("Ronda " + strconv.Itoa(IndiceRonda) + " - Elija un numero del 1 al 4")
			reader := bufio.NewReader(os.Stdin)
			entrada, _ := reader.ReadString('\n')
			eleccion := strings.TrimRight(entrada, "\r\n")
			eleccionInt, _ := strconv.Atoi(eleccion)
			flag_ronda := true
			for flag_ronda {
				response_jugada, err = c.Jugada(context.Background(), &chat.Message{Jugada: int32(eleccionInt), NumJuego: 3, NumRonda: int32(IndiceRonda), Jugador: int32(num_jug)})
				if err != nil {
					log.Fatalf("Error when calling Jugada: %s", err)
				}
				if response_jugada.Aux {
					flag_ronda = false
				}
				time.Sleep(3 * time.Second)
			}
			flagVerificarRonda := true
			for flagVerificarRonda {
				response_jugada, err = c.VerificarRonda(context.Background(), &chat.Message{NumJuego: 3, NumRonda: int32(IndiceRonda), Jugador: int32(num_jug)})
				if err != nil {
					log.Fatalf("Error when calling Jugada: %s", err)
				}
				if response_jugada.Aux {
					flagVerificarRonda = false
				}
				time.Sleep(3 * time.Second)
			}
			if response_jugada.Body == "Cagaste" {
				log.Printf("Se murio el jugador: %d", num_jug)
				response4, err4 := c.Muerto(context.Background(), &chat.Message{Jugador: int32(num_jug), Ronda: int32(IndiceRonda)})
				if err4 != nil {
					log.Fatalf("Error when calling Jugada: %s", err4)
				}
				log.Printf("%s", response4.Body)
				os.Exit(1)
			} else {
				// se premian a los ganadores
				log.Printf("El jugador %d es ganador de los juegos del calamar!", num_jug)
			}
			IndiceRonda++
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
	client_real(conn, 16)
}
