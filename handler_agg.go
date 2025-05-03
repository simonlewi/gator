package main

import (
	"context"
	"fmt"
)

func aggCommand(s *state, cmd command) error {
	ctx := context.Background()

	feedURL := cmd.Args[0]
	feed, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil

}
