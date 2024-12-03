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
		fmt.Printf("%s\n", feed.Name_2)
	}

	return nil
}

func HandlerFetch(s *commands.State, cmd commands.Command) error {
	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		log.Printf("ERROR: can't get user, %v\n", err)
		return err
	}

	feeds, err := s.DB.FetchFeeds(context.Background(), user.ID)
	if err != nil {
		log.Printf("ERROR: can't get feeds, %v\n", err)
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("Time: %s\n", feed.UpdatedAt)
		fmt.Printf("URL: %s\n", feed.Url)
	}

	return nil
}
