package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("* %s ", user.Name)
		if s.cfg.CurrentUserName == user.Name {
			fmt.Print("(current)")
		}
		fmt.Println()
	}
	return nil
}
