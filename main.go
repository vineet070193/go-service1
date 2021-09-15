package main

import (
	"net"
	"os"

	"github.com/vineet070193/service1/server"

	"github.com/hashicorp/go-hclog"
	protos "github.com/vineet070193/service1/protos/chat"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewChat(log)

	protos.RegisterChatServer(gs, cs)

	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(lis)

}