package grid

import (
	"testing"
)

func TestChunkCoord(t *testing.T) {
	cases := []struct {
		tx, ty, cx, cy int
	}{
		{0, 0, 0, 0},
		{31, 31, 0, 0},
		{32, 32, 1, 1},
		{32, -1, 1, -1},
		{-1, -1, -1, -1},
		{-32, -32, -1, -1},
		{-33, -32, -2, -1},
		{GridTiles, GridTiles, GridSize, GridSize},
	}
	for _, c := range cases {
		grid := New()
		chunkCoord := grid.chunkCoord(Coord(c.tx, c.ty))
		if chunkCoord.X() != c.cx || chunkCoord.Y() != c.cy {
			t.Errorf("%d / %d, want %d / %d", chunkCoord.X(), chunkCoord.Y(), c.cx, c.cy)
		}
	}
}

func TestChunkIndex(t *testing.T) {
	cases := []struct {
		cx, cy, ci int
	}{
		{0, 0, 500500},
		{31, 31, 531531},
		{32, 32, 532532},
		{32, -1, 532499},
		{-1, -1, 499499},
		{-32, -32, 468468},
		{-33, -32, 467468},
	}
	for _, c := range cases {
		grid := New()
		chunkIndex := grid.chunkIndex(c.cx, c.cy)
		if chunkIndex != c.ci {
			t.Errorf("%d, want %d", chunkIndex, c.ci)
		}
	}
}
