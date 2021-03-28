package conv

import "time"

// TimeFromUnixMs converts unix milliseconds to time.
// Returns zero time if v is 0 (instead of 1970-01-01).
func TimeFromUnixMs(v int64) time.Time {
	const mod = int64(time.Second / time.Millisecond)
	if v == 0 {
		return time.Time{}
	}
	return time.Unix(v/mod, (v%mod)*int64(time.Millisecond))
}

// UnixMsFromTime converts time to unix milliseconds.
// Returns 0 both for zero time (instead of -6795364578871) and zero unixtime (1970-01-01).
func UnixMsFromTime(v time.Time) int64 {
	if v.IsZero() {
		return 0
	}
	return v.UnixNano() / int64(time.Millisecond)
}
