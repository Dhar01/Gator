package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *commands.State, cmd commands.Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("USAGE: %s <name> <url>\n", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w\n", err)
	}

	fmt.Println("Feed created successfully!")

	if _, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("couldn't follow feed: %w\n", err)
	}

	fmt.Printf("%s is now following %s feed.\n", user.Name, name)

	return nil
}
