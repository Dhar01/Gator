package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAggregate(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("USAGE: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenReq, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenReq)

	ticker := time.NewTicker(timeBetweenReq)

	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *commands.State) {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get next feeds to fetch")
		return
	}

	log.Println("found a feed to fetch!")
	scrapeFeed(s.DB, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	if _, err := db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		log.Printf("couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}

		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			PublishedAt: publishedAt,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found\n", feed.Name, len(feedData.Channel.Item))
}
