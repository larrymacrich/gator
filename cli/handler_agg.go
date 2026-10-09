package cli

import (
	"context"
	"fmt"

	"github.com/larrymacrich/gator/internal/api"
)

const (
	url = "https://www.wagslane.dev/index.xml"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		errMsg := fmt.Errorf("no agrument(s) required")
		return errMsg
	}

	// fetch rss feed
	client := api.NewClient()
	feed, err := client.FetchFeed(context.Background(), url)
	if err != nil {
		errMsg := fmt.Errorf("fetching rss feed at '%s' failed: %w", url, err)
		return errMsg
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil

}
