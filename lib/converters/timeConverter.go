package converters

import (
	"time"
)

func ExpiresToTime(expire int, t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(expire))
}
