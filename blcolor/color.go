// Package blcolor contains color creation and manipulation tools.
package blcolor

import (
	"math"
	"math/rand"

	"github.com/bit101/bitlib/blmath"
)

// This file contains only Color creation methods

// Color holds rgba values for a color.
type Color struct {
	R float64
	G float64
	B float64
	A float64
}

// RGB creates a new Color struct with rgb values from 0.0 to 1.0 each (a = 1.0).
func RGB(r float64, g float64, b float64) Color {
	return Color{r, g, b, 1.0}
}

// RGBA creates a new Color struct with rgba values from 0.0 to 1.0 each.
func RGBA(r float64, g float64, b float64, a float64) Color {
	return Color{r, g, b, a}
}

// Number creates a new Color struct with a 24-bit int 0xRRGGBB (a = 1.0).
func Number(value int) Color {
	r := (value >> 16 & 0xff)
	g := (value >> 8 & 0xff)
	b := (value & 0xff)
	return RGBHex(r, g, b)
}

// NumberWithAlpha creates a new Color struct with a 32-bit int 0xAARRGGBB.
func NumberWithAlpha(value int) Color {
	a := (value >> 24)
	r := (value >> 16 & 0xff)
	g := (value >> 8 & 0xff)
	b := (value & 0xff)
	return RGBAHex(r, g, b, a)
}

// Lerp creates a new color by interpolating between two other colors
func Lerp(colorA, colorB Color, t float64) Color {
	r := colorA.R + (colorB.R-colorA.R)*t
	g := colorA.G + (colorB.G-colorA.G)*t
	b := colorA.B + (colorB.B-colorA.B)*t
	a := colorA.A + (colorB.A-colorA.A)*t
	return RGBA(r, g, b, a)
}

// RGBHex creates a Color struct with rgb values from 0 to 255 (a = 255).
func RGBHex(r int, g int, b int) Color {
	return RGBAHex(r, g, b, 255)
}

// RGBAHex creates a Color struct with rgba values from 0 to 255 each.
func RGBAHex(r int, g int, b int, a int) Color {
	return RGBA(float64(r)/255.0, float64(g)/255.0, float64(b)/255.0, float64(a)/255.0)
}

// RandomRGB creates a color struct with random rgb values (a = 1.0).
func RandomRGB() Color {
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	return RGB(r, g, b)
}

// HSV creates a Color struct using hue (0.0 - 360.0), value (0.0 - 1.0) and value (0.0 - 1.0) (a = 1.0).
func HSV(h float64, s float64, v float64) Color {
	h = blmath.ModPos(h, 360)
	h = h / 360.0
	i := math.Floor(h * 6.0)
	f := h*6.0 - i
	p := v * (1.0 - s)
	q := v * (1.0 - f*s)
	t := v * (1.0 - (1.0-f)*s)
	switch int(i) % 6 {
	case 0:
		return RGB(v, t, p)
	case 1:
		return RGB(q, v, p)
	case 2:
		return RGB(p, v, t)
	case 3:
		return RGB(p, q, v)
	case 4:
		return RGB(t, p, v)
	case 5:
		return RGB(v, p, q)
	default:
		return RGB(0.0, 0.0, 0.0)
	}
}

// HSVA creates a Color struct using hue (0.0 - 360.0), value (0.0 - 1.0), value (0.0 - 1.0) and alpha (0.0 - 1.0).
func HSVA(h float64, s float64, v float64, a float64) Color {
	c := HSV(h, s, v)
	c.A = a
	return c
}

// HSL creates a Color struct using hue (0.0 - 360.0), value (0.0 - 1.0) and lightness (0.0 - 1.0)
func HSL(h, s, l float64) Color {
	k := func(n float64) float64 {
		return math.Mod(n+h/30, 12)
	}
	a := s * math.Min(l, 1-l)
	f := func(n float64) float64 {
		return l - a*math.Max(-1, math.Min(k(n)-3, math.Min(9-k(n), 1)))
	}
	return RGB(f(0), f(8), f(4))
}

// HSLA creates a Color struct using hue (0.0 - 360.0), value (0.0 - 1.0), lightness (0.0 - 1.0), and alpha (0.0 - 1.0)
func HSLA(h, s, l, a float64) Color {
	c := HSL(h, s, l)
	c.A = a
	return c
}

// CYMK creates a color struct using cyan, yellow, magenta and black.
func CMYK(c, m, y, k float64) Color {
	c /= 100.0
	m /= 100.0
	y /= 100.0
	k /= 100.0
	r := (1 - c) * (1 - k)
	g := (1 - m) * (1 - k)
	b := (1 - y) * (1 - k)
	return RGB(r, g, b)
}

// Grey creates a new Color struct with rgb all equal to the same value from 0.0 to 1.0 (a = 1.0).
func Grey(shade float64) Color {
	return RGB(shade, shade, shade)
}

// GreyHex creates a new Color struct with rgb all equal to the same value from 0 to 255 (a = 1.0).
func GreyHex(shade int) Color {
	return Grey(float64(shade) / 255.0)
}

// RandomGrey creates a Color struct with a random grey shade from 0.0 to 1.0 (a = 1.0).
func RandomGrey() Color {
	return RandomGreyRange(0.0, 1.0)
}

// RandomGreyRange creates a Color struct with a random grey shade from min to max (a = 1.0).
func RandomGreyRange(min float64, max float64) Color {
	return Grey(min + rand.Float64()*(max-min))
}
