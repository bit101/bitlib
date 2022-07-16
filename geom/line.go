package geom

import (
	"math"

	"github.com/bit101/blgg/blgg"
)

type Line struct {
	Base      *Point
	Direction *Vector
}

func NewLine(base *Point, direction *Vector) *Line {
	return &Line{Base: base, Direction: direction}
}

func NewLineFromPoints(p0, p1 *Point) *Line {
	return NewLine(p0, NewVector(p1.X-p0.X, p1.Y-p0.Y))
}

func NewLineXY(x0, y0, x1, y1 float64) *Line {
	return NewLine(NewPoint(x0, y0), NewVector(x1-x0, y1-y0))
}

func (l *Line) IsParallelTo(m *Line) bool {
	return l.Direction.IsParallelTo(m.Direction)
}

func (l *Line) IsPerpendicularTo(m *Line) bool {
	return l.Direction.IsPerpendicularTo(m.Direction)
}

func (l *Line) PerpendicularThrough(p *Point) *Line {
	return NewLine(p, l.Direction.Perpendicular())
}

func (l *Line) ParallelThrough(p *Point) *Line {
	return NewLine(p, l.Direction)
}

func (l *Line) IntersectionWith(m *Line) *Point {
	if l.IsParallelTo(m) {
		return nil
	}
	d1, d2 := l.Direction, m.Direction
	crossProd := d1.Cross(d2)
	delta := NewVectorBetween(l.Base, m.Base)
	t1 := (delta.U*d2.V - delta.V*d2.U) / crossProd
	return l.Base.Displaced(d1, t1)
}

func (l *Line) Stroke(context *blgg.Context, length float64) {
	context.Push()
	context.Translate(l.Base.X, l.Base.Y)
	context.Rotate(math.Atan2(l.Direction.V, l.Direction.U))
	context.MoveTo(-length/2, 0)
	context.LineTo(length/2, 0)
	context.Pop()
	context.Stroke()
}
