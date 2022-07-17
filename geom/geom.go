package geom

import (
	"errors"
	"math"

	"github.com/bit101/bitlib/blmath"
)

func PointInRect(x, y, rx, ry, rw, rh float64) bool {
	return x >= rx && x <= rx+rw && y >= ry && y <= ry+rh
}

func PointInCircle(x, y, cx, cy, cr float64) bool {
	return math.Hypot(cx-x, cy-y) <= cr
}

func PointInPolygon(x, y float64, points []float64) bool {
	// TODO
	return false
}

// segment / rect
// segment / circle
// rect / rect
// circle / circle
// rect / circle

func SegmentOnLine(x0, y0, x1, y1, x2, y2, x3, y3 float64) (float64, float64, error) {
	v0 := VectorBetween(x0, y0, x1, y1)
	v1 := VectorBetween(x2, y2, x3, y3)
	crossProd := v0.CrossProduct(v1)
	if blmath.Equalish(crossProd, 0, 1e-5) {
		return 0, 0, errors.New("segments are parallel")
	}
	delta := VectorBetween(x0, y0, x2, y2)
	t1 := (delta.U*v1.V - delta.V*v1.U) / crossProd
	t2 := (delta.U*v0.V - delta.V*v0.U) / crossProd

	if tIsValid(t1) && tIsValid(t2) {
		return blmath.Lerp(t1, x0, x1), blmath.Lerp(t1, y0, y1), nil
	}
	return 0, 0, errors.New("segments do not intersect")
}

func SegmentOnSegment(x0, y0, x1, y1, x2, y2, x3, y3 float64) (float64, float64, error) {
	v0 := VectorBetween(x0, y0, x1, y1)
	v1 := VectorBetween(x2, y2, x3, y3)
	crossProd := v0.CrossProduct(v1)
	if blmath.Equalish(crossProd, 0, 1e-5) {
		return 0, 0, errors.New("segments are parallel")
	}
	delta := VectorBetween(x0, y0, x2, y2)
	t1 := (delta.U*v1.V - delta.V*v1.U) / crossProd
	t2 := (delta.U*v0.V - delta.V*v0.U) / crossProd

	if tIsValid(t1) && tIsValid(t2) {
		return blmath.Lerp(t1, x0, x1), blmath.Lerp(t1, y0, y1), nil
	}
	return 0, 0, errors.New("segments do not intersect")
}

func LineOnLine(x0, y0, x1, y1, x2, y2, x3, y3 float64) (float64, float64, error) {
	v0 := VectorBetween(x0, y0, x1, y1)
	v1 := VectorBetween(x2, y2, x3, y3)
	crossProd := v0.CrossProduct(v1)
	if blmath.Equalish(crossProd, 0, 1e-5) {
		return 0, 0, errors.New("segments are parallel")
	}
	delta := VectorBetween(x0, y0, x2, y2)
	t1 := (delta.U*v1.V - delta.V*v1.U) / crossProd
	return blmath.Lerp(t1, x0, x1), blmath.Lerp(t1, y0, y1), nil
}

func CircleOnCircle(x0, y0, r0, x1, y1, r1 float64) bool {
	dist := math.Hypot(x1-x0, y1-y0)
	return dist < (r0 + r1)
}
