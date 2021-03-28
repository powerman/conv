package conv_test

import (
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/powerman/check"

	"github.com/powerman/conv"
)

func TestMaybeInt8FromString(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	tests := []struct {
		given   *string
		want    *int8
		wantErr error
	}{
		{nil, nil, nil},
		{conv.NewString(""), nil, strconv.ErrSyntax},
		{conv.NewString("-42"), conv.NewInt8(-42), nil},
		{conv.NewString("1234"), nil, strconv.ErrRange},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			res, err := conv.MaybeInt8FromString(tc.given)
			t.Err(err, tc.wantErr)
			t.DeepEqual(res, tc.want)
		})
	}
}

func TestMaybeIntFromUInt64(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	tests := []struct {
		given   *uint64
		want    *int
		wantErr error
	}{
		{nil, nil, nil},
		{conv.NewUInt64(0), conv.NewInt(0), nil},
		{conv.NewUInt64(math.MaxInt64), conv.NewInt(math.MaxInt64), nil},
		{conv.NewUInt64(math.MaxInt64 + 1), nil, strconv.ErrRange},
		{conv.NewUInt64(math.MaxUint64), nil, strconv.ErrRange},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			res, err := conv.MaybeIntFromUInt64(tc.given)
			t.Err(err, tc.wantErr)
			t.DeepEqual(res, tc.want)
		})
	}
}

func TestMaybeUnixFromTime(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	var (
		tZero   time.Time
		tSecond = time.Date(2021, 1, 2, 3, 4, 5, 0, time.UTC)
		tMilli  = time.Date(2021, 1, 2, 3, 4, 5, 123000000, time.UTC)
		tNano   = time.Date(2021, 1, 2, 3, 4, 5, 123456789, time.UTC)
	)

	t.Nil(conv.MaybeUnixFromTime(tZero))
	t.Nil(conv.MaybeUnixMsFromTime(tZero))
	t.Nil(conv.MaybeUnixNanoFromTime(tZero))
	t.Zero(conv.MaybeTimeFromUnix(nil))
	t.Zero(conv.MaybeTimeFromUnixMs(nil))
	t.Zero(conv.MaybeTimeFromUnixNano(nil))

	t.Equal(conv.MaybeTimeFromUnix(conv.MaybeUnixFromTime(tSecond)), tSecond)
	t.NotEqual(conv.MaybeTimeFromUnix(conv.MaybeUnixFromTime(tMilli)), tMilli)
	t.NotEqual(conv.MaybeTimeFromUnix(conv.MaybeUnixFromTime(tNano)), tNano)

	t.Equal(conv.MaybeTimeFromUnixMs(conv.MaybeUnixMsFromTime(tSecond)), tSecond)
	t.Equal(conv.MaybeTimeFromUnixMs(conv.MaybeUnixMsFromTime(tMilli)), tMilli)
	t.NotEqual(conv.MaybeTimeFromUnixMs(conv.MaybeUnixMsFromTime(tNano)), tNano)

	t.Equal(conv.MaybeTimeFromUnixNano(conv.MaybeUnixNanoFromTime(tSecond)), tSecond)
	t.Equal(conv.MaybeTimeFromUnixNano(conv.MaybeUnixNanoFromTime(tMilli)), tMilli)
	t.Equal(conv.MaybeTimeFromUnixNano(conv.MaybeUnixNanoFromTime(tNano)), tNano)
}
