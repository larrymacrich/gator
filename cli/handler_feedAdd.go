package cli

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/larrymacrich/gator/internal/database"
)

// handlerAddFeed takes feed <name> and feed <url>,
// creates a feed and feed_follow record for the current user
// and prints feed and feed_follow record to the console
func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		errMsg := fmt.Errorf("usage: %s <name> <url>", cmd.name)
		return errMsg
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	// create new feed
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
	}
	newFeed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		// silence duplicate key violations in error log
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			errMsg := fmt.Errorf("you are already following feed %s.", feedName)
			return errMsg
		}
		errMsg := fmt.Errorf("creating feed %s failed: %w", feedName, err)
		return errMsg
	}
	fmt.Println("Feed created:")
	printFeed(&newFeed)

	// connect current user and new feed
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    newFeed.ID,
	}
	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		errMsg := fmt.Errorf("'%s' following '%s' failed: %w", user.Name, newFeed.Name, err)
		return errMsg
	}

	fmt.Printf("'%s' is now following '%s':\n", user.Name, newFeedFollow.FeedName)

	return nil
}
