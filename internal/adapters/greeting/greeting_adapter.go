package adapters

import (
	"context"
	"log"

	"github.com/jmarc101/grpc-go-client/internal/ports"
	"github.com/jmarc101/grpc-protobuf/protogen/go/greeting"
	"google.golang.org/grpc"
)

type GreetingAdapter struct {
	greetingClient ports.GreetingServicePort
}

func NewGreetingAdapter(conn *grpc.ClientConn) (*GreetingAdapter, error){
	client := greeting.NewGreetingServiceClient(conn)

	return &GreetingAdapter{
		greetingClient: client,
	}, nil
}

func (a *GreetingAdapter) Greet(ctx context.Context, name string) (*greeting.GreetingsResponse, error) {
	req := greeting.GreetingsRequest{
		Name: name,
	}

	greet, err := a.greetingClient.Greet(ctx, &req)
	if err != nil {
		log.Fatalln(err)
	}

	return greet, nil
}
