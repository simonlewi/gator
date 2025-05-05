package main

import (
	"context"
	"fmt"
)

func followingHandler(s *state, cmd command) error {
	// Get current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("You're not following any feeds")
		return nil
	}

	fmt.Println("Feeds you're following:")
	for _, follow := range feedFollows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}
