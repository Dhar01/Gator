package handlers

import (
	"errors"
	"fmt"

	"github.com/Dhar01/Gator/internal/config"
)

var (
	errNoCommandFound = errors.New("command not found")
	errNoArgument     = errors.New("no argument provided")
)

type State struct {
	config *config.Config
}

type command struct {
	Name     string
	Argument []string
}

type commands struct {
	handlers map[string]func(*State, command) error
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.Argument) > 1 {
		return errNoArgument
	}

	username := cmd.Argument[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("the user has been set")
	return nil
}

func (c *commands) register(name string, handler func(*State, command) error) {
	c.handlers[name] = handler
}

func (c *commands) run(s *State, cmd command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return errNoCommandFound
	}

	return handler(s, cmd)
}
