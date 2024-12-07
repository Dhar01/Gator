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
	"github.com/Dhar01/Gator/internal/database"
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

func scrapeFeeds(s *commands.State) {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get next feeds to fetch")
		return
	}

	log.Println("found a feed to fetch!")
	scrapeFeed(s.DB, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	if _, err := db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		log.Printf("couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	// for _, item := range feedData.Channel.Item {
	// 	publishedAt := sql.NullTime{}
	// 	if t, err := time.Parse()
	// 	_, err = db.CreatePost(context.Background(), database.CreatePostParams{
	// 		Url:         item.Link,
	// 		Description: item.Description,
	// 		Title:       item.Title,
	// 		PublishedAt: item.PubDate,
	// 	})
	// }

}
