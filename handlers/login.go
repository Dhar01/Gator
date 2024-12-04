package handlers

import (
	"context"
	"fmt"

	cmd "github.com/Dhar01/Gator/commands"
)

func HandlerLogin(s *cmd.State, cmd cmd.Command) error {
	if len(cmd.Args) < 1 {
		fmt.Printf("USAGE: %s <name>\n", cmd.Name)
		return errNoUsername
	}

	username := cmd.Args[0]

	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("ERROR: couldn't find user: %w\n", err)
	}

	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("ERROR: couldn't set current user: %w\n", err)
	}

	fmt.Printf("User %s login successfully!\n", username)

	return nil
}
