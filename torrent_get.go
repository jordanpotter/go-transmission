package transmission

import (
	"context"
	"strings"

	"github.com/pkg/errors"
)

const TorrentGet = "torrent-get"

type TorrentGetRequest struct {
	IDs    []interface{} `json:"ids,omitempty"`
	Fields []string      `json:"fields,omitempty"`
}

type TorrentGetResponse struct {
	Torrents []Torrent `json:"torrents"`
}

// Torrent retrieves torrent with id or hash.
func (t *Transmission) Torrent(ctx context.Context, id interface{}) (*Torrent, error) {
	req := TorrentGetRequest{
		IDs:    []interface{}{id},
		Fields: TorrentFields(),
	}

	var resp TorrentGetResponse
	if err := t.Do(ctx, TorrentGet, &req, &resp); err != nil {
		return nil, errors.Wrap(err, "failed request")
	}

	if len(resp.Torrents) == 0 {
		return nil, nil
	} else if len(resp.Torrents) >= 2 {
		return nil, errors.Errorf("unexpected number of torrents %d", len(resp.Torrents))
	}

	return &resp.Torrents[0], nil
}

// TorrentExists returns whether torrent with id or hash exists.
func (t *Transmission) TorrentExists(ctx context.Context, id interface{}) (bool, error) {
	req := TorrentGetRequest{
		IDs:    []interface{}{id},
		Fields: []string{"id"},
	}

	var resp TorrentGetResponse
	if err := t.Do(ctx, TorrentGet, &req, &resp); err != nil {
		return false, errors.Wrap(err, "failed request")
	}

	return len(resp.Torrents) != 0, nil
}

// TorrentExistsInPath returns whether a torrent exists that has a download
// directory in path.
func (t *Transmission) TorrentExistsInPath(ctx context.Context, path string) (bool, error) {
	req := TorrentGetRequest{
		Fields: []string{"downloadDir"},
	}

	var resp TorrentGetResponse
	if err := t.Do(ctx, TorrentGet, &req, &resp); err != nil {
		return false, errors.Wrap(err, "failed request")
	}

	for _, torrent := range resp.Torrents {
		if strings.Index(torrent.DownloadDir, path) == 0 {
			return true, nil
		}
	}

	return false, nil
}

// Torrents retrieves all torrents.
func (t *Transmission) Torrents(ctx context.Context) ([]Torrent, error) {
	req := TorrentGetRequest{
		Fields: TorrentFields(),
	}

	var resp TorrentGetResponse
	if err := t.Do(ctx, TorrentGet, &req, &resp); err != nil {
		return nil, errors.Wrap(err, "failed request")
	}

	return resp.Torrents, nil
}
