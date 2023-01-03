// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
)

// Circle is the struct representing a circle.
type Circle struct {
	X, Y, Radius float64
}

// NewCircle creates a new circle struct.
func NewCircle(x, y, r float64) *Circle {
	return &Circle{
		X:      x,
		Y:      y,
		Radius: r,
	}
}

// NewCircleFromPoint creates a new circle struct using a Point object as center.
func NewCircleFromPoint(center *Point, r float64) *Circle {
	return NewCircle(center.X, center.Y, r)
}

// Hit reports whether this circle intersects another circle.
func (c *Circle) Hit(d *Circle) bool {
	return CircleOnCircle(c.X, c.Y, c.Radius, d.X, d.Y, d.Radius)
}

// Intersection returns the points of intersection (if any) with another circles
// and a bool indicating whether or not they intersect at all.
func (c *Circle) Intersection(c1 *Circle) (*Point, *Point, bool) {
	if c.Hit(c1) {
		d := math.Sqrt(math.Pow(c.X-c1.X, 2) + math.Pow(c.Y-c1.Y, 2))
		l := (c.Radius*c.Radius - c1.Radius*c1.Radius + d*d) / (2 * d)
		h := math.Sqrt(c.Radius*c.Radius - l*l)
		x0 := l*(c1.X-c.X)/d + h*(c1.Y-c.Y)/d + c.X
		y0 := l*(c1.Y-c.Y)/d - h*(c1.X-c.X)/d + c.Y
		x1 := l*(c1.X-c.X)/d - h*(c1.Y-c.Y)/d + c.X
		y1 := l*(c1.Y-c.Y)/d + h*(c1.X-c.X)/d + c.Y
		return NewPoint(x0, y0), NewPoint(x1, y1), true
	}
	return nil, nil, false
}

// Contains reports whether a point is within this circle.
func (c *Circle) Contains(p *Point) bool {
	return PointInCircle(p.X, p.X, c.X, c.Y, c.Radius)
}

// InvertPoint inverts a point into our out of a circle.
func (c *Circle) InvertPoint(p *Point) *Point {
	x, y := c.InvertXY(p.X, p.Y)
	return NewPoint(x, y)
}

// InvertXY inverts a point into our out of a circle.
func (c *Circle) InvertXY(x, y float64) (float64, float64) {
	dx := x - c.X
	dy := y - c.Y
	dist0 := math.Sqrt(dx*dx + dy*dy)
	dist1 := c.Radius * c.Radius / dist0
	ratio := dist1 / dist0

	return c.X + dx*ratio, c.Y + dy*ratio
}

// OuterCircles returns a slice of circles arrange around the outside of the given circle.
func OuterCircles(c *Circle, count int, rotation float64) []*Circle {
	circles := make([]*Circle, count)
	countF := float64(count)

	angle := blmath.Tau / countF
	sinA2 := math.Sin(angle / 2)
	s := (c.Radius * sinA2) / (1 - sinA2)
	r := c.Radius + s
	a := rotation
	for i := 0; i < count; i++ {
		circles[i] = NewCircle(c.X+math.Cos(a)*r, c.Y+math.Sin(a)*r, s)
		a += angle
	}
	return circles
}

// InnerCircles returns a slice of circles arrange around the inside of the given circle.
func InnerCircles(c *Circle, count int, rotation float64) []*Circle {
	circles := make([]*Circle, count)
	countF := float64(count)

	angle := blmath.Tau / countF
	sinA2 := math.Sin(angle / 2)
	s := (c.Radius * sinA2) / (1 + sinA2)
	r := c.Radius - s
	a := rotation
	for i := 0; i < count; i++ {
		circles[i] = NewCircle(c.X+math.Cos(a)*r, c.Y+math.Sin(a)*r, s)
		a += angle
	}
	return circles
}

// func TangentPointToCircle(point *Point, circle *Circle, anticlockwise bool) *Point {
// 	d := math.Hypot(c.X-point.X, c.Y-point.Y)
// 	dir := -1.0
// 	if anticlockwise {
// 		dir = 1.0
// 	}
// 	angle := math.Cos(-circle.Radius/d) * dir
// 	baseAngle := math.Atan2(circle.Center.Y-point.Y, circle.Center.X-point.X)
// 	totalAngle := baseAngle + angle

// 	return &Point{
// 		circle.Center.X + math.Cos(totalAngle)*circle.Radius,
// 		circle.Center.Y + math.Sin(totalAngle)*circle.Radius,
// 	}
// }
