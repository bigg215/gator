package main

import (
	"context"
	"fmt"
	"os"
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
