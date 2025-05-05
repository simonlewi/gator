package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/simonlewi/gator/internal/database"
)

func followHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	feedURL := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
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
