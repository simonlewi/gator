package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/simonlewi/gator/internal/database"
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

	fmt.Printf("Found %d posts in feed: %s\n", len(rssFeed.Channel.Item), feed.Name)
	savedCount := 0

	for _, item := range rssFeed.Channel.Item {
		var publishedAt time.Time
		if item.PubDate != "" {
			parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				parsedTime, err = time.Parse(time.RFC822, item.PubDate)
				if err != nil {
					publishedAt = time.Now().UTC()
				} else {
					publishedAt = parsedTime
				}
			} else {
				publishedAt = parsedTime
			}
		} else {
			publishedAt = time.Now().UTC()
		}

		var description sql.NullString
		if item.Description != "" {
			description = sql.NullString{
				String: item.Description,
				Valid:  true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			// Check if it's a duplicate URL error - these should be ignored
			if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
				// Silently continue with the next post
				continue
			}
			// Log other errors but don't fail the whole scrape
			fmt.Printf("Error saving post %s: %v\n", item.Title, err)
		} else {
			savedCount++
		}
	}

	fmt.Printf("Saved %d new posts from feed: %s\n", savedCount, feed.Name)

	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	return nil
}
