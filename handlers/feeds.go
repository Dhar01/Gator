package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/Dhar01/Gator/commands"
)

func HandlerFeeds(s *commands.State, cmd commands.Command) error {
	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		log.Printf("ERROR: can't get feeds, %v\n", err)
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
		fmt.Printf("%s\n", feed.Url)
	}

	return nil
}
