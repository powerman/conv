package conv

// NewBool returns ref to v.
func NewBool(v bool) *bool { return &v }

// ValueBool returns dereference of v or zero value if nil.
func ValueBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// NewInt returns ref to v.
func NewInt(v int) *int { return &v }

// ValueInt returns dereference of v or zero value if nil.
func ValueInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// NewInt8 returns ref to v.
func NewInt8(v int8) *int8 { return &v }

// ValueInt8 returns dereference of v or zero value if nil.
func ValueInt8(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

// NewInt16 returns ref to v.
func NewInt16(v int16) *int16 { return &v }

// ValueInt16 returns dereference of v or zero value if nil.
func ValueInt16(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

// NewInt32 returns ref to v.
func NewInt32(v int32) *int32 { return &v }

// ValueInt32 returns dereference of v or zero value if nil.
func ValueInt32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

// NewInt64 returns ref to v.
func NewInt64(v int64) *int64 { return &v }

// ValueInt64 returns dereference of v or zero value if nil.
func ValueInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

// NewUInt returns ref to v.
func NewUInt(v uint) *uint { return &v }

// ValueUInt returns dereference of v or zero value if nil.
func ValueUInt(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

// NewUInt8 returns ref to v.
func NewUInt8(v uint8) *uint8 { return &v }

// ValueUInt8 returns dereference of v or zero value if nil.
func ValueUInt8(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

// NewUInt16 returns ref to v.
func NewUInt16(v uint16) *uint16 { return &v }

// ValueUInt16 returns dereference of v or zero value if nil.
func ValueUInt16(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

// NewUInt32 returns ref to v.
func NewUInt32(v uint32) *uint32 { return &v }

// ValueUInt32 returns dereference of v or zero value if nil.
func ValueUInt32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

// NewUInt64 returns ref to v.
func NewUInt64(v uint64) *uint64 { return &v }

// ValueUInt64 returns dereference of v or zero value if nil.
func ValueUInt64(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

// NewFloat32 returns ref to v.
func NewFloat32(v float32) *float32 { return &v }

// ValueFloat32 returns dereference of v or zero value if nil.
func ValueFloat32(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

// NewFloat64 returns ref to v.
func NewFloat64(v float64) *float64 { return &v }

// ValueFloat64 returns dereference of v or zero value if nil.
func ValueFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

// NewString returns ref to v.
func NewString(v string) *string { return &v }

// ValueString returns dereference of v or zero value if nil.
func ValueString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
