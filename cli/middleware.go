package cli

import (
	"context"
	"fmt"

	"github.com/larrymacrich/gator/internal/database"
)

// middlewareLoggedIn takes a function handle <handler>
// to ensure that the current user is logged in
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			errMsg := fmt.Errorf("'%s' failed: user '%s' not logged in", cmd.name, currentUser.Name)
			return errMsg
		}
		return handler(s, cmd, currentUser)
	}
}
