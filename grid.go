// A package for managing a grid based game world
package grid

type grid struct {
	size   int
	radius int
}

func New(size int) grid {
	if size < 0 {
		size = -size
	}
	if size%2 == 1 {
		size++
	}
	return grid{size, size / 2}
}

func (g grid) Size() int {
	return g.size
}

func (g grid) Radius() int {
	return g.radius
}
