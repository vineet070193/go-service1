package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/vineet070193/go-service1/protos/chat"
	// protos "github.com/vineet070193/go-service1/tree/main/protos/chat"
)

type Chat struct {
	log hclog.Logger
}

func NewChat(l hclog.Logger) *Chat {
	return &Chat{l}
}

func (c *Chat) Greet(ctx context.Context, r *protos.GreetRequest) (*protos.GreetResponse, error) {
	c.log.Info("Hadle Greeting", "first_name", r.FirstName, "last_name", r.LastName)

	return &protos.GreetResponse{
		Greeting: "Hello there",
		Name:     r.FirstName,
	}, nil
}
