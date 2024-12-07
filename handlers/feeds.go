package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
)

func HandlerFeeds(s *commands.State, cmd commands.Command) error {
	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds, %w\n", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found!")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
		fmt.Printf("%s\n", feed.Url)
	}

	return nil
}
