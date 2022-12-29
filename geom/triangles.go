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

// NewTriangle creates a new triangle from three x, y pairs
func NewTriangle(x0, y0, x1, y1, x2, y2 float64) *Triangle {
	return NewTriangleFromPoints(NewPoint(x0, y0), NewPoint(x1, y1), NewPoint(x2, y2))
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

// CircumCenter returns a point representing the circumcenter of the triangle
func (t *Triangle) CircumCenter() *Point {
	midA := LerpPoint(0.5, t.PointA, t.PointB)
	lineA := NewLine(t.PointA.X, t.PointA.Y, t.PointB.X, t.PointB.Y).Perpendicular(midA)

	midB := LerpPoint(0.5, t.PointA, t.PointC)
	lineB := NewLine(t.PointA.X, t.PointA.Y, t.PointC.X, t.PointC.Y).Perpendicular(midB)

	x, y, b := lineA.HitLine(lineB)
	if b {
		return NewPoint(x, y)
	}
	return nil
}

// CircumCircle returns a circle object representing the center and radius of the circumcircle.
func (t *Triangle) CircumCircle() *Circle {
	p := t.CircumCenter()
	if p == nil {
		return nil
	}
	radius := p.Distance(t.PointA)
	return NewCircle(p.X, p.Y, radius)
}

// InCenter returns the point representing the incenter of the triangle.
func (t *Triangle) InCenter() *Point {
	b := bisect(t.PointA, t.PointB, t.PointC)
	a := bisect(t.PointB, t.PointA, t.PointC)
	lineB := NewLine(b.X, b.Y, t.PointB.X, t.PointB.Y)
	lineA := NewLine(a.X, a.Y, t.PointA.X, t.PointA.Y)
	x, y, hit := lineA.HitLine(lineB)
	if hit {
		return NewPoint(x, y)
	}
	return nil
}

// InCircle returns a circle containing the center and radius of the incircle.
func (t *Triangle) InCircle() *Circle {
	center := t.InCenter()
	line := NewLine(t.PointA.X, t.PointA.Y, t.PointB.X, t.PointB.Y)
	radius := line.DistanceTo(center)
	return NewCircle(center.X, center.Y, radius)
}

func bisect(pA, pB, pC *Point) *Point {
	angleA := math.Atan2(pA.Y-pB.Y, pA.X-pB.X)
	angleC := math.Atan2(pC.Y-pB.Y, pC.X-pB.X)
	angleBi := angleA + (angleC-angleA)/2
	return NewPoint(pB.X+math.Cos(angleBi)*100, pB.Y+math.Sin(angleBi)*100)
}

// OrthoCenter returns a point representing the orthocenter of the triangle.
func (t *Triangle) OrthoCenter() *Point {
	bc := NewLine(t.PointB.X, t.PointB.Y, t.PointC.X, t.PointC.Y)
	pA := bc.ClosestPoint(t.PointA)

	ac := NewLine(t.PointA.X, t.PointA.Y, t.PointC.X, t.PointC.Y)
	pB := ac.ClosestPoint(t.PointB)

	line0 := NewLine(t.PointA.X, t.PointA.Y, pA.X, pA.Y)
	line1 := NewLine(t.PointB.X, t.PointB.Y, pB.X, pB.Y)

	x, y, _ := line0.HitLine(line1)
	return NewPoint(x, y)
}

// Area returns the area of the triangle
func (t *Triangle) Area() float64 {
	seg := NewSegment(t.PointA.X, t.PointA.Y, t.PointB.X, t.PointB.Y)
	h := seg.DistanceTo(t.PointC)
	b := seg.Length()
	return h * b / 2
}

// Points returns a list of points that make up this triangle
func (t *Triangle) Points() []*Point {
	return []*Point{t.PointA, t.PointB, t.PointC}
}

// Contains returns whether or not the given point is contained by the triangle.
func (t *Triangle) Contains(p *Point) bool {
	d1 := Clockwise(p, t.PointA, t.PointB)
	d2 := Clockwise(p, t.PointB, t.PointC)
	d3 := Clockwise(p, t.PointC, t.PointA)

	hasCCW := !(d1 && d2 && d3)
	hasCW := d1 || d2 || d3

	return !(hasCCW && hasCW)
}
