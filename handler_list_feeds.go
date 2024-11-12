package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		userID, _ := s.db.GetUserName(context.Background(), feed.UserID)
		fmt.Printf("Name: %s\tURL: %s\tUSER: %s\n", feed.Name, feed.Url, userID)
	}
	return nil
}
