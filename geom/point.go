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

//////////////////////////////
// Creation funcs
//////////////////////////////

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

//////////////////////////////
// Misc methods
//////////////////////////////

// Equals returns whether this point is equal to another point.
func (p *Point) Equals(other *Point) bool {
	d := 0.000001
	if p == other {
		return true
	}
	if !blmath.Equalish(p.X, other.X, d) {
		return false
	}
	if !blmath.Equalish(p.Y, other.Y, d) {
		return false
	}
	return true
}

// Clone returns a copy of this point.
func (p *Point) Clone() *Point {
	return NewPoint(p.X, p.Y)
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

// AngleBetween returns the angle formed from p0 to this point to p1.
func (p *Point) AngleBetween(p0, p1 *Point) float64 {
	d0 := p.Distance(p0)
	d1 := p.Distance(p1)
	d2 := p0.Distance(p1)
	num := (d0*d0 + d1*d1 - d2*d2)
	denom := (2 * d0 * d1)
	return math.Acos(num / denom)
}

// Clockwise returns whether or not the three points listed are in clockwise order
func Clockwise(p1, p2, p3 *Point) bool {
	return (p1.X-p3.X)*(p2.Y-p3.Y)-(p2.X-p3.X)*(p1.Y-p3.Y) > 0
}

// LinearBezierPoints creates a list of points roughly evenly distributed along the curve.
func LinearBezierPoints(count int, p0, p1, p2, p3 *Point) PointList {
	// An extra point will be added on the end of the path, so we'll account for that.
	count--
	fCount := float64(count)
	// Too few segments and the points will be approximated very innacurately.
	// So, we'll make a minimum of 50 segments.
	segCount := math.Max(fCount*2, 50)
	path := NewPointList()
	for i := 0.0; i <= segCount; i++ {
		t := i / segCount
		p := BezierPoint(t, p0, p1, p2, p3)
		path.Add(p)
	}

	points := NewPointList()
	for i := 0.0; i <= fCount; i++ {
		t := i / fCount
		p := path.Interpolate(t)
		points.Add(p)
	}
	return points
}

// BezLength returns the length of a bezier curve.
// count is how many segments the curve will be sliced into.
// The higher the count, the more accurate. 1000 is a decent start.
func BezLength(p0, p1, p2, p3 *Point, count float64) float64 {
	dist := 0.0
	pA := p0
	for i := 1.0; i < count; i++ {
		t := i / count
		pB := BezierPoint(t, p0, p1, p2, p3)
		dist += pA.Distance(pB)
		pA = pB
	}
	return dist
}

//////////////////////////////
// Transform in place
//////////////////////////////

// Normalize normalizes each component of this point, in place.
func (p *Point) Normalize() {
	mag := p.Magnitude()
	p.X /= mag
	p.Y /= mag
}

// Randomize randomizes this point.
func (p *Point) Randomize(rx, ry float64) {
	p.Translate(random.FloatRange(-rx, rx), random.FloatRange(-ry, ry))
}

// Rotate rotates this point around the origin.
func (p *Point) Rotate(angle float64) {
	x := p.X*math.Cos(angle) + p.Y*math.Sin(angle)
	y := p.Y*math.Cos(angle) - p.X*math.Sin(angle)
	p.X = x
	p.Y = y
}

// RotateFrom rotates this point around the the given x, y origin point.
func (p *Point) RotateFrom(x, y float64, angle float64) {
	p.Translate(-x, -y)

	x1 := p.X*math.Cos(angle) + p.Y*math.Sin(angle)
	y1 := p.Y*math.Cos(angle) - p.X*math.Sin(angle)
	p.X = x1
	p.Y = y1
	p.Translate(x, y)
}

// Scale scales this point on the x and y axes.
func (p *Point) Scale(scaleX float64, scaleY float64) {
	p.X *= scaleX
	p.Y *= scaleY
}

// ScaleFrom scales this point on the x and y axes, with the given x, y location as a center.
func (p *Point) ScaleFrom(x, y, scaleX float64, scaleY float64) {
	p.Translate(-x, -y)
	p.X *= scaleX
	p.Y *= scaleY
	p.Translate(x, y)
}

// Translate moves this point on the x and y axes.
func (p *Point) Translate(x float64, y float64) {
	p.X += x
	p.Y += y
}

// UniScale scales this point by the same amount on the x and y axes.
func (p *Point) UniScale(scale float64) {
	p.X *= scale
	p.Y *= scale
}

// UniScaleFrom scales this point by the same amount on the x and y axes,
// given the x, y location as a center
func (p *Point) UniScaleFrom(x, y, scale float64) {
	p.ScaleFrom(x, y, scale, scale)
}

//////////////////////////////
// Return transformed copy
//////////////////////////////

// Normalized returns a copy of this point, normalized.
func (p *Point) Normalized() *Point {
	p1 := p.Clone()
	p1.Normalize()
	return p1
}

// Randomized returns a new point, randomized from this.
func (p *Point) Randomized(rx, ry float64) *Point {
	p1 := p.Clone()
	p1.Randomize(rx, ry)
	return p1
}

// Rotated creates a new point, a rotated version of this point.
func (p *Point) Rotated(angle float64) *Point {
	p1 := p.Clone()
	p1.Rotate(angle)
	return p1
}

// RotatedFrom creates a new point, rotated from the given x, y location.
func (p *Point) RotatedFrom(x, y, angle float64) *Point {
	p1 := p.Clone()
	p1.RotateFrom(x, y, angle)
	return p1
}

// Scaled creates a new point, a scaled version of this point.
func (p *Point) Scaled(scaleX float64, scaleY float64) *Point {
	p1 := p.Clone()
	p1.Scale(scaleX, scaleY)
	return p1
}

// ScaledFrom creates a new point, scaled from a given x, y location.
func (p *Point) ScaledFrom(x, y, scaleX, scaleY float64) *Point {
	p1 := p.Clone()
	p1.ScaleFrom(x, y, scaleX, scaleY)
	return p1
}

// Translated creates a new point, a translated version of this point.
func (p *Point) Translated(tx, ty float64) *Point {
	p1 := p.Clone()
	p1.Translate(tx, ty)
	return p1
}

// UniScaled creates a new point, a scaled version of this point.
func (p *Point) UniScaled(scale float64) *Point {
	p1 := p.Clone()
	p1.UniScale(scale)
	return p1
}

// UniScaledFrom creates a new point, scaled from a given x, y location.
func (p *Point) UniScaledFrom(x, y, scale float64) *Point {
	p1 := p.Clone()
	p1.UniScaleFrom(x, y, scale)
	return p1
}
