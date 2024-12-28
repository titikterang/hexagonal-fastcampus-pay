package types

import "time"

type CurrentTime struct {
	Mock bool
	Now  time.Time
}

func (t *CurrentTime) CurrentTime() time.Time {
	if t.Mock {
		return t.Now
	}
	return time.Now()
}
