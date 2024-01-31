package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/KalimaraPeleteiro/RSS-Aggregator/internal/database"
)

func startScraping(database *database.Queries, nJobs int, timeBetweenRequest time.Duration) {
	log.Printf("Iniciando o Scraping com %v rotinas, a cada %v", nJobs, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := database.GetNextFeedsToFetch(context.Background(), int32(nJobs))
		if err != nil {
			log.Printf("Erro ao buscar feeds: %v", err)
			continue
		}

		group := &sync.WaitGroup{}
		for _, feed := range feeds {
			group.Add(1)

			go scrapeFeed(database, group, feed)
		}
		group.Wait()
	}
}

func scrapeFeed(database *database.Queries, group *sync.WaitGroup, feed database.Feed) {
	defer group.Done()

	_, err := database.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Erro ao marcar feed como visitado: %v", err)
		return
	}

	rssFeed, err := URLtoFeed(feed.Url)
	if err != nil {
		log.Printf("Erro ao buscar feed: %v", err)
	}

	for _, item := range rssFeed.Channel.Item {
		log.Printf("Post encontrado: %v em %v\n", item.Title, feed.Name)
	}
	log.Printf("Feed %v coletado, %v posts encontrados.", feed.Name, len(rssFeed.Channel.Item))
}
