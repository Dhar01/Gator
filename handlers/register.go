package handlers

import (
	"context"
	"fmt"
	"time"

	cmd "github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *cmd.State, cmd cmd.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("USAGE: %s <name>\n", cmd.Name)
	}

	username := cmd.Args[0]

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return errDuplicateUser
	}

	if err := s.Config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("User created successfully!")
	printUser(database.User(user))

	return nil
}


func printUser(user database.User) {
	fmt.Printf(" * ID:   %v\n", user.ID)
	fmt.Printf(" * Name: %v\n", user.Name)
}
