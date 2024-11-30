package commands

import (
	"errors"

	"github.com/Dhar01/Gator/internal/config"
	"github.com/Dhar01/Gator/internal/database"
)

var errNoCommandFound = errors.New("command not found")

type State struct {
	DB     *database.Queries
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
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
