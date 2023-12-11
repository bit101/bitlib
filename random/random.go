// Package random contains various methods for creating random values.
package random

import (
	"math"
	"math/rand"
	"time"

	"github.com/bit101/bitlib/blmath"
)

// You can use the random package in two ways now.
// 1. Call the methods directly as exported functions from the package.
// example:
//	 random.FloatRange(0, width)
//
// 2. Create an instance of random.Random and call methods on that object.
// example:
//	 r = random.NewRandom()
//   r.FloatRange(0, width)
//
// Using method 2, you can have different random generators with different seeds,
// unaffected by calls to the others.

// Random represents a single unique PRNG
type Random struct {
	rng *rand.Rand
}

// NewRandom creates a new Random instance.
func NewRandom() *Random {
	return &Random{rand.New(rand.NewSource(int64(time.Now().Nanosecond())))}
}

// rng is the default PRNG when method names are called directly via the package.
var rng = NewRandom()

// Seed sets the prng seed.
func Seed(seed int64) {
	rng.Seed(seed)
}

// RandSeed seeds the prng with a random seed.
func RandSeed() int64 {
	return rng.RandSeed()
}

// Float returns a random float from 0.0 to 1.0.
func Float() float64 {
	return rng.Float()
}

// Int returns a random integer.
func Int() int {
	return rng.Int()
}

// Angle returns a random angle from 0 to 2PI radians.
func Angle() float64 {
	return rng.Angle()
}

// FloatRange returns a random float from min to max.
func FloatRange(min float64, max float64) float64 {
	return rng.FloatRange(min, max)
}

// IntRange returns a random int from min to max.
func IntRange(min int, max int) int {
	return rng.IntRange(min, max)
}

// FloatArray returns an array of a given size filled with random floats from min to max.
func FloatArray(size int, min, max float64) []float64 {
	return rng.FloatArray(size, min, max)
}

// IntArray returns an array of a given size filled with random int from min to max.
func IntArray(size int, min, max int) []int {
	return rng.IntArray(size, min, max)
}

// Boolean returns a random boolean.
func Boolean() bool {
	return rng.Boolean()
}

// WeightedBool returns a weighted boolean.
func WeightedBool(weight float64) bool {
	return rng.WeightedBool(weight)
}

// Power returns a random number raised to a power.
func Power(min, max, power float64) float64 {
	return rng.Power(min, max, power)
}

// GaussRange returns a random number within a normal distribution (mostly) within a min/max range.
func GaussRange(min, max float64) float64 {
	return rng.GaussRange(min, max)
}

// Norm returns a random number within a normal distribution with the given mean and standard deviation.
func Norm(mean, std float64) float64 {
	return rng.Norm(mean, std)
}

// String returns a random string.
func String(length int) string {
	return rng.String(length)
}

// StringLower returns a random lower case string.
func StringLower(length int) string {
	return rng.StringLower(length)
}

// StringUpper returns a random upper case string.
func StringUpper(length int) string {
	return rng.StringUpper(length)
}

// StringAlpha returns a random string of letters.
func StringAlpha(length int) string {
	return rng.StringAlpha(length)
}

// ================================
// Instance methods
// ================================

// Seed sets the prng seed.
func (r *Random) Seed(seed int64) {
	r.rng = rand.New(rand.NewSource(seed))
}

// RandSeed seeds the prng with a random seed.
func (r *Random) RandSeed() int64 {
	seed := int64(time.Now().Nanosecond())
	r.rng = rand.New(rand.NewSource(seed))
	return seed
}

// Float returns a random float from 0.0 to 1.0.
func (r *Random) Float() float64 {
	return r.rng.Float64()
}

// Int returns a random integer.
func (r *Random) Int() int {
	return r.rng.Int()
}

// Angle returns a random angle from 0 to 2PI radians.
func (r *Random) Angle() float64 {
	return r.FloatRange(0, blmath.Tau)
}

// FloatRange returns a random float from min to max.
func (r *Random) FloatRange(min float64, max float64) float64 {
	return min + r.Float()*(max-min)
}

// IntRange returns a random int from min to max.
func (r *Random) IntRange(min int, max int) int {
	return int(r.FloatRange(float64(min), float64(max)))
}

// FloatArray returns an array of a given size filled with random floats from min to max.
func (r *Random) FloatArray(size int, min, max float64) []float64 {
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = r.FloatRange(min, max)
	}
	return arr
}

// IntArray returns an array of a given size filled with random int from min to max.
func (r *Random) IntArray(size int, min, max int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = r.IntRange(min, max)
	}
	return arr
}

// Boolean returns a random boolean.
func (r *Random) Boolean() bool {
	return r.WeightedBool(0.5)
}

// WeightedBool returns a weighted boolean.
func (r *Random) WeightedBool(weight float64) bool {
	return r.Float() < weight
}

// Power returns a random number raised to a power.
func (r *Random) Power(min, max, power float64) float64 {
	return min + math.Pow(r.Float(), power)*(max-min)
}

// GaussRange returns a random number within a normal distribution (mostly) within a min/max range.
func (r *Random) GaussRange(min, max float64) float64 {
	rng := (max - min) / 2.0
	mean := min + rng
	// a standard deviation of 1.0 will have 99.7% of its values between -3 and +3.
	// but for 100,000 samples, that's still around 300 points beyond that range.
	// 99.994% will be between -4 and 4.
	// for 100,000 samples, that's around 6 outside the range. better.
	// so we get the standard deviation by dividing the range by 4.0
	std := rng / 4.0
	return r.Norm(mean, std)
}

// Norm returns a random number within a normal distribution with the given mean and standard deviation.
func (r *Random) Norm(mean, std float64) float64 {
	return r.rng.NormFloat64()*std + mean
}

// String returns a random string.
func (r *Random) String(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(r.IntRange(33, 128))
		s += string(c)
	}
	return s
}

// StringLower returns a random lower case string.
func (r *Random) StringLower(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(r.IntRange(97, 123))
		s += string(c)
	}
	return s
}

// StringUpper returns a random upper case string.
func (r *Random) StringUpper(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(r.IntRange(65, 91))
		s += string(c)
	}
	return s
}

// StringAlpha returns a random string of letters.
func (r *Random) StringAlpha(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		c := rune(r.IntRange(65, 118))
		if c > 90 {
			c += 6
		}
		s += string(c)
	}
	return s
}

// StringFrom returns a random string of characters from the source string.
func (r *Random) StringFrom(length int, chars string) string {
	s := ""
	for i := 0; i < length; i++ {
		index := IntRange(0, len(chars))
		s += string(chars[index])
	}
	return s
}

// WeightedIndex returns a random int based on an array of weights.
func (r Random) WeightedIndex(weights []float64) int {
	total := 0.0
	for _, w := range weights {
		total += w
	}
	n := r.FloatRange(0, total)

	for i, w := range weights {
		if n < w {
			return i
		}
		n -= w
	}
	return -1 // should never happen
}
