package main

import (
	"context"
	"fmt"
	"os"
)

func handlerAGG(s *state, cmd command) error {
	feeds, err := fetchfeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Println(feeds)
	os.Exit(0)
	return nil
}
