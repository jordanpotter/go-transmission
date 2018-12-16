// Package transmission provides functions and types to simplify interacting
// with a Transmission daemon.
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

// New returns a new Transmission client
func New(url, username, password string) *Transmission {
	return &Transmission{
		httpClient: &http.Client{},
		url:        url + "/transmission/rpc",
		username:   username,
		password:   password,
	}
}
