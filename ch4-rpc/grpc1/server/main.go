package main

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/grpc1/api"
	"google.golang.org/grpc"
)

type HelloServiceImpl struct {
	api.UnimplementedHelloServiceServer
}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *api.String) (*api.String, error) {
	reply := &api.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream grpc.BidiStreamingServer[api.String, api.String]) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &api.String{Value: "hello: " + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}

	}
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
