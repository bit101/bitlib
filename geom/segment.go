// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
)

// Segment represents a line segment.
type Segment struct {
	PointA, PointB *Point
}

// NewSegment creates a new line segment.
func NewSegment(x0, y0, x1, y1 float64) *Segment {
	return &Segment{
		NewPoint(x0, y0),
		NewPoint(x1, y1),
	}
}

// NewSegmentFromPoints creates a new line segment.
func NewSegmentFromPoints(p0, p1 *Point) *Segment {
	return NewSegment(p0.X, p0.Y, p1.X, p1.Y)
}

// Points returns the two endpoints of this segment as Points.
func (s *Segment) Points() (*Point, *Point) {
	return s.PointA, s.PointB
}

// Parallel returns a Line that is parallel to this line, a certain distance away.
func (s *Segment) Parallel(dist float64) *Line {
	return NewLineFromSegment(s).Parallel(dist)
}

// HitSegment returns the point of intersection between this and another segment.
func (s *Segment) HitSegment(z *Segment) (float64, float64, bool) {
	return SegmentOnSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y, z.PointA.X, z.PointA.Y, z.PointB.X, z.PointA.Y)
}

// HitLine returns the point of intersection between this and another line.
func (s *Segment) HitLine(l *Line) (float64, float64, bool) {
	return SegmentOnLine(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y, l.PointA.X, l.PointA.Y, l.PointB.X, l.PointB.Y)
}

// HitRect returns whether this segment intersects a rectangle.
func (s *Segment) HitRect(r *Rect) bool {
	return SegmentOnRect(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y, r.X, r.Y, r.W, r.H)
}

// Length is the length of this segment.
func (s *Segment) Length() float64 {
	return math.Hypot(s.PointB.X-s.PointA.X, s.PointB.Y-s.PointA.Y)
}

// ClosestPoint returns the point on this segment closest to another point.
func (s *Segment) ClosestPoint(p *Point) *Point {
	v := VectorBetween(s.PointA.X, s.PointA.Y, p.X, p.Y)
	d := VectorBetween(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y).Normalized()
	vs := v.Project(d)
	if vs < 0 {
		return s.PointA.Clone()
	}
	if vs > s.Length() {
		return s.PointB.Clone()
	}
	t := vs / s.Length()
	return NewPoint(blmath.Lerp(t, s.PointA.X, s.PointB.X), blmath.Lerp(t, s.PointA.Y, s.PointB.Y))
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
	if blmath.Equalish(s.PointA.X, other.PointA.X, d) &&
		blmath.Equalish(s.PointA.Y, other.PointA.Y, d) &&
		blmath.Equalish(s.PointB.X, other.PointB.X, d) &&
		blmath.Equalish(s.PointB.Y, other.PointB.Y, d) {
		return true
	}

	if blmath.Equalish(s.PointA.X, other.PointB.X, d) &&
		blmath.Equalish(s.PointA.Y, other.PointB.Y, d) &&
		blmath.Equalish(s.PointB.X, other.PointA.X, d) &&
		blmath.Equalish(s.PointB.Y, other.PointA.Y, d) {
		return true
	}
	return false
}

//////////////////////////////
// Transform in place
//////////////////////////////

// Randomize randomizes the endpoints of the segment by the given amount.
func (s *Segment) Randomize(amount float64) {
	s.PointA.X += random.FloatRange(-amount, amount)
	s.PointA.Y += random.FloatRange(-amount, amount)
	s.PointB.X += random.FloatRange(-amount, amount)
	s.PointB.Y += random.FloatRange(-amount, amount)
}

// Scale scales a segment the given amount on each axis.
func (s *Segment) Scale(sx, sy float64) {
	s.PointA.X *= sx
	s.PointA.Y *= sy
	s.PointB.X *= sx
	s.PointB.Y *= sy
}

// ScaleFrom scales a segment the given amount on each axis, from the given x, y location.
func (s *Segment) ScaleFrom(x, y, sx, sy float64) {
	s.PointA.X = (s.PointA.X-x)*sx + x
	s.PointA.Y = (s.PointA.Y-y)*sy + y
	s.PointB.X = (s.PointB.X-x)*sx + x
	s.PointB.Y = (s.PointB.Y-y)*sy + y
}

// ScaleLocal scales a segment the given amount on each axis from its own center
func (s *Segment) ScaleLocal(sx, sy float64) {
	cx := (s.PointA.X + s.PointB.X) / 2
	cy := (s.PointA.Y + s.PointB.Y) / 2
	s.ScaleFrom(cx, cy, sx, sy)
}

// UniScale scales a segment the given amount on both axes equally.
func (s *Segment) UniScale(scale float64) {
	s.Scale(scale, scale)
}

// UniScaleFrom scales a segment using the x, y location as a center
func (s *Segment) UniScaleFrom(x, y, scale float64) {
	s.ScaleFrom(x, y, scale, scale)
}

// UniScaleLocal scales a segment the given amount on each axis.
func (s *Segment) UniScaleLocal(scale float64) {
	s.ScaleLocal(scale, scale)
}

// Translate translates a segment the given amount on each axis.
func (s *Segment) Translate(x, y float64) {
	s.PointA.X += x
	s.PointA.Y += y
	s.PointB.X += x
	s.PointB.Y += y
}

// Rotate rotates a segment around the origin.
func (s *Segment) Rotate(angle float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x0 := cos*s.PointA.X + sin*s.PointA.Y
	y0 := cos*s.PointA.Y - sin*s.PointA.X
	x1 := cos*s.PointB.X + sin*s.PointB.Y
	y1 := cos*s.PointB.Y - sin*s.PointB.X
	s.PointA.X = x0
	s.PointA.Y = y0
	s.PointB.X = x1
	s.PointB.Y = y1
}

// RotateFrom rotates a segment around the given x, y location.
func (s *Segment) RotateFrom(x, y, angle float64) {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x0 := cos*(s.PointA.X-x) + sin*(s.PointA.Y-y)
	y0 := cos*(s.PointA.Y-y) - sin*(s.PointA.X-x)
	x1 := cos*(s.PointB.X-x) + sin*(s.PointB.Y-y)
	y1 := cos*(s.PointB.Y-y) - sin*(s.PointB.X-x)
	s.PointA.X = x0 + x
	s.PointA.Y = y0 + y
	s.PointB.X = x1 + x
	s.PointB.Y = y1 + y
}

// RotateLocal rotates a segment around its own center.
func (s *Segment) RotateLocal(angle float64) {
	cx := (s.PointA.X + s.PointB.X) / 2
	cy := (s.PointA.Y + s.PointB.Y) / 2
	s.RotateFrom(cx, cy, angle)
}

//////////////////////////////
// Return new with transform
//////////////////////////////

// Randomized returns a randomized segment from this segment.
func (s *Segment) Randomized(amount float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.Randomize(amount)
	return s2
}

// Scaled returns a new segment scaled by the given amount.
func (s *Segment) Scaled(sx, sy float64) *Segment {
	return NewSegment(s.PointA.X*sx, s.PointA.Y*sy, s.PointB.X*sx, s.PointB.Y*sy)
}

// ScaledFrom returns a new segment scaled by the given amount.
func (s *Segment) ScaledFrom(x, y, sx, sy float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.ScaleFrom(x, y, sx, sy)
	return s2
}

// ScaledLocal returns a new segment scaled by the given amount.
func (s *Segment) ScaledLocal(sx, sy float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.ScaleLocal(sx, sy)
	return s2
}

// UniScaled returns a new segment scaled by the given amount - the same on both axes.
func (s *Segment) UniScaled(scale float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.UniScale(scale)
	return s2
}

// UniScaledFrom returns a new segment scaled from the given point.
func (s *Segment) UniScaledFrom(x, y, scale float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.ScaledFrom(x, y, scale, scale)
	return s2
}

// UniScaledLocal returns a new segment scaled by the given amount - the same on both axes.
func (s *Segment) UniScaledLocal(scale float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.UniScaleLocal(scale)
	return s2
}

// Translated returns a new segment translated from this.
func (s *Segment) Translated(x, y float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.Translate(x, y)
	return s2
}

// Rotated returns a new segment rotated from this.
func (s *Segment) Rotated(angle float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.Rotate(angle)
	return s2
}

// RotatedFrom returns a new segment rotated from this.
func (s *Segment) RotatedFrom(x, y, angle float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.RotateFrom(x, y, angle)
	return s2
}

// RotatedLocal returns a new segment rotated from this.
func (s *Segment) RotatedLocal(angle float64) *Segment {
	s2 := NewSegment(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y)
	s2.RotateLocal(angle)
	return s2
}
