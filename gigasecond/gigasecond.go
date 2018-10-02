package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	tAdd := t.Add(time.Second * 1000000000)
	return tAdd
}
