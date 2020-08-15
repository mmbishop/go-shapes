package shapes

import (
	"testing"
)

var b Box

func new_box_is_created(t *testing.T) {
	when_a_new_box_is_requested()
	then_a_new_box_is_created(t)
}

func when_a_new_box_is_requested() {
	b = NewBox(NewPoint(1, 2), NewPoint(5, 8))
}

func then_a_new_box_is_created(t *testing.T) {
	upperLeft := b.UpperLeft()
	lowerRight := b.LowerRight()
	if upperLeft.X() != 1 || upperLeft.Y() != 2 || lowerRight.X() != 5 || lowerRight.Y() != 8 {
		t.Errorf("Wanted Box{{1, 2}, {5, 8}, got {{%f, %f}, {%f, %f}}", upperLeft.X(), upperLeft.Y(), lowerRight.X(), lowerRight.Y())
	}
}

func TestNewBox(t *testing.T) {
	new_box_is_created(t)
}
