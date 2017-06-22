package kristianstad

import (
	"time"
	"fmt"
	"encoding/json"
)

type ResultSet struct {
	Name string
	StartTime time.Time
	EndTime time.Time
	Results []Result
	Log []string
}

func (rs *ResultSet) Write(p []byte) (n int, err error) {
	rs.Log = append(rs.Log, string(p));
	return len(p), nil
}

func (rs ResultSet) String() string {
	j, err := json.MarshalIndent(rs, "", "  ");
	if err != nil {
		return "failed to serialize the ResultSet"
	}
	return string(j)
}

type Result struct {
	Url string
	Comment string
	StartTime time.Time
	Duration time.Duration
	StatusCode int
	Size int
}

func (r Result) String() string {
	return fmt.Sprintf("%s: %s [%d] // %s", r.Url, r.Duration, r.StatusCode, r.Comment)
}

func (r Result) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}
