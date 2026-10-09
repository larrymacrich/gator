package cli

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		errMsg := fmt.Errorf("usage: %s no arguments required", cmd.name)
		return errMsg
	}

	// get feeds
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		errMsg := fmt.Errorf("get feeds failed: %w", err)
		return errMsg
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for i, feed := range feeds {
		name, err := s.db.GetUserNameByUserID(context.Background(), feed.UserID)
		if err != nil {
			errMsg := fmt.Errorf("get user '%s' failed: %w", feed.UserID, err)
			return errMsg
		}
		fmt.Printf("Feed #%v:\n", i)
		fmt.Printf("ID: %s\n", feed.ID)
		fmt.Printf("UserID: %s\n", feed.UserID)
		fmt.Printf("UserName: %s\n", name)
		fmt.Printf("CreatedAt: %v\n", feed.CreatedAt)
		fmt.Printf("UpdatedAt: %v\n", feed.UpdatedAt)
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("Url: %s\n", feed.Url)
		fmt.Println("=====================================")
	}

	return nil
}
