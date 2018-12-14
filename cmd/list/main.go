package main

import (
	"context"
	"flag"
	"fmt"

	transmission "github.com/jordanpotter/go-transmission"
)

var (
	url      string
	username string
	password string
)

func init() {
	flag.StringVar(&url, "url", "", "transmission url")
	flag.StringVar(&username, "username", "", "transmission username")
	flag.StringVar(&password, "password", "", "transmission password")
	flag.Parse()
}

func main() {
	t := transmission.New(url, username, password)

	torrents, err := t.Torrents(context.Background())
	if err != nil {
		panic(err)
	}

	for _, torrent := range torrents {
		fmt.Println(torrent.Name)
	}
}
