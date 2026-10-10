package cli

import (
	"context"
	"fmt"

	"github.com/larrymacrich/gator/internal/database"
)

// handlerFeedUnfollow takes a <url> feed
// and removes the follow record for the current user
func handlerFeedUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		errMsg := fmt.Errorf("usage: %s <url>", cmd.name)
		return errMsg
	}

	feedUrl := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		errMsg := fmt.Errorf("get feed failed: %w", err)
		return errMsg
	}

	// delete feed connected to current user
	deleteFeedFollowsForUserParams := database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	rowsAffected, err := s.db.DeleteFeedFollowsForUser(context.Background(), deleteFeedFollowsForUserParams)
	if err != nil {
		errMsg := fmt.Errorf("unfollowing feed failed: %w", err)
		return errMsg
	}
	if rowsAffected == 0 {
		return fmt.Errorf("you are not following '%s'", feed.Name)
	}

	fmt.Printf("'%s' unfollowed '%s'\n", user.Name, feed.Name)

	return nil
}
