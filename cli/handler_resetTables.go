package cli

import (
	"context"
	"fmt"
)

// handlerReset resets users table
// and cascades delete through related tables
func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		errMsg := fmt.Errorf("usage: %s no arguments required", cmd.name)
		return errMsg
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		errMsg := fmt.Errorf("resetting tables failed: %w", err)
		return errMsg
	}
	fmt.Printf("tables have been reset successfully.\n")

	return nil
}
