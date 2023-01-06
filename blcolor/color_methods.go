// Package blcolor contains color creation and manipulation tools.
package blcolor

import "math"

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

// ToHSL returns the hue, saturation and lightness of a color.
// hue = 0-360, s = 0-1, l = 0-1
func (c Color) ToHSL() (float64, float64, float64) {
	r, g, b := c.R, c.G, c.B
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	mmRange := max - min

	var h, s, l float64

	// lightness
	l = (max + min) / 2

	if min == max {
		return h, s, l
	}

	// saturation
	if l <= 0.5 {
		s = mmRange / (max + min)
	} else {
		s = mmRange / (2.0 - max - min)
	}

	// hue
	if max == r {
		h = (g - b) / mmRange
	} else if max == g {
		h = 2.0 + (b-r)/mmRange
	} else {
		h = 4.0 + (r-g)/mmRange
	}

	return h * 60, s, l
}

// ToHSV returns the hue, saturation and value of a color.
// hue = 0-360, s = 0-1, v = 0-1
func (c Color) ToHSV() (float64, float64, float64) {
	r, g, b := c.R, c.G, c.B
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	mmRange := max - min
	h := -1.0
	s := -1.0

	if max == min {
		h = 0.0
	} else if max == r {
		h = math.Mod(60*((g-b)/mmRange)+360, 360)
	} else if max == g {
		h = math.Mod(60*((b-r)/mmRange)+120, 360)
	} else if max == b {
		h = math.Mod(60*((r-g)/mmRange)+240, 360)
	}
	if max == 0 {
		s = 0
	} else {
		s = (mmRange / max) * 100
	}

	v := max * 100
	return h, s, v
}
