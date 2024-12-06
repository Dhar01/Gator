package handlers

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

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

	// feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	log.Printf("ERROR: can't get feed, %v\n", err)
	// 	return err
	// }
	// fmt.Printf("%+v\n", *feed)

	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("USAGE: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenReq, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenReq)

	ticker := time.NewTicker(timeBetweenReq)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var feed RSSFeed

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return nil, err
	}

	req.Header.Add("User-Agent", "gator")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR: response body error, %v\n", err)
		return nil, err
	}

	if err := xml.Unmarshal(body, &feed); err != nil {
		log.Printf("ERROR: unmarshal failed, %v\n", err)
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

func scrapeFeeds(s *commands.State) error {
	nextFeed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get the next feed, %w", err)
	}

	if _, err = s.DB.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return fmt.Errorf("can't mark as fetched")
	}

	feedData, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch the feed")
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("- %s\n", item.Title)
	}

	return nil
}
