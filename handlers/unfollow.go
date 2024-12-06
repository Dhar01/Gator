package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
)

func HandlerUnfollow(s *commands.State, cmd commands.Command, user database.User) error {
	if len(cmd.Args) < 1 {
		fmt.Printf("USAGE: %s <url>\n", cmd.Name)
		return fmt.Errorf("url needed to unfollow\n")
	}

	url := cmd.Args[0]

	if err := s.DB.FeedUnfollow(context.Background(), database.FeedUnfollowParams{
		UserID: user.ID,
		Url: url,
	}); err != nil {
		return fmt.Errorf("can't unfollow feed: %w\n", err)
	}

	fmt.Printf("Successfully unfollowed %s\n", url)

	return nil
}