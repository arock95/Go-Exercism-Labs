// Package gigasecond provides function to add 1 gigasecond to a date/time
package gigasecond

import "time"

// AddGigasecond takes a date/time and adds 1 gigasecond to it
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
