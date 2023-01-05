// Package blcolor contains color creation and manipulation tools.
package blcolor

import (
	"log"
	"sort"

	"github.com/bit101/bitlib/random"
)

// Palette is a list of colors.
type Palette struct {
	colors []Color
}

// NewPalette creates a new palette of the given size
func NewPalette() *Palette {
	return &Palette{
		[]Color{},
	}
}

// Size returns the number of colors in this palette.
func (p *Palette) Size() int {
	return len(p.colors)
}

// Get returns the color at the given index, if available.
func (p *Palette) Get(index int) Color {
	if index >= p.Size() || index < 0 {
		log.Fatalf("Can't get index %d for palette of size %d.", index, p.Size())
	}
	return p.colors[index]
}

// Add adds a color to the palette.
func (p *Palette) Add(color Color) {
	p.colors = append(p.colors, color)
}

// AddRGB adds a new color defined by the rgb channels.
func (p *Palette) AddRGB(r, g, b float64) {
	p.Add(RGB(r, g, b))
}

// AddRGBA adds a new color defined by the rgb channels.
func (p *Palette) AddRGBA(r, g, b, a float64) {
	p.Add(RGBA(r, g, b, a))
}

// GetRandom returns a random color from the palette.
func (p *Palette) GetRandom() Color {
	index := random.IntRange(0, p.Size())
	return p.Get(index)
}

// Len returns the size of the palette (used for sorting).
func (p *Palette) Len() int {
	return len(p.colors)
}

// Less returns if one color has lower luminance another (used for sorting).
func (p *Palette) Less(i, j int) bool {
	a := p.colors[i].Luminance()
	b := p.colors[j].Luminance()
	return a < b
}

// Swap swaps two colors in the palette (used for sorting).
func (p *Palette) Swap(i, j int) {
	p.colors[i], p.colors[j] = p.colors[j], p.colors[i]
}

// Sort sorts the palette based on luminance
func (p *Palette) Sort() {
	sort.Sort(p)
}
