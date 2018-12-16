# go-transmission [![](https://godoc.org/github.com/jordanpotter/go-transmission?status.svg)](https://godoc.org/github.com/jordanpotter/go-transmission) [![Go Report Card](https://goreportcard.com/badge/github.com/jordanpotter/go-transmission)](https://goreportcard.com/report/github.com/jordanpotter/go-transmission) [![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

Go client library to interface with Transmission

```golang
t := transmission.New(url, username, password)

// Add a torrent
torrent, err := t.TorrentAddWithURL(ctx,
        "http://releases.ubuntu.com/18.04/ubuntu-18.04.1-desktop-amd64.iso.torrent",
        "/download/path", false)

// Get a torrent
torrent, err := t.Torrent(ctx, id)

// Get all torrents
torrents, err := t.Torrents(ctx)

// Stop a torrent
err := t.TorrentStop(ctx, id)

// Remove a torrent
err := t.TorrentRemove(ctx, id, deleteData)

// Move a torrent's storage location
err := t.TorrentSetLocation(ctx, id, "/new/download/path", true)
```
