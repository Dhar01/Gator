package handlers

import (
	"errors"
	"fmt"

	cmd "github.com/Dhar01/Gator/commands"
)

var errNoUsername = errors.New("username is required!")

func HandlerLogin(s *cmd.State, cmd cmd.Command) error {
	if len(cmd.Args) < 1 {
		fmt.Printf("usage: %s <name>\n", cmd.Name)
		return errNoUsername
	}

	username := cmd.Args[0]

	if err := s.Config.SetUser(username); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("the user '%s' has been set\n", username)
	return nil
}
