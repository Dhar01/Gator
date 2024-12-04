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
		log.Printf("error reading config: %v", err)
	}

	dbURL := "postgres://postgres:postgres@localhost:5432/gator"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("ERROR: %v", err)
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
	cmd.Register("reset", handlers.HandlerReset)
	cmd.Register("users", handlers.HandlerUsers)
	cmd.Register("agg", handlers.HandlerAggregate)
	cmd.Register("addfeed", handlers.HandlerAddFeed)
	cmd.Register("feeds", handlers.HandlerFeeds)
	cmd.Register("fetch", handlers.HandlerFetch)
	cmd.Register("follow", handlers.HandlerFollow)

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
		log.Printf("ERROR: %v", err)
		os.Exit(1)
	}

}
