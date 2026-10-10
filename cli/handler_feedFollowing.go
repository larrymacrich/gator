package cli

import (
	"context"
	"fmt"

	"github.com/larrymacrich/gator/internal/database"
)

// handlerFeedFollowing has no input arguments
// and prints all feeds that the current user is following
func handlerFeedFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		errMsg := fmt.Errorf("usage: %s no arguments required", cmd.name)
		return errMsg
	}

	// get feeds connected to current user
	feedFollowsForUser, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		errMsg := fmt.Errorf("get feeds failed: %w", err)
		return errMsg
	}

	fmt.Printf("'%s' is following %v feed(s)\n", user.Name, len(feedFollowsForUser))
	printFeedFollowsForUser(feedFollowsForUser)

	return nil
}
