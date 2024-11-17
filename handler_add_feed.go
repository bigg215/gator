package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bigg215/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feedResp, err := s.db.CreateFeed(context.Background(), newFeed)

	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedResp.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Println("Feed follow created successfully.")
	fmt.Println(feedFollow)
	return nil
}
