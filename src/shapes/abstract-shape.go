package shapes

type AbstractShape struct {
	location Point
}

func (a AbstractShape) Location() Point {
	return a.location
}

func (a *AbstractShape) Move(dx float64, dy float64) {
	a.location = NewPoint(a.location.X()+dx, a.location.Y()+dy)
}
