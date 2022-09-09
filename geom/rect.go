package geom

type Rect struct {
	X, Y, W, H float64
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
}

func (r *Rect) Contains(p *Point) bool {
	return PointInRect(p.X, p.Y, r.X, r.Y, r.W, r.H)
}

func (r *Rect) HitRect(s *Rect) bool {
	return RectOnRect(r.X, r.Y, r.W, r.H, s.X, s.Y, s.W, s.H)
}

func (r *Rect) HitSegment(s *Segment) bool {
	return SegmentOnRect(s.X0, s.Y0, s.X1, s.Y1, r.X, r.Y, r.W, r.H)
}
