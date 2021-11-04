package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"example.com/m/pozo/github.com/Benja-99/tarea1_SD/pozo"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func pozo1(conn *grpc.ClientConn) {

	c := chat.NewChatServiceClient(conn)
	c_pozo := pozo.NewPozoServiceClient(conn)

	var response *pozo.Monto
	var err error
	response, err = c_pozo.GetMonto(context.Background(), &pozo.Monto{CantidadTotal: 1})
	if err != nil {
		log.Fatalf("Error when calling Peticion: %s", err)
	}
	log.Printf("Response from server: %d", response.CantidadTotal)

	var response1 *chat.Message
	var err1 error
	response1, err1 = c.EsperarPeticion(context.Background(), &chat.Message{Monto: int32(response.CantidadTotal)})
	if err1 != nil {
		log.Fatalf("Error when calling Peticion: %s", err)
	}
	if response1.Body == "Pidiendo monto" {
		var response2 *chat.Message
		var err2 error
		response2, err2 = c.EsperarPeticion(context.Background(), &chat.Message{Monto: int32(response.CantidadTotal)})
		if err2 != nil {
			log.Fatalf("Error when calling Peticion: %s", err)
		}
		log.Printf("Response from server: %s", response2.Body)
	}

}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	var monto int = 0

	conn1, err1 := amqp.Dial("amqp://guest:guest@localhost:5672/")

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

	forever := make(chan bool)

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

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	go pozo1(conn)

	fmt.Scanln()
}
