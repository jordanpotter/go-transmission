package transmission

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Time is a wrapper around time.Time with custom JSON serialization logic.
type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Time.Unix(), 10)), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	val := strings.Trim(string(b), "\"")

	unix, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return errors.Wrapf(err, "failed to parse int %q", val)
	}

	t.Time = time.Unix(unix, 0)
	return nil
}
