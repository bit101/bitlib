package geom

type AffineTransform struct {
	sx  float64
	sy  float64
	tx  float64
	ty  float64
	shx float64
	shy float64
}

func NewAffineTransform(sx, sy, tx, ty, shx, shy float64) *AffineTransform {
	return &AffineTransform{
		sx:  sx,
		sy:  sy,
		tx:  tx,
		ty:  ty,
		shx: shx,
		shy: shy,
	}
}

func NewAffineTransformIdentity() *AffineTransform {
	return NewAffineTransform(1, 1, 0, 0, 0, 0)
}

func (a *AffineTransform) ApplyToPoint(p *Point) *Point {
	return NewPoint(
		(a.sx*p.X)+(a.shx*p.Y)+a.tx,
		(a.sy*p.Y)+(a.shy*p.X)+a.ty,
	)
}

func (a *AffineTransform) ApplyToSegment(s *Segment) *Segment {
	return NewSegment(
		a.ApplyToPoint(s.Start),
		a.ApplyToPoint(s.End),
	)
}

func (a *AffineTransform) ApplyToPolygon(p *Polygon) *Polygon {
	vertices := []*Point{}
	for _, v := range p.Vertices {
		vertices = append(vertices, a.ApplyToPoint(v))
	}
	return NewPolygon(vertices)
}

func (a *AffineTransform) ApplyToRect(r *Rect) *Polygon {
	return a.ApplyToPolygon(r.ToPolygon())
}

func (a *AffineTransform) ApplyToCircl(c *Circle, divs int) *Polygon {
	return a.ApplyToPolygon(c.ToPolygon(divs))
}

func (a *AffineTransform) Then(b *AffineTransform) *AffineTransform {
	return NewAffineTransform(
		b.sx*a.sx+b.shx*a.shy,
		b.shy*a.shx+b.sy*a.sy,
		b.sx*a.tx+b.shx*a.ty+b.tx,
		b.shy*a.tx+b.sy*a.ty+b.ty,
		b.sx*a.shx+b.shx*a.sy,
		b.shy*a.sx+b.sy*a.shy,
	)
}

func (a *AffineTransform) Equals(b *AffineTransform) bool {
	if a == b {
		return true
	}
	return AreClose(a.sx, b.sx) &&
		AreClose(a.sy, b.sy) &&
		AreClose(a.shx, b.shx) &&
		AreClose(a.ty, b.ty) &&
		AreClose(a.ty, b.ty) &&
		AreClose(a.shy, b.shy)
}

func (a *AffineTransform) Inverse() *AffineTransform {
	denom := a.sx*a.sy - a.shx*a.shy
	return NewAffineTransform(
		a.sy/denom,
		a.sx/denom,
		(a.ty*a.shx-a.sy*a.tx)/denom,
		(a.tx*a.shy-a.sx*a.ty)/denom,
		-a.shx/denom,
		-a.shy/denom,
	)
}
