package shapes

type Shape interface {
	GetLocation() Point

	GetBoundingBox() Box

	Move(dx float64, dy float64)
}
