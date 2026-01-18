package main

import (
	"context"
	"net"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/RESTful/api"
	"google.golang.org/grpc"
)

type RestServiceImpl struct {
	api.UnimplementedRestServiceServer
}

func (r *RestServiceImpl) Get(ctx context.Context, message *api.StringMessage) (*api.StringMessage, error) {
	return &api.StringMessage{Value: "Get hi: " + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *api.StringMessage) (*api.StringMessage, error) {
	return &api.StringMessage{Value: "Post hi: " + message.Value + "@"}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	api.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	lis, _ := net.Listen("tcp", ":5000")
	grpcServer.Serve(lis)
}
