package handlers

import (
	"context"
	"errors"
	"fmt"

	cmd "github.com/Dhar01/Gator/commands"
	db "github.com/Dhar01/Gator/internal/database"
)

var (
	errNoUsername    = errors.New("username is required")
	errNoUserFound   = errors.New("username doesn't exist")
	errDuplicateUser = errors.New("duplicate username found")
)

func HandlerReset(s *cmd.State, cmd cmd.Command) error {
	if err := s.DB.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("Couldn't reset: %v", err)
	}

	fmt.Println("database reset successfully!")

	return nil
}

func printUser(user db.User) {
	fmt.Printf(" * ID:   %v\n", user.ID)
	fmt.Printf(" * Name: %v\n", user.Name)
}
