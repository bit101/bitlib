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
		NewPoint(x0, y0),
		NewPoint(x1, y1),
	}
}

// NewLineFromPoints creates a new line from two Points.
func NewLineFromPoints(p0, p1 *Point) *Line {
	return NewLine(p0.X, p0.Y, p1.X, p1.Y)
}

// NewLineFromSegment creates a new line from an existing Segment.
func NewLineFromSegment(seg *Segment) *Line {
	return NewLine(seg.PointA.X, seg.PointA.Y, seg.PointB.X, seg.PointB.Y)
}

// HitSegment reports the point of intersection of a line and a line segment.
func (l *Line) HitSegment(s *Segment) (float64, float64, bool) {
	return SegmentOnLine(s.PointA.X, s.PointA.Y, s.PointB.X, s.PointB.Y, l.PointA.X, l.PointA.Y, l.PointB.X, l.PointB.Y)
}

// HitLine reports the point of intersection of two lines.
func (l *Line) HitLine(m *Line) (float64, float64, bool) {
	return LineOnLine(l.PointA.X, l.PointA.Y, l.PointB.X, l.PointB.Y, m.PointA.X, m.PointA.Y, m.PointB.X, m.PointB.Y)
}

// ClosestPoint reports the point on a line closest to another given point.
func (l *Line) ClosestPoint(p *Point) *Point {

	v := VectorBetween(l.PointA.X, l.PointA.Y, p.X, p.Y)
	d := VectorBetween(l.PointA.X, l.PointA.Y, l.PointB.X, l.PointB.Y).Normalized()
	vs := v.Project(d)
	t := vs / math.Hypot(l.PointB.X-l.PointA.X, l.PointB.Y-l.PointA.Y)
	return NewPoint(blmath.Lerp(t, l.PointA.X, l.PointB.X), blmath.Lerp(t, l.PointA.Y, l.PointB.Y))
}

// DistanceTo reports the distance from a given point to this line.
func (l *Line) DistanceTo(p *Point) float64 {
	return l.ClosestPoint(p).Distance(p)
}

// Angle returns the angle of rthis line (from the first point to the second)
func (l *Line) Angle() float64 {
	return math.Atan2(l.PointB.Y-l.PointA.Y, l.PointB.X-l.PointA.X) - math.Pi/2
}

// Perpendicular returns a line that is perpendicular to the line and crosses through the given point.
func (l *Line) Perpendicular(p *Point) *Line {
	dx := l.PointB.X - l.PointA.X
	dy := l.PointB.Y - l.PointA.Y
	return NewLine(p.X, p.Y, p.X-dy, p.Y+dx)
}

// Parallel returns a line parallel to this line, at a certain distance.
func (l *Line) Parallel(dist float64) *Line {
	angle := l.Angle()
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return NewLine(
		l.PointA.X+cos*dist,
		l.PointA.Y+sin*dist,
		l.PointB.X+cos*dist,
		l.PointB.Y+sin*dist,
	)
}
