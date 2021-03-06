package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	chatprotos "github.com/vineet070193/go-service1/protos/chat"
)

type Chat struct {
	log hclog.Logger
}

func NewChat(l hclog.Logger) *Chat {
	return &Chat{l}
}

func (c *Chat) Greet(ctx context.Context, r *chatprotos.GreetRequest) (*chatprotos.GreetResponse, error) {
	c.log.Info("Hadle Greeting", "first_name", r.FirstName, "last_name", r.LastName)

	return &chatprotos.GreetResponse{
		Greeting: "Hello there",
		Name:     r.FirstName,
	}, nil
}
