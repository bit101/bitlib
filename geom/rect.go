// Package geom has geometry related structs and funcs.
package geom

// Rect represents a rectangle
type Rect struct {
	X, Y, W, H float64
}

// NewRect creates a new rectangle.
func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

// NewRectFromPoints creates a new rectangle.
func NewRectFromPoints(p0, p1 *Point) *Rect {
	x0 := p0.X
	y0 := p0.Y
	x1 := p1.X
	y1 := p1.Y
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return NewRect(x0, y0, x1-x0, y1-y0)
}

// Contains reports whether a point is in this rectangle.
func (r *Rect) Contains(p *Point) bool {
	return PointInRect(p.X, p.Y, r.X, r.Y, r.W, r.H)
}

// HitRect reports whether or not another rectangle is intersecting this one.
func (r *Rect) HitRect(s *Rect) bool {
	return RectOnRect(r.X, r.Y, r.W, r.H, s.X, s.Y, s.W, s.H)
}

// HitSegment reports whether a line segment is interesecting this rectangle.
func (r *Rect) HitSegment(s *Segment) bool {
	return SegmentOnRect(s.X0, s.Y0, s.X1, s.Y1, r.X, r.Y, r.W, r.H)
}
