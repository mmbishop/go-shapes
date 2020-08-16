package tests

import (
	. "shapes"
	"testing"
)

var p Point

func new_point_is_created(t *testing.T) {
	when_a_new_point_is_requested()
	then_a_new_point_is_created(t)
}

func when_a_new_point_is_requested() {
	p = NewPoint(3, 4)
}

func then_a_new_point_is_created(t *testing.T) {
	if p.X() != 3 || p.Y() != 4 {
		t.Errorf("Wanted {3, 4}, got {%f, %f}\n", p.X(), p.Y())
	}
}

func TestNewPoint(t *testing.T) {
	new_point_is_created(t)
}
