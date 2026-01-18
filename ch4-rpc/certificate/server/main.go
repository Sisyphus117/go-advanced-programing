package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/subscribe/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type PubsubService struct {
	pub *pubsub.Publisher
	pubsub.UnimplementedPubsubServiceServer
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *pubsub.String) (*pubsub.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pubsub.String{}, nil
}

func (p *PubsubService) Subscribe(arg *pubsub.String, stream pubsub.PubsubService_SubscribeServer) error {
	ch := p.pub.Subscribe(func(v any) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&pubsub.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(grpc.Creds(creds))
	svc := NewPubsubService()
	pubsub.RegisterPubsubServiceServer(server, svc)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	server.Serve(lis)

}
