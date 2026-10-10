package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/larrymacrich/gator/internal/database"
)

// handlerFeedFollow takes a <url> feed
// and creates a new follow record for the current user
func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		errMsg := fmt.Errorf("usage: %s <url>", cmd.name)
		return errMsg
	}

	feedUrl := cmd.args[0]

	// get feed
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		errMsg := fmt.Errorf("get feed '%s' failed: %w", feedUrl, err)
		return errMsg
	}

	// connect current user and feedFollow
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		errMsg := fmt.Errorf("create follow failed: %w", err)
		return errMsg
	}
	fmt.Printf("'%s' is now following '%s' feed:\n", user.Name, feed.Name)
	printFeedFollow(&newFeedFollow)

	return nil
}
