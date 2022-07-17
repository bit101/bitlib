package geom

import (
	"math"

	"github.com/bit101/bitlib/random"
)

type Circle struct {
	Center *Point
	Radius float64
}

func NewCircle(center *Point, radius float64) *Circle {
	return &Circle{
		Center: center,
		Radius: radius,
	}
}

func NewCircleXY(x, y, r float64) *Circle {
	return NewCircle(NewPoint(x, y), r)
}

func RandomCircle(x, y, w, h, minRadius, maxRadius float64) *Circle {
	return NewCircle(RandomPoint(x, y, w, h), random.FloatRange(minRadius, maxRadius))
}

func NewCircleFromPoints(a, b, c *Point) *Circle {
	bisecAB := NewSegment(a, b).Bisector()
	bisecBC := NewSegment(b, c).Bisector()
	center := bisecAB.IntersectionWith(bisecBC)
	radius := center.Distance(a)
	return NewCircle(center, radius)

}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Contains(p *Point) bool {
	return p.Distance(c.Center) < c.Radius
}

func (c *Circle) ToPolygon(divs int) *Polygon {
	angle := 0.0
	delta := math.Pi * 2 / float64(divs)
	verts := []*Point{}
	for i := 0; i < divs; i++ {
		verts = append(verts, NewPoint(c.Center.X+math.Cos(angle)*c.Radius, c.Center.Y+math.Sin(angle)*c.Radius))
		angle += delta
	}
	return NewPolygon(verts)
}

func (c *Circle) Equals(d *Circle) bool {
	if c == d {
		return true
	}
	if c.Center.Equals(d.Center) {
		return c.Radius == d.Radius
	}
	return false
}
