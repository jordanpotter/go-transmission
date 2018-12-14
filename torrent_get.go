package transmission

import (
	"context"

	"github.com/pkg/errors"
)

const TorrentGet = "torrent-get"

type TorrentGetRequest struct {
	IDs    []string `json:"ids,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

type TorrentGetResponse struct {
	Torrents []Torrent `json:"torrents"`
}

func (t *Transmission) Torrent(ctx context.Context, id string) (*Torrent, error) {
	req := TorrentGetRequest{
		IDs:    []string{id},
		Fields: torrentFields(),
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

func (t *Transmission) Torrents(ctx context.Context) ([]Torrent, error) {
	req := TorrentGetRequest{
		Fields: torrentFields(),
	}

	var resp TorrentGetResponse
	if err := t.Do(ctx, TorrentGet, &req, &resp); err != nil {
		return nil, errors.Wrap(err, "failed request")
	}

	return resp.Torrents, nil
}
