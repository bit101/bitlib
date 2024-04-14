// Package geom has geometry related structs and funcs.
package geom

// SegmentList is a slice of segments
type SegmentList []*Segment

// NewSegmentList creates a new segment list
func NewSegmentList() SegmentList {
	return SegmentList{}
}

// Add adds a segment to the list
func (s *SegmentList) Add(segment *Segment) {
	*s = append(*s, segment)
}

// AddXY adds a segment to the list
func (s *SegmentList) AddXY(x0, y0, x1, y1 float64) {
	*s = append(*s, NewSegment(x0, y0, x1, y1))
}

// Rotate rotates all the segments in a list.
func (s *SegmentList) Rotate(angle float64) {
	for _, segment := range *s {
		segment.Rotate(angle)
	}
}

// RotateLocal rotates each segment from its own center.
func (s *SegmentList) RotateLocal(angle float64) {
	for _, segment := range *s {
		segment.RotateLocal(angle)
	}
}

// RotateFrom rotates all the segments in a list using the x, y location as a center.
func (s *SegmentList) RotateFrom(x, y float64, angle float64) {
	for _, segment := range *s {
		segment.RotateFrom(x, y, angle)
	}
}

// Translate translates all the segments in a list.
func (s *SegmentList) Translate(x, y float64) {
	for _, segment := range *s {
		segment.Translate(x, y)
	}
}

// Scale scales all the segments in a list.
func (s *SegmentList) Scale(sx, sy float64) {
	for _, segment := range *s {
		segment.Scale(sx, sy)
	}
}

// ScaleLocal scales each segment from its own center.
func (s *SegmentList) ScaleLocal(sx, sy float64) {
	for _, segment := range *s {
		segment.ScaleLocal(sx, sy)
	}
}

// ScaleFrom scales all the segments in a list using the x, y location as a center.
func (s *SegmentList) ScaleFrom(x, y, sx, sy float64) {
	for _, segment := range *s {
		segment.ScaleFrom(x, y, sx, sy)
	}
}

// Randomize randomizes the position of both ends of the segment by up to a given amount on each axis.
func (s *SegmentList) Randomize(amount float64) {
	for _, seg := range *s {
		seg.Randomize(amount)
	}
}
