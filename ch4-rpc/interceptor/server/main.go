package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic:%v", r)
		}
	}()
	return handler(ctx, req)
}

func main() {
	_ = grpc.NewServer(grpc.UnaryInterceptor(filter))
}
