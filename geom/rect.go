package geom

import (
	"log"
	"math"
)

type Rect struct {
	Origin *Point
	Size   *Size
}

func NewRect(p *Point, s *Size) *Rect {
	return &Rect{
		Origin: p,
		Size:   s,
	}
}

func NewRectFromPoints(p0, p1 *Point) *Rect {
	return NewRect(p0, NewSize(p1.X-p0.X, p1.Y-p0.Y))
}

func NewRectXY(x, y, w, h float64) *Rect {
	return NewRect(NewPoint(x, y), NewSize(w, h))
}

func NewRectContaining(points []*Point) *Rect {
	if len(points) < 1 {
		log.Fatalln("Rectangle should contain at least one point.")
	}
	first := points[0]
	minX, maxX := first.X, first.X
	minY, maxY := first.Y, first.Y

	for i := 1; i < len(points); i++ {
		p := points[i]
		minX, maxX = math.Min(minX, p.X), math.Max(maxX, p.X)
		minY, maxY = math.Min(minY, p.Y), math.Max(maxY, p.Y)
	}
	return NewRectFromPoints(NewPoint(minX, minY), NewPoint(maxX, maxY))
}

func NewRectContainingWithMargin(points []*Point, margin float64) *Rect {
	rect := NewRectContaining(points)
	return NewRect(
		NewPoint(rect.Origin.X-margin, rect.Origin.Y-margin),
		NewSize(rect.Size.Width+margin*2, rect.Size.Height+margin*2),
	)
}

func RandomRect(x, y, w, h, minW, maxW, minH, maxH float64) *Rect {
	return NewRect(RandomPoint(x, y, w, h), RandomSize(minW, maxW, minH, maxH))
}

func (r *Rect) Left() float64 {
	return r.Origin.X
}

func (r *Rect) Right() float64 {
	return r.Origin.X + r.Size.Width
}

func (r *Rect) Top() float64 {
	return r.Origin.Y
}

func (r *Rect) Bottom() float64 {
	return r.Origin.Y + r.Size.Height
}

func (r *Rect) Area() float64 {
	return r.Size.Width + r.Size.Height
}

func (r *Rect) Perimeter() float64 {
	return 2 * (r.Size.Width + r.Size.Height)
}

func (r *Rect) Contains(p *Point) bool {
	return r.Left() < p.X && p.X < r.Right() && r.Top() < p.Y && p.Y < r.Bottom()
}

func (r *Rect) Intersection(s *Rect) *Rect {
	hOverlap := r.horizOverlap(s)
	if hOverlap == nil {
		return nil
	}
	vOverlap := r.vertOverlap(s)
	if vOverlap == nil {
		return nil
	}
	return NewRect(NewPoint(hOverlap.Start, vOverlap.Start), NewSize(hOverlap.Length(), vOverlap.Length()))
}

func (r *Rect) horizOverlap(s *Rect) *OpenInterval {
	selfInterval := NewOpenInterval(r.Left(), r.Right())
	otherInterval := NewOpenInterval(s.Left(), s.Right())
	return selfInterval.ComputeOverlap(otherInterval)
}

func (r *Rect) vertOverlap(s *Rect) *OpenInterval {
	selfInterval := NewOpenInterval(r.Top(), r.Bottom())
	otherInterval := NewOpenInterval(s.Top(), s.Bottom())
	return selfInterval.ComputeOverlap(otherInterval)
}

func (r *Rect) ToPolygon() *Polygon {
	return NewPolygon([]*Point{
		r.Origin,
		NewPoint(r.Right(), r.Top()),
		NewPoint(r.Right(), r.Bottom()),
		NewPoint(r.Left(), r.Bottom()),
	})
}

func (r *Rect) Equals(s *Rect) bool {
	if r == s {
		return true
	}
	return r.Origin.Equals(s.Origin) && r.Size.Equals(s.Size)
}

