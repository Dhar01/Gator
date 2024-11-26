package handlers

import (
	"errors"
	"fmt"

	cmd "github.com/Dhar01/Gator/commands"
)

var errNoUsername = errors.New("username is required!")

func HandlerLogin(s *cmd.State, cmd cmd.Command) error {
	if len(cmd.Argument) < 1 {
		return errNoUsername
	}

	username := cmd.Argument[0]
	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("the user '%s' has been set\n", username)
	return nil
}
