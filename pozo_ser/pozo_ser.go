package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func pozo1(conn *grpc.ClientConn, mont int) {

	c := chat.NewChatServiceClient(conn)
	var response *chat.Message
	var err error
	var flag bool = true
	for flag {
		response, err = c.EsperarPeticionPozo(context.Background(), &chat.Message{Monto: int32(mont)})
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		if response.Monto != int32(mont) {
			flag = false
		}
		time.Sleep(3 * time.Second)
	}
	log.Printf("Monto entregado")
}

func muerto(conn *grpc.ClientConn, mont int) {
	c := chat.NewChatServiceClient(conn)
	var response *chat.Message
	var err error
	var flag bool = true
	for flag {
		response, err = c.SacarMuerto(context.Background(), &chat.Message{Body: "Esperando muertos"})
		if err != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		if response.Body == "Lo encontramos muerto" {
			flag = false
		}
		time.Sleep(10)
	}
	a := "Jugador_" + string(response.Jugador) + " Ronda_" + string(response.Ronda) + " " + strconv.Itoa(mont)
	b := []byte(a)
	err1 := ioutil.WriteFile("pozo.txt", b, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}

}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.185:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	var monto int = 0

	log.Printf("El pozo se esta ejecutando")
	conn1, err1 := amqp.Dial("amqp://guest:guest@172.17.0.1:5672/")

	if err1 != nil {
		log.Fatal(err1)
	}
	defer conn1.Close()

	ch, err := conn1.Channel()

	if err != nil {
		log.Fatal(err1)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello1", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			intVar, err := strconv.Atoi(string(d.Body))
			if err != nil {
				log.Fatal(err)
			}
			monto = monto + intVar
			log.Printf("Monto actualizado: %d", monto)
		}
	}()

	for i := 0; i < 4; i++ {
		go pozo1(conn, monto)
	}
	for i := 0; i < 16; i++ {
		go muerto(conn, monto)
	}
	fmt.Scanln()
}
