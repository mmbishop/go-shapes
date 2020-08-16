package main

import (
	"fmt"
	. "shapes"
)

func main() {
	var shapes [4]Shape
	shapes[0] = NewCircle(NewPoint(3, 4), 5)
	shapes[1] = NewRectangle(NewPoint(7, 2), 3, 6)
	shapes[2] = NewCircle(NewPoint(10, 8), 6)
	shapes[3] = NewRectangle(NewPoint(15, 12), 5, 5)

	fmt.Println("As constructed:")
	for i, shape := range shapes {
		fmt.Printf("%d: %v with bounds %v\n", i, shape, shape.GetBoundingBox())
	}

	shapes[0].Move(3, 2)
	shapes[1].Move(-1.5, 3.5)
	shapes[2].Move(-3, 2.5)
	shapes[3].Move(-3, -2.5)

	fmt.Println("After moves:")
	for i, shape := range shapes {
		fmt.Printf("%d: %v with bounds %v\n", i, shape, shape.GetBoundingBox())
	}
}
