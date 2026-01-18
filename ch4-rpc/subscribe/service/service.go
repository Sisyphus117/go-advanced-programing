package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/subscribe/pubsub"
	"google.golang.org/grpc"
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
	grpcServer := grpc.NewServer()
	svc := NewPubsubService()
	pubsub.RegisterPubsubServiceServer(grpcServer, svc)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
