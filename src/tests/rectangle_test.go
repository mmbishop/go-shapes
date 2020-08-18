package tests

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, NewPoint(3, 4), location)
	assert.Equal(t, 5.0, r.Width())
	assert.Equal(t, 6.0, r.Height())
}

func then_the_rectangle_is_at_the_current_new_location(t *testing.T) {
	location := r.Location()
	assert.Equal(t, NewPoint(5, 7), location)
}

func then_the_rectangle_is_resized_to_the_requested_area(t *testing.T) {
	area := r.Width() * r.Height()
	assert.Equal(t, 60.0, area)
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
