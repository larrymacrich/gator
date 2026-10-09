package cli

import (
	"context"
	"fmt"

	"github.com/larrymacrich/gator/internal/api"
)

const (
	feedUrl = "https://www.wagslane.dev/index.xml"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		errMsg := fmt.Errorf("no agrument(s) required")
		return errMsg
	}

	// fetch rss feed
	client := api.NewClient()
	feed, err := client.FetchFeed(context.Background(), feedUrl)
	if err != nil {
		errMsg := fmt.Errorf("fetching rss feed at '%s' failed: %w", feedUrl, err)
		return errMsg
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil

}
