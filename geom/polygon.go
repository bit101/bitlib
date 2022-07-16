package geom

import (
	"log"

	"github.com/bit101/blgg/blgg"
)

type Polygon struct {
	Vertices []*Point
	Sides    []*Segment
}

func NewPolygon(vertices []*Point) *Polygon {
	if len(vertices) < 3 {
		log.Fatalln("A Polygon needs 3 or more vertices.")
	}

	p := &Polygon{
		Vertices: vertices,
	}
	p.makeSides()
	return p
}

func (p *Polygon) makeSides() {
	numVertices := len(p.Vertices)
	for i := 0; i < numVertices; i++ {
		a := i
		b := (i + 1) % numVertices
		p.Sides = append(p.Sides, NewSegment(p.Vertices[a], p.Vertices[b]))
	}
}

func (p *Polygon) Draw(context *blgg.Context) {
	for _, s := range p.Sides {
		s.Stroke(context)
	}
}

func (p *Polygon) Centroid() *Point {
	vSum := NewPoint(0, 0)
	for _, v := range p.Vertices {
		vSum.X += v.X
		vSum.Y += v.Y
	}
	count := float64(len(p.Vertices))
	vSum.X /= count
	vSum.Y /= count
	return vSum
}

func (p *Polygon) DrawVertices(context *blgg.Context, r float64) {
	for _, v := range p.Vertices {
		v.Fill(context, r)
	}
}
