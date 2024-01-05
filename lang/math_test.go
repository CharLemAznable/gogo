package lang_test

import (
	"github.com/CharLemAznable/gogo/lang"
	"testing"
)

func TestMinMax(t *testing.T) {
	tests := []struct {
		x, y, min, max int64
	}{
		{3, 5, 3, 5},
		{5, 3, 3, 5},
		{0, 0, 0, 0},
		{-5, -3, -5, -3},
		{-3, -5, -5, -3},
		{9223372036854775807, 9223372036854775807, 9223372036854775807, 9223372036854775807},
		{-9223372036854775808, -9223372036854775808, -9223372036854775808, -9223372036854775808},
		{9223372036854775807, -9223372036854775808, -9223372036854775808, 9223372036854775807},
		{-9223372036854775808, 9223372036854775807, -9223372036854775808, 9223372036854775807},
	}

	for _, test := range tests {
		min := lang.Min(test.x, test.y)
		if min != test.min {
			t.Errorf("Min(%d, %d) = %d; want %d", test.x, test.y, min, test.min)
		}
		max := lang.Max(test.x, test.y)
		if max != test.max {
			t.Errorf("Max(%d, %d) = %d; want %d", test.x, test.y, max, test.max)
		}
	}
}
