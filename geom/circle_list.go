// Package geom has geometry related structs and funcs.
package geom

// CircleList is a slice of circles
type CircleList []*Circle

// NewCircleList creates a new circle list
func NewCircleList() CircleList {
	return CircleList{}
}

// Add adds a circle to the list
func (c *CircleList) Add(circle *Circle) {
	*c = append(*c, circle)
}

// AddXY adds a circle to the list
func (c *CircleList) AddXY(x, y, r float64) {
	*c = append(*c, NewCircle(x, y, r))
}

// RandomizePosition randomizes the position of all the circles in a list.
func (c *CircleList) RandomizePosition(amount float64) {
	for _, circle := range *c {
		circle.RandomizePosition(amount)
	}
}

// RandomizeRadius randomizes the position of all the circles in a list.
func (c *CircleList) RandomizeRadius(amount float64) {
	for _, circle := range *c {
		circle.RandomizeRadius(amount)
	}
}

// Scale scales all the circles in a list.
func (c *CircleList) Scale(scale float64) {
	for _, circle := range *c {
		circle.Scale(scale)
	}
}

// ScaleFrom scales all the circles in a list using the x, y location as a center.
func (c *CircleList) ScaleFrom(x, y, scale float64) {
	for _, circle := range *c {
		circle.ScaleFrom(x, y, scale)
	}
}

// ScaleLocal scales all the circles in a list.
func (c *CircleList) ScaleLocal(scale float64) {
	for _, circle := range *c {
		circle.ScaleLocal(scale)
	}
}

// Translate translates all the circles in a list.
func (c *CircleList) Translate(x, y float64) {
	for _, circle := range *c {
		circle.Translate(x, y)
	}
}

// Rotate rotates all the circles in a list.
func (c *CircleList) Rotate(angle float64) {
	for _, circle := range *c {
		circle.Rotate(angle)
	}
}

// RotateFrom rotates all the circles in a list using the x, y location as a center.
func (c *CircleList) RotateFrom(x, y float64, angle float64) {
	for _, circle := range *c {
		circle.RotateFrom(x, y, angle)
	}
}
