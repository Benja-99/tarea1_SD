package chat

import (
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/streadway/amqp"
	"golang.org/x/net/context"
)

type Server struct {
	pidiendo      int
	monto         int32
	pedido        int
	array_jugador [16]int32

	// Variables Juego 1
	inicio_juego         bool
	array_jugadas1       [16]int
	array_jugadas1_comp  [16]int
	num_lider_1          int
	array_jugadas_enviar [4]int32

	juegos_iniciados [3]bool
	ronda_iniciada   bool
	jugador_pidiendo int
	lider_pidiendo   int
	muerto           int32
	ronda            int32

	jugadores [16]int32

	// variables juego 2
	equipo1          [8]int32
	equipo2          [8]int32
	jugados_equipo2  [8]int32
	jugados_equipo1  [8]int32
	suma_equipo1     int32
	suma_equipo2     int32
	largo_equipo     int32
	num_lider_juego2 int32
	asesinado        int32

	registro     int
	jugador_name int32
	ronda_name   int32
	jugada_name  []int32

	//variables juego 3
	equipos_juego3        [8]int32
	num_lider_juego3      int32
	num_elegido_jugadores [8]int32
	ganadores             [8]int32
	perdedores            [4]int32
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	if in.Body == "Iniciar" {
		s.inicio_juego = true
	}
	return &Message{Body: "Se inicio el juego, esperando a los jugadores"}, nil
}

func (s *Server) IniciarJuego(ctx context.Context, in *Message) (*Message, error) {
	if in.Body == "Iniciar" {
		s.inicio_juego = true

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
							s.asesinado = s.array_jugador[i]
						}
					}
				}
			}
		}
		s.largo_equipo = int32(vivos / 2)
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
	if in.NumJuego == 3 {
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
							s.asesinado = s.array_jugador[i]
						}
					}
				}
			}
		}
		var pos int32
		pos = 0
		for i := 0; i < 16; i++ {
			if s.array_jugador[i] == 1 {
				s.equipos_juego3[pos] = int32(i)
				pos = pos + 1
			}
		}
	}
	return &Message{Body: "Se inicio el juego numero " + strconv.Itoa(int(in.NumJuego))}, nil
}

func (s *Server) IniciarRonda(ctx context.Context, in *Message) (*Message, error) {
	s.ronda_iniciada = true
	if in.NumJuego == 1 {
		s.num_lider_1 = rand.Intn(4) + 6
	} else if in.NumJuego == 2 {
		s.num_lider_juego2 = int32(rand.Intn(3) + 1)
	}
	return &Message{Body: "Se inicio la ronda"}, nil
}

func (s *Server) TerminarRonda(ctx context.Context, in *Message) (*Message, error) {
	s.ronda_iniciada = false
	return &Message{Body: "Se termino la ronda", Jugadores: s.array_jugador[:], Jugadas: s.array_jugadas_enviar[:]}, nil
}

func (s *Server) Verificar(ctx context.Context, in *Message) (*Message, error) {
	for i := 0; i < 16; i++ {
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
			log.Printf("Procesando la respuesta del jugador %d: %d - %d", i+1, s.array_jugadas1_comp[i], in.NumRonda+1)
			if s.array_jugador[i] == 1 && s.array_jugadas1_comp[i] != int(in.NumRonda+1) {
				log.Printf("Falta que el jugador %d escoga su numero", i+1)
				return &Message{Aux: true}, nil
			}
		}
		return &Message{Aux: false}, nil
	} else if in.NumJuego == 2 {
		for i := 0; i < int(s.largo_equipo); i++ {
			if s.jugados_equipo1[i] != 1 || s.jugados_equipo2[i] != 1 {
				return &Message{Aux: false}, nil
			}
		}
		paridad_lider := s.num_lider_juego2 % 2
		paridad_equipo1 := s.suma_equipo1 % 2
		paridad_equipo2 := s.suma_equipo2 % 2
		if paridad_lider == paridad_equipo1 && paridad_lider != paridad_equipo2 {
			for i := 0; i < int(s.largo_equipo); i++ {
				if in.Jugador == s.equipo2[i] {
					return &Message{Body: "Cagaste", Aux: true, NumRonda: in.NumRonda, NumJuego: 2}, nil
				}
			}
			return &Message{Body: "Eres del equipo ganador", Aux: true}, nil
		} else if paridad_lider != paridad_equipo1 && paridad_lider == paridad_equipo2 {
			for i := 0; i < int(s.largo_equipo); i++ {
				if in.Jugador == s.equipo1[i] {
					return &Message{Body: "Cagaste", Aux: true, NumRonda: in.NumRonda, NumJuego: 2}, nil
				}
			}
			return &Message{Body: "Eres del equipo ganador", Aux: true}, nil
		} else if paridad_lider != paridad_equipo1 && paridad_lider != paridad_equipo2 {
			equipo_eliminado := rand.Intn(1) + 1
			if equipo_eliminado == 1 {
				for i := 0; i < int(s.largo_equipo); i++ {
					if in.Jugador == s.equipo2[i] {
						return &Message{Body: "Cagaste", Aux: true, NumRonda: in.NumRonda, NumJuego: 2}, nil
					}
				}
				return &Message{Body: "Eres del equipo ganador", Aux: true, NumRonda: in.NumRonda, NumJuego: 2}, nil
			} else {
				for i := 0; i < int(s.largo_equipo); i++ {
					if in.Jugador == s.equipo1[i] {
						return &Message{Body: "Cagaste", Aux: true, NumRonda: in.NumRonda, NumJuego: 2}, nil
					}
				}
				return &Message{Body: "Eres del equipo ganador", Aux: true, NumRonda: in.NumRonda, NumJuego: 2}, nil
			}
		}
	} else if in.NumJuego == 3 {
		for i := 0; i < 8; i++ {
			if s.num_lider_juego3 == 0 || s.num_elegido_jugadores[i] == 0 {
				return &Message{Aux: false}, nil
			}
		}
		for i := 0; i < 4; i++ {
			abs_jug1 := math.Abs(float64(s.num_elegido_jugadores[2*i] - s.num_lider_juego3))
			abs_jug2 := math.Abs(float64(s.num_elegido_jugadores[2*i+1] - s.num_lider_juego3))
			if abs_jug1 == abs_jug2 {
				// ambos ganan
				for j := 0; j < 8; i++ {
					if s.ganadores[j] == 0 {
						s.ganadores[j] = s.equipos_juego3[2*i]
						s.ganadores[j+1] = s.equipos_juego3[2*i+1]
						break
					}
				}
			} else if abs_jug1 < abs_jug2 {
				// gana el jug 1 de ese equipo
				for j := 0; j < 8; i++ {
					if s.ganadores[j] == 0 {
						s.ganadores[j] = s.equipos_juego3[2*i]
						for k := 0; k < 4; k++ {
							if s.perdedores[k] == 0 {
								s.perdedores[k] = s.equipos_juego3[2*i+1]
							}
						}
						break
					}
				}
			} else if abs_jug2 < abs_jug1 {
				// gana el jug 2 de ese equipo
				for j := 0; j < 8; i++ {
					if s.ganadores[j] == 0 {
						s.ganadores[j] = s.equipos_juego3[2*i+1]
						for k := 0; k < 4; k++ {
							if s.perdedores[k] == 0 {
								s.perdedores[k] = s.equipos_juego3[2*i]
							}
						}
						break
					}
				}
			}
		}
		for i := 0; i < 8; i++ {
			if in.Jugador == s.ganadores[i] {
				return &Message{Body: "Eres el ganador de los juegos", Aux: true}, nil
			} else if i < 4 && in.Jugador == s.perdedores[i] {
				return &Message{Body: "Cagaste", Aux: true, NumRonda: in.NumRonda, NumJuego: 3}, nil
			}
		}
	}
	return &Message{Body: "No se escogio ningun juego"}, nil
}

func (s *Server) Jugada(ctx context.Context, in *Message) (*Message, error) {
	if in.NumJuego == 1 {
		if s.juegos_iniciados[in.NumJuego-1] && s.ronda_iniciada {
			if in.Jugada >= int32(s.num_lider_1) {
				s.array_jugadas1_comp[in.Jugador-1] = s.array_jugadas1_comp[in.Jugador-1] + 1
				s.array_jugador[in.Jugador-1] = 0
				return &Message{Body: "Cagaste", Aux: true, NumRonda: int32(s.array_jugadas1_comp[in.Jugador-1])}, nil
			} else {
				s.array_jugadas1[in.Jugador-1] = s.array_jugadas1[in.Jugador-1] + int(in.Jugada)
				s.array_jugadas1_comp[in.Jugador-1]++
				if s.array_jugadas1_comp[in.Jugador-1] == 3 {
					if s.array_jugadas1_comp[in.Jugador-1] < 21 {
						s.array_jugador[in.Jugador-1] = 0
						return &Message{Body: "Cagaste", Aux: true, NumRonda: int32(s.array_jugadas1_comp[in.Jugador-1])}, nil
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
				return &Message{Body: "Cagaste", Aux: true, NumJuego: 2}, nil
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
	} else if in.NumJuego == 3 {
		if s.juegos_iniciados[1] && s.ronda_iniciada {
			if in.Jugador == s.asesinado {
				return &Message{Body: "Cagaste", Aux: true, NumJuego: 2}, nil
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

	if s.pidiendo == 1 {
		s.pidiendo = 0
		return &Message{Body: "Jugador pidiendo monto"}, nil
	} else {
		return &Message{Body: "Nadie pidiendo"}, nil
	}
}

func (s *Server) PedirMonto(ctx context.Context, in *Message) (*Message, error) {

	s.pidiendo = 1
	return &Message{Body: "Pidiendo monto"}, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *Server) Muerto(ctx context.Context, in *Message) (*Message, error) {

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

	s.jugadores = [16]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, numero := range s.jugadores {
		if numero == s.muerto {
			s.muerto = 1000
			s.ronda = 1000
			return &Message{Body: "Lo encontramos muerto", Jugador: s.muerto, Ronda: s.ronda}, nil
		}
	}
	return &Message{Body: "Nadie muerto"}, nil
}

func (s *Server) PedirMontoPozo(ctx context.Context, in *Message) (*Message, error) {

	s.pedido = 1
	return &Message{Body: "Pidiendo monto"}, nil
}

func (s *Server) EsperarPeticionPozo(ctx context.Context, in *Message) (*Message, error) {

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

	if s.jugador_pidiendo == 1 {
		s.jugador_pidiendo = 0
		return &Message{Monto: s.monto}, nil
	} else {
		return &Message{Body: "Nada aun"}, nil
	}
}

func (s *Server) EntregarLider(ctx context.Context, in *Message) (*Message, error) {

	if s.lider_pidiendo == 1 {
		s.lider_pidiendo = 0
		return &Message{Monto: s.monto}, nil
	} else {
		return &Message{Monto: -100}, nil
	}
}

func (s *Server) EsperandoRegistro(ctx context.Context, in *Message) (*Message, error) {
	if s.registro == 1 {
		s.registro = 0
		return &Message{Body: "Registro nuevo", Jugador: s.jugador_name, Ronda: s.ronda_name, Jugadas: s.jugada_name}, nil
	} else {
		return &Message{Body: "Aun nada"}, nil
	}
}

func (s *Server) DandoRegistro(ctx context.Context, in *Message) (*Message, error) {

	s.jugador_name = in.Jugador
	s.ronda_name = in.Ronda
	s.jugada_name = in.Jugadas
	s.registro = 1

	return &Message{Body: "Enviando registro"}, nil
}
