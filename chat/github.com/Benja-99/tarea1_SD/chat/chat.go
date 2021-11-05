package chat

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/streadway/amqp"
	"golang.org/x/net/context"
)

type Server struct {
	pidiendo              int
	monto                 int32
	pedido                int
	array_jugador         [16]int32
	inicio_juego          bool
	array_jugadas1        [16]int
	array_jugadas1_comp   [16]int
	num_lider             [3]int
	juegos_iniciados      [3]bool
	ronda_iniciada        bool
	jugador_pidiendo      int
	lider_pidiendo        int
	muerto                int32
	ronda                 int32
	jugadores             [16]int32
	equipo1               [8]int32
	equipo2               [8]int32
	jugados_equipo2       [8]int32
	jugados_equipo1       [8]int32
	suma_equipo1          int32
	suma_equipo2          int32
	largo_equipo          int32
	asesinado             int32
	registro              int
	jugador_name          int32
	ronda_name            int32
	jugada_name           []int32
	array_jugadas_enviar [4]int32
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	if in.Body == "Iniciar" {
		s.inicio_juego = true
	}
	return &Message{Body: "Se inicio el juego"}, nil
}

func (s *Server) IniciarJuego(ctx context.Context, in *Message) (*Message, error) {
	if in.Body == "Iniciar" {
		s.inicio_juego = true
		s.num_lider[0] = rand.Intn(4) + 6
		s.num_lider[1] = rand.Intn(3) + 1
		s.num_lider[2] = rand.Intn(9) + 1
	}
	return &Message{Body: "Se inicio el juego"}, nil
}

func (s *Server) IniciarXJuego(ctx context.Context, in *Message) (*Message, error) {
	s.juegos_iniciados[in.NumJuego-1] = true
	if in.NumJuego == 2 {
		vivos := 0
		for i := 0; i < 16; i++ {
			if s.array_jugador[i] == 1 {
				vivos = vivos + 1
			}
		}
		if vivos%2 == 1 {
			pos_seleccionada := rand.Intn(vivos) + 1
			flagAsesinado := true
			for flagAsesinado {
				vivos_revisado := 0
				for i := 0; i < 16; i++ {
					if s.array_jugador[i] == 1 {
						vivos_revisado = vivos_revisado + 1
						if vivos_revisado == pos_seleccionada {
							//matar al wn
						}
					}
				}
			}
		}
		s.largo_equipo = vivos / 2
		j := 0
		pos := 0
		for i := 0; i < 16; i++ {
			if s.array_jugador[i] == 1 {
				if j == 0 {
					s.equipo1[pos] = int32(i)
				} else {
					s.equipo2[pos] = int32(i)
					j = 0
					pos = pos + 1
				}
			}
		}
	}
	return &Message{Body: "Se inicio el juego numero " + strconv.Itoa(int(in.NumJuego))}, nil
}

func (s *Server) IniciarRonda(ctx context.Context, in *Message) (*Message, error) {
	s.ronda_iniciada = true
	if in.NumJuego == 2 {
		s.num_lider = rand.Intn(3) + 1
	}
	return &Message{Body: "Se inicio la ronda"}, nil
}

func (s *Server) TerminarRonda(ctx context.Context, in *Message) (*Message, error) {
	s.ronda_iniciada = false
	return &Message{Body: "Se termino la ronda", Jugadores: s.array_jugador, Jugadas: s.array_jugadas1_enviar}, nil
}

func (s *Server) Verificar(ctx context.Context, in *Message) (*Message, error) {
	for i := 0; i <= 16; i++ {
		if s.array_jugador[i] == 0 {
			jug_faltante := strconv.Itoa(i)
			return &Message{Body: "Aun no " + jug_faltante}, nil
		}
	}
	return &Message{Body: "Ahora si"}, nil
}

func (s *Server) Peticion(ctx context.Context, in *Message) (*Message, error) {
	if s.inicio_juego {
		s.array_jugador[in.Peticion-1] = 1
		return &Message{Body: "Jugador ingresado"}, nil
	} else {
		return &Message{Body: "Jugador no ingresado"}, nil
	}

}

func (s *Server) VerificarRonda(ctx context.Context, in *Message) (*Message, error) {
	if in.NumJuego == 1 {
		for i := 0; i < 16; i++ {
			if s.array_jugadas1_comp[i] != int(in.NumRonda) {
				return &Message{Aux: false}, nil
			}
		}
		return &Message{Aux: true}, nil
	} else if in.NumJuego == 2 {
		for i := 0; i < int(s.largo_equipo); i++{
			if s.jugados_equipo1 != 1 or s.jugados_equipo2 != 1{
				return &Message{Aux: false}, nil
			} 
		}
		paridad_lider := s.num_lider % 2
		paridad_equipo1 := s.suma_equipo1 % 2
		paridad_equipo2 := s.suma_equipo2 % 2
		if paridad_lider == paridad_equipo1 && paridad_lider != paridad_equipo2{
			for i := 0; i < int32(largo_equipo); i++ {
				if in.Jugador == equipo2[i]{
					return &Message{Body: "Cagaste", Aux: true}
				}
			}
			return &Message{Body: "Eres del equipo ganador", Aux: true}
		}else if paridad_lider != paridad_equipo1 && paridad_lider == paridad_equipo2{
			for i := 0; i < int32(largo_equipo); i++ {
				if in.Jugador == equipo1[i]{
					return &Message{Body: "Cagaste", Aux: true}
				}
			}
			return &Message{Body: "Eres del equipo ganador", Aux: true}
		}else if paridad_lider != paridad_equipo1 && paridad_lider != paridad_equipo2{
			equipo_eliminado := rand.Intn(1) + 1
			if equipo_eliminado == 1{
				for i := 0; i < int32(largo_equipo); i++ {
					if in.Jugador == equipo2[i]{
						return &Message{Body: "Cagaste", Aux: true}
					}
				}
				return &Message{Body: "Eres del equipo ganador", Aux: true}
			}else{
				for i := 0; i < int32(largo_equipo); i++ {
					if in.Jugador == equipo1[i]{
						return &Message{Body: "Cagaste", Aux: true}
					}
				}
				return &Message{Body: "Eres del equipo ganador", Aux: true}
			}
		}
	}
}

func (s *Server) Jugada(ctx context.Context, in *Message) (*Message, error) {
	if in.NumJuego == 1 {
		if s.juegos_iniciados[in.NumJuego-1] && s.ronda_iniciada {
			if in.Jugada >= int32(s.num_lider[in.NumJuego-1]) {
				s.array_jugadas1_comp[in.Jugador-1]++
				s.array_jugador[in.Jugador-1] = 0
				return &Message{Body: "Cagaste", Aux: true, NumRonda: s.array_jugadas1_comp[in.Jugador-1]}, nil
			} else {
				s.array_jugadas1[in.Jugador-1] = s.array_jugadas1[in.Jugador-1] + int(in.Jugada)
				s.array_jugadas1_comp[in.Jugador-1]++
				if s.array_jugadas1_comp[in.Jugador-1] == 3 {
					if s.array_jugadas1_comp[in.Jugador-1] < 21 {
						s.array_jugador[in.Jugador-1] = 0

						return &Message{Body: "Cagaste", Aux: true, NumRonda: s.array_jugadas1_comp[in.Jugador-1]}, nil
					}
				}
				return &Message{Body: "Ronda pasada", Aux: true, SumJuego: int32(s.array_jugadas1[in.Jugador-1])}, nil
			}
		} else {
			return &Message{Aux: false}, nil
		}
	} else if in.NumJuego == 2 {
		if s.juegos_iniciados[1] && s.ronda_iniciada {
			if in.Jugador == s.asesinado {
				return &Message{Body: "Cagaste", Aux: true, NumRonda: s.array_jugadas1_comp[in.Jugador-1]}, nil
			}
			flagJugada := true
			for flagJugada {
				for i := 0; i < 8; i++ {
					if s.equipo1[i] == in.Jugador {
						s.suma_equipo1 = s.suma_equipo1 + in.Jugada
						s.jugados_equipo1[i] = 1
						flagJugada = false
						break
					} else if s.equipo2[i] == in.Jugador {
						s.suma_equipo2 = s.suma_equipo2 + in.Jugada
						s.jugados_equipo2[i] = 1
						flagJugada = false
						break
					}
				}
			}
		}
		return &Message{Body: "Jugada registrada exitosamente del " + strconv.Itoa(int(in.Jugador))}, nil
	}
	return &Message{Body: "ultimo retorno"}, nil
}


func (s *Server) EsperarPeticion(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)
	if s.pidiendo == 1 {
		s.pidiendo = 0
		return &Message{Body: "Jugador pidiendo monto"}, nil
	} else {
		return &Message{Body: "Nadie pidiendo"}, nil
	}
}

func (s *Server) PedirMonto(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)

	s.pidiendo = 1
	return &Message{Body: "Pidiendo monto"}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *Server) Muerto(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello1", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	failOnError(err, "Failed to declare a queue")

	aumento := "100000000"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(aumento),
		})

	failOnError(err, "Failed to publish a message")

	s.muerto = in.Jugador
	s.ronda = in.Ronda

	return &Message{Jugador: in.Jugador, Ronda: in.Ronda}, nil
}

func (s *Server) SacarMuerto(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)

	s.jugadores = [16]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, numero := range s.jugadores {
		if numero == s.muerto {
			return &Message{Body: "Lo encontramos muerto", Jugador: s.muerto, Ronda: s.ronda}, nil
		}
	}
	return &Message{Body: "Nadie muerto"}, nil
}

func (s *Server) PedirMontoPozo(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)

	s.pedido = 1
	return &Message{Body: "Pidiendo monto"}, nil
}

func (s *Server) EsperarPeticionPozo(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)
	if s.pedido == 1 {
		s.monto = in.Monto
		s.jugador_pidiendo = 1
		s.lider_pidiendo = 1
		s.pedido = 0
		return &Message{Monto: in.Monto}, nil
	} else {
		return &Message{Monto: 0}, nil
	}
}

func (s *Server) Entregar(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)
	if s.jugador_pidiendo == 1 {
		s.jugador_pidiendo = 0
		return &Message{Monto: s.monto}, nil
	} else {
		return &Message{Body: "Nada aun"}, nil
	}
}

func (s *Server) EntregarLider(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)
	if s.lider_pidiendo == 1 {
		s.lider_pidiendo = 0
		return &Message{Monto: s.monto}, nil
	} else {
		return &Message{Monto: -100}, nil
	}
}

func (s *Server) EsperandoRegistro(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)
	if s.registro == 1 {
		s.registro = 0
		return &Message{Body: "Registro nuevo", Jugador: s.jugador_name, Ronda: s.ronda_name, Jugadas: s.jugada_name}, nil
	} else {
		return &Message{Body: "Aun nada"}, nil
	}
}

func (s *Server) DandoRegistro(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %d", in.Monto)
	
	s.jugador_name = in.Jugador
	s.ronda_name = in.Ronda
	s.jugada_name = in.Jugadas
	s.registro = 1
	
	return &Message{Body: "Enviando registro"}, nil
}
