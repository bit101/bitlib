package geom

import "math"

func IsCloseToOne(val float64) bool {
	return AreClose(val, 1.0)
}

func IsCloseToZero(val float64) bool {
	return AreClose(val, 0)
}

func AreClose(val0, val1 float64) bool {
	return math.Abs(val0-val1) < 1e-5
}

func tIsValid(t float64) bool {
	return t >= 0.0 && t <= 1.0
}
