// Package blcolor contains color creation and manipulation tools.
package blcolor

import (
	"math"

	"github.com/bit101/bitlib/blmath"
)

// Equals returns whether or not this color is equal to another channel.
func (c Color) Equals(colorB Color) bool {
	return blmath.Equalish(c.R, colorB.R, 0.00001) &&
		blmath.Equalish(c.G, colorB.G, 0.00001) &&
		blmath.Equalish(c.B, colorB.B, 0.00001) &&
		blmath.Equalish(c.A, colorB.A, 0.00001)
}

// ColorDiff returns the euclidian distance between two colors.
func (c Color) ColorDiff(colorB Color) float64 {
	r := colorB.R - c.R
	g := colorB.G - c.G
	b := colorB.B - c.B
	return math.Sqrt(r*r + g*g + b*b)
}

// ColorDiffPercept returns the euclidian distance between two colors.
func (c Color) ColorDiffPercept(colorB Color) float64 {
	r := colorB.R - c.R
	g := colorB.G - c.G
	b := colorB.B - c.B
	if (c.R+colorB.R)/2 < 0.5 {
		return math.Sqrt(2*r*r + 4*g*g + 3*b*b)
	}
	return math.Sqrt(3*r*r + 4*g*g + 2*b*b)
}

// Luminance returnss the Luminance of a given color
func (c Color) Luminance() float64 {
	adjust := func(val float64) float64 {
		if val <= 0.03928 {
			return val / 12.92
		}
		return math.Pow((val+0.055)/1.055, 2.4)
	}
	r := adjust(c.R)
	g := adjust(c.G)
	b := adjust(c.B)
	return r*0.2126 + g*0.7152 + b*0.0722
}

// Contrast returns the contrast between this color and another color.
// ref: https://www.w3.org/TR/2012/NOTE-WCAG20-TECHS-20120103/G17.html
func (c Color) Contrast(c2 Color) float64 {
	l1 := c.Luminance()
	l2 := c2.Luminance()
	if l1 < l2 {
		return (l2 + 0.05) / (l1 + 0.05)
	}
	return (l1 + 0.05) / (l2 + 0.05)
}

// Invert returns a color with the inverted RGB values of the original.
func (c Color) Invert() Color {
	return RGBA(1.0-c.R, 1.0-c.G, 1.0-c.B, c.A)
}

// Rotate returns a color with the hue rotated by the specified amount.
func (c Color) Rotate(degrees float64) Color {
	h, s, l := c.ToHSL()
	return HSLA(h+degrees, s, l, c.A)
}

// Scale scales the r, g, b channels by a percent.
// This will lighten or darken the color.
// Scale(1.05) will lighten the color by 5% of its current value.
// Scale(0.95) will darken the color by 5% of its current value.
// This is virtually the same as getting h, s, l and scaling the l
// But probably simpler in most cases.
func (c Color) Scale(mult float64) Color {
	// no change
	if mult == 1.0 {
		return c
	}

	// it's black
	if mult <= 0.0 {
		return RGBA(0, 0, 0, c.A)
	}

	// it's between 0 and 1. darken it.
	if mult < 1.0 {
		r := c.R * mult
		g := c.G * mult
		b := c.B * mult
		return RGBA(r, g, b, c.A)
	}

	// it's over 1. lighten it.
	r := c.R * mult
	g := c.G * mult
	b := c.B * mult

	// need to check if any channel has gone over 1.0 and adjust.
	max := math.Max(r, math.Max(g, b))
	if max <= 1.0 {
		// everything is in range. use the multiplied values.
		return RGBA(r, g, b, c.A)
	}

	if r >= 1.0 && g >= 1.0 && b >= 1.0 {
		// brightness is maxed out. return white.
		return RGBA(1, 1, 1, c.A)
	}

	// at least one channel has maxed out, but not all.
	// map each channel value from the range of avg->max to the range avg->1.0
	// this is really just blmath.Map, inlined and optimized
	avg := (r + g + b) / 3.0
	ratio := (1.0 - avg) / (max - avg)
	r = avg + ratio*(r-avg)
	g = avg + ratio*(g-avg)
	b = avg + ratio*(b-avg)

	return RGBA(r, g, b, c.A)
}

// ToGray returns a new grayscale color from this color.
func (c Color) ToGray() Color {
	// One would think using luminance as defined above would be better here.
	// But the result tends to be too dark.
	// The same goes for the "lightness" calculation in ToHSL below.
	// This works better.
	g := 0.3*c.R + 0.59*c.G + 0.11*c.B
	return RGBA(g, g, g, c.A)
}

// maxRGB is a helper method to get the largest rgb component.
func (c Color) maxRGB() float64 {
	return math.Max(c.R, math.Max(c.G, c.B))
}

// minRGB is a helper method to get the smallest rgb component.
func (c Color) minRGB() float64 {
	return math.Min(c.R, math.Min(c.G, c.B))
}

// ToHSL returns the hue, saturation and lightness of a color.
// hue = 0-360, s = 0-1, l = 0-1
func (c Color) ToHSL() (float64, float64, float64) {
	r, g, b := c.R, c.G, c.B
	max := c.maxRGB()
	min := c.minRGB()
	rng := max - min

	var h, s, l float64

	// lightness
	l = (max + min) / 2

	if min == max {
		return h, s, l
	}

	// saturation
	if l <= 0.5 {
		s = rng / (max + min)
	} else {
		s = rng / (2.0 - max - min)
	}

	// hue
	if max == r {
		h = (g - b) / rng
	} else if max == g {
		h = 2.0 + (b-r)/rng
	} else {
		h = 4.0 + (r-g)/rng
	}

	return h * 60, s, l
}

// ToHSV returns the hue, saturation and value of a color.
// hue = 0-360, s = 0-1, v = 0-1
func (c Color) ToHSV() (float64, float64, float64) {
	r, g, b := c.R, c.G, c.B
	max := c.maxRGB()
	min := c.minRGB()
	rng := max - min
	h := -1.0
	s := -1.0

	if max == min {
		h = 0.0
	} else if max == r {
		h = math.Mod(60*((g-b)/rng)+360, 360)
	} else if max == g {
		h = math.Mod(60*((b-r)/rng)+120, 360)
	} else if max == b {
		h = math.Mod(60*((r-g)/rng)+240, 360)
	}
	if max == 0 {
		s = 0
	} else {
		s = (rng / max) * 100
	}

	v := max * 100
	return h, s, v
}

// ToCMYK returns the cyan, magenta, yellow and black of a color.
func (c Color) ToCMYK() (float64, float64, float64, float64) {
	k := 1.0 - c.maxRGB()
	cy := (1.0 - c.R - k) / (1.0 - k)
	m := (1.0 - c.G - k) / (1.0 - k)
	y := (1.0 - c.B - k) / (1.0 - k)
	return cy * 100, m * 100, y * 100, k * 100
}

// Quant quantizes each color channel into a specified number of bands.
func (c Color) Quant(quant int) Color {
	return RGB(
		quantChannel(c.R, quant),
		quantChannel(c.G, quant),
		quantChannel(c.B, quant),
	)
}

// AddRGB adds a value to each channel, returning a new Color.
func (c Color) AddRGB(r, g, b float64) Color {
	return RGB(c.R+r, c.G+g, c.B+b)
}

// AddColor adds a Color to this Color, returning a new Color.
func (c Color) AddColor(b Color) Color {
	return RGB(c.R+b.R, c.G+b.G, c.B+b.B)
}

func quantChannel(value float64, quant int) float64 {
	return blmath.Clamp(math.Floor(value*float64(quant))/float64(quant-1), 0, 1)
}
