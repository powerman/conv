package conv_test

import (
	"testing"

	"github.com/powerman/check"
	"github.com/shopspring/decimal"

	"github.com/powerman/conv"
)

func TestValueDecimal(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	d := conv.ValueDecimal(nil)
	t.Zero(d)
	t.Equal(d.String(), "0")
	t.True(d.IsZero())
	t.Equal(d.Cmp(decimal.Zero), 0)
}
