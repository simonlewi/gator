package main

import (
	"context"
	"fmt"
)

func scrapeFeedHandler(s *state, cmd command) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed content from %s: %w", feed.Url, err)
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	return nil
}
