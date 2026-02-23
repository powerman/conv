//nolint:nilnil // Valid in this context.
package conv

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

// MaybeIntFromInt8 converts Go types.
func MaybeIntFromInt8(v *int8) *int {
	if v == nil {
		return nil
	}
	m := int(*v)
	return &m
}

// MaybeInt8FromInt converts Go types or returns [strconv.ErrRange].
func MaybeInt8FromInt(v *int) (*int8, error) {
	if v == nil {
		return nil, nil
	}
	m := int8(*v) //nolint:gosec // False positive.
	if int(m) != *v {
		return nil, fmt.Errorf("int8 %w: %d", strconv.ErrRange, *v)
	}
	return &m, nil
}

// MustMaybeInt8FromInt converts Go types or panics on overflow.
func MustMaybeInt8FromInt(v *int) *int8 {
	m, err := MaybeInt8FromInt(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeIntFromInt16 converts Go types.
func MaybeIntFromInt16(v *int16) *int {
	if v == nil {
		return nil
	}
	m := int(*v)
	return &m
}

// MaybeInt16FromInt converts Go types or returns [strconv.ErrRange].
func MaybeInt16FromInt(v *int) (*int16, error) {
	if v == nil {
		return nil, nil
	}
	m := int16(*v) //nolint:gosec // False positive.
	if int(m) != *v {
		return nil, fmt.Errorf("int16 %w: %d", strconv.ErrRange, *v)
	}
	return &m, nil
}

// MustMaybeInt16FromInt converts Go types or panics on overflow.
func MustMaybeInt16FromInt(v *int) *int16 {
	m, err := MaybeInt16FromInt(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeIntFromInt32 converts Go types.
func MaybeIntFromInt32(v *int32) *int {
	if v == nil {
		return nil
	}
	m := int(*v)
	return &m
}

// MaybeInt32FromInt converts Go types or returns [strconv.ErrRange].
func MaybeInt32FromInt(v *int) (*int32, error) {
	if v == nil {
		return nil, nil
	}
	m := int32(*v) //nolint:gosec // False positive.
	if int(m) != *v {
		return nil, fmt.Errorf("int32 %w: %d", strconv.ErrRange, *v)
	}
	return &m, nil
}

// MustMaybeInt32FromInt converts Go types or panics on overflow.
func MustMaybeInt32FromInt(v *int) *int32 {
	m, err := MaybeInt32FromInt(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeIntFromInt64 converts Go types or returns [strconv.ErrRange].
func MaybeIntFromInt64(v *int64) (*int, error) {
	if v == nil {
		return nil, nil
	}
	m := int(*v)
	if int64(m) != *v {
		return nil, fmt.Errorf("int %w: %d", strconv.ErrRange, *v)
	}
	return &m, nil
}

// MustMaybeIntFromInt64 converts Go types or panics on overflow.
func MustMaybeIntFromInt64(v *int64) *int {
	m, err := MaybeIntFromInt64(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeInt64FromInt converts Go types.
func MaybeInt64FromInt(v *int) *int64 {
	if v == nil {
		return nil
	}
	m := int64(*v)
	return &m
}

// MaybeIntFromUInt8 converts Go types.
func MaybeIntFromUInt8(v *uint8) *int {
	if v == nil {
		return nil
	}
	m := int(*v)
	return &m
}

// MaybeIntFromUInt16 converts Go types.
func MaybeIntFromUInt16(v *uint16) *int {
	if v == nil {
		return nil
	}
	m := int(*v)
	return &m
}

// MaybeIntFromUInt32 converts Go types.
func MaybeIntFromUInt32(v *uint32) (*int, error) {
	if v == nil {
		return nil, nil
	}
	m := int(*v)
	if m < 0 {
		return nil, fmt.Errorf("uint32 %w: %d", strconv.ErrRange, *v)
	}
	return &m, nil
}

// MustMaybeIntFromUInt32 converts Go types or panics on overflow.
func MustMaybeIntFromUInt32(v *uint32) *int {
	m, err := MaybeIntFromUInt32(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeIntFromUInt64 converts Go types.
func MaybeIntFromUInt64(v *uint64) (*int, error) {
	if v == nil {
		return nil, nil
	}
	m := int(*v) //nolint:gosec // False positive.
	if m < 0 || uint64(m) != *v {
		return nil, fmt.Errorf("uint64 %w: %d", strconv.ErrRange, *v)
	}
	return &m, nil
}

// MustMaybeIntFromUInt64 converts Go types or panics on overflow.
func MustMaybeIntFromUInt64(v *uint64) *int {
	m, err := MaybeIntFromUInt64(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeStringFromInt converts Go types.
func MaybeStringFromInt(v *int) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(int64(*v), 10)
	return &m
}

// MaybeIntFromString converts Go types.
func MaybeIntFromString(v *string) (*int, error) {
	if v == nil {
		return nil, nil
	}
	m, err := strconv.Atoi(*v)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// MaybeStringFromInt8 converts Go types.
func MaybeStringFromInt8(v *int8) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(int64(*v), 10)
	return &m
}

// MaybeInt8FromString converts Go types.
func MaybeInt8FromString(v *string) (*int8, error) {
	if v == nil {
		return nil, nil
	}
	i, err := strconv.ParseInt(*v, 10, 8)
	if err != nil {
		return nil, err
	}
	m := int8(i)
	return &m, nil
}

// MustMaybeInt8FromString converts Go types or panics on overflow/parse error.
func MustMaybeInt8FromString(v *string) *int8 {
	m, err := MaybeInt8FromString(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeStringFromInt16 converts Go types.
func MaybeStringFromInt16(v *int16) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(int64(*v), 10)
	return &m
}

// MaybeInt16FromString converts Go types.
func MaybeInt16FromString(v *string) (*int16, error) {
	if v == nil {
		return nil, nil
	}
	i, err := strconv.ParseInt(*v, 10, 16)
	if err != nil {
		return nil, err
	}
	m := int16(i)
	return &m, nil
}

// MustMaybeInt16FromString converts Go types or panics on overflow/parse error.
func MustMaybeInt16FromString(v *string) *int16 {
	m, err := MaybeInt16FromString(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeStringFromInt32 converts Go types.
func MaybeStringFromInt32(v *int32) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(int64(*v), 10)
	return &m
}

// MaybeInt32FromString converts Go types.
func MaybeInt32FromString(v *string) (*int32, error) {
	if v == nil {
		return nil, nil
	}
	i, err := strconv.ParseInt(*v, 10, 32)
	if err != nil {
		return nil, err
	}
	m := int32(i)
	return &m, nil
}

// MustMaybeInt32FromString converts Go types or panics on overflow/parse error.
func MustMaybeInt32FromString(v *string) *int32 {
	m, err := MaybeInt32FromString(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeStringFromInt64 converts Go types.
func MaybeStringFromInt64(v *int64) *string {
	if v == nil {
		return nil
	}
	m := strconv.FormatInt(*v, 10)
	return &m
}

// MaybeInt64FromString converts Go types.
func MaybeInt64FromString(v *string) (*int64, error) {
	if v == nil {
		return nil, nil
	}
	m, err := strconv.ParseInt(*v, 10, 64)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// MustMaybeInt64FromString converts Go types or panics on overflow/parse error.
func MustMaybeInt64FromString(v *string) *int64 {
	m, err := MaybeInt64FromString(v)
	if err != nil {
		panic(err)
	}
	return m
}

// MaybeUnixFromTime converts time to seconds or nil for zero time.
func MaybeUnixFromTime(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	vSec := t.Unix()
	return &vSec
}

// MaybeTimeFromUnix converts seconds to time or nil to zero time.
func MaybeTimeFromUnix(v *int64) (t time.Time) {
	if v == nil {
		return t
	}
	return time.Unix(*v, 0)
}

// MaybeUnixMsFromTime converts time to milliseconds or nil for zero time.
func MaybeUnixMsFromTime(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	v := t.UnixNano() / int64(time.Millisecond)
	return &v
}

// MaybeTimeFromUnixMs converts milliseconds to time or nil to zero time.
func MaybeTimeFromUnixMs(v *int64) (t time.Time) {
	if v == nil {
		return t
	}
	const mod = int64(time.Second / time.Millisecond)
	return time.Unix(*v/mod, *v%mod*int64(time.Millisecond))
}

// MaybeUnixNanoFromTime converts time to nanoseconds or nil for zero time.
func MaybeUnixNanoFromTime(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	vNano := t.UnixNano()
	return &vNano
}

// MaybeTimeFromUnixNano converts nanoseconds to time or nil to zero time.
func MaybeTimeFromUnixNano(v *int64) (t time.Time) {
	if v == nil {
		return t
	}
	const mod = int64(time.Second / time.Nanosecond)
	return time.Unix(*v/mod, *v%mod)
}

// MaybeNullBool converts to sql Null.
func MaybeNullBool(v *bool) sql.NullBool {
	if v == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  *v,
		Valid: true,
	}
}

// MaybeBoolFromNull converts from sql Null.
func MaybeBoolFromNull(n sql.NullBool) *bool {
	if !n.Valid {
		return nil
	}
	return &n.Bool
}

// MaybeNullFloat64 converts to sql Null.
func MaybeNullFloat64(v *float64) sql.NullFloat64 {
	if v == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: *v,
		Valid:   true,
	}
}

// MaybeFloat64FromNull converts from sql Null.
func MaybeFloat64FromNull(n sql.NullFloat64) *float64 {
	if !n.Valid {
		return nil
	}
	return &n.Float64
}

// MaybeNullInt32 converts to sql Null.
func MaybeNullInt32(v *int32) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: *v,
		Valid: true,
	}
}

// MaybeInt32FromNull converts from sql Null.
func MaybeInt32FromNull(n sql.NullInt32) *int32 {
	if !n.Valid {
		return nil
	}
	return &n.Int32
}

// MaybeNullInt64 converts to sql Null.
func MaybeNullInt64(v *int64) sql.NullInt64 {
	if v == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: *v,
		Valid: true,
	}
}

// MaybeInt64FromNull converts from sql Null.
func MaybeInt64FromNull(n sql.NullInt64) *int64 {
	if !n.Valid {
		return nil
	}
	return &n.Int64
}

// MaybeNullString converts to sql Null.
func MaybeNullString(v *string) sql.NullString {
	if v == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *v,
		Valid:  true,
	}
}

// MaybeStringFromNull converts from sql Null.
func MaybeStringFromNull(n sql.NullString) *string {
	if !n.Valid {
		return nil
	}
	return &n.String
}

// MaybeNullTime converts to sql Null.
func MaybeNullTime(v time.Time) sql.NullTime {
	if v.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  v,
		Valid: true,
	}
}

// MaybeTimeFromNull converts from sql Null.
func MaybeTimeFromNull(n sql.NullTime) time.Time {
	return n.Time
}
