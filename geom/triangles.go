// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
)

// Triangle represents a triangle.
type Triangle struct {
	PointA, PointB, PointC *Point
}

// NewTriangleFromPoints creates a new triangle from three points.
func NewTriangleFromPoints(pA, pB, pC *Point) *Triangle {
	return &Triangle{pA, pB, pC}
}

// NewTriangleXY creates a new triangle from three x, y pairs
func NewTriangleXY(x0, y0, x1, y1, x2, y2 float64) *Triangle {
	return &Triangle{
		NewPoint(x0, y0),
		NewPoint(x1, y1),
		NewPoint(x2, y2),
	}
}

// EquilateralTriangleFromCenterAndPoint creates a new equilateral triangle from a centroid point and another point.
func EquilateralTriangleFromCenterAndPoint(c, p0 *Point) *Triangle {
	angle := math.Atan2(p0.Y-c.Y, p0.X-c.X)
	dist := math.Hypot(p0.Y-c.Y, p0.X-c.X)
	p1 := NewPoint(
		c.X+math.Cos(angle+blmath.Tau/3)*dist,
		c.Y+math.Sin(angle+blmath.Tau/3)*dist,
	)
	p2 := NewPoint(
		c.X+math.Cos(angle-blmath.Tau/3)*dist,
		c.Y+math.Sin(angle-blmath.Tau/3)*dist,
	)
	return NewTriangleFromPoints(p0, p1, p2)
}

// EquilateralTriangleFromTwoPoints creates a new equilateral triangle from a two points.
func EquilateralTriangleFromTwoPoints(p0, p1 *Point, clockwise bool) *Triangle {
	angle := math.Atan2(p1.Y-p0.Y, p1.X-p0.X)
	dist := math.Hypot(p1.Y-p0.Y, p1.X-p0.X)
	if clockwise {
		angle += math.Pi / 3
	} else {
		angle -= math.Pi / 3
	}
	p2 := NewPoint(
		p0.X+math.Cos(angle)*dist,
		p0.Y+math.Sin(angle)*dist,
	)
	return NewTriangleFromPoints(p0, p1, p2)
}

// Centroid returns the centroid of a triangle.
func (t *Triangle) Centroid() *Point {
	midA := LerpPoint(0.5, t.PointB, t.PointC)
	midB := LerpPoint(0.5, t.PointA, t.PointC)

	segA := NewSegment(t.PointA.X, t.PointA.Y, midA.X, midA.Y)
	segB := NewSegment(t.PointB.X, t.PointB.Y, midB.X, midB.Y)
	x, y, _ := segA.HitSegment(segB)
	return NewPoint(x, y)
}

// Area returns the area of the triangle
func (t *Triangle) Area() float64 {
	seg := NewSegment(t.PointA.X, t.PointA.Y, t.PointB.X, t.PointB.Y)
	h := seg.DistanceTo(t.PointC)
	b := seg.Length()
	return h * b / 2
}
