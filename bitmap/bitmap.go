// Package bitmap creates bitmap images.
package bitmap

import "math"

// Bitmap represents a bitmap image.
type Bitmap struct {
	Width, Height int
	Pixels        []float64
}

// NewBitmap creates a new Bitmap.
func NewBitmap(w, h int) *Bitmap {
	c := &Bitmap{
		w, h,
		make([]float64, w*h*3),
	}
	c.Clear(0, 0, 0)
	return c
}

// Clear clears the bitmap.
func (c *Bitmap) Clear(r, g, b float64) {
	for i := 0; i < len(c.Pixels); i += 3 {
		c.Pixels[i] = b
		c.Pixels[i+1] = g
		c.Pixels[i+2] = r
	}
}

// GetPixel returns the rgb values of the pixel at the given coords.
func (c *Bitmap) GetPixel(x, y int) (float64, float64, float64) {
	if x >= c.Width || x < 0 || y >= c.Height || y < 0 {
		return 0, 0, 0
	}
	index := (y*c.Width + x) * 3
	return c.Pixels[index+2], c.Pixels[index+1], c.Pixels[index]
}

// SetPixel sets the rgb values of the pixel at the given coords.
func (c *Bitmap) SetPixel(x, y int, r, g, b float64) {
	if x >= c.Width || x < 0 || y >= c.Height || y < 0 {
		return
	}
	index := (y*c.Width + x) * 3
	c.Pixels[index] = clamp(b)
	c.Pixels[index+1] = clamp(g)
	c.Pixels[index+2] = clamp(r)
}

// SetPixelGray sets the the pixel at the given coords to the gray value given.
func (c *Bitmap) SetPixelGray(x, y int, val float64) {
	c.SetPixel(x, y, val, val, val)
}

// SaveImage saves the bitmap to a file.
func (c *Bitmap) SaveImage(filename string) {
	EncodeBmp(c.Pixels, c.Width, -c.Height, filename)
}

// clamp clamps a channel value between 0 and 1.
func clamp(val float64) float64 {
	return math.Min(1.0, math.Max(0.0, val))
}
