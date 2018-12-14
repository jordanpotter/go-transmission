package transmission

const TorrentSet = "torrent-set"

type TorrentSetRequest struct {
	IDs                 []int    `json:"ids,omitempty"`
	BandwidthPriority   int      `json:"bandwidthPriority,omitempty"`
	DownloadLimit       int      `json:"downloadLimit,omitempty"`
	DownloadLimited     bool     `json:"downloadLimited"`
	FilesWanted         []int    `json:"files-wanted,omitempty"`
	FilesUnwanted       []int    `json:"files-unwanted,omitempty"`
	HonorsSessionLimits bool     `json:"honorsSessionLimits"`
	Location            string   `json:"location,omitempty"`
	PeerLimit           int      `json:"peer-limit,omitempty"`
	PriorityHigh        []int    `json:"priority-high,omitempty"`
	PriorityLow         []int    `json:"priority-low,omitempty"`
	PriorityNormal      []int    `json:"priority-normal,omitempty"`
	QueuePosition       int      `json:"queuePosition,omitempty"`
	SeedIdleLimit       int      `json:"seedIdleLimit,omitempty"`
	SeedIdleMode        int      `json:"seedIdleMode,omitempty"`
	SeedRatioLimit      float64  `json:"seedRatioLimit,omitempty"`
	SeedRatioMode       int      `json:"seedRatioMode,omitempty"`
	TrackersAdd         []string `json:"trackerAdd,omitempty"`
	TrackersRemove      []int    `json:"trackerRemove,omitempty"`
	TrackersReplace     []string `json:"trackerReplace,omitempty"`
	UploadLimit         int      `json:"uploadLimit,omitempty"`
	UploadLimited       bool     `json:"uploadLimited"`
}
