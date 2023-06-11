package utils

import "time"

// NowMillis return current time
func NowMillis() int64 {
	nanosecond := time.Now().UnixNano()
	millis := nanosecond / 1000000
	return millis
}
