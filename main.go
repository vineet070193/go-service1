package main

import (
	"net"
	"os"

	"github.com/vineet070193/go-service1/server"

	hclog "github.com/hashicorp/go-hclog"
	chatprotos "github.com/vineet070193/go-service1/protos/chat"
	feedprotos "github.com/vineet070193/go-service1/protos/feed"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewChat(log)
	fs := server.NewFeedService(log)

	chatprotos.RegisterChatServer(gs, cs)
	feedprotos.RegisterFeedServiceServer(gs, fs)

	// Disable this flag in production as it is used to test the service using grpcurl locally.
	// This can be added behind an environment flag for safer testing.
	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(lis)

}
