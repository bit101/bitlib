package blcolor

import (
	"testing"

	"github.com/bit101/bitlib/blmath"
)

func TestQuantChannel(t *testing.T) {
	type test struct {
		val      float64
		quant    int
		expected float64
	}

	tests := []test{
		{0.0, 2, 0.0},
		{0.0001, 2, 0.0},
		{0.25, 2, 0.0},
		{0.499999, 2, 0.0},
		{0.5, 2, 1.0},
		{0.75, 2, 1.0},
		{0.999, 2, 1.0},
		{1, 2, 1.0},

		{0.0, 3, 0.0},
		{0.0001, 3, 0.0},
		{0.25, 3, 0.0},
		{0.33, 3, 0.0},
		{0.34, 3, 0.5},
		{0.5, 3, 0.5},
		{0.666, 3, 0.5},
		{0.667, 3, 1.0},
		{0.999, 3, 1.0},
		{1.0, 3, 1.0},

		{0.0, 4, 0.0},
		{0.0001, 4, 0.0},
		{0.24, 4, 0.0},
		{0.25, 4, 0.333333333},
		{0.26, 4, 0.333333333},
		{0.49, 4, 0.333333333},
		{0.5, 4, 0.666666666},
		{0.74, 4, 0.666666666},
		{0.75, 4, 1.0},
		{0.999, 4, 1.0},
		{1.0, 4, 1.0},

		{0.0, 16, 0.0},
		{0.0001, 16, 0.0},
		{0.08, 16, 1.0 / 15.0},
		{0.15, 16, 2.0 / 15.0},
		{0.22, 16, 3.0 / 15.0},
		{0.29, 16, 4.0 / 15.0},
		{0.36, 16, 5.0 / 15.0},
		{0.43, 16, 6.0 / 15.0},
		{0.49, 16, 7.0 / 15.0},
		{0.56, 16, 8.0 / 15.0},
		{0.62, 16, 9.0 / 15.0},
		{0.67, 16, 10.0 / 15.0},
		{0.74, 16, 11.0 / 15.0},
		{0.81, 16, 12.0 / 15.0},
		{0.87, 16, 13.0 / 15.0},
		{0.93, 16, 14.0 / 15.0},
		{0.999, 16, 1.0},
		{1.0, 16, 1.0},
	}

	for _, tc := range tests {
		result := quantChannel(tc.val, tc.quant)
		if !blmath.Equalish(result, tc.expected, 0.0001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestEqual(t *testing.T) {
	type test struct {
		colorA   Color
		colorB   Color
		expected bool
	}

	tests := []test{
		{RGB(1, 1, 1), RGB(1, 1, 1), true},
		{White, RGB(1, 1, 1), true},
		{Black, RGB(0, 0, 0), true},
		{RGB(0.15, 0.43, 0.87).Quant(16), RGB(2.0/15.0, 6.0/15.0, 13.0/15.0), true},
	}

	for _, tc := range tests {
		result := tc.colorA.Equals(tc.colorB)
		if result != tc.expected {
			t.Errorf("Expected %t, got %t\n", tc.expected, result)
		}
	}
}

func TestAddColor(t *testing.T) {
	type test struct {
		colorA   Color
		colorB   Color
		expected Color
	}

	tests := []test{
		{RGB(0.5, 0.5, 0.5), RGB(0.25, 0.25, 0.25), RGB(0.75, 0.75, 0.75)},
		{RGB(0.6, 0.6, 0.6), RGB(0.5, 0.5, 0.5), RGB(1, 1, 1)},
	}

	for _, tc := range tests {
		result := tc.colorA.AddColor(tc.colorB)
		if !result.Equals(tc.expected) {
			t.Errorf("Expected %v, got %v\n", tc.expected, result)
		}
	}
}
