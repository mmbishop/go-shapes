package tests

import (
	. "shapes"
	"testing"
)

var c *Circle

func new_circle_is_created(t *testing.T) {
	when_a_new_circle_is_requested()
	then_a_new_circle_is_created(t)
}

func circle_is_moved_by_a_specified_offset(t *testing.T) {
	given_a_circle()
	when_the_circle_is_moved()
	then_the_circle_is_at_the_current_new_location(t)
}

func circle_is_resized_by_a_specified_area_ratio(t *testing.T) {
	given_a_circle()
	when_the_circle_is_resized()
	then_the_circle_is_resized_to_the_requested_area(t)
}

func given_a_circle() {
	c = NewCircle(NewPoint(3, 4), 5)
}

func when_a_new_circle_is_requested() {
	c = NewCircle(NewPoint(3, 4), 5)
}

func when_the_circle_is_moved() {
	c.Move(2, 3)
}

func when_the_circle_is_resized() {
	c.Resize(1.96)
}

func then_a_new_circle_is_created(t *testing.T) {
	center := c.Location()
	if center.X() != 3 || center.Y() != 4 || c.Radius() != 5 {
		t.Errorf("Wanted Circle{center: {3, 4}, radius: 5}, got Circle{center: {%f, %f}, radius: %f}", center.X(), center.Y(), c.Radius())
	}
}

func then_the_circle_is_at_the_current_new_location(t *testing.T) {
	center := c.Location()
	if center.X() != 5 || center.Y() != 7 {
		t.Errorf("Wanted circle at {5, 7}, got circle at {%f, %f}", center.X(), center.Y())
	}
}

func then_the_circle_is_resized_to_the_requested_area(t *testing.T) {
	if c.Radius() > 7.000001 || c.Radius() < 6.999999 {
		t.Errorf("Wanted circle with radius 7, got circle with radius %f", c.Radius())
	}
}

func TestNewCircle(t *testing.T) {
	new_circle_is_created(t)
}

func TestMoveCircle(t *testing.T) {
	circle_is_moved_by_a_specified_offset(t)
}

func TestResizeCircle(t *testing.T) {
	circle_is_resized_by_a_specified_area_ratio(t)
}
