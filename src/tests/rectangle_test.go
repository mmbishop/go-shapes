package tests

import (
	. "shapes"
	"testing"
)

var r *Rectangle

func new_rectangle_is_created(t *testing.T) {
	when_a_new_rectangle_is_requested()
	then_a_new_rectangle_is_created(t)
}

func rectangle_is_moved_by_specified_offset(t *testing.T) {
	given_a_rectangle()
	when_the_rectangle_is_moved()
	then_the_rectangle_is_at_the_current_new_location(t)
}

func rectangle_is_resized_by_a_specified_area_ratio(t *testing.T) {
	given_a_rectangle()
	when_the_rectangle_is_resized()
	then_the_rectangle_is_resized_to_the_requested_area(t)
}

func given_a_rectangle() {
	r = NewRectangle(NewPoint(3, 4), 5, 6)
}

func when_a_new_rectangle_is_requested() {
	r = NewRectangle(NewPoint(3, 4), 5, 6)
}

func when_the_rectangle_is_moved() {
	r.Move(2, 3)
}

func when_the_rectangle_is_resized() {
	r.Resize(2)
}

func then_a_new_rectangle_is_created(t *testing.T) {
	location := r.Location()
	if location.X() != 3 || location.Y() != 4 || r.Width() != 5 || r.Height() != 6 {
		t.Errorf("Wanted Rectangle{location: {3, 4}, width: 5, height: 6}, got Rectangle{location: {%f, %f}, width: %f, height: %f}",
			location.X(), location.Y(), r.Width(), r.Height())
	}
}

func then_the_rectangle_is_at_the_current_new_location(t *testing.T) {
	location := r.Location()
	if location.X() != 5 || location.Y() != 7 {
		t.Errorf("Wanted rectangle at {5, 6}, got rectangle at {%f, %f}", location.X(), location.Y())
	}
}

func then_the_rectangle_is_resized_to_the_requested_area(t *testing.T) {
	area := r.Width() * r.Height()
	if area != 60 {
		t.Errorf("Wanted a rectangle with area 15, got a rectangle with area %f (width: %f, height = %f)",
			area, r.Width(), r.Height())
	}
}

func TestNewRectangle(t *testing.T) {
	new_rectangle_is_created(t)
}

func TestMoveRectangle(t *testing.T) {
	rectangle_is_moved_by_specified_offset(t)
}

func TestResizeRectangle(t *testing.T) {
	rectangle_is_resized_by_a_specified_area_ratio(t)
}
