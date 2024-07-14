package blmath

import (
	"math"
	"testing"
)

func TestConstants(t *testing.T) {
	type test struct {
		val, expected float64
	}
	tests := []test{
		{TwoPi, 6.283185},
		{Tau, 6.283185},
		{Tau, TwoPi},
		{HalfPi, 1.5707963},
	}
	for _, tc := range tests {
		if !Equalish(tc.val, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, tc.val)
		}
	}
}

func TestDifference(t *testing.T) {
	type test struct {
		x, y, expected float64
	}

	tests := []test{
		{1.1, 1.2, 0.1},
		{1.3, 1.2, 0.1},
		{-1.1, -1.2, 0.1},
		{-1.1, 1.2, 2.3},
		{1.1, -1.2, 2.3},
		{0.0, 0.0, 0.0},
	}

	for _, tc := range tests {
		result := Difference(tc.x, tc.y)
		if !Equalish(result, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestNorm(t *testing.T) {
	type test struct {
		val, min, max, expected float64
	}

	tests := []test{
		{100, 100, 150, 0.0},
		{125, 100, 150, 0.5},
		{150, 100, 150, 1.0},
		{75, 100, 150, -0.5},
		{175, 100, 150, 1.5},

		{-100, -100, -150, 0.0},
		{-125, -100, -150, 0.5},
		{-150, -100, -150, 1.0},
		{-75, -100, -150, -0.5},
		{-175, -100, -150, 1.5},

		{0.0, -10, 10, 0.5},
		{0.0, -5, 15, 0.25},
		{0.0, -15, 5, 0.75},

		{7, 5, 9, 0.5},
		{0.3, 0.0, 0.5, 0.6},
		{math.Pi, 0.0, Tau, 0.5},
		{HalfPi, 0.0, Tau, 0.25},
	}

	for _, tc := range tests {
		result := Norm(tc.val, tc.min, tc.max)
		if result != tc.expected {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestLerp(t *testing.T) {
	type test struct {
		val, min, max, expected float64
	}

	tests := []test{
		{0.0, 100, 150, 100},
		{0.5, 100, 150, 125},
		{1.0, 100, 150, 150},
		{-0.5, 100, 150, 75},
		{1.5, 100, 150, 175},

		{0.0, -100, -150, -100},
		{0.5, -100, -150, -125},
		{1.0, -100, -150, -150},
		{-0.5, -100, -150, -75},
		{1.5, -100, -150, -175},

		{0.5, -10, 10, 0.0},
		{0.25, -5, 15, 0.0},
		{0.75, -15, 5, 0.0},

		{0.5, 5, 9, 7},
		{0.6, 0, 0.5, 0.3},
		{0.5, 0.0, Tau, math.Pi},
		{0.25, 0.0, Tau, HalfPi},
	}

	for _, tc := range tests {
		result := Lerp(tc.val, tc.min, tc.max)
		if result != tc.expected {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestMap(t *testing.T) {

	type test struct {
		val, sMin, sMax, dMin, dMax, expected float64
	}

	tests := []test{
		// positives
		{5, 5, 15, 200, 300, 200},
		{10, 5, 15, 200, 300, 250},
		{15, 5, 15, 200, 300, 300},
		{0, 5, 15, 200, 300, 150},
		{20, 5, 15, 200, 300, 350},

		// negative source
		{-5, -5, -15, 200, 300, 200},
		{-10, -5, -15, 200, 300, 250},
		{-15, -5, -15, 200, 300, 300},
		{0, -5, -15, 200, 300, 150},
		{-20, -5, -15, 200, 300, 350},

		// negative dest
		{5, 5, 15, -200, -300, -200},
		{10, 5, 15, -200, -300, -250},
		{15, 5, 15, -200, -300, -300},
		{0, 5, 15, -200, -300, -150},
		{20, 5, 15, -200, -300, -350},

		// both negative
		{-5, -5, -15, -200, -300, -200},
		{-10, -5, -15, -200, -300, -250},
		{-15, -5, -15, -200, -300, -300},
		{0, -5, -15, -200, -300, -150},
		{-20, -5, -15, -200, -300, -350},

		// split source
		{-5, -5, 5, 200, 300, 200},
		{0, -5, 5, 200, 300, 250},
		{5, -5, 5, 200, 300, 300},
		{-10, -5, 5, 200, 300, 150},
		{10, -5, 5, 200, 300, 350},

		// consts
		{0, 0, Tau, 200, 300, 200},
		{math.Pi, 0, Tau, 200, 300, 250},
		{HalfPi, 0, Tau, 200, 300, 225},
		{Tau, 0, Tau, 200, 300, 300},
		{200, 200, 300, 0, Tau, 0},
		{250, 200, 300, 0, Tau, math.Pi},
		{225, 200, 300, 0, Tau, HalfPi},
		{300, 200, 300, 0, Tau, Tau},

		// min/max equal
		{5, 100, 100, 300, 400, 350},
	}

	for _, tc := range tests {
		result := Map(tc.val, tc.sMin, tc.sMax, tc.dMin, tc.dMax)
		if result != tc.expected {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestWrap(t *testing.T) {
	type test struct {
		val, min, max, expected float64
	}

	tests := []test{
		// zero based
		{5, 0, 10, 5},
		{-1, 0, 10, 9},
		{-9, 0, 10, 1},
		{-10, 0, 10, 0},
		{-11, 0, 10, 9},
		{11, 0, 10, 1},
		{19, 0, 10, 9},

		// non zero
		{2, 5, 20, 17},
		{4, 5, 20, 19},
		{6, 5, 20, 6},
		{22, 5, 20, 7},

		// reversed
		{2, 20, 5, 17},
		{4, 20, 5, 19},
		{6, 20, 5, 6},
		{22, 20, 5, 7},

		// split
		{-10, -5, 5, 0},
		{-6, -5, 5, 4},
		{6, -5, 5, -4},

		// negative
		{-12, -10, -5, -7},
		{3, -10, -5, -7},
	}

	for _, tc := range tests {
		result := Wrap(tc.val, tc.min, tc.max)
		if result != tc.expected {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestClamp(t *testing.T) {
	type test struct {
		val, min, max, expected float64
	}

	tests := []test{
		// positive
		{5, 2, 10, 5},
		{0, 2, 10, 2},
		{-5, 2, 10, 2},
		{10, 2, 10, 10},
		{15, 2, 10, 10},

		// reversed
		{5, 10, 2, 5},
		{0, 10, 2, 2},
		{-5, 10, 2, 2},
		{10, 10, 2, 10},
		{15, 10, 2, 10},

		// negative
		{-15, -2, -10, -10},
		{5, -2, -10, -2},

		// split
		{-7, -5, 5, -5},
		{7, -5, 5, 5},
	}

	for _, tc := range tests {
		result := Clamp(tc.val, tc.min, tc.max)
		if result != tc.expected {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestRoundTo(t *testing.T) {
	type test struct {
		val      float64
		decimals int
		expected float64
	}

	tests := []test{
		// positive
		{1.23456, 0, 1},
		{1.23456, 1, 1.2},
		{1.23456, 2, 1.23},
		{1.23456, 3, 1.235},
		{1.23456, 4, 1.2346},

		// negative
		{-1.23456, 0, -1},
		{-1.23456, 1, -1.2},
		{-1.23456, 2, -1.23},
		{-1.23456, 3, -1.235},
		{-1.23456, 4, -1.2346},

		// less than 0 decimals
		{123456, -1, 123460},
		{123456, -2, 123500},
		{123456, -3, 123000},
		{123456, -4, 120000},
		{123456, -5, 100000},

		// negative less than 0 decimals
		{-123456, -1, -123460},
		{-123456, -2, -123500},
		{-123456, -3, -123000},
		{-123456, -4, -120000},
		{-123456, -5, -100000},
	}

	for _, tc := range tests {
		result := RoundTo(tc.val, tc.decimals)
		if !Equalish(result, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestRoundToNearest(t *testing.T) {
	type test struct {
		val      float64
		mult     float64
		expected float64
	}

	tests := []test{
		// positive
		{123457, 2, 123458},
		{123457, 3, 123456},
		{123457, 4, 123456},
		{123457, 5, 123455},
		{123457, 7, 123459},
		{123457, 10, 123460},

		// negative
		{-123457, 2, -123458},
		{-123457, 3, -123456},
		{-123457, 4, -123456},
		{-123457, 5, -123455},
		{-123457, 7, -123459},
		{-123457, 10, -123460},

		// fraction
		{123.456, 0.2, 123.4},
		{123.456, 0.5, 123.5},
		{123.456, 0.05, 123.45},
	}

	for _, tc := range tests {
		result := RoundToNearest(tc.val, tc.mult)
		if !Equalish(result, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestSinRange(t *testing.T) {
	type test struct {
		angle    float64
		min      float64
		max      float64
		expected float64
	}

	tests := []test{
		{0, 10, 20, 15},
		{HalfPi, 10, 20, 20},
		{math.Pi, 10, 20, 15},
		{math.Pi + HalfPi, 10, 20, 10},
		{Tau, 10, 20, 15},
		{2 * Tau, 10, 20, 15},
		{-Tau, 10, 20, 15},
		{-HalfPi, 10, 20, 10},
		{-math.Pi, 10, 20, 15},
	}

	for _, tc := range tests {
		result := SinRange(tc.angle, tc.min, tc.max)
		if !Equalish(result, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestCosRange(t *testing.T) {
	type test struct {
		angle    float64
		min      float64
		max      float64
		expected float64
	}

	tests := []test{
		{0, 10, 20, 20},
		{HalfPi, 10, 20, 15},
		{math.Pi, 10, 20, 10},
		{math.Pi + HalfPi, 10, 20, 15},
		{Tau, 10, 20, 20},
		{2 * Tau, 10, 20, 20},
		{-Tau, 10, 20, 20},
		{-HalfPi, 10, 20, 15},
		{-math.Pi, 10, 20, 10},
	}

	for _, tc := range tests {
		result := CosRange(tc.angle, tc.min, tc.max)
		if !Equalish(result, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestFract(t *testing.T) {
	type test struct {
		val      float64
		expected float64
	}

	tests := []test{
		{123.456, 0.456},
		{123.4, 0.4},
		{123, 0},
		{-123.456, -0.456},
		{-123.4, -0.4},
		{0, 0},
		{-123, 0},
	}

	for _, tc := range tests {
		result := Fract(tc.val)
		if !Equalish(result, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestLerpSin(t *testing.T) {
	type test struct {
		t        float64
		min      float64
		max      float64
		expected float64
	}

	tests := []test{
		{0, 10, 20, 10},
		{0.25, 10, 20, 15},
		{0.5, 10, 20, 20},
		{0.75, 10, 20, 15},
		{1, 10, 20, 10},
		{-0.25, 10, 20, 15},
		{-0.5, 10, 20, 20},
		{1.25, 10, 20, 15},
		{1.5, 10, 20, 20},
		{1.75, 10, 20, 15},
	}

	for _, tc := range tests {
		result := LoopSin(tc.t, tc.min, tc.max)
		if !Equalish(result, tc.expected, 0.00001) {
			t.Errorf("Expected %f, got %f\n", tc.expected, result)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b, exp int
	}{
		{12, 9, 3},
		{28, 7, 7},
		{33, 7, 1},
		{33, 0, 0},
		{32, -8, 8},
		{-32, -8, 8},
		{-32, 8, 8},
		{117, 72, 9},
	}

	for _, test := range tests {
		gcd := GCD(test.a, test.b)
		if gcd != test.exp {
			t.Errorf("Expected %d, got %d", test.exp, gcd)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b, exp int
	}{
		{12, 9, 36},
		{28, 7, 28},
		{33, 7, 231},
		{32, -8, 32},
		{-32, -8, 32},
		{-32, 8, 32},
		{117, 72, 936},
	}

	for i, test := range tests {
		lcm := LCM(test.a, test.b)
		if lcm != test.exp {
			t.Errorf("%d. Expected %d, got %d", i, test.exp, lcm)
		}
	}
}

func TestSimplify(t *testing.T) {
	tests := []struct {
		a, b, expA, expB int
	}{
		{12, 9, 4, 3},
		{28, 7, 4, 1},
		{33, 7, 33, 7},
		{32, -8, 4, -1},
		{-32, -8, -4, -1},
		{-32, 8, -4, 1},
		{117, 72, 13, 8},
	}

	for i, test := range tests {
		a, b := Simplify(test.a, test.b)
		if a != test.expA {
			t.Errorf("%d. Expected %d, got %d", i, test.expA, a)
		}
		if b != test.expB {
			t.Errorf("%d. Expected %d, got %d", i, test.expB, b)
		}
	}
}

func TestAbs(t *testing.T) {
	x := Abs(5)
	exp := 5
	if x != exp {
		t.Errorf("Expected %d, got %d", exp, x)
	}

	x = Abs(-5)
	exp = 5
	if x != exp {
		t.Errorf("Expected %d, got %d", exp, x)
	}

	y := Abs(5.0)
	expf := 5.0
	if y != expf {
		t.Errorf("Expected %f, got %f", expf, y)
	}

	y = Abs(-5.0)
	expf = 5.0
	if y != expf {
		t.Errorf("Expected %f, got %f", expf, y)
	}
}

func TestMinMax(t *testing.T) {
	x := Min(5, 4)
	exp := 4
	if x != exp {
		t.Errorf("Expected %d, got %d", exp, x)
	}

	x = Max(5, 4)
	exp = 5
	if x != exp {
		t.Errorf("Expected %d, got %d", exp, x)
	}

	y := Min(5.0, 4.0)
	expf := 4.0
	if y != expf {
		t.Errorf("Expected %f, got %f", expf, y)
	}

	y = Max(5.0, 4.0)
	expf = 5.0
	if y != expf {
		t.Errorf("Expected %f, got %f", expf, y)
	}
}

func TestWrapTau(t *testing.T) {
	tests := []struct {
		x, exp float64
	}{
		{1.23, 1.23},
		{7.89, 7.89 - Tau},
		{17.89, 17.89 - 2*Tau},
		{19.89, 19.89 - 3*Tau},
		{-1.23, Tau - 1.23},
		{-7.89, 2*Tau - 7.89},
		{-17.89, 3*Tau - 17.89},
		{0, 0},
		{Tau, 0},
		{-Tau, 0},
		{2 * Tau, 0},
		{-2 * Tau, 0},
		{20 * Tau, 0},
		{-20 * Tau, 0},
		{math.Pi, math.Pi},
		{2 * math.Pi, 0},
		{-math.Pi, math.Pi},
		{-2 * math.Pi, 0},
		{3 * math.Pi, math.Pi},
		{-3 * math.Pi, math.Pi},
	}

	for i, test := range tests {
		x := WrapTau(test.x)
		if !Equalish(x, test.exp, 0.0001) {
			t.Errorf("%d. Expected %f, got %f", i, test.exp, x)
		}
	}
}

func TestWrapPi(t *testing.T) {
	tests := []struct {
		x, exp float64
	}{
		{1.23, 1.23},
		{7.89, 7.89 - math.Pi*2},
		{17.89, 17.89 - 6*math.Pi},
		{19.89, 19.89 - 6*math.Pi},
		{-1.23, -1.23},
		{-7.89, -7.89 + 2*math.Pi},
		{-17.89, 3*Tau - 17.89},
		{0, 0},
		{Tau, 0},
		{-Tau, 0},
		{2 * Tau, 0},
		{-2 * Tau, 0},
		{20 * Tau, 0},
		{-20 * Tau, 0},
		{math.Pi, -math.Pi},
		{2 * math.Pi, 0},
		{-math.Pi, -math.Pi},
		{-2 * math.Pi, 0},
		{3 * math.Pi, -math.Pi},
		{-3 * math.Pi, -math.Pi},
	}

	for i, test := range tests {
		x := WrapPi(test.x)
		if !Equalish(x, test.exp, 0.0001) {
			t.Errorf("%d. Expected %f, got %f", i, test.exp, x)
		}
	}
}

func TestDigRoot(t *testing.T) {
	tests := []struct {
		x, exp int
	}{
		{1, 1},
		{0, 0},
		{11, 2},
		{44, 8},
		{55, 1},
		{59, 5},
		{123, 6},
		{923, 5},
		{999888777, 9},
		{999888778, 1},
	}

	for i, test := range tests {
		x := DigRoot(test.x)
		if x != test.exp {
			t.Errorf("%d. Expected %d, got %d", i, x, test.exp)
		}
	}
}
