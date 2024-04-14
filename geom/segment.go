// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
)

// Segment represents a line segment.
type Segment struct {
	X0, Y0, X1, Y1 float64
}

// NewSegment creates a new line segment.
func NewSegment(x0, y0, x1, y1 float64) *Segment {
	return &Segment{
		X0: x0,
		Y0: y0,
		X1: x1,
		Y1: y1,
	}
}

// NewSegmentFromPoints creates a new line segment.
func NewSegmentFromPoints(p0, p1 *Point) *Segment {
	return NewSegment(p0.X, p0.Y, p1.X, p1.Y)
}

// HitSegment returns the point of intersection between this and another segment.
func (s *Segment) HitSegment(z *Segment) (float64, float64, bool) {
	return SegmentOnSegment(s.X0, s.Y0, s.X1, s.Y1, z.X0, z.Y0, z.X1, z.Y1)
}

// HitLine returns the point of intersection between this and another line.
func (s *Segment) HitLine(l *Line) (float64, float64, bool) {
	return SegmentOnLine(s.X0, s.Y0, s.X1, s.Y1, l.X0, l.Y0, l.X1, l.Y1)
}

// HitRect returns whether this segment intersects a rectangle.
func (s *Segment) HitRect(r *Rect) bool {
	return SegmentOnRect(s.X0, s.Y0, s.X1, s.Y1, r.X, r.Y, r.W, r.H)
}

// Length is the length of this segment.
func (s *Segment) Length() float64 {
	return math.Hypot(s.X1-s.X0, s.Y1-s.Y0)
}

// ClosestPoint returns the point on this segment closest to another point.
func (s *Segment) ClosestPoint(p *Point) *Point {
	v := VectorBetween(s.X0, s.Y0, p.X, p.Y)
	d := VectorBetween(s.X0, s.Y0, s.X1, s.Y1).Normalized()
	vs := v.Project(d)
	if vs < 0 {
		return NewPoint(s.X0, s.Y0)
	}
	if vs > s.Length() {
		return NewPoint(s.X1, s.Y1)
	}
	t := vs / s.Length()
	return NewPoint(blmath.Lerp(t, s.X0, s.X1), blmath.Lerp(t, s.Y0, s.Y1))
}

// DistanceTo returns the distance from this segment to another point.
func (s *Segment) DistanceTo(p *Point) float64 {
	return s.ClosestPoint(p).Distance(p)
}

// Equals returns whether or not this segment is roughly equal to another segment.
func (s *Segment) Equals(other *Segment) bool {
	if s == other {
		return true
	}
	d := 0.000001
	if blmath.Equalish(s.X0, other.X0, d) &&
		blmath.Equalish(s.Y0, other.Y0, d) &&
		blmath.Equalish(s.X1, other.X1, d) &&
		blmath.Equalish(s.Y1, other.Y1, d) {
		return true
	}

	if blmath.Equalish(s.X0, other.X1, d) &&
		blmath.Equalish(s.Y0, other.Y1, d) &&
		blmath.Equalish(s.X1, other.X0, d) &&
		blmath.Equalish(s.Y1, other.Y0, d) {
		return true
	}
	return false
}

// Randomize randomizes the endpoints of the segment by the given amount.
func (s *Segment) Randomize(amount float64) {
	s.X0 += random.FloatRange(-amount, amount)
	s.Y0 += random.FloatRange(-amount, amount)
	s.X1 += random.FloatRange(-amount, amount)
	s.Y1 += random.FloatRange(-amount, amount)
}

// Rotate rotates a segment around the origin.
func (s *Segment) Rotate(angle float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x0 := cos*s.X0 + sin*s.Y0
	y0 := cos*s.Y0 - sin*s.X0
	x1 := cos*s.X1 + sin*s.Y1
	y1 := cos*s.Y1 - sin*s.X1
	s.X0 = x0
	s.Y0 = y0
	s.X1 = x1
	s.Y1 = y1
}

// RotateLocal rotates a segment around its own center.
func (s *Segment) RotateLocal(angle float64) {
	cx := (s.X0 + s.X1) / 2
	cy := (s.Y0 + s.Y1) / 2
	s.RotateFrom(cx, cy, angle)
}

// RotateFrom rotates a segment around the given x, y location.
func (s *Segment) RotateFrom(x, y, angle float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x0 := cos*(s.X0-x) + sin*(s.Y0-y)
	y0 := cos*(s.Y0-y) - sin*(s.X0-x)
	x1 := cos*(s.X1-x) + sin*(s.Y1-y)
	y1 := cos*(s.Y1-y) - sin*(s.X1-x)
	s.X0 = x0 + x
	s.Y0 = y0 + y
	s.X1 = x1 + x
	s.Y1 = y1 + y
}

// Translate translates a segment the given amount on each axis.
func (s *Segment) Translate(x, y float64) {
	s.X0 += x
	s.Y0 += y
	s.X1 += x
	s.Y1 += y
}

// Scale scales a segment the given amount on each axis.
func (s *Segment) Scale(sx, sy float64) {
	s.X0 *= sx
	s.Y0 *= sy
	s.X1 *= sx
	s.Y1 *= sy
}

// ScaleLocal scales a segment the given amount on each axis.
func (s *Segment) ScaleLocal(sx, sy float64) {
	cx := (s.X0 + s.X1) / 2
	cy := (s.Y0 + s.Y1) / 2
	s.ScaleFrom(cx, cy, sx, sy)
}

// ScaleFrom scales a segment the given amount on each axis, from the given x, y location.
func (s *Segment) ScaleFrom(x, y, sx, sy float64) {
	s.X0 = (s.X0-x)*sx + x
	s.Y0 = (s.Y0-y)*sy + y
	s.X1 = (s.X1-x)*sx + x
	s.Y1 = (s.Y1-y)*sy + y
}
