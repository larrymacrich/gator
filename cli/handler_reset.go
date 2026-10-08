package cli

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		errMsg := fmt.Errorf("no arguments required")
		return errMsg
	}

	// resetting users table
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		errMsg := fmt.Errorf("resetting user table failed: %w", err)
		return errMsg
	}
	fmt.Printf("user table has been reset successfully.\n")

	return nil
}
