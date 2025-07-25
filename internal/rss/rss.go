package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/database"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/state"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Desctiption string    `xml:"description"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Desctiption string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedUrl string) (RssFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	req.Header.Set("User-Agent", "gator")
	if err != nil {
		return RssFeed{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return RssFeed{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RssFeed{}, err
	}
	var rssFeed RssFeed
	if err = xml.Unmarshal(data, &rssFeed); err != nil {
		return RssFeed{}, err
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Desctiption = html.UnescapeString(rssFeed.Channel.Desctiption)
	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Desctiption = html.UnescapeString(rssFeed.Channel.Item[i].Desctiption)
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
	}
	return rssFeed, nil
}

func ScrapeFeeds(ctx context.Context, s *state.State) error {
	feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}
	_, err = s.Db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}
	rssFeed, err := FetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}
	for _, item := range rssFeed.Channel.Item {
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Desctiption,
			PublishedAt: time.Now(),
			FeedID:      feed.ID,
		}
		s.Db.CreatePost(ctx, params)
	}
	return nil
}
