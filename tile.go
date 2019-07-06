package grid

// Tile is the data contained in any given coordinate in the grid
type tile struct {
	coordinates coord
	values      map[string]int
}

func (t tile) X() int {
	return t.coordinates.X()
}

func (t tile) Y() int {
	return t.coordinates.Y()
}
