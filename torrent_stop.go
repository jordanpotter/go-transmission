package transmission

import (
	"context"

	"github.com/pkg/errors"
)

const TorrentStop = "torrent-stop"

type TorrentStopRequest struct {
	IDs []int `json:"ids,omitempty"`
}

func (t *Transmission) TorrentStop(ctx context.Context, id int) error {
	req := TorrentStopRequest{
		IDs: []int{id},
	}

	err := t.Do(ctx, TorrentStop, &req, nil)
	return errors.Wrap(err, "failed request")
}
