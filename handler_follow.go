package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/simonlewi/gator/internal/database"
)

func followHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	feedURL := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		rssFeed, err := fetchFeed(context.Background(), feedURL)
		if err != nil {
			return fmt.Errorf("error fetching feed content from %s: %w", feedURL, err)
		}

		feedParams := database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      rssFeed.Channel.Title,
			Url:       feedURL,
			UserID:    user.ID,
		}
		feed, err = s.db.CreateFeed(context.Background(), feedParams)
		if err != nil {
			return fmt.Errorf("error creating feed: %w", err)
		}
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("You are now following %s\n", feedFollow.FeedName)
	return nil
}
