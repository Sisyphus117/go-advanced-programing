package main

import (
	"context"

	hello "github.com/sisyphus/go-advanced-programming/ch4-rpc/token/api"
	"github.com/sisyphus/go-advanced-programming/ch4-rpc/token/internal/auth"
)

type grpcServer struct{ auth *auth.Authentication }

func (p *grpcServer) SomeMethod(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}
	return &hello.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {}
