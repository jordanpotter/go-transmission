package transmission

import (
	"context"

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

// Torrent retrieves torrent by id or hash.
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
