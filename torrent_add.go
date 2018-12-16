package transmission

import (
	"context"
	"encoding/base64"

	"github.com/pkg/errors"
)

const TorrentAdd = "torrent-add"

type TorrentAddRequest struct {
	Cookies           string `json:"cookies,omitempty"`
	DownloadDir       string `json:"download-dir,omitempty"`
	Filename          string `json:"filename,omitempty"`
	MetaInfo          string `json:"metainfo,omitempty"`
	Paused            bool   `json:"paused"`
	PeerLimit         int    `json:"peer-limit,omitempty"`
	BandwidthPriority int    `json:"bandwidthPriority,omitempty"`
	FilesWanted       []int  `json:"files-wanted,omitempty"`
	FilesUnwanted     []int  `json:"files-unwanted,omitempty"`
	PriorityHigh      []int  `json:"priority-high,omitempty"`
	PriorityLow       []int  `json:"priority-low,omitempty"`
	PriorityNormal    []int  `json:"priority-normal,omitempty"`
}

type TorrentAddResponse struct {
	TorrentAdded     *Torrent `json:"torrent-added"`
	TorrentDuplicate *Torrent `json:"torrent-duplicate"`
}

// TorrentAddWithURL adds a torrent given the URL.
func (t *Transmission) TorrentAddWithURL(ctx context.Context, url, path string, paused bool) (*Torrent, bool, error) {
	req := TorrentAddRequest{
		Filename:    url,
		DownloadDir: path,
		Paused:      paused,
	}

	var resp TorrentAddResponse
	if err := t.Do(ctx, TorrentAdd, &req, &resp); err != nil {
		return nil, false, errors.Wrap(err, "failed request")
	}

	if resp.TorrentAdded != nil {
		return resp.TorrentAdded, false, nil
	} else if resp.TorrentDuplicate != nil {
		return resp.TorrentDuplicate, true, nil
	} else {
		return nil, false, errors.New("received no torrent data")
	}
}

// TorrentAddWithFile adds a torrent given the .torrent file data.
func (t *Transmission) TorrentAddWithFile(ctx context.Context, data []byte, path string, paused bool) (*Torrent, bool, error) {
	b := base64.StdEncoding.EncodeToString(data)

	req := TorrentAddRequest{
		MetaInfo:    b,
		DownloadDir: path,
		Paused:      paused,
	}

	var resp TorrentAddResponse
	if err := t.Do(ctx, TorrentAdd, &req, &resp); err != nil {
		return nil, false, errors.Wrap(err, "failed request")
	}

	if resp.TorrentAdded != nil {
		return resp.TorrentAdded, false, nil
	} else if resp.TorrentDuplicate != nil {
		return resp.TorrentDuplicate, true, nil
	} else {
		return nil, false, errors.New("received no torrent data")
	}
}
