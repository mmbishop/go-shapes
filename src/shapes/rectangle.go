package shapes

type Rectangle struct {
	*AbstractShape
	width  float64
	height float64
}

func NewRectangle(location Point, width float64, height float64) Rectangle {
	return Rectangle{&AbstractShape{location}, width, height}
}

func (r Rectangle) Width() float64 {
	return r.width
}

func (r Rectangle) Height() float64 {
	return r.height
}

func (r Rectangle) GetBoundingBox() Box {
	upperLeft := r.GetLocation()
	return NewBox(upperLeft, NewPoint(upperLeft.X()+r.width, upperLeft.Y()+r.height))
}
