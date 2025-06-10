package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DanilShapilov/gator/internal/database"
)

func handlerAdd(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("failed to create a feed: %w", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:       %v\n", feed.ID)
	fmt.Printf(" * Name:     %v\n", feed.Name)
	fmt.Printf(" * URL:      %v\n", feed.Url)
	fmt.Printf(" * UserID:   %v\n", feed.UserID)
	fmt.Printf(" * Created:   %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated:   %v\n", feed.UpdatedAt)
}
