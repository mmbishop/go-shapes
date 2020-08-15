package shapes

type Point struct {
	x float64
	y float64
}

// NewPoint creates a Point using the given x and y coordinates.
func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}
