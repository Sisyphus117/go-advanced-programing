package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &api.String{Value: "client"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
