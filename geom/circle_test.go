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

func TestCircleThroughPointsWithHeight(t *testing.T) {
	x0, y0 := -100.0, 0.0
	x1, y1 := 100.0, 0.0
	height := 100.0
	circle, err := CircleThroughPointsWithArcHeight(x0, y0, x1, y1, height)
	if err != nil {
		t.Error("Should not produce error")
	}
	if circle.X != 0 || circle.Y != 0 || circle.Radius != 100 {
		t.Errorf("Expected circle: x: %f, y: %f, r: %f. Got x: %f, y:%f, r: %f", 0.0, 0.0, 100.0, circle.X, circle.Y, circle.Radius)
	}

	x0, y0 = -50.0, 0.0
	x1, y1 = 50.0, 0.0
	height = 100.0
	circle, err = CircleThroughPointsWithArcHeight(x0, y0, x1, y1, height)
	if err != nil {
		t.Error("Should not produce error")
	}
	if !blmath.Equalish(circle.X, 0, 0.0001) {
		t.Errorf("Expected circle.X: %f, got: %f", 0.0, circle.X)
	}
	if !blmath.Equalish(circle.Y, -37.5, 0.0001) {
		t.Errorf("Expected circle.Y: %f, got: %f", 0.0, circle.X)
	}
	if !blmath.Equalish(circle.Radius, 62.5, 0.0001) {
		t.Errorf("Expected circle.Radius: %f, got: %f", 62.5, circle.Radius)
	}

	x0, y0 = -50.0, 0.0
	x1, y1 = 50.0, 0.0
	height = 0.0
	circle, err = CircleThroughPointsWithArcHeight(x0, y0, x1, y1, height)
	if err == nil {
		t.Error("Zero height should create an error")
	}
}

func TestCircleContains(t *testing.T) {
	// contains
	c := NewCircle(200, 200, 100)
	p := NewPoint(130, 130)
	b := c.Contains(p)
	if !b {
		t.Errorf("Expected %t, got %t\n", false, b)
	}

	// does not contain
	p = NewPoint(120, 120)
	b = c.Contains(p)
	if b {
		t.Errorf("Expected %t, got %t\n", true, b)
	}

	// does not contain
	p = NewPoint(130, 120)
	b = c.Contains(p)
	if b {
		t.Errorf("Expected %t, got %t\n", true, b)
	}

	// does not contain
	p = NewPoint(120, 130)
	b = c.Contains(p)
	if b {
		t.Errorf("Expected %t, got %t\n", true, b)
	}

}
