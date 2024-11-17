package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bigg215/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAGG(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <time between reqs>", cmd.Name)
	}

	parsedTimeDuration, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}

	time_between_reqs := time.Duration(parsedTimeDuration)

	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())

	if err != nil {
		fmt.Printf("error getting next feed: %s\n", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)

	if err != nil {
		fmt.Printf("error marking fetched: %s\n", err)
	}

	RSSFeed, err := fetchfeed(context.Background(), nextFeed.Url)

	if err != nil {
		fmt.Printf("error fetching rss feed: %s\n", err)
	}

	fmt.Printf("START:\t%s\n\n", nextFeed.Url)
	for _, item := range RSSFeed.Channel.Item {
		publishedAt := sql.NullTime{}

		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	fmt.Printf("END:\t%s\nPOSTS:\t%v\n\n", nextFeed.Name, len(RSSFeed.Channel.Item))
}
