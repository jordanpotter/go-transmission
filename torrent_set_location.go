package transmission

import (
	"context"

	"github.com/pkg/errors"
)

const TorrentSetLocation = "torrent-set-location"

type TorrentSetLocationRequest struct {
	IDs      []int  `json:"ids,omitempty"`
	Location string `json:"location,omitempty"`
	Move     bool   `json:"move"`
}

func (t *Transmission) TorrentSetLocation(ctx context.Context, id int, path string, moveData bool) error {
	req := TorrentSetLocationRequest{
		IDs:      []int{id},
		Location: path,
		Move:     moveData,
	}

	err := t.Do(ctx, TorrentSetLocation, &req, nil)
	return errors.Wrap(err, "failed request")
}
