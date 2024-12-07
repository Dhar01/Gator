package handlers

import (
	"context"
	"fmt"

	"github.com/Dhar01/Gator/commands"
)

func HandlerLogin(s *commands.State, cmd commands.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("USAGE: %s <name>\n", cmd.Name)
	}

	username := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w\n", err)
	}

	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("couldn't set current user: %w\n", err)
	}

	fmt.Printf("User %s login successfully!\n", username)

	return nil
}
