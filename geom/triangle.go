// Package geom has geometry related structs and funcs.
package geom

import (
	"fmt"
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

// AngleFromTriangleSideLengths returns one angle of a triangle where the lengths of each side are known.
// It returns the angle between a and b.
func AngleFromTriangleSideLengths(a, b, c float64) float64 {
	return math.Acos((a*a + b*b - c*c) / (2 * a * b))
}

func (t *Triangle) String() string {
	return fmt.Sprintf("[triangle: \n {%0.3f, %0.3f},\n {%0.3f, %0.3f},\n {%0.3f, %0.3f}\n]",
		t.PointA.X, t.PointA.Y,
		t.PointB.X, t.PointB.Y,
		t.PointC.X, t.PointC.Y,
	)
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
func (t *Triangle) Points() PointList {
	return PointList{t.PointA, t.PointB, t.PointC}
}

// Edges returns the three segments that make up the triangle.
func (t *Triangle) Edges() SegmentList {
	edges := NewSegmentList()
	edges.Add(NewSegmentFromPoints(t.PointA, t.PointB))
	edges.Add(NewSegmentFromPoints(t.PointB, t.PointC))
	edges.Add(NewSegmentFromPoints(t.PointC, t.PointA))
	return edges
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

// Equals returns whether or not this triangle is equal to another triangle
func (t *Triangle) Equals(other *Triangle) bool {
	if t == other {
		return true
	}
	testTwo := func(pA, pB, pC, pD *Point) bool {
		return (pA.Equals(pC) && pB.Equals(pD)) || (pA.Equals(pD) && pB.Equals(pC))
	}

	if t.PointA.Equals(other.PointA) {
		return testTwo(t.PointB, t.PointC, other.PointB, other.PointC)
	}
	if t.PointA.Equals(other.PointB) {
		return testTwo(t.PointB, t.PointC, other.PointA, other.PointC)
	}
	if t.PointA.Equals(other.PointC) {
		return testTwo(t.PointB, t.PointC, other.PointA, other.PointB)
	}
	return false
}

//////////////////////////////
// Transform in place
//////////////////////////////

// Randomize randomizes the position of the three points of the triangle.
func (t *Triangle) Randomize(amount float64) {
	t.PointA.Randomize(amount, amount)
	t.PointB.Randomize(amount, amount)
	t.PointC.Randomize(amount, amount)
}

// Rotate rotates a triangle around the origin.
func (t *Triangle) Rotate(angle float64) {
	t.PointA.Rotate(angle)
	t.PointB.Rotate(angle)
	t.PointC.Rotate(angle)
}

// RotateLocal rotates a triangle around its own center (centroid)
func (t *Triangle) RotateLocal(angle float64) {
	center := t.Centroid()
	t.RotateFrom(center.X, center.Y, angle)
}

// RotateFrom rotates a triangle around the given x, y location.
func (t *Triangle) RotateFrom(x, y, angle float64) {
	t.PointA.RotateFrom(x, y, angle)
	t.PointB.RotateFrom(x, y, angle)
	t.PointC.RotateFrom(x, y, angle)
}

// Translate translates a triangle the given amount on each axis.
func (t *Triangle) Translate(x, y float64) {
	t.PointA.Translate(x, y)
	t.PointB.Translate(x, y)
	t.PointC.Translate(x, y)
}

// Scale scales a triangle the given amount on each axis.
func (t *Triangle) Scale(sx, sy float64) {
	t.PointA.Scale(sx, sy)
	t.PointB.Scale(sx, sy)
	t.PointC.Scale(sx, sy)
}

// ScaleLocal scales a triangle the given amount on each axis.
func (t *Triangle) ScaleLocal(sx, sy float64) {
	center := t.Centroid()
	t.ScaleFrom(center.X, center.Y, sx, sy)
}

// ScaleFrom scales a triangle the given amount on each axis, from the given x, y location.
func (t *Triangle) ScaleFrom(x, y, sx, sy float64) {
	t.PointA.ScaleFrom(x, y, sx, sy)
	t.PointB.ScaleFrom(x, y, sx, sy)
	t.PointC.ScaleFrom(x, y, sx, sy)
}

// UniScale scales a triangle the given amount on each axis.
func (t *Triangle) UniScale(scale float64) {
	t.Scale(scale, scale)
}

// UniScaleFrom scales a triangle the given amount on each axis, from the given x, y location.
func (t *Triangle) UniScaleFrom(x, y, scale float64) {
	t.ScaleFrom(x, y, scale, scale)
}

//////////////////////////////
// Return transformed copy
//////////////////////////////

// Scaled returns another triangle that is scaled from this one.
func (t *Triangle) Scaled(sx, sy float64) *Triangle {
	return NewTriangle(
		t.PointA.X*sx, t.PointA.Y*sy,
		t.PointB.X*sx, t.PointB.Y*sy,
		t.PointC.X*sx, t.PointC.Y*sy,
	)
}

// ScaledFrom returns a new triangle scaled by the given amount.
func (t *Triangle) ScaledFrom(x, y, sx, sy float64) *Triangle {
	t2 := NewTriangleFromPoints(t.PointA, t.PointB, t.PointC)
	t2.ScaleFrom(x, y, sx, sy)
	return t2
}

// UniScaled returns a new triangle scaled by the given amount - the same on both axes.
func (t *Triangle) UniScaled(scale float64) *Triangle {
	return t.Scaled(scale, scale)
}

// UniScaledFrom returns a new segment scaled from the given point.
func (t *Triangle) UniScaledFrom(x, y, scale float64) *Triangle {
	return t.ScaledFrom(x, y, scale, scale)
}
