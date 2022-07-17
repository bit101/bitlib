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
