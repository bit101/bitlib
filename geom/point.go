package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

func MidPoint(p0, p1 *Point) *Point {
	return NewPoint((p0.X+p1.X)/2, (p0.Y+p1.Y)/2)
}

func LerpPoint(t float64, p0, p1 *Point) *Point {
	return NewPoint(
		blmath.Lerp(t, p0.X, p1.X),
		blmath.Lerp(t, p0.Y, p1.Y),
	)
}

func RandomPointInRect(x, y, w, h float64) *Point {
	return NewPoint(
		random.FloatRange(x, x+w),
		random.FloatRange(y, y+h),
	)
}

func RandomPointInCircle(x, y, r float64) *Point {
	angle := random.FloatRange(0, math.Pi*2)
	radius := random.FloatRange(0, r)
	return NewPoint(x+math.Cos(angle)*radius, y+math.Sin(angle)*radius)
}
