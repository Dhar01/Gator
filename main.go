package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/handlers"
	"github.com/Dhar01/Gator/internal/config"
	"github.com/Dhar01/Gator/internal/database"
	_ "github.com/lib/pq"
)

var errLessArg = errors.New("not enough arguments provided")

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	dbURL := "postgres://postgres:postgres@localhost:5432/gator"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	dbQueries := database.New(db)

	state := &commands.State{
		DB:     dbQueries,
		Config: &cfg,
	}

	cmd := commands.Commands{
		Handlers: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmd.Register("login", handlers.HandlerLogin)
	cmd.Register("register", handlers.HandlerRegister)

	if len(os.Args) < 2 {
		fmt.Println(errLessArg)
		os.Exit(1)
	}

	command := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmd.Run(state, command)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}

}
