package main

import (
	"fmt"
	"log"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/token/internal/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = "1234"

func main() {
	auth := auth.Authentication{
		User:     "gopher",
		Password: "password",
	}

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("ok")
}
