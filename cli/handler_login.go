package cli

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		errMsg := fmt.Errorf("missing username for command: %s", cmd.name)
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
