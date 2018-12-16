package transmission

import (
	"reflect"
	"strings"
)

// Torrent contains all of the fields returned by Transmission for a torrent
type Torrent struct {
	ID                      int                  `json:"id"`
	ActivityDate            int                  `json:"activityDate"`
	AddedDate               int                  `json:"addedDate"`
	BandwidthPriority       int                  `json:"bandwidthPriority"`
	Comment                 string               `json:"comment"`
	CorruptEver             int                  `json:"corruptEver"`
	Creator                 string               `json:"creator"`
	DateCreated             int                  `json:"dateCreated"`
	DesiredAvailable        int                  `json:"desiredAvailable"`
	DoneDate                int                  `json:"doneDate"`
	DownloadDir             string               `json:"downloadDir"`
	DownloadedEver          int                  `json:"downloadedEver"`
	DownloadLimit           int                  `json:"downloadLimit"`
	DownloadLimited         bool                 `json:"downloadLimited"`
	Error                   int                  `json:"error"`
	ErrorString             string               `json:"errorString"`
	Eta                     int                  `json:"eta"`
	EtaIdle                 int                  `json:"etaIdle"`
	Files                   []TorrentFile        `json:"files"`
	FileStats               []TorrentFileStat    `json:"fileStats"`
	HashString              string               `json:"hashString"`
	HaveUnchecked           int                  `json:"haveUnchecked"`
	HaveValid               int                  `json:"haveValid"`
	HonorsSessionLimits     bool                 `json:"honorsSessionLimits"`
	IsFinished              bool                 `json:"isFinished"`
	IsPrivate               bool                 `json:"isPrivate"`
	IsStalled               bool                 `json:"isStalled"`
	LeftUntilDone           int                  `json:"leftUntilDone"`
	MagnetLink              string               `json:"magnetLink"`
	ManualAnnounceTime      Time                 `json:"manualAnnounceTime"`
	MaxConnectedPeers       int                  `json:"maxConnectedPeers"`
	MetadataPercentComplete float64              `json:"metadataPercentComplete"`
	Name                    string               `json:"name"`
	PeerLimit               int                  `json:"peer-limit"`
	Peers                   []TorrentPeer        `json:"peers"`
	PeersConnected          int                  `json:"peersConnected"`
	PeersFrom               *TorrentPeersFrom    `json:"peersFrom"`
	PeersGettingFromUs      int                  `json:"peersGettingFromUs"`
	PeersSendingToUs        int                  `json:"peersSendingToUs"`
	PercentDone             float64              `json:"percentDone"`
	Pieces                  string               `json:"pieces"`
	PieceCount              int                  `json:"pieceCount"`
	PieceSize               int                  `json:"pieceSize"`
	Priorities              []int                `json:"priorities"`
	QueuePosition           int                  `json:"queuePosition"`
	RateDownload            int                  `json:"rateDownload"`
	RateUpload              int                  `json:"rateUpload"`
	RecheckProgress         float64              `json:"recheckProgress"`
	SecondsDownloading      int                  `json:"secondsDownloading"`
	SecondsSeeding          int                  `json:"secondsSeeding"`
	SeedIdleLimit           int                  `json:"seedIdleLimit"`
	SeedIdleMode            int                  `json:"seedIdleMode"`
	SeedRatioLimit          float64              `json:"seedRatioLimit"`
	SeedRatioMode           int                  `json:"seedRatioMode"`
	SizeWhenDone            int                  `json:"sizeWhenDone"`
	StartDate               int                  `json:"startDate"`
	Status                  int                  `json:"status"`
	Trackers                []TorrentTracker     `json:"trackers"`
	TrackerStats            []TorrentTrackerStat `json:"trackerStats"`
	TotalSize               int                  `json:"totalSize"`
	TorrentFile             string               `json:"torrentFile"`
	UploadedEver            int                  `json:"uploadedEver"`
	UploadLimit             int                  `json:"uploadLimit"`
	UploadLimited           bool                 `json:"uploadLimited"`
	UploadRatio             float64              `json:"uploadRatio"`
	Wanted                  []int                `json:"wanted"`
	Webseeds                []string             `json:"webseeds"`
	WebseedsSendingToUs     int                  `json:"webseedsSendingToUs"`
}

type TorrentFile struct {
	BytesCompleted int    `json:"bytesCompleted"`
	Length         int    `json:"length"`
	Name           string `json:"name"`
}

type TorrentFileStat struct {
	BytesCompleted int  `json:"bytesCompleted"`
	Wanted         bool `json:"wanted"`
	Priority       int  `json:"priority"`
}

type TorrentPeer struct {
	Address            string  `json:"address"`
	ClientName         string  `json:"clientName"`
	ClientIsChoked     bool    `json:"clientIsChoked"`
	ClientIsInterested bool    `json:"clientIsInterested"`
	FlagStr            string  `json:"flagStr"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom"`
	IsEncrypted        bool    `json:"isEncrypted"`
	IsComing           bool    `json:"isIncoming"`
	IsUploadingTo      bool    `json:"isUploadingTo"`
	IsUTP              bool    `json:"isUTP"`
	PeerIsChoked       bool    `json:"peerIsChoked"`
	PeerIsInterested   bool    `json:"peerIsInterested"`
	Port               int     `json:"port"`
	Progress           float64 `json:"progress"`
	RateToClient       int     `json:"rateToClient"`
	RateToPeer         int     `json:"rateToPeer"`
}

type TorrentPeersFrom struct {
	FromCache    int `json:"fromCache"`
	FromDHT      int `json:"fromDht"`
	FromIncoming int `json:"fromIncoming"`
	FromLPD      int `json:"fromLpd"`
	FromLTEP     int `json:"fromLtep"`
	FromPEX      int `json:"fromPex"`
	FromTracker  int `json:"fromTracker"`
}

type TorrentTracker struct {
	ID       int    `json:"id"`
	Announce string `json:"announce"`
	Scrape   string `json:"scrape"`
	Tier     int    `json:"tier"`
}

type TorrentTrackerStat struct {
	ID                    int    `json:"id"`
	Announce              string `json:"announce"`
	AnnounceState         int    `json:"announceState"`
	DownloadCount         int    `json:"downloadCount"`
	HasAnnounced          bool   `json:"hasAnnounced"`
	HasScraped            bool   `json:"hasScraped"`
	Host                  string `json:"host"`
	IsBackup              bool   `json:"isBackup"`
	LastAnnouncePeerCount int    `json:"lastAnnouncePeerCount"`
	LastAnnounceResult    string `json:"lastAnnounceResult"`
	LastAnnounceStartTime Time   `json:"lastAnnounceStartTime"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded"`
	LastAnnounceTime      Time   `json:"lastAnnounceTime"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut"`
	LastScrapeResult      string `json:"lastScrapeResult"`
	LastScrapeStartTime   Time   `json:"lastScrapeStartTime"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded"`
	LastScrapeTime        Time   `json:"lastScrapeTime"`
	LastScrapeTimedOut    int    `json:"lastScrapeTimedOut"`
	LeecherCount          int    `json:"leecherCount"`
	NextAnnounceTime      Time   `json:"nextAnnounceTime"`
	NextScrapeTime        Time   `json:"nextScrapeTime"`
	Scrape                string `json:"scrape"`
	ScrapeState           int    `json:"scrapeState"`
	SeederCount           int    `json:"seederCount"`
	Tier                  int    `json:"tier"`
}

// TorrentFields returns all of the different fields that can be requested when
// retrieving torrents
func TorrentFields() []string {
	torrentType := reflect.TypeOf(Torrent{})

	var jsonKeys []string
	for k := 0; k < torrentType.NumField(); k++ {
		field := torrentType.Field(k)
		jsonTag := field.Tag.Get("json")
		jsonKey := strings.Split(jsonTag, ",")[0]
		jsonKeys = append(jsonKeys, jsonKey)
	}

	return jsonKeys
}
