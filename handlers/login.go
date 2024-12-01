package handlers

import (
	"context"
	"fmt"

	cmd "github.com/Dhar01/Gator/commands"
)

func HandlerLogin(s *cmd.State, cmd cmd.Command) error {
	if len(cmd.Args) < 1 {
		fmt.Printf("usage: %s <name>\n", cmd.Name)
		return errNoUsername
	}

	username := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s login successfully!\n", username)

	return nil
}
