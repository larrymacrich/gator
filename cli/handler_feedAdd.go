package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/larrymacrich/gator/internal/database"
)

// handlerAddFeed takes feed <name> and feed <url>,
// creates a feed and feed_follow record for the current user
// and prints feed and feed_follow record to the console
func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		errMsg := fmt.Errorf("usage: %s <name> <url>", cmd.name)
		return errMsg
	}

	userName := s.cfg.CurrentUserName
	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	// get current user
	currentUser, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		errMsg := fmt.Errorf("creating feed failed: User '%s' not registered.", userName)
		return errMsg
	}

	// create new feed
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    currentUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
	}
	newFeed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		errMsg := fmt.Errorf("creating feed %s failed: %w", feedName, err)
		return errMsg
	}
	fmt.Println("Feed created:")
	printFeed(&newFeed)

	// connect current user and new feed
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    newFeed.ID,
	}
	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		errMsg := fmt.Errorf("'%s' following '%s' failed: %w", currentUser.Name, newFeed.Name, err)
		return errMsg
	}

	fmt.Printf("'%s' is now following '%s':\n", currentUser.Name, newFeed.Name)
	printFeedFollow(&newFeedFollow)

	return nil
}
