package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
)

func HandlerUnfollow(s *commands.State, cmd commands.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("USAGE: %s <url>\n", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w\n", err)
	}

	if err := s.DB.FeedUnfollow(context.Background(), database.FeedUnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w\n", err)
	}

	fmt.Printf("Successfully unfollowed %s\n", url)

	return nil
}
