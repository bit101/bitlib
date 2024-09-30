// Package geom has geometry related structs and funcs.
package geom

import "github.com/bit101/bitlib/blmath"

// BezierCurve represents a Bezier curve.
// In addition to containing the control points of the curve,
// it precalculates a segmented path that represents the curve as a series of points.
// This can be used to linearly interpolate along the curve, to find points
// or the slope of the curve at those positions.
// Changing the segment count will recalculate the path.
type BezierCurve struct {
	P0, P1, P2, P3 *Point
	path           PointList
	segCount       int
}

// NewBezierCurve creates a new BezierCurve object.
func NewBezierCurve(p0, p1, p2, p3 *Point) *BezierCurve {
	bez := &BezierCurve{
		p0, p1, p2, p3,
		nil,
		200,
	}
	bez.makeSegPath()
	return bez
}

// MakeSegPath creates a segmented path for this curve.
func (b *BezierCurve) makeSegPath() {
	path := NewPointList()
	count := float64(b.segCount)
	for i := 0.0; i <= count; i++ {
		t := i / count
		path.Add(b.Point(t))
	}
	b.path = path
}

// Point returns a point on the curve interpolated from t = 0.0 to 1.0.
func (b *BezierCurve) Point(t float64) *Point {
	return BezierPoint(t, b.P0, b.P1, b.P2, b.P3)
}

// SetSegmentCount sets the count of segments to use in a path of this curve, recreates path.
func (b *BezierCurve) SetSegmentCount(count int) {
	b.segCount = count
	b.makeSegPath()
}

// GetSegmentCount returns the current segment count.
func (b *BezierCurve) GetSegmentCount() int {
	return b.segCount
}

// GetPath returns the point list representing this path.
func (b *BezierCurve) GetPath() PointList {
	return b.path
}

// LinearPoints creates a list of points roughly evenly distributed along the curve.
func (b *BezierCurve) LinearPoints(count int) PointList {
	// An extra point will be added on the end of the path, so we'll account for that.
	count--
	fCount := float64(count)
	path := NewPointList()
	for i := 0.0; i <= fCount; i++ {
		t := i / fCount
		p := b.path.Interpolate(t)
		path.Add(p)
	}
	return path
}

// LerpPoint returns a linearly interpolated point along the BezierCurve.
func (b *BezierCurve) LerpPoint(t float64) *Point {
	return b.path.Interpolate(t)
}

// SlopeAtLinearT returns the slope at a given linearly interpolated point.
func (b *BezierCurve) SlopeAtLinearT(t float64) float64 {
	t = blmath.Clamp(t, 0, 1)
	delta := 0.001
	if t+delta > 1.0 {
		p0 := b.LerpPoint(1 - delta)
		p1 := b.LerpPoint(1)
		return p0.AngleTo(p1)
	}
	if t-delta < 0 {
		p0 := b.LerpPoint(0)
		p1 := b.LerpPoint(delta)
		return p0.AngleTo(p1)
	}
	p0 := b.LerpPoint(t - delta)
	p1 := b.LerpPoint(t + delta)
	return p0.AngleTo(p1)
}

// SlopeAtBezierT returns the slope at a given t value of the Bezier curve.
func (b *BezierCurve) SlopeAtBezierT(t float64) float64 {
	t = blmath.Clamp(t, 0, 1)
	delta := 0.001
	if t+delta > 1.0 {
		p0 := b.Point(1 - delta)
		p1 := b.Point(1)
		return p0.AngleTo(p1)
	}
	if t-delta < 0 {
		p0 := b.Point(0)
		p1 := b.Point(delta)
		return p0.AngleTo(p1)
	}
	p0 := b.Point(t - delta)
	p1 := b.Point(t + delta)
	return p0.AngleTo(p1)
}
