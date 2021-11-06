package main

import (
	"context"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"example.com/m/chat/github.com/Benja-99/tarea1_SD/chat"
	"google.golang.org/grpc"
)

func distribuirData() string {
	var datas [3]string
	datas[0] = "10.0.1.10"
	datas[1] = "10.0.1.11"
	datas[2] = "10.0.1.12"

	rand.Seed(int64(time.Now().UnixNano()))
	aleatorio := rand.Intn(3)

	return datas[aleatorio]
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	log.Printf("El namenode se esta ejecutando")
	c := chat.NewChatServiceClient(conn)
	var flag bool = true
	for flag {
		aleatorio := distribuirData()
		var response *chat.Message
		var err error
		var flag1 bool = true
		for flag1 {
			response, err = c.EsperandoRegistro(context.Background(), &chat.Message{Body: "Esperando registro"})
			if err != nil {
				log.Fatalf("Error when calling Peticion: %s", err)
			}
			if response.Body == "Registro nuevo" {
				flag = false
			}
			time.Sleep(3 * time.Second)
		}
		a := "Jugador_" + string(response.Jugador) + " Ronda_" + string(response.Ronda) + " " + aleatorio
		b := []byte(a)
		err1 := ioutil.WriteFile("pozo.txt", b, 0644)
		if err1 != nil {
			log.Fatal(err1)
		}
	}

}
