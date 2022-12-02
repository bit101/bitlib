// Package geom has geometry related structs and funcs.
package geom

import (
	"math"
	"testing"

	"github.com/bit101/bitlib/blmath"
)

func TestNewTri(t *testing.T) {
	tri0 := NewTriangleFromPoints(NewPoint(100, 100), NewPoint(200, 100), NewPoint(100, 300))
	if tri0.PointA.X != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointA.X)
	}
	if tri0.PointA.Y != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointA.Y)
	}

	if tri0.PointB.X != 200 {
		t.Errorf("Expected %f, got %f\n", 200.0, tri0.PointB.X)
	}
	if tri0.PointB.Y != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointB.Y)
	}

	if tri0.PointC.X != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointC.X)
	}
	if tri0.PointC.Y != 300 {
		t.Errorf("Expected %f, got %f\n", 300.0, tri0.PointC.Y)
	}

	tri1 := NewTriangleXY(100, 100, 200, 100, 100, 300)
	if tri1.PointA.X != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointA.X)
	}
	if tri1.PointA.Y != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointA.Y)
	}

	if tri1.PointB.X != 200 {
		t.Errorf("Expected %f, got %f\n", 200.0, tri0.PointB.X)
	}
	if tri1.PointB.Y != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointB.Y)
	}

	if tri1.PointC.X != 100 {
		t.Errorf("Expected %f, got %f\n", 100.0, tri0.PointC.X)
	}
	if tri1.PointC.Y != 300 {
		t.Errorf("Expected %f, got %f\n", 300.0, tri0.PointC.Y)
	}
}

func TestCentroid(t *testing.T) {
	tri := NewTriangleXY(100, 100, 200, 300, -50, -100)
	centroid := tri.Centroid()

	if !blmath.Equalish(centroid.X, 83.33333, 0.00001) {
		t.Errorf("Expected %f, got %f\n", 83.33333, centroid.X)
	}

	if !blmath.Equalish(centroid.Y, 100, 0.00001) {
		t.Errorf("Expected %f, got %f\n", 100.0, centroid.Y)
	}
}

func TestNewEqTri(t *testing.T) {
	c := NewPoint(0, 0)
	p0 := NewPoint(100, 0)
	tri := EquilateralTriangleFromCenterAndPoint(c, p0)
	p1 := tri.PointB
	p2 := tri.PointC

	expX := math.Cos(blmath.Tau/3) * 100
	expY := math.Sin(blmath.Tau/3) * 100
	if !blmath.Equalish(p1.X, expX, 0.00001) {
		t.Errorf("Expected %f, got %f\n", expX, p1.X)
	}
	if !blmath.Equalish(p1.Y, expY, 0.00001) {
		t.Errorf("Expected %f, got %f\n", expY, p1.Y)
	}

	expX = math.Cos(-blmath.Tau/3) * 100
	expY = math.Sin(-blmath.Tau/3) * 100
	if !blmath.Equalish(p2.X, expX, 0.00001) {
		t.Errorf("Expected %f, got %f\n", expX, p2.X)
	}
	if !blmath.Equalish(p2.Y, expY, 0.00001) {
		t.Errorf("Expected %f, got %f\n", expY, p2.Y)
	}
}

func TestNewEqTriPoints(t *testing.T) {
	p0 := NewPoint(0, 0)
	p1 := NewPoint(100, 0)
	tri := EquilateralTriangleFromTwoPoints(p0, p1, true)
	p2 := tri.PointC

	expX := math.Cos(math.Pi/3) * 100
	expY := math.Sin(math.Pi/3) * 100
	if !blmath.Equalish(p2.X, expX, 0.00001) {
		t.Errorf("Expected %f, got %f\n", expX, p2.X)
	}
	if !blmath.Equalish(p2.Y, expY, 0.00001) {
		t.Errorf("Expected %f, got %f\n", expY, p2.Y)
	}

	tri = EquilateralTriangleFromTwoPoints(p0, p1, false)
	p2 = tri.PointC

	if !blmath.Equalish(p2.X, expX, 0.00001) {
		t.Errorf("Expected %f, got %f\n", expX, p2.X)
	}
	if !blmath.Equalish(p2.Y, -expY, 0.00001) {
		t.Errorf("Expected %f, got %f\n", -expY, p2.Y)
	}
}

func TestArea(t *testing.T) {
	tri := NewTriangleXY(100, 100, 200, 100, 150, 0)
	area := tri.Area()

	if area != 5000 {
		t.Errorf("Expected %f, got %f\n", 5000.0, area)
	}

}
