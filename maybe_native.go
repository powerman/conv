package conv

import (
	"fmt"
	"strconv"
	"time"
)

// MaybeIntFromInt32 converts Go types.
func MaybeIntFromInt32(v *int32) *int {
	if v == nil {
		return nil
	}
	m := int(*v)
	return &m
}

// MaybeInt32FromInt converts Go types or panics on overflows.
func MaybeInt32FromInt(v *int) *int32 {
	if v == nil {
		return nil
	}
	m := int32(*v)
	if int(m) != *v {
		panic(fmt.Sprintf("value %d overflows int32", *v))
	}
	return &m
}

// MustMaybeIntFromInt64 converts Go types or panics on overflow.
func MustMaybeIntFromInt64(v *int64) *int {
	if v == nil {
		return nil
	}
	m := int(*v)
	if int64(m) != *v {
		panic(fmt.Sprintf("value %d overflows int", *v))
	}
	return &m
}

// MaybeStringFromInt converts Go types.
func MaybeStringFromInt(v *int) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(int64(*v), 10)
	return &m
}

// MaybeStringFromInt32 converts Go types.
func MaybeStringFromInt32(v *int32) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(int64(*v), 10)
	return &m
}

// MaybeStringFromInt64 converts Go types.
func MaybeStringFromInt64(v *int64) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(*v, 10)
	return &m
}

// MaybeUnixFromTime converts time to seconds or nil for zero time.
func MaybeUnixFromTime(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	v := t.Unix()
	return &v
}

// MaybeUnixMsFromTime converts time to milliseconds or nil for zero time.
func MaybeUnixMsFromTime(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	v := t.UnixNano() / int64(time.Millisecond)
	return &v
}

// MaybeUnixNanoFromTime converts time to nanoseconds or nil for zero time.
func MaybeUnixNanoFromTime(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	v := t.UnixNano()
	return &v
}
