package main

import (
	"context"
	"fmt"

	"github.com/simonlewi/gator/internal/database"
)

func unfollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		fmt.Println("Error: URL is required for the unfollow command")
		return nil
	}

	feedURL := cmd.Args[0]

	params := database.UnfollowFeedParams{
		Url:    feedURL,
		UserID: user.ID,
	}

	_, err := s.db.UnfollowFeed(context.Background(), params)
	if err != nil {
		fmt.Println("Error unfollowing feed:", err)
		return nil
	}

	fmt.Println("Feed unfollowed successfully!")
	return nil
}
