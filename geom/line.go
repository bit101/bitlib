package geom

// A Line is the same as a segment, but we'll keep them separate for collision purposes
type Line Segment

func NewLine(x0, y0, x1, y1 float64) *Line {
	return &Line{
		X0: x0,
		Y0: y0,
		X1: x1,
		Y1: y1,
	}
}

func (l *Line) HitSegment(s *Segment) (float64, float64, error) {
	return SegmentOnSegment(s.X0, s.Y0, s.X1, s.Y1, l.X0, l.Y0, l.X1, l.Y1)
}

func (l *Line) HitLine(m *Line) (float64, float64, error) {
	return SegmentOnLine(l.X0, l.Y0, l.X1, l.Y1, m.X0, m.Y0, m.X1, m.Y1)
}
