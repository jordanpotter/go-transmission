package transmission

import (
	"net/http"
)

type Transmission struct {
	httpClient *http.Client
	url        string
	username   string
	password   string
	token      string
}

func New(url, username, password string) *Transmission {
	return &Transmission{
		httpClient: &http.Client{},
		url:        url + "/transmission/rpc",
		username:   username,
		password:   password,
	}
}
