package ports

import (
	"context"

	"github.com/jmarc101/grpc-protobuf/protogen/go/greeting"
	"google.golang.org/grpc"
)

// GreetingPort is the interface that wraps the Greet method.
type GreetingServicePort interface {
	Greet(ctx context.Context, in *greeting.GreetingsRequest, opts ...grpc.CallOption) (*greeting.GreetingsResponse, error)
}
