package shapes

import (
	"fmt"
	"math"
)

type Rectangle struct {
	*AbstractShape
	width  float64
	height float64
}

func NewRectangle(location Point, width float64, height float64) *Rectangle {
	return &Rectangle{&AbstractShape{location}, width, height}
}

func (r Rectangle) Width() float64 {
	return r.width
}

func (r Rectangle) Height() float64 {
	return r.height
}

func (r *Rectangle) Resize(newAreaRatio float64) {
	area := r.width * r.height
	widthRatio := r.width / math.Sqrt(area)
	heightRatio := r.height / math.Sqrt(area)
	area *= newAreaRatio
	r.width = widthRatio * math.Sqrt(area)
	r.height = heightRatio * math.Sqrt(area)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{location: %v, width: %f, height: %f}", r.GetLocation(), r.width, r.height)
}
