package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-ex/chat"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chat.Message{Subject: "Client hello", Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Client error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
