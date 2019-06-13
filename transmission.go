// Package transmission provides functions and types to simplify interacting
// with a Transmission daemon.
package transmission

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/pkg/errors"
)

type Transmission struct {
	httpClient *http.Client
	endpoint   string
	username   string
	password   string
	token      string
	tokenLock  sync.Mutex
}

// New returns a new Transmission client
func New(endpoint, username, password string) (*Transmission, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint cannot be empty")
	} else if _, err := url.Parse(endpoint); err != nil {
		return nil, errors.Wrap(err, "failed to parse endpoint")
	}

	return &Transmission{
		httpClient: &http.Client{},
		endpoint:   endpoint + "/transmission/rpc/",
		username:   username,
		password:   password,
	}, nil
}
