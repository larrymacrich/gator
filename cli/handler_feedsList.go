package cli

import (
	"context"
	"fmt"
)

// handlerListFeeds has no input arguments
// and prints all feed records and its creator to the console
func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		errMsg := fmt.Errorf("usage: %s no arguments required", cmd.name)
		return errMsg
	}

	// get all feeds
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		errMsg := fmt.Errorf("get feeds failed: %w", err)
		return errMsg
	}

	if len(feeds) == 0 {
		fmt.Println("no feed(s) found.")
		return nil
	}

	fmt.Printf("found %d feed(s):\n", len(feeds))
	fmt.Println("=====================================")
	for i, feed := range feeds {
		name, err := s.db.GetUserNameByUserID(context.Background(), feed.UserID)
		if err != nil {
			errMsg := fmt.Errorf("get user '%s' failed: %w", feed.UserID, err)
			return errMsg
		}
		fmt.Printf("feed #%v created by '%s':\n", i, name)
		printFeed(&feed)
	}

	return nil
}
