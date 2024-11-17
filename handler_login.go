package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bigg215/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name)

	if err != nil {
		return fmt.Errorf("user doesnt exist")
	}

	s.cfg.SetUser(user.Name)
	fmt.Println("User switched successfully!")
	os.Exit(0)
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)

	if err != nil {
		return fmt.Errorf("feed not found: %w", err)
	}

	ffrow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("unable to create follower: %w", err)
	}

	fmt.Println("Feed follow created:")
	fmt.Printf("User:\t%s\n", ffrow.UserName)
	fmt.Printf("Feed:\t%s\n", ffrow.FeedName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {

	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("error retrieving feeds followed: %w", err)
	}

	if len(feedsFollowed) == 0 {
		fmt.Printf("0 feeds are followed by user: %s\n", s.cfg.CurrentUserName)
	} else {
		fmt.Printf("%d feeds followed by user: %s\n", len(feedsFollowed), s.cfg.CurrentUserName)
	}

	for _, ff := range feedsFollowed {
		fmt.Printf("Feed:\t%s\n", ff.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	err := s.db.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		Url:    url,
	})

	if err != nil {
		return fmt.Errorf("unable to delete feed follow error: %w", err)
	}

	return nil
}
