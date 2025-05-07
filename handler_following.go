package main

import (
	"context"
	"fmt"

	"github.com/simonlewi/gator/internal/database"
)

func followingHandler(s *state, cmd command, user database.User) error {
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
