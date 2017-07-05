package kristianstad

import (
	"time"
	"fmt"
)

type ResultSet struct {
	Name string
	User string
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
	return fmt.Sprintf("ResultSet{Name: %s, User: %s, StartTime: %s, EndTime: %s, Results: %d, Logs: %d }",
		rs.Name, rs.User, rs.StartTime, rs.EndTime, len(rs.Results), len(rs.Log));
}

type Result struct {
	Url string
	Comment string
	StartTime time.Time
	Duration time.Duration
	StatusCode int
	Size int
	Iteration int
}

func (r Result) String() string {
	return fmt.Sprintf("%d{ %s: %s @ %s [%d] [%d] // %s}", r.Iteration, r.Url, r.Duration, r.StartTime, r.StatusCode, r.Size, r.Comment)
}
