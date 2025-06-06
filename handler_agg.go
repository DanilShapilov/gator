package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	const tmpUrl = "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), tmpUrl)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v", *feed)
	return nil
}
