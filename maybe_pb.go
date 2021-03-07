package conv

import (
	"time"

	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// MaybeBool converts from protobuf WKT.
func MaybeBool(wkt *wrapperspb.BoolValue) *bool {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybePbBool converts to protobuf WKT.
func MaybePbBool(v *bool) *wrapperspb.BoolValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Bool(*v)
}

// MaybeFloat32 converts from protobuf WKT.
func MaybeFloat32(wkt *wrapperspb.FloatValue) *float32 {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybePbFloat converts to protobuf WKT.
func MaybePbFloat(v *float32) *wrapperspb.FloatValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Float(*v)
}

// MaybeFloat64 converts from protobuf WKT.
func MaybeFloat64(wkt *wrapperspb.DoubleValue) *float64 {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybePbDouble converts to protobuf WKT.
func MaybePbDouble(v *float64) *wrapperspb.DoubleValue {
	if v == nil {
		return nil
	}
	return wrapperspb.Double(*v)
}

// MaybeInt32 converts from protobuf WKT.
func MaybeInt32(wkt *wrapperspb.Int32Value) *int32 {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybePbInt32 converts to protobuf WKT.
func MaybePbInt32(v *int32) *wrapperspb.Int32Value {
	if v == nil {
		return nil
	}
	return wrapperspb.Int32(*v)
}

// MaybeInt64 converts from protobuf WKT.
func MaybeInt64(wkt *wrapperspb.Int64Value) *int64 {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybePbInt64 converts to protobuf WKT.
func MaybePbInt64(v *int64) *wrapperspb.Int64Value {
	if v == nil {
		return nil
	}
	return wrapperspb.Int64(*v)
}

// MaybeString converts from protobuf WKT.
func MaybeString(wkt *wrapperspb.StringValue) *string {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybePbString converts to protobuf WKT.
func MaybePbString(v *string) *wrapperspb.StringValue {
	if v == nil {
		return nil
	}
	return wrapperspb.String(*v)
}

// MaybeUInt32 converts from protobuf WKT.
func MaybeUInt32(wkt *wrapperspb.UInt32Value) *uint32 {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybeUInt32Value converts to protobuf WKT.
func MaybeUInt32Value(v *uint32) *wrapperspb.UInt32Value {
	if v == nil {
		return nil
	}
	return wrapperspb.UInt32(*v)
}

// MaybeUInt64 converts from protobuf WKT.
func MaybeUInt64(wkt *wrapperspb.UInt64Value) *uint64 {
	if wkt == nil {
		return nil
	}
	return &wkt.Value
}

// MaybeUInt64Value converts to protobuf WKT.
func MaybeUInt64Value(v *uint64) *wrapperspb.UInt64Value {
	if v == nil {
		return nil
	}
	return wrapperspb.UInt64(*v)
}

// MaybeDuration converts from protobuf WKT.
func MaybeDuration(wkt *durationpb.Duration) *time.Duration {
	if wkt == nil {
		return nil
	}
	v := wkt.AsDuration()
	return &v
}

// MaybePbDuration converts to protobuf WKT.
func MaybePbDuration(v *time.Duration) *durationpb.Duration {
	if v == nil {
		return nil
	}
	return durationpb.New(*v)
}

// MaybeTime converts from protobuf WKT, returns zero time for nil.
func MaybeTime(wkt *timestamppb.Timestamp) time.Time {
	if wkt == nil {
		return time.Time{}
	}
	return wkt.AsTime()
}

// MaybePbTimestamp converts to protobuf WKT, returns nil for zero time.
func MaybePbTimestamp(v time.Time) *timestamppb.Timestamp {
	if v.IsZero() {
		return nil
	}
	return timestamppb.New(v)
}

// MaybeDecimal converts from protobuf WKT.
func MaybeDecimal(wkt *wrapperspb.StringValue) (*decimal.Decimal, error) {
	if wkt == nil {
		return nil, nil
	}
	d, err := decimal.NewFromString(wkt.Value)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// MaybePbStringFromDecimal converts to protobuf WKT.
func MaybePbStringFromDecimal(d *decimal.Decimal) *wrapperspb.StringValue {
	if d == nil {
		return nil
	}
	return wrapperspb.String(d.String())
}
