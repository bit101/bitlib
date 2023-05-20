// Package random contains various methods for creating random values.
package random

import (
	"math"
	"math/rand"
	"time"

	"github.com/bit101/bitlib/blmath"
)

// Seed sets the prng seed.
func Seed(seed int64) {
	rand.Seed(seed)
}

// RandSeed seeds the prng with a random seed.
func RandSeed() int64 {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)
	return seed
}

// Float returns a random float from 0.0 to 1.0.
func Float() float64 {
	return rand.Float64()
}

// Int returns a random integer.
func Int() int {
	return rand.Int()
}

// Angle returns a random angle from 0 to 2PI radians.
func Angle() float64 {
	return FloatRange(0, blmath.Tau)
}

// FloatRange returns a random float from min to max.
func FloatRange(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// IntRange returns a random int from min to max.
func IntRange(min int, max int) int {
	return int(FloatRange(float64(min), float64(max)))
}

// FloatArray returns an array of a given size filled with random floats from min to max.
func FloatArray(size int, min, max float64) []float64 {
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = FloatRange(min, max)
	}
	return arr
}

// IntArray returns an array of a given size filled with random int from min to max.
func IntArray(size int, min, max int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = IntRange(min, max)
	}
	return arr
}

// Boolean returns a random boolean.
func Boolean() bool {
	return WeightedBool(0.5)
}

// WeightedBool returns a weighted boolean.
func WeightedBool(weight float64) bool {
	return rand.Float64() < weight
}

// Power returns a random number raised to a power.
func Power(min, max, power float64) float64 {
	return min + math.Pow(Float(), power)*(max-min)
}

// GaussRange returns a random number within a normal distribution (mostly) within a min/max range.
func GaussRange(min, max float64) float64 {
	rng := (max - min) / 2.0
	mean := min + rng
	// a standard deviation of 1.0 will have 99.7% of its values between -3 and +3.
	// but for 100,000 samples, that's still around 300 points beyond that range.
	// 99.994% will be between -4 and 4.
	// for 100,000 samples, that's around 6 outside the range. better.
	// so we get the standard deviation by dividing the range by 4.0
	std := rng / 4.0
	return rand.NormFloat64()*std + mean
}

// Norm returns a random number within a normal distrubution with the given mean and standard deviation.
func Norm(mean, std float64) float64 {
	return rand.NormFloat64()*std + mean
}

// String returns a random string.
func String(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(IntRange(33, 128))
		s += string(c)
	}
	return s
}

// StringLower returns a random lower case string.
func StringLower(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(IntRange(97, 123))
		s += string(c)
	}
	return s
}

// StringUpper returns a random upper case string.
func StringUpper(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(IntRange(65, 91))
		s += string(c)
	}
	return s
}

// StringAlpha returns a random string of letters.
func StringAlpha(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(IntRange(65, 118))
		if c > 90 {
			c += 6
		}
		s += string(c)
	}
	return s
}
