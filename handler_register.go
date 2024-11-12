package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bigg215/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)

	if err == nil {
		return fmt.Errorf("user already exists")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	_, err = s.db.CreateUser(context.Background(), newUser)

	if err != nil {
		return err
	}

	s.cfg.SetUser(name)
	fmt.Printf("User created successfully!: %s\n", name)
	os.Exit(0)
	return nil
}
