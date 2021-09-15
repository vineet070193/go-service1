package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/vineet070193/service1/protos/chat"
)

type Chat struct {
	log hclog.Logger
}

func NewChat(l hclog.Logger) *Chat {
	return &Chat{l}
}

func (c *Chat) Greeting(ctx context.Context, r *protos.GreetRequest) (*protos.GreetREsponse, error) {
	c.log.Info("Hadle Greeting", "first_name", r.FirstName(), "last_name", r.LastName())

	return &protos.GreetResponse{
		Greeting: "Hello there",
		Name:     r.FirstName(),
	}, nil
}
