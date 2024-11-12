package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bigg215/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	currentUserName, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return err
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUserName.ID,
	}

	feedResp, err := s.db.CreateFeed(context.Background(), newFeed)

	if err != nil {
		return err
	}

	fmt.Println(feedResp)
	return nil
}
