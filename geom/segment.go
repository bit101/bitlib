package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
)

type Segment struct {
	X0, Y0, X1, Y1 float64
}

func NewSegment(x0, y0, x1, y1 float64) *Segment {
	return &Segment{
		X0: x0,
		Y0: y0,
		X1: x1,
		Y1: y1,
	}
}

func (s *Segment) HitSegment(z *Segment) (float64, float64, bool) {
	return SegmentOnSegment(s.X0, s.Y0, s.X1, s.Y1, z.X0, z.Y0, z.X1, z.Y1)
}

func (s *Segment) HitLine(l *Line) (float64, float64, bool) {
	return SegmentOnLine(s.X0, s.Y0, s.X1, s.Y1, l.X0, l.Y0, l.X1, l.Y1)
}

func (s *Segment) HitRect(r *Rect) bool {
	return SegmentOnRect(s.X0, s.Y0, s.X1, s.Y1, r.X, r.Y, r.W, r.H)
}

func (s *Segment) Length() float64 {
	return math.Hypot(s.X1-s.X0, s.Y1-s.Y0)
}

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
