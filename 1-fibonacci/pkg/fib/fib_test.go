package fib

import (
	"testing"
)

func TestCalc(t *testing.T) {
	samples := []struct {
		step     int
		expected int
	}{
		{
			step:     1,
			expected: 0,
		},
		{
			step:     2,
			expected: 1,
		},
		{
			step:     3,
			expected: 1,
		},
		{
			step:     20,
			expected: 4181,
		},
	}

	for _, sample := range samples {
		if got := Calc(sample.step); got != sample.expected {
			t.Errorf("got: %v, expected: %v", got, sample.expected)
		}
	}
}
