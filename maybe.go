//nolint:nilnil // Valid in this context.
package conv

import "github.com/shopspring/decimal"

// MaybeStringFromDecimal converts decimal.
func MaybeStringFromDecimal(d *decimal.Decimal) *string {
	if d == nil {
		return nil
	}
	return NewString(d.String())
}

// MaybeDecimalFromString converts decimal.
func MaybeDecimalFromString(v *string) (*decimal.Decimal, error) {
	if v == nil {
		return nil, nil
	}
	d, err := decimal.NewFromString(*v)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// MustMaybeDecimalFromString converts decimal or panics if v is not decimal.
func MustMaybeDecimalFromString(v *string) *decimal.Decimal {
	if v == nil {
		return nil
	}
	return MustNewDecimal(*v)
}
