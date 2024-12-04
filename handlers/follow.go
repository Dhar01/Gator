package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
)

func HandlerFollow(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) < 1 {
		fmt.Printf("USAGE: %s <URL>\n", cmd.Name)
		return fmt.Errorf("URL missing\n")
	}

	url := cmd.Args[0]
	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("can't get user!\n")
	}

	// check if a feed exists
	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		fmt.Printf("%s feed not found, creating...\n", url)
		return err
	}

	followFeed, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	// fmt.Printf("Feed Name: %s", followFeed.FeedName)
	// fmt.Printf("User Name: %s", followFeed.UserName)
	fmt.Println(followFeed)

	return nil
}
