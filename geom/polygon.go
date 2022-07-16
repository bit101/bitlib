package geom

import (
	"log"
	"math"

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

func NewPolygonFromCoords(coords ...float64) *Polygon {
	if len(coords)%2 != 0 {
		log.Fatalln("Polygon needs an even number of coords.")
	}

	vertices := []*Point{}
	for i := 0; i < len(coords)-1; i += 2 {
		vertices = append(vertices, NewPoint(coords[i], coords[i+1]))
	}
	return NewPolygon(vertices)
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

func (p *Polygon) Contains(point *Point, context *blgg.Context) bool {
	vecs := []*Vector{}
	for _, v := range p.Vertices {
		if v.Equals(point) {
			return true
		}
		vecs = append(vecs, NewVectorBetween(point, v))
	}
	angle := 0.0
	for i := 0; i < len(vecs); i++ {
		v1 := vecs[i]
		v2 := vecs[(i+1)%len(vecs)]
		angle += v1.AngleTo(v2)
	}
	return AreClose(angle, math.Pi*2)
}

func (p *Polygon) Equals(q *Polygon) bool {
	if p == q {
		return true
	}
	if len(p.Vertices) != len(q.Vertices) {
		return false
	}

	for i := 0; i < len(p.Vertices); i++ {
		if !p.Vertices[i].Equals(q.Vertices[i]) {
			return false
		}
	}
	return true
}

func (p *Polygon) DrawVertices(context *blgg.Context, r float64) {
	for _, v := range p.Vertices {
		v.Fill(context, r)
	}
}
