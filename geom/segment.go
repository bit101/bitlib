package geom

type Segment struct {
	Start *Point
	End   *Point
}

func NewSegment(p0, p1 *Point) *Segment {
	return &Segment{p0, p1}
}

func NewSegmentXY(x0, y0, x1, y1 float64) *Segment {
	return &Segment{NewPoint(x0, y0), NewPoint(x1, y1)}
}

func LerpSegment(t float64, s0 *Segment, s1 *Segment) *Segment {
	return NewSegment(LerpPoint(t, s0.Start, s1.Start), LerpPoint(t, s0.End, s1.End))
}

func RandomSegment(x, y, w, h float64) *Segment {
	return NewSegment(RandomPoint(x, y, w, h), RandomPoint(x, y, w, h))
}

func (s *Segment) DirectionVector() *Vector {
	return NewVectorBetween(s.Start, s.End)
}

func (s *Segment) DirectionVersor() *Vector {
	return NewVersorBetween(s.Start, s.End)
}

func (s *Segment) NormalVersor() *Vector {
	return s.DirectionVersor().Perpendicular()
}

func (s *Segment) Length() float64 {
	return s.Start.Distance(s.End)
}

func (s *Segment) PointAt(t float64) *Point {
	return LerpPoint(t, s.Start, s.End)
}

func (s *Segment) Middle() *Point {
	return MidPoint(s.Start, s.End)
}

func (s *Segment) ClosestPoint(p *Point) *Point {
	v := NewVectorBetween(s.Start, p)
	d := s.DirectionVersor()
	vs := v.Project(d)
	if vs < 0 {
		return s.Start
	}
	if vs > s.Length() {
		return s.End
	}

	return s.Start.Displaced(d, vs)
}

func (s *Segment) DistanceTo(p *Point) float64 {
	return s.ClosestPoint(p).Distance(p)
}

func (s *Segment) Intersection(z *Segment) *Point {
	d1, d2 := s.DirectionVector(), z.DirectionVector()
	if d1.IsParallelTo(d2) {
		return nil
	}

	crossProd := d1.Cross(d2)
	delta := z.Start.Sub(s.Start)
	t1 := (delta.U*d2.V - delta.V*d2.U) / crossProd
	t2 := (delta.U*d1.V - delta.V*d1.U) / crossProd

	if tIsValid(t1) && tIsValid(t2) {
		return s.PointAt(t1)
	}
	return nil
}

func (s *Segment) Equals(z *Segment) bool {
	if s == z {
		return true
	}
	return s.Start.Equals(z.Start) && s.End.Equals(z.End)
}

func (s *Segment) Translate(tx, ty float64) {
	s.Start.Translate(tx, ty)
	s.End.Translate(tx, ty)
}

func (s *Segment) Scale(sx, sy float64) {
	s.Start.Scale(sx, sy)
	s.End.Scale(sx, sy)
}

func (s *Segment) Rotate(radians float64) {
	s.Start.Rotate(radians)
	s.End.Rotate(radians)
}

func (s *Segment) RotateAround(p *Point, radians float64) {
	s.Start.RotateAround(p, radians)
	s.End.RotateAround(p, radians)
}

func (s *Segment) Bisector() *Line {
	return NewLine(s.Middle(), s.NormalVersor())
}
