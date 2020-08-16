package tests

import (
	. "shapes"
	"testing"
)

var circleBoundingBox Box
var c Circle

func new_circle_is_created(t *testing.T) {
	when_a_new_circle_is_requested()
	then_a_new_circle_is_created(t)
}

func circle_returns_its_bounding_box(t *testing.T) {
	given_a_circle()
	when_the_bounding_box_of_the_circle_is_requested()
	then_the_bounding_box_of_the_circle_is_returned(t)
}

func circle_is_moved_by_a_specified_offset(t *testing.T) {
	given_a_circle()
	when_the_circle_is_moved()
	then_the_circle_is_at_the_current_new_location(t)
}

func given_a_circle() {
	c = NewCircle(NewPoint(3, 4), 5)
}

func when_a_new_circle_is_requested() {
	c = NewCircle(NewPoint(3, 4), 5)
}

func when_the_bounding_box_of_the_circle_is_requested() {
	circleBoundingBox = c.GetBoundingBox()
}

func when_the_circle_is_moved() {
	c.Move(2, 3)
}

func then_a_new_circle_is_created(t *testing.T) {
	center := c.GetLocation()
	if center.X() != 3 || center.Y() != 4 || c.GetRadius() != 5 {
		t.Errorf("Wanted Circle{center: {3, 4}, radius: 5}, got Circle{center: {%f, %f}, radius: %f}", center.X(), center.Y(), c.GetRadius())
	}
}

func then_the_bounding_box_of_the_circle_is_returned(t *testing.T) {
	upperLeft := circleBoundingBox.UpperLeft()
	lowerRight := circleBoundingBox.LowerRight()
	if upperLeft.X() != -2 || upperLeft.Y() != -1 || lowerRight.X() != 8 || lowerRight.Y() != 9 {
		t.Errorf("Wanted BoundingBox{{-2, -1}, {8, 9}}, got BoundingBox{{%f, %f}, {%f, %f}}",
			upperLeft.X(), upperLeft.Y(), lowerRight.X(), lowerRight.Y())
	}
}

func then_the_circle_is_at_the_current_new_location(t *testing.T) {
	center := c.GetLocation()
	if center.X() != 5 || center.Y() != 7 {
		t.Errorf("Wanted circle at {5, 7}, got circle at {%f, %f}", center.X(), center.Y())
	}
}

func TestNewCircle(t *testing.T) {
	new_circle_is_created(t)
}

func TestGetCircleBoundingBox(t *testing.T) {
	circle_returns_its_bounding_box(t)
}

func TestMoveCircle(t *testing.T) {
	circle_is_moved_by_a_specified_offset(t)
}
