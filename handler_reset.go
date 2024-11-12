package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAll(context.Background())

	if err != nil {
		return err
	}

	fmt.Println("All users deleted successfully")
	os.Exit(0)
	return nil
}
