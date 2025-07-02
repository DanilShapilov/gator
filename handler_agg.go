package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DanilShapilov/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <time_between_reqs>", cmd.Name)
	}

	durationStr := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s", durationStr)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get feed to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}
func scrapeFeed(db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(
		context.Background(),
		database.MarkFeedFetchedParams{
			LastFetchedAt: sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			},
			UpdatedAt: time.Now().UTC(),
			ID:        feed.ID,
		},
	)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
		pubDate := sql.NullTime{}
		if t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate); err == nil {
			pubDate = sql.NullTime{
				Time:  t.UTC(),
				Valid: true,
			}
		}

		post, err := db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				Title:     item.Title,
				Url:       item.Link,
				Description: sql.NullString{
					String: item.Description,
					Valid:  true,
				},
				PublishedAt: pubDate,
				FeedID:      feed.ID,
			},
		)
		if err != nil && !strings.Contains(err.Error(), `duplicate key value violates unique constraint "posts_url_key"`) {
			log.Printf("ERROR ON POST CREATE: %v", err)
		}
		if err == nil {
			fmt.Printf("Post %v added, URL:%v\n", post.Title, post.Url)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
