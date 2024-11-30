package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	cmd "github.com/Dhar01/Gator/commands"
	db "github.com/Dhar01/Gator/internal/database"
	"github.com/google/uuid"
)

var (
	errNoUsername    = errors.New("username is required")
	errNoUserFound   = errors.New("username doesn't exist")
	errDuplicateUser = errors.New("duplicate username found")
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

func HandlerRegister(s *cmd.State, cmd cmd.Command) error {
	if len(cmd.Args) < 1 {
		fmt.Printf("usage: %s <name>\n", cmd.Name)
		return errNoUsername
	}

	username := cmd.Args[0]

	user := db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	data, err := s.DB.CreateUser(context.Background(), user)
	if err != nil {
		return errDuplicateUser
	}

	if err := s.Config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("the user %s was created\n", username)
	printUser(db.User(user))
	log.Println(data)


	return nil
}

func printUser(user db.User) {
	fmt.Printf(" * ID:   %v\n", user.ID)
	fmt.Printf(" * Name: %v\n", user.Name)
}
