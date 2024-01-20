// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
)

// PointList is a slice of Points
type PointList []*Point

// NewPointList creates a new point list
func NewPointList() PointList {
	return PointList{}
}

// RandomPointList returns a list of random points in a given rect.
func RandomPointList(count int, x, y, w, h float64) PointList {
	points := NewPointList()
	for i := 0; i < count; i++ {
		points.Add(RandomPointInRect(x, y, w, h))
	}
	return points
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

// Add adds a point to the list
func (p *PointList) Add(point *Point) {
	*p = append(*p, point)
}

// AddXY adds a point to the list
func (p *PointList) AddXY(x, y float64) {
	*p = append(*p, NewPoint(x, y))
}

// Rotate rotates all the points in a list.
func (p *PointList) Rotate(angle float64) {
	for _, point := range *p {
		point.Rotate(angle)
	}
}

// RotateFrom rotates all the points in a list using the x, y location as a center.
func (p *PointList) RotateFrom(x, y float64, angle float64) {
	for _, point := range *p {
		point.RotateFrom(x, y, angle)
	}
}

// Translate translates all the points in a list.
func (p *PointList) Translate(x, y float64) {
	for _, point := range *p {
		point.Translate(x, y)
	}
}

// Scale scales all the points in a list.
func (p *PointList) Scale(sx, sy float64) {
	for _, point := range *p {
		point.Scale(sx, sy)
	}
}

// ScaleFrom scales all the points in a list using the x, y location as a center.
func (p *PointList) ScaleFrom(x, y, sx, sy float64) {
	for _, point := range *p {
		point.ScaleFrom(x, y, sx, sy)
	}
}

// Randomize randomizes all the points in a list.
func (p *PointList) Randomize(rx, ry float64) {
	for _, point := range *p {
		point.Randomize(rx, ry)
	}
}

// Noisify warps the point locations with simplex noise
func (p *PointList) Noisify(sx, sy, z, offset float64) {
	for _, point := range *p {
		t := noise.Simplex3(point.X*sx, point.Y*sy, z) * blmath.Tau
		point.X += math.Cos(t) * offset
		point.Y += math.Sin(t) * offset
	}
}

// PointListCullTest is a function that takes a point and returns a bool.
// Used for culling points from a list.
type PointListCullTest func(*Point) bool

// Cull returns a new point list of points from this list that match a test
func (p PointList) Cull(test PointListCullTest) PointList {
	out := NewPointList()
	for _, point := range p {
		if test(point) {
			out.Add(point)
		}
	}
	return out
}
