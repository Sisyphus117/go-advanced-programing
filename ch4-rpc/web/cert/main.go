package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

	_ = grpc.NewServer(grpc.Creds(creds))
}
