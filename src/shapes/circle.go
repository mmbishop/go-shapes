package shapes

import "fmt"

type Circle struct {
	*AbstractShape
	radius float64
}

func NewCircle(center Point, radius float64) Circle {
	return Circle{&AbstractShape{center}, radius}
}

func (c Circle) GetRadius() float64 {
	return c.radius
}

func (c Circle) GetBoundingBox() Box {
	center := c.GetLocation()
	return NewBox(NewPoint(center.X()-c.radius, center.Y()-c.radius), NewPoint(center.X()+c.radius, center.Y()+c.radius))
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{center: %v, radius: %f}", c.GetLocation(), c.radius)
}
