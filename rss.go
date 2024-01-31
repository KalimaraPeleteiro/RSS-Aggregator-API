package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func URLtoFeed(url string) (RSSFeed, error) {
	client := http.Client{
		Timeout: 10 * time.Second, // Tamanho máximo de espera de resposta
	}

	response, err := client.Get(url)
	if err != nil {
		log.Printf("Erro ao realizar requisição GET ao feed %v: %v", url, err)
		return RSSFeed{}, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Erro ao ler o feed %v: %v", url, err)
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		log.Printf("Erro ao traduzir o feed %v para XML: %v", url, err)
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
