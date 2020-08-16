package tests

import (
	. "shapes"
	"testing"
)

var boundingBox Box
var r Rectangle

func new_rectangle_is_created(t *testing.T) {
	when_a_new_rectangle_is_requested()
	then_a_new_rectangle_is_created(t)
}

func rectangle_returns_its_bounding_box(t *testing.T) {
	given_a_rectangle()
	when_its_bounding_box_is_requested()
	then_its_bounding_box_is_returned(t)
}

func rectangle_is_moved_by_specified_offset(t *testing.T) {
	given_a_rectangle()
	when_the_rectangle_is_moved()
	then_the_rectangle_is_at_the_current_new_location(t)
}

func given_a_rectangle() {
	r = NewRectangle(NewPoint(3, 4), 5, 6)
}

func when_a_new_rectangle_is_requested() {
	r = NewRectangle(NewPoint(3, 4), 5, 6)
}

func when_its_bounding_box_is_requested() {
	boundingBox = r.GetBoundingBox()
}

func when_the_rectangle_is_moved() {
	r.Move(2, 3)
}

func then_a_new_rectangle_is_created(t *testing.T) {
	location := r.GetLocation()
	if location.X() != 3 || location.Y() != 4 || r.Width() != 5 || r.Height() != 6 {
		t.Errorf("Wanted Rectangle{location: {3, 4}, width: 5, height: 6}, got Rectangle{location: {%f, %f}, width: %f, height: %f}",
			location.X(), location.Y(), r.Width(), r.Height())
	}
}

func then_its_bounding_box_is_returned(t *testing.T) {
	upperLeft := boundingBox.UpperLeft()
	lowerRight := boundingBox.LowerRight()
	if upperLeft.X() != 3 || upperLeft.Y() != 4 || lowerRight.X() != 8 || lowerRight.Y() != 10 {
		t.Errorf("Wanted BoundingBox{{3, 4}, {8, 10}}, got BoundingBox{{%f, %f}, {%f, %f}}", upperLeft.X(), upperLeft.Y(), lowerRight.X(), lowerRight.Y())
	}
}

func then_the_rectangle_is_at_the_current_new_location(t *testing.T) {
	location := r.GetLocation()
	if location.X() != 5 || location.Y() != 7 {
		t.Errorf("Wanted rectangle at {5, 6}, got rectangle at {%f, %f}", location.X(), location.Y())
	}
}

func TestNewRectangle(t *testing.T) {
	new_rectangle_is_created(t)
}

func TestGetRectangleBoundingBox(t *testing.T) {
	rectangle_returns_its_bounding_box(t)
}

func TestMoveRectangle(t *testing.T) {
	rectangle_is_moved_by_specified_offset(t)
}
