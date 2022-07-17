package geom

import (
	"log"
	"math"
)

type OpenInterval struct {
	Start, End float64
}

func NewOpenInterval(start, end float64) *OpenInterval {
	if start > end {
		log.Fatalln("OpenInterval: start should be smaller than end.")
	}
	return &OpenInterval{
		Start: start,
		End:   end,
	}
}

func (o *OpenInterval) Length() float64 {
	return o.End - o.Start
}

func (o *OpenInterval) Contains(value float64) bool {
	return o.Start < value && value < o.End
}

func (o *OpenInterval) Overlaps(p *OpenInterval) bool {
	if AreClose(o.Start, p.Start) && AreClose(o.End, p.End) {
		return true
	}
	return o.Contains(p.Start) || o.Contains(p.End) || p.Contains(o.Start) || p.Contains(o.End)
}

func (o *OpenInterval) ComputeOverlap(p *OpenInterval) *OpenInterval {
	if !o.Overlaps(p) {
		return nil
	}
	return NewOpenInterval(math.Max(o.Start, p.Start), math.Min(o.End, p.End))
}
