// Package blcolor contains color creation and manipulation tools.
package blcolor

import (
	"math"
	"math/rand"

	"github.com/bit101/bitlib/blmath"
)

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

// HSLA creates a Color struct using hue (0.0 - 360.0), value (0.0 - 1.0), lightness (0.0 - 1.0), and alpha (0.0 - 1.0)
func HSLA(h, s, l, alpha float64) Color {
	k := func(n float64) float64 {
		return math.Mod(n+h/30, 12)
	}
	a := s * math.Min(l, 1-l)
	f := func(n float64) float64 {
		return l - a*math.Max(-1, math.Min(k(n)-3, math.Min(9-k(n), 1)))
	}
	return RGBA(f(0), f(8), f(4), alpha)
}

// HSL creates a Color struct using hue (0.0 - 360.0), value (0.0 - 1.0) and lightness (0.0 - 1.0)
func HSL(h, s, l float64) Color {
	return HSLA(h, s, l, 1)
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

// Blueviolet creates a Blueviolet Color struct.
func Blueviolet() Color { return RGBHex(138, 43, 226) }

// Brown creates a Brown Color struct.
func Brown() Color { return RGBHex(165, 42, 42) }

// Aliceblue creates a Aliceblue Color struct.
func Aliceblue() Color { return RGBHex(240, 248, 255) }

// Antiquewhite creates a Antiquewhite Color struct.
func Antiquewhite() Color { return RGBHex(250, 235, 215) }

// Aqua creates a Aqua Color struct.
func Aqua() Color { return RGBHex(0, 255, 255) }

// Aquamarine creates a Aquamarine Color struct.
func Aquamarine() Color { return RGBHex(127, 255, 212) }

// Azure creates a Azure Color struct.
func Azure() Color { return RGBHex(240, 255, 255) }

// Beige creates a Beige Color struct.
func Beige() Color { return RGBHex(245, 245, 220) }

// Bisque creates a Bisque Color struct.
func Bisque() Color { return RGBHex(255, 228, 196) }

// Black creates a Black Color struct.
func Black() Color { return RGBHex(0, 0, 0) }

// Blanchedalmond creates a Blanchedalmond Color struct.
func Blanchedalmond() Color { return RGBHex(255, 235, 205) }

// Blue creates a Blue Color struct.
func Blue() Color { return RGBHex(0, 0, 255) }

// Burlywood creates a Burlywood Color struct.
func Burlywood() Color { return RGBHex(222, 184, 135) }

// Cadetblue creates a Cadetblue Color struct.
func Cadetblue() Color { return RGBHex(95, 158, 160) }

// Chartreuse creates a Chartreuse Color struct.
func Chartreuse() Color { return RGBHex(127, 255, 0) }

// Chocolate creates a Chocolate Color struct.
func Chocolate() Color { return RGBHex(210, 105, 30) }

// Coral creates a Coral Color struct.
func Coral() Color { return RGBHex(255, 127, 80) }

// Cornflowerblue creates a Cornflowerblue Color struct.
func Cornflowerblue() Color { return RGBHex(100, 149, 237) }

// Cornsilk creates a Cornsilk Color struct.
func Cornsilk() Color { return RGBHex(255, 248, 220) }

// Crimson creates a Crimson Color struct.
func Crimson() Color { return RGBHex(220, 20, 60) }

// Cyan creates a Cyan Color struct.
func Cyan() Color { return RGBHex(0, 255, 255) }

// Darkblue creates a Darkblue Color struct.
func Darkblue() Color { return RGBHex(0, 0, 139) }

// Darkcyan creates a Darkcyan Color struct.
func Darkcyan() Color { return RGBHex(0, 139, 139) }

// Darkgoldenrod creates a Darkgoldenrod Color struct.
func Darkgoldenrod() Color { return RGBHex(184, 134, 11) }

// Darkgray creates a Darkgray Color struct.
func Darkgray() Color { return RGBHex(169, 169, 169) }

// Darkgreen creates a Darkgreen Color struct.
func Darkgreen() Color { return RGBHex(0, 100, 0) }

// Darkgrey creates a Darkgrey Color struct.
func Darkgrey() Color { return RGBHex(169, 169, 169) }

// Darkkhaki creates a Darkkhaki Color struct.
func Darkkhaki() Color { return RGBHex(189, 183, 107) }

// Darkmagenta creates a Darkmagenta Color struct.
func Darkmagenta() Color { return RGBHex(139, 0, 139) }

// Darkolivegreen creates a Darkolivegreen Color struct.
func Darkolivegreen() Color { return RGBHex(85, 107, 47) }

// Darkorange creates a Darkorange Color struct.
func Darkorange() Color { return RGBHex(255, 140, 0) }

// Darkorchid creates a Darkorchid Color struct.
func Darkorchid() Color { return RGBHex(153, 50, 204) }

// Darkred creates a Darkred Color struct.
func Darkred() Color { return RGBHex(139, 0, 0) }

// Darksalmon creates a Darksalmon Color struct.
func Darksalmon() Color { return RGBHex(233, 150, 122) }

// Darkseagreen creates a Darkseagreen Color struct.
func Darkseagreen() Color { return RGBHex(143, 188, 143) }

// Darkslateblue creates a Darkslateblue Color struct.
func Darkslateblue() Color { return RGBHex(72, 61, 139) }

// Darkslategray creates a Darkslategray Color struct.
func Darkslategray() Color { return RGBHex(47, 79, 79) }

// Darkslategrey creates a Darkslategrey Color struct.
func Darkslategrey() Color { return RGBHex(47, 79, 79) }

// Darkturquoise creates a Darkturquoise Color struct.
func Darkturquoise() Color { return RGBHex(0, 206, 209) }

// Darkviolet creates a Darkviolet Color struct.
func Darkviolet() Color { return RGBHex(148, 0, 211) }

// Deeppink creates a Deeppink Color struct.
func Deeppink() Color { return RGBHex(255, 20, 147) }

// Deepskyblue creates a Deepskyblue Color struct.
func Deepskyblue() Color { return RGBHex(0, 191, 255) }

// Dimgray creates a Dimgray Color struct.
func Dimgray() Color { return RGBHex(105, 105, 105) }

// Dimgrey creates a Dimgrey Color struct.
func Dimgrey() Color { return RGBHex(105, 105, 105) }

// Dodgerblue creates a Dodgerblue Color struct.
func Dodgerblue() Color { return RGBHex(30, 144, 255) }

// Firebrick creates a Firebrick Color struct.
func Firebrick() Color { return RGBHex(178, 34, 34) }

// Floralwhite creates a Floralwhite Color struct.
func Floralwhite() Color { return RGBHex(255, 250, 240) }

// Forestgreen creates a Forestgreen Color struct.
func Forestgreen() Color { return RGBHex(34, 139, 34) }

// Fuchsia creates a Fuchsia Color struct.
func Fuchsia() Color { return RGBHex(255, 0, 255) }

// Gainsboro creates a Gainsboro Color struct.
func Gainsboro() Color { return RGBHex(220, 220, 220) }

// Ghostwhite creates a Ghostwhite Color struct.
func Ghostwhite() Color { return RGBHex(248, 248, 255) }

// Gold creates a Gold Color struct.
func Gold() Color { return RGBHex(255, 215, 0) }

// Goldenrod creates a Goldenrod Color struct.
func Goldenrod() Color { return RGBHex(218, 165, 32) }

// Gray creates a Gray Color struct.
func Gray() Color { return RGBHex(128, 128, 128) }

// Green creates a Green Color struct.
func Green() Color { return RGBHex(0, 128, 0) }

// Greenyellow creates a Greenyellow Color struct.
func Greenyellow() Color { return RGBHex(173, 255, 47) }

// Honeydew creates a Honeydew Color struct.
func Honeydew() Color { return RGBHex(240, 255, 240) }

// Hotpink creates a Hotpink Color struct.
func Hotpink() Color { return RGBHex(255, 105, 180) }

// Indianred creates a Indianred Color struct.
func Indianred() Color { return RGBHex(205, 92, 92) }

// Indigo creates a Indigo Color struct.
func Indigo() Color { return RGBHex(75, 0, 130) }

// Ivory creates a Ivory Color struct.
func Ivory() Color { return RGBHex(255, 255, 240) }

// Khaki creates a Khaki Color struct.
func Khaki() Color { return RGBHex(240, 230, 140) }

// Lavender creates a Lavender Color struct.
func Lavender() Color { return RGBHex(230, 230, 250) }

// Lavenderblush creates a Lavenderblush Color struct.
func Lavenderblush() Color { return RGBHex(255, 240, 245) }

// Lawngreen creates a Lawngreen Color struct.
func Lawngreen() Color { return RGBHex(124, 252, 0) }

// Lemonchiffon creates a Lemonchiffon Color struct.
func Lemonchiffon() Color { return RGBHex(255, 250, 205) }

// Lightblue creates a Lightblue Color struct.
func Lightblue() Color { return RGBHex(173, 216, 230) }

// Lightcoral creates a Lightcoral Color struct.
func Lightcoral() Color { return RGBHex(240, 128, 128) }

// Lightcyan creates a Lightcyan Color struct.
func Lightcyan() Color { return RGBHex(224, 255, 255) }

// Lightgoldenrodyellow creates a Lightgoldenrodyellow Color struct.
func Lightgoldenrodyellow() Color { return RGBHex(250, 250, 210) }

// Lightgray creates a Lightgray Color struct.
func Lightgray() Color { return RGBHex(211, 211, 211) }

// Lightgreen creates a Lightgreen Color struct.
func Lightgreen() Color { return RGBHex(144, 238, 144) }

// Lightgrey creates a Lightgrey Color struct.
func Lightgrey() Color { return RGBHex(211, 211, 211) }

// Lightpink creates a Lightpink Color struct.
func Lightpink() Color { return RGBHex(255, 182, 193) }

// Lightsalmon creates a Lightsalmon Color struct.
func Lightsalmon() Color { return RGBHex(255, 160, 122) }

// Lightseagreen creates a Lightseagreen Color struct.
func Lightseagreen() Color { return RGBHex(32, 178, 170) }

// Lightskyblue creates a Lightskyblue Color struct.
func Lightskyblue() Color { return RGBHex(135, 206, 250) }

// Lightslategray creates a Lightslategray Color struct.
func Lightslategray() Color { return RGBHex(119, 136, 153) }

// Lightslategrey creates a Lightslategrey Color struct.
func Lightslategrey() Color { return RGBHex(119, 136, 153) }

// Lightsteelblue creates a Lightsteelblue Color struct.
func Lightsteelblue() Color { return RGBHex(176, 196, 222) }

// Lightyellow creates a Lightyellow Color struct.
func Lightyellow() Color { return RGBHex(255, 255, 224) }

// Lime creates a Lime Color struct.
func Lime() Color { return RGBHex(0, 255, 0) }

// Limegreen creates a Limegreen Color struct.
func Limegreen() Color { return RGBHex(50, 205, 50) }

// Linen creates a Linen Color struct.
func Linen() Color { return RGBHex(250, 240, 230) }

// Magenta creates a Magenta Color struct.
func Magenta() Color { return RGBHex(255, 0, 255) }

// Maroon creates a Maroon Color struct.
func Maroon() Color { return RGBHex(128, 0, 0) }

// Mediumaquamarine creates a Mediumaquamarine Color struct.
func Mediumaquamarine() Color { return RGBHex(102, 205, 170) }

// Mediumblue creates a Mediumblue Color struct.
func Mediumblue() Color { return RGBHex(0, 0, 205) }

// Mediumorchid creates a Mediumorchid Color struct.
func Mediumorchid() Color { return RGBHex(186, 85, 211) }

// Mediumpurple creates a Mediumpurple Color struct.
func Mediumpurple() Color { return RGBHex(147, 112, 219) }

// Mediumseagreen creates a Mediumseagreen Color struct.
func Mediumseagreen() Color { return RGBHex(60, 179, 113) }

// Mediumslateblue creates a Mediumslateblue Color struct.
func Mediumslateblue() Color { return RGBHex(123, 104, 238) }

// Mediumspringgreen creates a Mediumspringgreen Color struct.
func Mediumspringgreen() Color { return RGBHex(0, 250, 154) }

// Mediumturquoise creates a Mediumturquoise Color struct.
func Mediumturquoise() Color { return RGBHex(72, 209, 204) }

// Mediumvioletred creates a Mediumvioletred Color struct.
func Mediumvioletred() Color { return RGBHex(199, 21, 133) }

// Midnightblue creates a Midnightblue Color struct.
func Midnightblue() Color { return RGBHex(25, 25, 112) }

// Mintcream creates a Mintcream Color struct.
func Mintcream() Color { return RGBHex(245, 255, 250) }

// Mistyrose creates a Mistyrose Color struct.
func Mistyrose() Color { return RGBHex(255, 228, 225) }

// Moccasin creates a Moccasin Color struct.
func Moccasin() Color { return RGBHex(255, 228, 181) }

// Navajowhite creates a Navajowhite Color struct.
func Navajowhite() Color { return RGBHex(255, 222, 173) }

// Navy creates a Navy Color struct.
func Navy() Color { return RGBHex(0, 0, 128) }

// Oldlace creates a Oldlace Color struct.
func Oldlace() Color { return RGBHex(253, 245, 230) }

// Olive creates a Olive Color struct.
func Olive() Color { return RGBHex(128, 128, 0) }

// Olivedrab creates a Olivedrab Color struct.
func Olivedrab() Color { return RGBHex(107, 142, 35) }

// Orange creates a Orange Color struct.
func Orange() Color { return RGBHex(255, 165, 0) }

// Orangered creates a Orangered Color struct.
func Orangered() Color { return RGBHex(255, 69, 0) }

// Orchid creates a Orchid Color struct.
func Orchid() Color { return RGBHex(218, 112, 214) }

// Palegoldenrod creates a Palegoldenrod Color struct.
func Palegoldenrod() Color { return RGBHex(238, 232, 170) }

// Palegreen creates a Palegreen Color struct.
func Palegreen() Color { return RGBHex(152, 251, 152) }

// Paleturquoise creates a Paleturquoise Color struct.
func Paleturquoise() Color { return RGBHex(175, 238, 238) }

// Palevioletred creates a Palevioletred Color struct.
func Palevioletred() Color { return RGBHex(219, 112, 147) }

// Papayawhip creates a Papayawhip Color struct.
func Papayawhip() Color { return RGBHex(255, 239, 213) }

// Peachpuff creates a Peachpuff Color struct.
func Peachpuff() Color { return RGBHex(255, 218, 185) }

// Peru creates a Peru Color struct.
func Peru() Color { return RGBHex(205, 133, 63) }

// Pink creates a Pink Color struct.
func Pink() Color { return RGBHex(255, 192, 203) }

// Plum creates a Plum Color struct.
func Plum() Color { return RGBHex(221, 160, 221) }

// Powderblue creates a Powderblue Color struct.
func Powderblue() Color { return RGBHex(176, 224, 230) }

// Purple creates a Purple Color struct.
func Purple() Color { return RGBHex(128, 0, 128) }

// Rebeccapurple creates a Rebeccapurple Color struct.
func Rebeccapurple() Color { return RGBHex(102, 51, 153) }

// Red creates a Red Color struct.
func Red() Color { return RGBHex(255, 0, 0) }

// Rosybrown creates a Rosybrown Color struct.
func Rosybrown() Color { return RGBHex(188, 143, 143) }

// Royalblue creates a Royalblue Color struct.
func Royalblue() Color { return RGBHex(65, 105, 225) }

// Saddlebrown creates a Saddlebrown Color struct.
func Saddlebrown() Color { return RGBHex(139, 69, 19) }

// Salmon creates a Salmon Color struct.
func Salmon() Color { return RGBHex(250, 128, 114) }

// Sandybrown creates a Sandybrown Color struct.
func Sandybrown() Color { return RGBHex(244, 164, 96) }

// Seagreen creates a Seagreen Color struct.
func Seagreen() Color { return RGBHex(46, 139, 87) }

// Seashell creates a Seashell Color struct.
func Seashell() Color { return RGBHex(255, 245, 238) }

// Sienna creates a Sienna Color struct.
func Sienna() Color { return RGBHex(160, 82, 45) }

// Silver creates a Silver Color struct.
func Silver() Color { return RGBHex(192, 192, 192) }

// Skyblue creates a Skyblue Color struct.
func Skyblue() Color { return RGBHex(135, 206, 235) }

// Slateblue creates a Slateblue Color struct.
func Slateblue() Color { return RGBHex(106, 90, 205) }

// Slategray creates a Slategray Color struct.
func Slategray() Color { return RGBHex(112, 128, 144) }

// Slategrey creates a Slategrey Color struct.
func Slategrey() Color { return RGBHex(112, 128, 144) }

// Snow creates a Snow Color struct.
func Snow() Color { return RGBHex(255, 250, 250) }

// Springgreen creates a Springgreen Color struct.
func Springgreen() Color { return RGBHex(0, 255, 127) }

// Steelblue creates a Steelblue Color struct.
func Steelblue() Color { return RGBHex(70, 130, 180) }

// Tan creates a Tan Color struct.
func Tan() Color { return RGBHex(210, 180, 140) }

// Teal creates a Teal Color struct.
func Teal() Color { return RGBHex(0, 128, 128) }

// Thistle creates a Thistle Color struct.
func Thistle() Color { return RGBHex(216, 191, 216) }

// Tomato creates a Tomato Color struct.
func Tomato() Color { return RGBHex(255, 99, 71) }

// Turquoise creates a Turquoise Color struct.
func Turquoise() Color { return RGBHex(64, 224, 208) }

// Violet creates a Violet Color struct.
func Violet() Color { return RGBHex(238, 130, 238) }

// Wheat creates a Wheat Color struct.
func Wheat() Color { return RGBHex(245, 222, 179) }

// White creates a White Color struct.
func White() Color { return RGBHex(255, 255, 255) }

// Whitesmoke creates a Whitesmoke Color struct.
func Whitesmoke() Color { return RGBHex(245, 245, 245) }

// Yellow creates a Yellow Color struct.
func Yellow() Color { return RGBHex(255, 255, 0) }

// Yellowgreen creates a Yellowgreen Color struct.
func Yellowgreen() Color { return RGBHex(154, 205, 50) }

//////////////////////////////////////
// Color object methods
//////////////////////////////////////

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
