package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
)

func HandlerFollowing(s *commands.State, cmd commands.Command, user database.User) error {
	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("can't get feeds for user %s, %v", user.Name, err)
	}

	fmt.Printf("User: %s\n", user.Name)

	for _, feed := range feeds {
		fmt.Printf("- %s\n", feed.Name)
	}

	return nil
}
