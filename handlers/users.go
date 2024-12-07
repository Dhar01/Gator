package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
)

func HandlerUsers(s *commands.State, cmd commands.Command) error {
	users, err := s.DB.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't find users: %v\n", err)
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Println("*", user.Name)
		}
	}

	return nil
}
