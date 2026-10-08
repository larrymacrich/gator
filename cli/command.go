package cli

import (
	"fmt"
)

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

type command struct {
	name string
	args []string
}

func (c *commands) run(s *state, cmd command) error {
	cmdHandler, exists := c.registeredCommands[cmd.name]
	if !exists {
		errMsg := fmt.Errorf("invalid command: %s", cmd.name)
		return errMsg
	}

	err := cmdHandler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
