package shapes

type Shape interface {

	// Location returns the location of a Shape. The interpretation of the location (e.g., center or a vertex) is
	// up to the Shape subclass.
	Location() Point

	// Move translates a Shape by the given x and y offsets.
	Move(dx float64, dy float64)

	// Resize changes the area of a Shape based on the given size ratio. For example, a newAreaRatio of 0.5 will reduce the
	// area by half; a ratio of 2.0 will double the area.
	Resize(newAreaRatio float64)
}
