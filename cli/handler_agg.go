package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/larrymacrich/gator/internal/api"
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

	// print the rss feed to the console
	fetchedFeed.PrintRSSFeed()

	return nil

}
