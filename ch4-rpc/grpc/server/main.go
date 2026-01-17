package main

import (
	"context"
	"log"
	"net"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/grpc/api"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct {
	api.UnimplementedHelloServiceServer
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *api.String) (*api.String, error) {
	reply := &api.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	api.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
