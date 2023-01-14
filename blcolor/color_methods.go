// Package blcolor contains color creation and manipulation tools.
package blcolor

import (
	"math"
)

// ColorDiff returns the euclidian distance between two colors.
func (c Color) ColorDiff(colorB Color) float64 {
	r := colorB.R - c.R
	g := colorB.G - c.G
	b := colorB.B - c.B
	return math.Sqrt(r*r + g*g + b*b)
}

// ColorDiffPercept returns the euclidian distance between two colors.
func (c Color) ColorDiffPercept(colorB Color) float64 {
	rr := (c.R + colorB.R) / 2
	r := colorB.R - c.R
	g := colorB.G - c.G
	b := colorB.B - c.B
	return math.Sqrt((2+rr)*r*r + 4*g*g + (3-rr)*b*b)
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

// Lighten lightens a color by a percent.
// i.e. Lighten(0.05) will lighten the color by 5% of its current value.
func (c Color) Lighten(mult float64) Color {
	if mult <= 0 {
		return c.Darken(-mult)
	}
	mult += 1.0

	r := c.R * mult
	g := c.G * mult
	b := c.B * mult
	max := math.Max(r, math.Max(g, b))
	if max <= 1.0 {
		return RGBA(r, g, b, c.A)
	}
	total := r + g + b
	if total > 3.0 {
		return RGBA(1, 1, 1, c.A)
	}
	x := (3.0 - total) / (3.0*max - total)
	gray := 1.0 - x*max
	return RGBA(gray+x*r, gray+x*g, gray+x*b, c.A)
}

// Darken darkens a color by a percent.
// i.e. Darken(0.05) will darken the color by 5% of its current value.
func (c Color) Darken(mult float64) Color {
	if mult <= 0 {
		return c.Lighten(-mult)
	}
	mult = 1.0 - mult

	r := c.R * mult
	g := c.G * mult
	b := c.B * mult
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
