package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
)

func HandlerFeeds(s *commands.State, cmd commands.Command) error {
	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
		fmt.Printf("%s\n", feed.Url)
		fmt.Printf("%s\n", feed.Name_2)
	}

	return nil
}
