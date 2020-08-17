package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	*AbstractShape
	radius float64
}

func NewCircle(center Point, radius float64) *Circle {
	return &Circle{&AbstractShape{center}, radius}
}

func (c Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) Resize(newAreaRatio float64) {
	area := math.Pi * c.radius * c.radius * newAreaRatio
	c.radius = math.Sqrt(area / math.Pi)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{center: %v, radius: %f}", c.Location(), c.radius)
}
