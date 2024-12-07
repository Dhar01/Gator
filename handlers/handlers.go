package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dhar01/Gator/commands"
)

var (
	errNoUsername    = errors.New("username is required")
	errNoUserFound   = errors.New("username doesn't exist")
	errDuplicateUser = errors.New("duplicate username found")
)

func HandlerReset(s *commands.State, cmd commands.Command) error {
	if err := s.DB.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("ERROR: Couldn't reset, %v\n", err)
	}

	fmt.Println("Database reset successfully!")

	return nil
}
