package cli

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		errMsg := fmt.Errorf("usage: %s <name>", cmd.name)
		return errMsg
	}

	// get user data
	userName := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		errMsg := fmt.Errorf("login user '%s' failed: User not registered.", userName)
		return errMsg
	}

	// login user
	err = s.cfg.SetUser(userName)
	if err != nil {
		errMsg := fmt.Errorf("login user '%s' failed: %w", userName, err)
		return errMsg
	}
	fmt.Printf("login user '%s' successful.\n", userName)

	return nil
}
