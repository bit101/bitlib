// Package geom has geometry related structs and funcs.
package geom

import "math"

// PolyInnerRadius returns the radius of the largest circle
// that can fit inside the given regular polygon.
func PolyInnerRadius(radius float64, sides int) float64 {
	angle := math.Pi / float64(sides)
	return math.Cos(angle) * radius
}
