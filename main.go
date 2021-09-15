package main

import (
	"net"
	"os"

	"github.com/vineet070193/go-service1/server"
	// "github.com/vineet070193/go-service1/tree/main/server"

	"github.com/hashicorp/go-hclog"
	protos "github.com/vineet070193/go-service1/protos/chat"

	// protos "github.com/vineet070193/go-service1/tree/main/protos/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewChat(log)

	protos.RegisterChatServer(gs, cs)

	// Disable this flag in production as it is used to test the service using grpcurl locally.
	reflection.Register(gs)

	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(lis)

}
