// Package sdf defines signed distance functions
package sdf

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
)

// Min returns the minimum of a variable number of args.
func Min(shapes ...float64) float64 {
	min := math.MaxFloat64
	for _, v := range shapes {
		min = math.Min(min, v)
	}
	return min
}

// Max returns the maximum of a variable number of args.
func Max(shapes ...float64) float64 {
	max := -math.MaxFloat64
	for _, v := range shapes {
		max = math.Max(max, v)
	}
	return max
}

func Repeat(val, dist float64) float64 {
	val /= dist
	val -= math.Floor(val) + 0.5
	val *= dist
	return val
}

func Hline(x, y, pos float64) float64 {
	return y - pos
}

func Vline(x, y, pos float64) float64 {
	return x - pos
}

// Circle computes the signed distance to a circle.
func Circle(x, y, cx, cy, radius float64) float64 {
	v := geom.NewVector(x-cx, y-cy)
	return v.Magnitude() - radius
}

// Box computes the signed distance to an axis-aligned box.
func Box(x, y, bx, by, bw, bh float64) float64 {
	p := geom.NewVector(x-bx, y-by)
	b := geom.NewVector(bw, bh)
	d := p.Abs().Subtract(b)
	return d.Max(0.0).Magnitude() + math.Min(math.Max(d.U, d.V), 0.0)
}

// RoundBox computes the signed distance to an axis-aligned rounded box.
func RoundBox(x, y, bx, by, bw, bh, r float64) float64 {
	p := geom.NewVector(x-bx, y-by)
	b := geom.NewVector(bw, bh)
	rv := geom.NewVector(r, r)
	q := p.Abs().Subtract(b).Add(rv)
	return math.Min(math.Max(q.U, q.V), 0.0) + q.Max(0.0).Magnitude() - r
}

func Segment(x, y, ax, ay, bx, by float64) float64 {
	p := geom.NewVector(x, y)
	a := geom.NewVector(ax, ay)
	b := geom.NewVector(bx, by)

	pa := p.Subtract(a)
	ba := b.Subtract(a)
	h := blmath.Clamp(pa.DotProduct(ba)/ba.DotProduct(ba), 0, 1)
	return pa.Subtract(ba.Scaled(h)).Magnitude()
}

// float sdSegment( in vec2 p, in vec2 a, in vec2 b )
// {
//     vec2 pa = p-a, ba = b-a;
//     float h = clamp( dot(pa,ba)/dot(ba,ba), 0.0, 1.0 );
//     return length( pa - ba*h );
// }
// func OrientedBox(x, y, ax, ay, bx, by, bw, bh, t float64) float64 {
// 	p := geom.NewVector(x-bx, y-by)
// 	a := geom.NewVector(ax, ay)
// 	b := geom.NewVector(bw, bh)
// 	l := b.Subtract(a).Magnitude()
// 	q := p.Subtract(a.Add(b).Scaled(0.5))
// 	q =

// }

// float sdOrientedBox( in vec2 p, in vec2 a, in vec2 b, float th )
// {
//     float l = length(b-a);
//     vec2  d = (b-a)/l;
//     vec2  q = (p-(a+b)*0.5);
//           q = mat2(d.x,-d.y,d.y,d.x)*q;
//           q = abs(q)-vec2(l,th)*0.5;
//     return length(max(q,0.0)) + min(max(q.x,q.y),0.0);
// }
