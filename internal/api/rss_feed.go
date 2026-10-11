package api

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (c *Client) FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// create request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		errMsg := fmt.Errorf("creating request failed: %w", err)
		return nil, errMsg
	}

	// set header
	req.Header.Set("User-Agent", "gator")

	// send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		errMsg := fmt.Errorf("sending request failed: %w", err)
		return nil, errMsg
	}
	defer res.Body.Close()

	// check return codes
	if res.StatusCode != http.StatusOK {
		errMsg := fmt.Errorf("response failed with status code: %s", res.Status)
		return nil, errMsg
	}

	// read response
	var data RSSFeed
	rawBytes, err := io.ReadAll(res.Body)
	if err != nil {
		errMsg := fmt.Errorf("reading response body failed: %w", err)
		return nil, errMsg
	}

	// unmarshal raw XML
	if err := xml.Unmarshal(rawBytes, &data); err != nil {
		errMsg := fmt.Errorf("unmarshaling xml failed: %w", err)
		return nil, errMsg
	}

	// sanitize escaped sequences
	cleanResponse(&data)

	return &data, nil

}

func cleanResponse(data *RSSFeed) {
	data.Channel.Title = html.UnescapeString(data.Channel.Title)
	data.Channel.Description = html.UnescapeString(data.Channel.Description)
	for i := range data.Channel.Item {
		data.Channel.Item[i].Title = html.UnescapeString(data.Channel.Item[i].Title)
		data.Channel.Item[i].Description = html.UnescapeString(data.Channel.Item[i].Description)
	}
}

func (rssFeed *RSSFeed) PrintRSSFeed() {
	fmt.Println("RSSFeed Channel: ")
	fmt.Printf("Title: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Link: %s\n", rssFeed.Channel.Link)
	fmt.Printf("Description: %s\n", rssFeed.Channel.Description)
	fmt.Println("=====================================")
	for i := range rssFeed.Channel.Item {
		fmt.Printf("	RSSItem: #%v\n", i)
		fmt.Printf("	Title: %s\n", rssFeed.Channel.Item[i].Title)
		fmt.Printf("	Link: %s\n", rssFeed.Channel.Item[i].Link)
		fmt.Printf("	Description: %s\n", rssFeed.Channel.Item[i].Description)
		fmt.Printf("	PubDate: %s\n", rssFeed.Channel.Item[i].PubDate)
		fmt.Println("	=====================================")
	}
}
