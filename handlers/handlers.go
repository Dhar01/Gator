package handlers

import (
	"errors"
	"fmt"

	"github.com/Dhar01/Gator/internal/config"
)

var (
	errNoCommandFound = errors.New("command not found")
	errNoUsername     = errors.New("username is required!")
)

type State struct {
	Config *config.Config
}

type Command struct {
	Name     string
	Argument []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {
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

func (c *Commands) Register(name string, handler func(*State, Command) error) {
	c.Handlers[name] = handler
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return errNoCommandFound
	}

	return handler(s, cmd)
}
