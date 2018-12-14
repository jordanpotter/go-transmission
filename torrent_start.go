package transmission

import (
	"context"

	"github.com/pkg/errors"
)

const TorrentStart = "torrent-start"

type TorrentStartRequest struct {
	IDs []int `json:"ids,omitempty"`
}

func (t *Transmission) TorrentStart(ctx context.Context, id int) error {
	req := TorrentStartRequest{
		IDs: []int{id},
	}

	err := t.Do(ctx, TorrentStart, &req, nil)
	return errors.Wrap(err, "failed request")
}
