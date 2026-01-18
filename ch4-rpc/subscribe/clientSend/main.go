package main

import (
	"context"
	"log"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/subscribe/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pubsub.NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(), &pubsub.String{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &pubsub.String{Value: "docker: hello docker"})
	if err != nil {
		log.Fatal(err)
	}
}
