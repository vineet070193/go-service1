package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	feedprotos "github.com/vineet070193/go-service1/protos/feed"
)

type FeedService struct {
	log hclog.Logger
}

func NewFeedService(l hclog.Logger) *FeedService {
	return &FeedService{l}
}

func (f *FeedService) AddFeed(ctx context.Context, r *feedprotos.AddFeedRequest) (*feedprotos.AddFeedResponse, error) {
	f.log.Info("Handle AddFeed", "feed_id", r.Feed.FeedID, "snippet", r.Feed.Snippet, "Content", r.Feed.Content, "Timestamp", r.Feed.Timestamp)

	return &feedprotos.AddFeedResponse{
		Status: &feedprotos.Status{
			Error:  nil,
			Status: true,
		},
	}, nil
}

func (f *FeedService) ListFeed(ctx context.Context, r *feedprotos.ListFeedRequest) (*feedprotos.ListFeedResponse, error) {
	f.log.Info("Handle AddFeed", "feed_id", r.FeedID)

	return &feedprotos.ListFeedResponse{
		Status: &feedprotos.Status{
			Error:  nil,
			Status: true,
		},
		FeedList: []*feedprotos.Feed{
			{
				FeedID:    "FeedID1",
				Snippet:   "Snippet1",
				Content:   "Content",
				Timestamp: 123456,
			},
		},
	}, nil
}
