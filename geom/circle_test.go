// Package geom has geometry related structs and funcs.
package geom

import (
	"testing"

	"github.com/bit101/bitlib/blmath"
)

func TestInversionPoint(t *testing.T) {

	c := NewCircle(0, 0, 100)
	p := NewPoint(131.5, 51)

	i := c.InvertPoint(p)

	x, y := i.X, i.Y
	ex := 66.102824
	ey := 25.636837

	if !blmath.Equalish(x, ex, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ex, x)
	}
	if !blmath.Equalish(y, ey, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ey, y)
	}

	p = NewPoint(-54.4, -36.5)

	i = c.InvertPoint(p)

	x, y = i.X, i.Y
	ex = -126.758955
	ey = -85.049667

	if !blmath.Equalish(x, ex, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ex, x)
	}
	if !blmath.Equalish(y, ey, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ey, y)
	}
}

func TestInversionXY(t *testing.T) {

	c := NewCircle(0, 0, 100)
	x, y := 131.5, 51.0

	x, y = c.InvertXY(x, y)

	ex := 66.102824
	ey := 25.636837

	if !blmath.Equalish(x, ex, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ex, x)
	}
	if !blmath.Equalish(y, ey, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ey, y)
	}

	x, y = -54.4, -36.5

	x, y = c.InvertXY(x, y)

	ex = -126.758955
	ey = -85.049667

	if !blmath.Equalish(x, ex, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ex, x)
	}
	if !blmath.Equalish(y, ey, 0.00001) {
		t.Errorf("Expected %f, got %f\n", ey, y)
	}
}
