package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) < 2 {
		fmt.Printf("USAGE: addfeed <name> <feed_link>\n")
		return fmt.Errorf("%s command, wrong structure\n", cmd.Name)
	}

	username := s.Config.CurrentUserName
	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("ERROR: couldn't find user!\n")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	data := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), data)
	if err != nil {
		return fmt.Errorf("can't create feed: %v\n", err)
	}


	followFeed, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("can't follow feed: %w", err)
	}

	fmt.Printf("Adding feed %s is done!\n", name)
	fmt.Printf("%s is now following %s feed.\n", user.Name, name)

	fmt.Println(followFeed)

	return nil
}
