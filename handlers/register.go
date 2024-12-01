package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	cmd "github.com/Dhar01/Gator/commands"
	db "github.com/Dhar01/Gator/internal/database"
	"github.com/google/uuid"
)

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
