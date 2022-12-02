// Package geom has geometry related structs and funcs.
package geom

import "math"

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

// Hit reports whether this circle intersects another circle.
func (c *Circle) Hit(d *Circle) bool {
	return CircleOnCircle(c.X, c.Y, c.Radius, d.X, d.Y, d.Radius)
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
