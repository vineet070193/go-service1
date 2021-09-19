package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	feedprotos "github.com/vineet070193/go-service1/protos/feed"
	"google.golang.org/grpc"
)

// pingRequest sends a new gRPC ping request to the server configured in the connection.
func sendAddFeedRequest(conn *grpc.ClientConn, p *feedprotos.AddFeedRequest) (*feedprotos.AddFeedResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := feedprotos.NewFeedServiceClient(conn)
	return client.AddFeed(ctx, p)
}

// NewConn creates a new gRPC connection.
// host should be of the form domain:port, e.g., example.com:443
func NewConn(host string, insecure bool) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}

	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	return grpc.Dial(host, opts...)
}

func invokeClient() {
	conn, err := NewConn("fleet-breaker-325818.el.r.appspot.com:8080", true)
	if err != nil {
		fmt.Printf("Error creating the client connection: %v\n", err)
		os.Exit(1)
	}

	r := &feedprotos.AddFeedRequest{
		Feed: &feedprotos.Feed{
			FeedID:    "feed_id1",
			Snippet:   "This is a feed snippet",
			Content:   "This is a feed snippet. This would be added as a feed",
			Timestamp: 11223,
		},
	}

	resp, err := sendAddFeedRequest(conn, r)
	fmt.Printf("Response: %v\n", resp)
	fmt.Printf("Error: %v\n", err)

}
