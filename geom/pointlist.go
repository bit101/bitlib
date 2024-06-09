// Package geom has geometry related structs and funcs.
package geom

import (
	"errors"
	"math"
	"slices"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/noise"
)

// PointList is a slice of Points
type PointList []*Point

//////////////////////////////
// Creation funcs
//////////////////////////////

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

// PointGrid creats a list of points arranged in a grid.
func PointGrid(x, y, w, h, xres, yres float64) PointList {
	list := NewPointList()

	for i := x; i < x+w; i += xres {
		for j := y; j < y+h; j += yres {
			list.AddXY(i, j)
		}
	}
	return list
}

//////////////////////////////
// Misc methods
//////////////////////////////

// Clone returns a deep copy of this PointList.
func (p PointList) Clone() PointList {
	temp := NewPointList()
	for _, p := range p {
		temp.Add(p.Clone())
	}
	return temp
}

// Add adds a point to the list
func (p *PointList) Add(point *Point) {
	*p = append(*p, point)
}

// AddXY adds a point to the list
func (p *PointList) AddXY(x, y float64) {
	*p = append(*p, NewPoint(x, y))
}

// Insert inserts a point at the given index.
func (p *PointList) Insert(index int, point *Point) error {
	if index < 0 || index >= len(*p) {
		return errors.New("index out of range")
	}

	(*p) = append((*p)[:index+1], (*p)[index:]...)
	(*p)[index] = point
	return nil
}

// First gets the first point in this list
func (p PointList) First() *Point {
	return p[0]
}

// Last gets the last point in this list
func (p PointList) Last() *Point {
	return p[len(p)-1]
}

// Get returns the point at the given index.
// If the index is negative, it counts from the end of the list.
func (p PointList) Get(index int) *Point {
	if index < 0 {
		index = len(p) + index
	}
	return p[index]
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// Cull returns a new point list of points from this list that match a test
func (p PointList) Cull(test func(*Point) bool) PointList {
	out := NewPointList()
	for _, point := range p {
		if test(point) {
			out.Add(point)
		}
	}
	return out
}

// Normalize normalizes all the points in this list.
func (p *PointList) Normalize() {
	for _, point := range *p {
		point.Normalize()
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

// Randomize randomizes all the points in a list.
func (p *PointList) Randomize(rx, ry float64) {
	for _, point := range *p {
		point.Randomize(rx, ry)
	}
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

// Translate translates all the points in a list.
func (p *PointList) Translate(x, y float64) {
	for _, point := range *p {
		point.Translate(x, y)
	}
}

// Unique returns a new PointList with any duplicate points removed.
func (p PointList) Unique() PointList {
	temp := NewPointList()
	for i := 0; i < len(p); i++ {
		found := false
		for j := 0; j < len(temp); j++ {
			if p[i].Equals(temp[j]) {
				found = true
				break
			}
		}
		if !found {
			temp.Add(p[i])
		}
	}
	return temp
}

// UniScale scales all the points in a list.
func (p *PointList) UniScale(scale float64) {
	for _, point := range *p {
		point.UniScale(scale)
	}
}

// UniScaleFrom scales all the points in a list using the x, y location as a center.
func (p *PointList) UniScaleFrom(x, y, scale float64) {
	for _, point := range *p {
		point.UniScaleFrom(x, y, scale)
	}
}

//////////////////////////////
// Sort in place
//////////////////////////////

// SortXY sorts the list by x position, deciding matches with y
func (p PointList) SortXY() PointList {
	temp := p.Clone()
	slices.SortFunc(temp, func(a, b *Point) int {
		if a.X < b.X {
			return -1
		}
		if a.X > b.X {
			return 1
		}
		if a.Y < b.Y {
			return -1
		}
		if a.Y > b.Y {
			return 1
		}
		return 0
	})
	return temp
}

// SortYX sorts the list by y position, deciding matches with x
func (p PointList) SortYX() PointList {
	temp := p.Clone()
	slices.SortFunc(temp, func(a, b *Point) int {
		if a.Y < b.Y {
			return -1
		}
		if a.Y > b.Y {
			return 1
		}
		if a.X < b.X {
			return -1
		}
		if a.X > b.X {
			return 1
		}
		return 0
	})
	return temp
}

// SortDistFrom sorts the list by the distance to a given x, y location
func (p PointList) SortDistFrom(x, y float64) PointList {
	temp := p.Clone()
	point := NewPoint(x, y)
	slices.SortFunc(temp, func(a, b *Point) int {
		da := a.Distance(point)
		db := b.Distance(point)
		if da < db {
			return -1
		}
		if db < da {
			return 1
		}
		return 0
	})
	return temp
}
