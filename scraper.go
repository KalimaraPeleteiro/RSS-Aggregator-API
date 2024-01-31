package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, nJobs int, timeBetweenRequest time.Duration) {
	log.Printf("Iniciando o Scraping com %v rotinas, a cada %v", nJobs, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(nJobs))
		if err != nil {
			log.Printf("Erro ao buscar feeds: %v", err)
			continue
		}

		group := &sync.WaitGroup{}
		for _, feed := range feeds {
			group.Add(1)

			go scrapeFeed(db, group, feed)
		}
		group.Wait()

	}
}

func scrapeFeed(db *database.Queries, group *sync.WaitGroup, feed database.Feed) {
	defer group.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Erro ao marcar feed como visitado: %v", err)
		return
	}

	rssFeed, err := URLtoFeed(feed.Url)
	if err != nil {
		log.Printf("Erro ao buscar feed: %v", err)
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Erro ao transformar data %v: %v", item.PubDate, err)
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Erro ao criar post: %v", err)
		}
	}
	log.Printf("Feed %v coletado, %v posts encontrados.", feed.Name, len(rssFeed.Channel.Item))
}
