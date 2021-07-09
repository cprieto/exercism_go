// Package gigasecond calculates stuff with gigaseconds
package gigasecond

import "time"

const gigasecond = 1e9

// AddGigasecond add a gigasecond amount to a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigasecond * 1e9))
}
