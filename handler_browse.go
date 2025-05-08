package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/simonlewi/gator/internal/database"
)

func browseHandler(s *state, cmd command, user database.User) error {
	limit := int32(2)

	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err == nil && parsedLimit > 0 {
			limit = int32(parsedLimit)
		} else if err != nil {
			return fmt.Errorf("invalid limit: %s", cmd.Args[0])
		}
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error fetching posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found. Try following some feeds!")
		return nil
	}

	for _, post := range posts {
		fmt.Println("Title:", post.Title)
		fmt.Println("URL:", post.Url)
		fmt.Println("Published:", post.PublishedAt.Format("2006-01-02 15:04:05"))
		fmt.Println("---")
	}

	return nil
}
