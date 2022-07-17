package geom

import "math"

type Vector struct {
	U, V float64
}

func VectorBetween(x0, y0, x1, y1 float64) *Vector {
	return &Vector{
		U: x1 - x0,
		V: y1 - y0,
	}
}

func (v *Vector) DotProduct(w *Vector) float64 {
	return v.U*w.U + v.V*w.V
}

func (v *Vector) CrossProduct(w *Vector) float64 {
	return v.U*w.V - v.V*w.U
}

func (v *Vector) Norm() float64 {
	return math.Hypot(v.U, v.V)
}

func (v *Vector) AngleValueTo(w *Vector) float64 {
	dotProduct := v.DotProduct(w)
	normProduct := v.Norm() * w.Norm()
	return math.Acos(dotProduct / normProduct)
}
func tIsValid(t float64) bool {
	return t >= 0.0 && t <= 1.0
}
