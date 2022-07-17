package geom

import "github.com/bit101/bitlib/random"

type Size struct {
	Width, Height float64
}

func NewSize(w, h float64) *Size {
	return &Size{
		Width:  w,
		Height: h,
	}
}

func RandomSize(minW, maxW, minH, maxH float64) *Size {
	return NewSize(random.FloatRange(minW, maxW), random.FloatRange(minH, maxH))
}

func (s *Size) Equals(z *Size) bool {
	if s == z {
		return true
	}
	return AreClose(s.Width, z.Width) && AreClose(s.Height, z.Height)
}
