# Transmission

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
