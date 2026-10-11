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

func printFeedsRow(feed *database.GetFeedsRow) {
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("UserID: %s\n", feed.UserID)
	fmt.Printf("CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("Url: %s\n", feed.Url)
	fmt.Printf("LastFetchedAt: %v\n", feed.LastFetchedAt)
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

func printPost(post *database.GetPostsForUserByUserIDRow) {
	fmt.Printf("ID: %s\n", post.ID)
	fmt.Printf("CreatedAt: %v\n", post.CreatedAt)
	fmt.Printf("UpdatedAt: %v\n", post.UpdatedAt)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Url: %s\n", post.Url)
	fmt.Printf("Description: %s\n", post.Description.String)
	fmt.Printf("PublishedAt: %v\n", post.PublishedAt.Time)
	fmt.Printf("FeedID: %s\n", post.FeedID)
	fmt.Println("=====================================")
}
