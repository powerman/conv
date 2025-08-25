package conv_test

import (
	"testing"
	"time"

	"github.com/powerman/check"

	"github.com/powerman/conv"
)

func TestTimeFromUnixMs(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	t.Equal(conv.TimeFromUnixMs(1616522747123).UTC(), time.Date(2021, time.March, 23, 18, 5, 47, 123000000, time.UTC))
	t.Equal(conv.TimeFromUnixMs(1).UTC(), time.Date(1970, time.January, 1, 0, 0, 0, 1000000, time.UTC))
	t.Equal(conv.TimeFromUnixMs(0).UTC(), time.Time{})
}

func TestUnixMsFromTime(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	t.Equal(conv.UnixMsFromTime(time.Date(2021, time.March, 23, 18, 5, 47, 123987654, time.UTC)), int64(1616522747123))
	t.Equal(conv.UnixMsFromTime(time.Date(1970, time.January, 1, 0, 0, 0, 1987654, time.UTC)), int64(1))
	t.Equal(conv.UnixMsFromTime(time.Date(1970, time.January, 1, 0, 0, 0, 987654, time.UTC)), int64(0))
	t.Equal(conv.UnixMsFromTime(time.Time{}), int64(0))
}
