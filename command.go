package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/larrymacrich/gator/internal/config"
	"github.com/larrymacrich/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type commands struct {
	mapCommands map[string]func(*state, command) error
}

type command struct {
	name string
	args []string
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		errMsg := fmt.Errorf("missing username for command: %s\n", cmd.name)
		return errMsg
	}

	// get user data
	userName := cmd.args[0]
	sqlUserName := sql.NullString{
		String: userName,
		Valid:  true,
	}
	_, err := s.db.GetUser(context.Background(), sqlUserName)
	if err != nil {
		errMsg := fmt.Errorf("login user '%s' failed: User not registered. \n", userName)
		return errMsg
	}

	// login user
	err = s.cfg.SetUser(userName)
	if err != nil {
		errMsg := fmt.Errorf("login user '%s' failed: %w\n", userName, err)
		return errMsg
	}
	fmt.Printf("login user '%s' successful.\n", userName)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		errMsg := fmt.Errorf("missing username for command: %s\n", cmd.name)
		return errMsg
	}

	// create new user data
	userName := cmd.args[0]
	sqlUserName := sql.NullString{
		String: userName,
		Valid:  true,
	}
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sqlUserName,
	}
	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		errMsg := fmt.Errorf("creating user '%s' failed: %w\n", userName, err)
		return errMsg
	}
	fmt.Printf("user '%s' has been created successfully:\n", userName)

	// login user
	err = s.cfg.SetUser(userName)
	if err != nil {
		errMsg := fmt.Errorf("login user '%s' failed: %w\n", userName, err)
		return errMsg
	}
	fmt.Printf("login user '%s' successful.\n", userName)

	return nil
}
