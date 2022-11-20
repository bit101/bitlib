// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
)

// Line is the same as a segment, but we'll keep them separate for collision purposes
type Line Segment

// NewLine creates a new line.
func NewLine(x0, y0, x1, y1 float64) *Line {
	return &Line{
		X0: x0,
		Y0: y0,
		X1: x1,
		Y1: y1,
	}
}

// HitSegment reports the point of intersection of a line and a line segment.
func (l *Line) HitSegment(s *Segment) (float64, float64, bool) {
	return SegmentOnSegment(s.X0, s.Y0, s.X1, s.Y1, l.X0, l.Y0, l.X1, l.Y1)
}

// HitLine reports the point of intersection of two lines.
func (l *Line) HitLine(m *Line) (float64, float64, bool) {
	return SegmentOnLine(l.X0, l.Y0, l.X1, l.Y1, m.X0, m.Y0, m.X1, m.Y1)
}

// ClosestPoint reports the point on a line closest to another given point.
func (l *Line) ClosestPoint(p *Point) *Point {

	v := VectorBetween(l.X0, l.Y0, p.X, p.Y)
	d := VectorBetween(l.X0, l.Y0, l.X1, l.Y1).Normalized()
	vs := v.Project(d)
	t := vs / math.Hypot(l.X1-l.X0, l.Y1-l.Y0)
	return NewPoint(blmath.Lerp(t, l.X0, l.X1), blmath.Lerp(t, l.Y0, l.Y1))
}

// DistanceTo reports the distance from a given point to this line.
func (l *Line) DistanceTo(p *Point) float64 {
	return l.ClosestPoint(p).Distance(p)
}
