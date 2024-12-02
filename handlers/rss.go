package handlers

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"

	"github.com/Dhar01/Gator/commands"
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

func HandlerAggregate(s *commands.State, cmd commands.Command) error {
	result, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Printf("fetch ERROR: %v", err)
		return err
	}

	// title := html.UnescapeString(result.Channel.Title)
	// description := html.UnescapeString(result.Channel.Description)

	// fmt.Println(title)
	// fmt.Println(description)

	for _, item := range result.Channel.Item {
		title := html.UnescapeString(item.Title)
		description := html.UnescapeString(item.Description)
		if title == "The Zen of Proverbs" || title == "Optimize for simplicity" {
			fmt.Println(title)
			fmt.Println("Result: ", description)
		}
	}

	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var feed RSSFeed

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return &feed, err
	}

	req.Header.Add("User-Agent", "gator")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
		return &feed, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Resp Body Error: %v", err)
		return &feed, err
	}

	if err := xml.Unmarshal(body, &feed); err != nil {
		log.Printf("Unmarshal Error: %v", err)
		return &feed, err
	}

	return &feed, nil
}
