package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func aggHandler(_ *state, _ command) error {
	// if len(cmd.Args) < 1 {
	// 	return fmt.Errorf("command agg requires 1 arg as FeedUrl")
	// }
	// feedURL := cmd.Args[0]

	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	fmt.Printf("Feed Description: %s\n", feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		fmt.Printf("Item Title: %s\n", item.Title)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Println()
	}
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var feed RSSFeed
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}

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
