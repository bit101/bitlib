// Package blmath contains numeric and math related functions.
package blmath

import (
	"log"
	"math"
)

// TwoPi 2 pi
const TwoPi = math.Pi * 2

// Tau 2 pi
const Tau = math.Pi * 2

// HalfPi pi / 2
const HalfPi = math.Pi / 2

// Difference returns the absolute value of the difference between two numbers.
func Difference(a, b float64) float64 {
	return math.Abs(a - b)
}

// Norm returns a normalized value in a min/max range.
func Norm(value float64, min float64, max float64) float64 {
	return (value - min) / (max - min)
}

// Lerp is linear interpolation within a min/max range.
func Lerp(t float64, min float64, max float64) float64 {
	return min + (max-min)*t
}

// Map maps a value within one min/max range to a value within another range.
func Map(srcValue float64, srcMin float64, srcMax float64, dstMin float64, dstMax float64) float64 {
	norm := Norm(srcValue, srcMin, srcMax)
	return Lerp(norm, dstMin, dstMax)
}

// MapLinExp maps a value within a linear min/max range to another exponential range.
// min and max values will be swapped if min is greater than max.
// The resulting dstMin value cannot be 0.
func MapLinExp(srcValue, srcMin, srcMax, dstMin, dstMax float64) float64 {
	if dstMin == 0 {
		log.Fatal("blmath.MapLinExp: dstMin parameter cannot be 0")
	}
	if srcMin > srcMax {
		srcMin, srcMax = srcMax, srcMin
	}
	if dstMin > dstMax {
		dstMin, dstMax = dstMax, dstMin
	}
	// taken from SuperCollider's linexp function
	return math.Pow(dstMax/dstMin, (srcValue-srcMin)/(srcMax-srcMin)) * dstMin
}

// MapExpLin maps a value within an exponential min/max range to another linear range.
// min and max values will be swapped if min is greater than max.
// The resulting srcMin value cannot be 0.
func MapExpLin(srcValue, srcMin, srcMax, dstMin, dstMax float64) float64 {
	if srcMin == 0 {
		log.Fatal("blmath.MapExpLin: srcMin parameter cannot be 0")
	}
	if srcMin > srcMax {
		srcMin, srcMax = srcMax, srcMin
	}
	if dstMin > dstMax {
		dstMin, dstMax = dstMax, dstMin
	}
	// taken from SuperCollider's explin function
	return (math.Log(srcValue/srcMin))/(math.Log(srcMax/srcMin))*(dstMax-dstMin) + dstMin
}

// Wrap wraps a value around so it remains between min (inclusive) and max (exclusive).
func Wrap(value float64, min float64, max float64) float64 {
	rng := max - min
	return min + math.Mod((math.Mod(value-min, rng)+rng), rng)
}

// WrapTau wraps a number to be within 0 (inclusive) to 2 * Pi (exclusive).
func WrapTau(value float64) float64 {
	return math.Mod(math.Mod(value, Tau)+Tau, Tau)
}

// WrapPi wraps a number to be within -Pi (inclusive) and +Pi (exclusive).
func WrapPi(value float64) float64 {
	return Wrap(value, -math.Pi, math.Pi)
}

// Clamp enforces a value does not go beyond a min/max range.
func Clamp(value float64, min float64, max float64) float64 {
	// let min and max be reversed and still work.
	realMin := min
	realMax := max
	if min > max {
		realMin = max
		realMax = min
	}
	result := value
	if value < realMin {
		result = realMin
	}
	if value > realMax {
		result = realMax
	}
	return result
}

// RoundTo rounds a number to the nearest decimal value.
func RoundTo(value float64, decimal int) float64 {
	mult := math.Pow(10.0, float64(decimal))
	return math.Round(value*mult) / mult
}

// RoundToNearest rounds a number to the nearest multiple of a value.
func RoundToNearest(value float64, mult float64) float64 {
	return math.Round(value/mult) * mult
}

// SinRange returns the sin of an angle mapped to a min/max range.
func SinRange(angle float64, min float64, max float64) float64 {
	return Map(math.Sin(angle), -1, 1, min, max)
}

// CosRange returns the cos of an angle mapped to a min/max range.
func CosRange(angle float64, min float64, max float64) float64 {
	return Map(math.Cos(angle), -1, 1, min, max)
}

// Fract returns the fractional part of a floating point number.
func Fract(n float64) float64 {
	if n < 0 {
		return n + math.Floor(math.Abs(n))
	}
	return n - math.Floor(n)
}

// LerpSin maps a normal value to min and max values with a sine wave.
func LerpSin(value, min, max float64) float64 {
	// subtracting pi/2 makes value 0 and 1 return min. value 0.5 will return max.
	return SinRange(value*math.Pi*2-math.Pi/2, min, max)
}

// Equalish returns whether the two values are approximately equal.
func Equalish(a float64, b float64, delta float64) bool {
	return math.Abs(a-b) <= delta
}

// ComplexImagAbs returns a complex number with the real component and the abolute value of the imaginary component.
// Useful for certain types of fractals, such as "duck fractals"
func ComplexImagAbs(z complex128) complex128 {
	if imag(z) < 0 {
		return complex(real(z), -imag(z))
	}
	return z
}

// ComplexMagnitude returns the magnitude of a complex number
func ComplexMagnitude(z complex128) float64 {
	return math.Hypot(real(z), imag(z))
}

// Gamma increases/decreases the given value by an amount specified in gamma.
// Usually val is a pixel brightness value from 0.0 - 1.0.
// gamma of 1.0 makes no change. Higher is brighter, lower is darker.
func Gamma(val, gamma float64) float64 {
	return math.Pow(val, 1/gamma)
}

// GCD returns the greatest common denominator of two integers.
func GCD(a, b int) int {
	result := Min(Abs(a), Abs(b))

	for result > 0 {
		if a%result == 0 && b%result == 0 {
			break
		}
		result--
	}
	return result
}

// LCM returns the least common multiple of two integers.
func LCM(x, y int) int {
	x = Abs(x)
	y = Abs(y)
	return x * y / GCD(x, y)
}

// Simplify reduces an int/int fraction to its simplest form.
func Simplify(x, y int) (int, int) {
	g := GCD(x, y)
	return x / g, y / g
}

// Number can be an int or float
type Number interface {
	int | int64 | float32 | float64
}

// Abs is a generic absolute value function.
func Abs[T Number](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

// Min is a generic min value function
func Min[T Number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

// Max is a generic max value function
func Max[T Number](a, b T) T {
	if a < b {
		return b
	}
	return a
}

// ModPos computes a % b for float64, and ensures the result is positive
func ModPos(a, b float64) float64 {
	val := math.Mod(a, b)
	if (val < 0 && b > 0) || (val > 0 && b < 0) {
		val += b
	}
	return val
}

// ModPosInt computes a % b for int, and ensures the result is positive
func ModPosInt(a, b int) int {
	val := a % b
	if (val < 0 && b > 0) || (val > 0 && b < 0) {
		val += b
	}
	return val
}
