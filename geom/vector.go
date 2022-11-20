// Package geom has geometry related structs and funcs.
package geom

import "math"

// Vector is a struct representing a 2D vector.
type Vector struct {
	U, V float64
}

// VectorBetween returns the vector between two points.
func VectorBetween(x0, y0, x1, y1 float64) *Vector {
	return &Vector{
		U: x1 - x0,
		V: y1 - y0,
	}
}

// DotProduct returns the dot product between this and another vector.
func (v *Vector) DotProduct(w *Vector) float64 {
	return v.U*w.U + v.V*w.V
}

// CrossProduct returns the cross product between this and another vector.
func (v *Vector) CrossProduct(w *Vector) float64 {
	return v.U*w.V - v.V*w.U
}

// Norm returns a normalized vector.
func (v *Vector) Norm() float64 {
	return math.Hypot(v.U, v.V)
}

// AngleValueTo returns the angle to another vector.
func (v *Vector) AngleValueTo(w *Vector) float64 {
	dotProduct := v.DotProduct(w)
	normProduct := v.Norm() * w.Norm()
	return math.Acos(dotProduct / normProduct)
}
func tIsValid(t float64) bool {
	return t >= 0.0 && t <= 1.0
}

// Normalized returns this vector normalized.
func (v *Vector) Normalized() *Vector {
	return v.Scaled(1.0 / v.Norm())
}

// Scaled returns this vector scaled.
func (v *Vector) Scaled(factor float64) *Vector {
	return &Vector{U: v.U * factor, V: v.V * factor}
}

// Project returns the value by projecting this vector onto another vector.
func (v *Vector) Project(w *Vector) float64 {
	return v.DotProduct(w.Normalized())
}
