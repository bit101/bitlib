// Package geom has geometry related structs and funcs.
package geom

import (
	"math"

	"github.com/bit101/bitlib/blmath"
)

// PointInRect returns whether or not a point is within a rectangle.
func PointInRect(x, y, rx, ry, rw, rh float64) bool {
	return x >= rx && x <= rx+rw && y >= ry && y <= ry+rh
}

// PointInCircle returns whether or not a point is within a circle.
func PointInCircle(x, y, cx, cy, cr float64) bool {
	return math.Hypot(cx-x, cy-y) <= cr
}

// PointInPolygon returns whether or not a point is within a polygon.
func PointInPolygon(x, y float64, points []float64) bool {
	// TODO
	return false
}

// segment / circle
// rect / circle

// SegmentOnLine returns whether or not a segment intersects an infinite length line and returns the point it intersects if it does.
// x0, y0 -> x1, y1 = segment. x2, y2 -> x3, y3 = line
func SegmentOnLine(x0, y0, x1, y1, x2, y2, x3, y3 float64) (float64, float64, bool) {
	v0 := VectorBetween(x0, y0, x1, y1)
	v1 := VectorBetween(x2, y2, x3, y3)
	crossProd := v0.CrossProduct(v1)
	if blmath.Equalish(crossProd, 0, 1e-5) {
		return 0, 0, false
	}
	delta := VectorBetween(x0, y0, x2, y2)
	t1 := (delta.U*v1.V - delta.V*v1.U) / crossProd

	if tIsValid(t1) {
		return blmath.Lerp(t1, x0, x1), blmath.Lerp(t1, y0, y1), true
	}
	return 0, 0, false
}

// SegmentOnSegment returns whether or not two line segments intersect and the point they intersect if they do.
func SegmentOnSegment(x0, y0, x1, y1, x2, y2, x3, y3 float64) (float64, float64, bool) {
	v0 := VectorBetween(x0, y0, x1, y1)
	v1 := VectorBetween(x2, y2, x3, y3)
	crossProd := v0.CrossProduct(v1)
	if blmath.Equalish(crossProd, 0, 1e-5) {
		return 0, 0, false
	}
	delta := VectorBetween(x0, y0, x2, y2)
	t1 := (delta.U*v1.V - delta.V*v1.U) / crossProd
	t2 := (delta.U*v0.V - delta.V*v0.U) / crossProd

	if tIsValid(t1) && tIsValid(t2) {
		return blmath.Lerp(t1, x0, x1), blmath.Lerp(t1, y0, y1), true
	}
	return 0, 0, false
}

// LineOnLine returns whether or not two infinite length lines intersect and the point they intersect if they do.
func LineOnLine(x0, y0, x1, y1, x2, y2, x3, y3 float64) (float64, float64, bool) {
	v0 := VectorBetween(x0, y0, x1, y1)
	v1 := VectorBetween(x2, y2, x3, y3)
	crossProd := v0.CrossProduct(v1)
	if blmath.Equalish(crossProd, 0, 1e-5) {
		return 0, 0, false
	}
	delta := VectorBetween(x0, y0, x2, y2)
	t1 := (delta.U*v1.V - delta.V*v1.U) / crossProd
	return blmath.Lerp(t1, x0, x1), blmath.Lerp(t1, y0, y1), true
}

// CircleOnCircle returns whether or not two circles intersect.
// TODO: return the points of intersection.
func CircleOnCircle(x0, y0, r0, x1, y1, r1 float64) bool {
	dist := math.Hypot(x1-x0, y1-y0)
	return dist < (r0 + r1)
}

// RectOnRect returns whether or not two rectangles intersect.
// TODO: return the rect of intersection.
func RectOnRect(x0, y0, w0, h0, x1, y1, w1, h1 float64) bool {
	return !(x1 > x0+w0 ||
		x1+w1 < x0 ||
		y1 > y0+h0 ||
		y1+h1 < y0)
}

// SegmentOnRect returns whether or not a segment intersects a rectangle.
// TODO: return the points of intersection.
func SegmentOnRect(x0, y0, x1, y1, rx, ry, rw, rh float64) bool {
	_, _, hit := SegmentOnSegment(x0, y0, x1, y1, rx, ry, rx+rw, ry)
	if hit {
		return true
	}
	_, _, hit = SegmentOnSegment(x0, y0, x1, y1, rx+rw, ry, rx+rw, ry+rh)
	if hit {
		return true
	}
	_, _, hit = SegmentOnSegment(x0, y0, x1, y1, rx+rw, ry+rh, rx, ry+rh)
	if hit {
		return true
	}
	_, _, hit = SegmentOnSegment(x0, y0, x1, y1, rx, ry+rh, rx, ry)
	if hit {
		return true
	}
	return false
}

// PointDistanceToSegment reports the distance from a point to a line segment
func PointDistanceToSegment(px, py, x0, y0, x1, y1 float64) float64 {
	// https://stackoverflow.com/questions/849211/shortest-distance-between-a-point-and-a-line-segment
	a := px - x0
	b := py - y0
	c := x1 - x0
	d := y1 - y0

	dot := a*c + b*d
	lenSq := c*c + d*d
	param := -1.0
	if lenSq != 0 {
		//in case of 0 length line
		param = dot / lenSq
	}

	var xx, yy float64
	if param < 0 {
		xx = x0
		yy = y0
	} else if param > 1 {
		xx = x1
		yy = y1
	} else {
		xx = x0 + param*c
		yy = y0 + param*d
	}

	return math.Hypot(px-xx, py-yy)
}

// PointDistanceToLine reports the distance from a point to a line.
func PointDistanceToLine(px, py, x0, y0, x1, y1 float64) float64 {
	dx := x1 - x0
	dy := y1 - y0
	numerator := math.Abs(dx*(y0-py) - (x0-px)*dy)
	denominator := math.Sqrt(dx*dx + dy*dy)
	return numerator / denominator
}

// CircleToLine reports the points of intersection between a line and a circle.
func CircleToLine(x0, y0, x1, y1, cx, cy, r float64) (*Point, *Point, bool) {
	p := NewPoint(cx, cy)
	l := NewLine(x0, y0, x1, y1)
	lp := l.ClosestPoint(p)
	d := lp.Distance(p)

	// circle/line intersection forms a chord
	if d < r {
		// height of chord
		h := r - d
		// half length of chord
		c := math.Sqrt(r*r - math.Pow(r-h, 2))
		angle := math.Atan2(y1-y0, x1-x0)
		cos := math.Cos(angle)
		sin := math.Sin(angle)
		// two points on line c distance away from closest point.
		return NewPoint(lp.X+cos*c, lp.Y+sin*c), NewPoint(lp.X-cos*c, lp.Y-sin*c), true
	}
	return nil, nil, false
}
