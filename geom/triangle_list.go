// Package geom has geometry related structs and funcs.
package geom

// TriangleList is a slice of triangles
type TriangleList []*Triangle

// NewTriangleList creates a new triangle list
func NewTriangleList() TriangleList {
	return TriangleList{}
}

// Add adds a Triangle to the list
func (t *TriangleList) Add(triangle *Triangle) {
	*t = append(*t, triangle)
}

// Rotate rotates all the segments in a list.
func (t *TriangleList) Rotate(angle float64) {
	for _, segment := range *t {
		segment.Rotate(angle)
	}
}

// RotateLocal rotates each segment from its own center.
func (t *TriangleList) RotateLocal(angle float64) {
	for _, segment := range *t {
		segment.RotateLocal(angle)
	}
}

// RotateFrom rotates all the segments in a list using the x, y location as a center.
func (t *TriangleList) RotateFrom(x, y float64, angle float64) {
	for _, segment := range *t {
		segment.RotateFrom(x, y, angle)
	}
}

// Translate translates all the segments in a list.
func (t *TriangleList) Translate(x, y float64) {
	for _, segment := range *t {
		segment.Translate(x, y)
	}
}

// Scale scales all the segments in a list.
func (t *TriangleList) Scale(sx, sy float64) {
	for _, segment := range *t {
		segment.Scale(sx, sy)
	}
}

// ScaleLocal scales each segment from its own center.
func (t *TriangleList) ScaleLocal(sx, sy float64) {
	for _, segment := range *t {
		segment.ScaleLocal(sx, sy)
	}
}

// ScaleFrom scales all the segments in a list using the x, y location as a center.
func (t *TriangleList) ScaleFrom(x, y, sx, sy float64) {
	for _, segment := range *t {
		segment.ScaleFrom(x, y, sx, sy)
	}
}

// Randomize randomizes the position of both ends of the segment by up to a given amount on each axis.
func (t *TriangleList) Randomize(amount float64) {
	for _, seg := range *t {
		seg.Randomize(amount)
	}
}

// Edges returns a list of unique edges from the triangle list.
func (t *TriangleList) Edges() SegmentList {
	edges := NewSegmentList()
	for _, t := range *t {
		sides := t.Edges()
		for _, side := range sides {
			found := false
			for _, e := range edges {
				if side.Equals(e) {
					found = true
					break
				}
			}
			if !found {
				edges = append(edges, side)
			}
		}
	}
	return edges
}
