package cli

import (
	"context"
	"fmt"
)

func handlerList(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		errMsg := fmt.Errorf("no arguments required")
		return errMsg
	}

	// get users
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		errMsg := fmt.Errorf("get users: %w", err)
		return errMsg
	}

	for _, user := range users {
		if user.Name.String == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name.String)
		} else {
			fmt.Printf("* %s\n", user.Name.String)
		}
	}

	return nil
}
