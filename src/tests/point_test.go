package tests

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, 3.0, p.X())
	assert.Equal(t, 4.0, p.Y())
}

func TestNewPoint(t *testing.T) {
	new_point_is_created(t)
}
