package main

import (
	"context"
	"fmt"
	"io"
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

	stream, err := client.Subscribe(context.Background(), &pubsub.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Println(reply.GetValue())
	}
}
