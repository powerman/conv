package conv

import (
	"github.com/powerman/sensitive"
	"github.com/shopspring/decimal"
)

// MustNewDecimal returns ref to v or panics if v is not decimal.
func MustNewDecimal(v string) *decimal.Decimal {
	d := decimal.RequireFromString(v)
	return &d
}

// ValueDecimal returns dereference of v or zero value if nil.
func ValueDecimal(d *decimal.Decimal) decimal.Decimal {
	if d == nil {
		return decimal.Decimal{}
	}
	return *d
}

// NewSensitiveString returns ref to v.
func NewSensitiveString(v sensitive.String) *sensitive.String { return &v }

// ValueSensitiveString returns dereference of v or zero value if nil.
func ValueSensitiveString(v *sensitive.String) sensitive.String {
	if v == nil {
		return ""
	}
	return *v
}
