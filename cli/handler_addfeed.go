package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/larrymacrich/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		errMsg := fmt.Errorf("missing arguemnt(s) for command: %s", cmd.name)
		return errMsg
	}

	// get current user
	userName := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		errMsg := fmt.Errorf("creating feed failed: User '%s' not registered.", userName)
		return errMsg
	}

	// create new feed data
	feedName := cmd.args[0]
	feedUrl := cmd.args[1]
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    currentUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
	}

	// add feed to db
	newFeed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		errMsg := fmt.Errorf("creating feed %s failed: %w", feedName, err)
		return errMsg
	}

	fmt.Println("Feed created:")
	fmt.Printf("ID: %s\n", newFeed.ID)
	fmt.Printf("UserID: %s\n", newFeed.UserID)
	fmt.Printf("CreatedAt: %v\n", newFeed.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", newFeed.UpdatedAt)
	fmt.Printf("Name: %s\n", newFeed.Name)
	fmt.Printf("Url: %s\n", newFeed.Url)

	return nil
}
