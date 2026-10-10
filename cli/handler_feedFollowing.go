package cli

import (
	"context"
	"fmt"
)

// handlerFeedFollowing has no input arguments
// and prints all feeds that the current user is following
func handlerFeedFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		errMsg := fmt.Errorf("usage: %s no arguments required", cmd.name)
		return errMsg
	}

	userName := s.cfg.CurrentUserName

	// get current user
	currentUser, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		errMsg := fmt.Errorf("get current user '%s' failed: %w", userName, err)
		return errMsg
	}

	// get feeds connected to current user
	feedFollowsForUser, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		errMsg := fmt.Errorf("get feeds failed: %w", err)
		return errMsg
	}

	fmt.Printf("current user '%s' is following %v feed(s)\n", userName, len(feedFollowsForUser))
	printFeedFollowsForUser(feedFollowsForUser)

	return nil
}
