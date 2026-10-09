package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/larrymacrich/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		errMsg := fmt.Errorf("usage: %s <name>", cmd.name)
		return errMsg
	}

	// create new user data
	userName := cmd.args[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}
	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		errMsg := fmt.Errorf("creating user '%s' failed: %w", userName, err)
		return errMsg
	}
	fmt.Printf("user '%s' has been created successfully:\n", userName)

	// login user
	err = s.cfg.SetUser(userName)
	if err != nil {
		errMsg := fmt.Errorf("login user '%s' failed: %w", userName, err)
		return errMsg
	}
	fmt.Printf("login user '%s' successful.\n", userName)

	return nil
}
