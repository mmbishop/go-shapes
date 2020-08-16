package shapes

import "fmt"

type Box struct {
	upperLeft  Point
	lowerRight Point
}

// NewBox creates a Box from the given upper-left and lower-right Points.
func NewBox(upperLeft Point, lowerRight Point) Box {
	return Box{upperLeft, lowerRight}
}

func (b Box) UpperLeft() Point {
	return b.upperLeft
}

func (b Box) LowerRight() Point {
	return b.lowerRight
}

func (b Box) String() string {
	return fmt.Sprintf("BoundingBox{upperLeft: %v, lowerRight: %v}", b.UpperLeft(), b.LowerRight())
}
