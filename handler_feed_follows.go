package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DanilShapilov/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <url>", cmd.Name)
	}
	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			UserID:    user.ID,
			FeedID:    feed.ID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed followed:")
	fmt.Printf(" * User:     %v\n", feedFollow.UserName)
	fmt.Printf(" * Feed:     %v\n", feedFollow.FeedName)

	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <url>", cmd.Name)
	}
	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}
	fmt.Printf("%s unfollowed successfully!\n", feed.Name)
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("User %s is following:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf(" * %v\n", feedFollow.FeedName)
	}

	return nil
}
