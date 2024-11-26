package commands

import (
	"errors"

	"github.com/Dhar01/Gator/internal/config"
)

var errNoCommandFound = errors.New("command not found")

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
