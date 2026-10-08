package main

import (
	"fmt"

	"github.com/larrymacrich/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		errMsg := fmt.Errorf("missing argument(s) for command: %s\n", cmd.name)
		return errMsg
	}

	userName := cmd.args[0]
	err := s.cfg.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Printf("User %s has been set successfully.\n", userName)

	return nil
}

type commands struct {
	mapCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdHandler, exists := c.mapCommands[cmd.name]
	if !exists {
		errMsg := fmt.Errorf("invalid command: %s\n", cmd.name)
		return errMsg
	}

	err := cmdHandler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.mapCommands[name] = f
}
