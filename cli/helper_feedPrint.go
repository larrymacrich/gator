package cli

import (
	"fmt"

	"github.com/larrymacrich/gator/internal/database"
)

func printFeed(feed *database.Feed) {
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("UserID: %s\n", feed.UserID)
	fmt.Printf("CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("Url: %s\n", feed.Url)
	fmt.Println("=====================================")
}

func printFeedFollow(feedFollow *database.CreateFeedFollowRow) {
	fmt.Printf("ID: %s\n", feedFollow.ID)
	fmt.Printf("CreatedAt: %v\n", feedFollow.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", feedFollow.UpdatedAt)
	fmt.Printf("UserID: %s\n", feedFollow.UserID)
	fmt.Printf("FeedID: %s\n", feedFollow.FeedID)
	fmt.Println("=====================================")
}

func printFeedFollowsForUser(feedFollowsForUser []database.GetFeedFollowsForUserRow) {
	for i, userFollowsFeed := range feedFollowsForUser {
		fmt.Printf("#%v: '%s'\n", i, userFollowsFeed.FeedName)
	}
	fmt.Println("=====================================")
}
