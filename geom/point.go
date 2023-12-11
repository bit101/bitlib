// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
)

// Point is the structure representing a 2D point.
type Point struct {
	X, Y float64
}

// NewPoint creates a new point.
func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

// MidPoint returns a point exactly midway between two points.
func MidPoint(p0, p1 *Point) *Point {
	return LerpPoint(0.5, p0, p1)
}

// LerpPoint linearly interpolates between two points.
func LerpPoint(t float64, p0, p1 *Point) *Point {
	return NewPoint(
		blmath.Lerp(t, p0.X, p1.X),
		blmath.Lerp(t, p0.Y, p1.Y),
	)
}

// RandomPointInRect creates a random point within the defined rectangle.
func RandomPointInRect(x, y, w, h float64) *Point {
	return NewPoint(
		random.FloatRange(x, x+w),
		random.FloatRange(y, y+h),
	)
}

// RandomPointInCircle creates a random point within the defined circle.
// if dist is true, it will evenly distribute them in the circle.
func RandomPointInCircle(x, y, r float64, dist bool) *Point {
	var radius float64
	angle := random.FloatRange(0, math.Pi*2)
	if dist {
		radius = math.Sqrt(random.Float()) * r
	} else {
		radius = random.Float() * r
	}
	return NewPoint(x+math.Cos(angle)*radius, y+math.Sin(angle)*radius)
}

// RandomPointInTriangle returns a randomly generated point within the triangle described by the given points.
func RandomPointInTriangle(A, B, C *Point) *Point {
	s := random.Float()
	t := random.Float()
	a := 1.0 - math.Sqrt(t)
	b := (1.0 - s) * math.Sqrt(t)
	c := s * math.Sqrt(t)
	return NewPoint(a*A.X+b*B.X+c*C.X, a*A.Y+b*B.Y+c*C.Y)
}

// PointList is a slice of Points
type PointList []*Point

// NewPointList creates a new point list
func NewPointList() PointList {
	return PointList{}
}

// Add adds a point to the list
func (p *PointList) Add(point *Point) {
	*p = append(*p, point)
}

// AddXY adds a point to the list
func (p *PointList) AddXY(x, y float64) {
	*p = append(*p, NewPoint(x, y))
}

// Coords returns the x, y coords of this point.
func (p *Point) Coords() (float64, float64) {
	return p.X, p.Y
}

// Distance between this point and another point
func (p *Point) Distance(p1 *Point) float64 {
	return math.Hypot(p.X-p1.X, p.Y-p1.Y)
}

// Magnitude is distance from origin to this point
func (p *Point) Magnitude() float64 {
	return math.Hypot(p.X, p.Y)
}

// Angle returns the angle from the origin to this point.
func (p *Point) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

// AngleTo returns the angle from this point to another point.
func (p *Point) AngleTo(o *Point) float64 {
	return math.Atan2(o.Y-p.Y, o.X-p.X)
}

// BezierPoint calculates a point along a Bezier curve.
func BezierPoint(t float64, p0 *Point, p1 *Point, p2 *Point, p3 *Point) *Point {
	oneMinusT := 1.0 - t
	m0 := oneMinusT * oneMinusT * oneMinusT
	m1 := 3.0 * oneMinusT * oneMinusT * t
	m2 := 3.0 * oneMinusT * t * t
	m3 := t * t * t
	return &Point{
		m0*p0.X + m1*p1.X + m2*p2.X + m3*p3.X,
		m0*p0.Y + m1*p1.Y + m2*p2.Y + m3*p3.Y,
	}
}

// QuadraticPoint calculated a point along a quadratic Bezier curve.
func QuadraticPoint(t float64, p0 *Point, p1 *Point, p2 *Point) *Point {
	oneMinusT := 1.0 - t
	m0 := oneMinusT * oneMinusT
	m1 := 2.0 * oneMinusT * t
	m2 := t * t
	return &Point{
		m0*p0.X + m1*p1.X + m2*p2.X,
		m0*p0.Y + m1*p1.Y + m2*p2.Y,
	}
}

// Translate moves this point on the x and y axes.
func (p *Point) Translate(x float64, y float64) {
	p.X += x
	p.Y += y
}

// Scale scales this point on the x and y axes.
func (p *Point) Scale(scaleX float64, scaleY float64) {
	p.X *= scaleX
	p.Y *= scaleY
}

// UniScale scales this point by the same amount on the x and y axes.
func (p *Point) UniScale(scale float64) {
	p.X *= scale
	p.Y *= scale
}

// Scaled creates a new point, a scaled version of this point.
func (p *Point) Scaled(scaleX float64, scaleY float64) *Point {
	return NewPoint(p.X*scaleX, p.Y*scaleY)
}

// UniScaled creates a new point, a scaled version of this point.
func (p *Point) UniScaled(scale float64) *Point {
	return NewPoint(p.X*scale, p.Y*scale)
}

// Rotate rotates this point around the origin.
func (p *Point) Rotate(angle float64) {
	x := p.X*math.Cos(angle) + p.Y*math.Sin(angle)
	y := p.Y*math.Cos(angle) - p.X*math.Sin(angle)
	p.X = x
	p.Y = y
}

// Clockwise returns whether or not the three points listed are in clockwise order
func Clockwise(p1, p2, p3 *Point) bool {
	return (p1.X-p3.X)*(p2.Y-p3.Y)-(p2.X-p3.X)*(p1.Y-p3.Y) > 0
}

// ConvexHull returns a list of points that form a convex hull around the given set of points.
func ConvexHull(points PointList) PointList {
	hull := NewPointList()
	pointOnHull := points[0]
	for _, p := range points {
		if p.X < pointOnHull.X {
			pointOnHull = p
		}
	}
	for true {
		hull.Add(pointOnHull)
		endpoint := points[0]
		for _, p := range points {
			if endpoint == pointOnHull || Clockwise(p, endpoint, pointOnHull) {
				endpoint = p
			}
		}
		pointOnHull = endpoint
		if endpoint == hull[0] {
			break
		}
	}
	return hull
}
