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

// NewLineFromPoints creates a new line from two Points.
func NewLineFromPoints(p0, p1 *Point) *Line {
	return NewLine(p0.X, p0.Y, p1.X, p1.Y)
}

// NewLineFromSegment creates a new line from an existing Segment.
func NewLineFromSegment(seg *Segment) *Line {
	return NewLine(seg.X0, seg.Y0, seg.X1, seg.Y1)
}

// HitSegment reports the point of intersection of a line and a line segment.
func (l *Line) HitSegment(s *Segment) (float64, float64, bool) {
	return SegmentOnLine(s.X0, s.Y0, s.X1, s.Y1, l.X0, l.Y0, l.X1, l.Y1)
}

// HitLine reports the point of intersection of two lines.
func (l *Line) HitLine(m *Line) (float64, float64, bool) {
	return LineOnLine(l.X0, l.Y0, l.X1, l.Y1, m.X0, m.Y0, m.X1, m.Y1)
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

// Angle returns the angle of rthis line (from the first point to the second)
func (l *Line) Angle() float64 {
	return math.Atan2(l.Y1-l.Y0, l.X1-l.X0) - math.Pi/2
}

// Perpendicular returns a line that is perpendicular to the line and crosses through the given point.
func (l *Line) Perpendicular(p *Point) *Line {
	dx := l.X1 - l.X0
	dy := l.Y1 - l.Y0
	return NewLine(p.X, p.Y, p.X-dy, p.Y+dx)
}

// Parallel returns a line parallel to this line, at a certain distance.
func (l *Line) Parallel(dist float64) *Line {
	angle := l.Angle()
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return NewLine(
		l.X0+cos*dist,
		l.Y0+sin*dist,
		l.X1+cos*dist,
		l.Y1+sin*dist,
	)
}
