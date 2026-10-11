package cli

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/larrymacrich/gator/internal/api"
	"github.com/larrymacrich/gator/internal/database"
)

// handlerAgg takes a time frame <timeBetweenRequests>
// and scrapes rss feed after that time frame passes
func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		errMsg := fmt.Errorf("usage: %s <timeBetweenRequests>", cmd.name)
		return errMsg
	}

	timeFrame := cmd.args[0]

	time_between_reqs, err := time.ParseDuration(timeFrame)
	if err != nil {
		errMsg := fmt.Errorf("inavlid time format: %w", err)
		return errMsg
	}
	if time_between_reqs <= 0 {
		return fmt.Errorf("<timeBetweenRequests> must be greater than 0")
	}

	// scrape feed
	ticker := time.NewTicker(time_between_reqs)
	fmt.Printf("Collecting feeds every %s\n", timeFrame)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Printf("scraping feed failed: %v\n", err)
		}
	}

}

// scrapeFeeds has no input arguments
// mark, fetch, and prints feed to the console
func scrapeFeeds(s *state) error {
	client := api.NewClient()

	// get next feed to fetch
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		errMsg := fmt.Errorf("get feed to fetch next failed: %w", err)
		return errMsg
	}

	// mark that feed to be fetched next
	markedFeed, err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		errMsg := fmt.Errorf("mark feed as fetched failed: %w", err)
		return errMsg
	}

	// fetch the feed
	fetchedFeed, err := client.FetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		errMsg := fmt.Errorf("fetching marked feed '%s' at  failed: %w", markedFeed.Url, err)
		return errMsg
	}

	// save the posts
	fmt.Printf("fetching post from feed: '%s'\n", markedFeed.Url)
	fmt.Println("============FEED START===========")
	for _, item := range fetchedFeed.Channel.Item {
		creatPostParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: convertDateStringToTime(item.PubDate),
			FeedID:      markedFeed.ID,
		}
		post, err := s.db.CreatePost(context.Background(), creatPostParams)
		if err != nil {
			// silence duplicate key violations
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("saving post '%s' form feed '%s' failed: %v\n", item.Link, markedFeed.Name, err)
			continue
		}
		fmt.Printf("saving post '%s'\n", post.Url)
	}
	fmt.Println("============FEED END=============")

	return nil
}

// convertDateStringToTime takes a time format <dateString>
// and converts it into sql.NullTime using
// most common time formats used in RSS 2.0 feeds:
// ttime.RFC1123, time.RFC1123Z
func convertDateStringToTime(dateString string) sql.NullTime {
	formats := []string{
		time.RFC3339,
		time.RFC1123,
		time.RFC1123Z,
	}

	for _, format := range formats {
		t, err := time.Parse(format, dateString)
		if err == nil {
			return sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
	}
	fmt.Printf("unable to parse time: %s", dateString)

	return sql.NullTime{}
}
