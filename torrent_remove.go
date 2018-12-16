package transmission

import (
	"context"

	"github.com/pkg/errors"
)

const TorrentRemove = "torrent-remove"

type TorrentRemoveRequest struct {
	IDs             []int `json:"ids,omitempty"`
	DeleteLocalData bool  `json:"delete-local-data"`
}

// TorrentRemove removes a torrent. If deleteData is true, will remove all
// data associated with that torrent.
func (t *Transmission) TorrentRemove(ctx context.Context, id int, deleteData bool) error {
	req := TorrentRemoveRequest{
		IDs:             []int{id},
		DeleteLocalData: deleteData,
	}

	err := t.Do(ctx, TorrentRemove, &req, nil)
	return errors.Wrap(err, "failed request")

}
