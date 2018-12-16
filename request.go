package transmission

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const csrfHeader = "X-Transmission-Session-Id"

type request struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

type response struct {
	Result    string          `json:"result"`
	Arguments json.RawMessage `json:"arguments"`
}

// Do performs a Transmission operation. If provided a non-nil result
// parameter, will use the response from Transmission to hydrate result.
func (t *Transmission) Do(ctx context.Context, method string, arguments, result interface{}) error {
	reqPayload := request{method, arguments}
	b, err := json.Marshal(reqPayload)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	req, err := http.NewRequest("POST", t.url, bytes.NewBuffer(b))
	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	req = req.WithContext(ctx)
	req.Header.Add(csrfHeader, t.token)
	req.SetBasicAuth(t.username, t.password)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to perform request")
	}
	defer resp.Body.Close()

	t.token = resp.Header.Get("X-Transmission-Session-Id")

	if resp.StatusCode == http.StatusConflict {
		return t.Do(ctx, method, arguments, result)
	} else if resp.StatusCode != http.StatusOK {
		return errors.Errorf("received unexpected status code %d", resp.StatusCode)
	}

	var respPayload response
	if err = json.NewDecoder(resp.Body).Decode(&respPayload); err != nil {
		return errors.Wrap(err, "failed to unmarshal response")
	}

	if respPayload.Result != "success" {
		return errors.New(respPayload.Result)
	}

	if result != nil {
		if err = json.Unmarshal(respPayload.Arguments, &result); err != nil {
			return errors.Wrap(err, "failed to unmarshal response arguments")
		}
	}

	return nil
}
