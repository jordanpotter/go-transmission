package main

import (
	"context"
	"flag"
	"fmt"

	transmission "github.com/jordanpotter/go-transmission"
)

var (
	endpoint string
	username string
	password string
)

func init() {
	flag.StringVar(&endpoint, "endpoint", "", "transmission endpoint")
	flag.StringVar(&username, "username", "", "transmission username")
	flag.StringVar(&password, "password", "", "transmission password")
	flag.Parse()
}

func main() {
	t, err := transmission.New(endpoint, username, password)
	if err != nil {
		panic(err)
	}

	torrents, err := t.Torrents(context.Background())
	if err != nil {
		panic(err)
	}

	for _, torrent := range torrents {
		fmt.Println(torrent.Name)
	}
}
