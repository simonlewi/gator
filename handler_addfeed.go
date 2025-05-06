package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/simonlewi/gator/internal/database"
)

func addFeedHandler(s *state, cmd command, user database.User) error {
	// Existing feed creation code to validate arguments and get name/url
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	// Create the feed with your existing code
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
	}

	/* Get current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}*/

	// Set the feed's creator
	feedParams.UserID = user.ID

	// Create the feed
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Printf("Feed '%s' added successfully!\n", feedName)

	// Now automatically create a feed follow for this user
	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("You are now following '%s'\n", feedName)
	return nil
}
