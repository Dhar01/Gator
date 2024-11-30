package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/handlers"
	"github.com/Dhar01/Gator/internal/config"
)

var errLessArg = errors.New("not enough arguments provided")

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	state := commands.State{
		Config: &cfg,
	}

	cmd := commands.Commands{
		Handlers: make(map[string]func(*commands.State, commands.Command) error),
	}

	cmd.Register("login", handlers.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println(errLessArg)
		os.Exit(1)
	}

	command := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmd.Run(&state, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
