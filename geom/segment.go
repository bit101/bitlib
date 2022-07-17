package geom

type Segment struct {
	X0, Y0, X1, Y1 float64
}

func NewSegment(x0, y0, x1, y1 float64) *Segment {
	return &Segment{
		X0: x0,
		Y0: y0,
		X1: x1,
		Y1: y1,
	}
}

func (s *Segment) HitSegment(z *Segment) (float64, float64, error) {
	return SegmentOnSegment(s.X0, s.Y0, s.X1, s.Y1, z.X0, z.Y0, z.X1, z.Y1)
}

func (s *Segment) HitLine(l *Line) (float64, float64, error) {
	return SegmentOnLine(s.X0, s.Y0, s.X1, s.Y1, l.X0, l.Y0, l.X1, l.Y1)
}
