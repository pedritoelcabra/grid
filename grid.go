// A package for managing a grid based game world
package grid

import (
	"log"
)

const (
	ChunkSize  = 32
	GridSize   = 1000
	GridOffset = 500
	GridTiles  = 32000
	TileOffset = 16000
)

type grid struct {
	chunks map[int]*chunk
}

func New() *grid {
	arraySize := GridSize * GridSize
	arrayChunks := make(map[int]*chunk, arraySize)
	return &grid{arrayChunks}
}

func (g *grid) Tile(tileCoord coord) *tile {
	chunkCoord := g.chunkCoord(tileCoord)
	chunkIndex := g.chunkIndex(chunkCoord.X(), chunkCoord.Y())
	if aChunk, chunkExists := g.chunks[chunkIndex]; chunkExists {
		return aChunk.Tile(tileCoord)
	}
	g.chunks[chunkIndex] = NewChunk(chunkCoord)
	return g.chunks[chunkIndex].Tile(tileCoord)
}

func (g *grid) chunkIndex(x, y int) int {
	x += GridOffset
	y += GridOffset
	if x < 1 || x > GridSize || y < 1 || y > GridSize {
		log.Fatalf("Grid.Tile() requested invalid chunk %d / %d", x, y)
	}
	return (x * GridSize) + y
}

func (g *grid) chunkCoord(tileCoord coord) coord {
	x := ((tileCoord.X() + TileOffset) / ChunkSize) - GridOffset
	y := ((tileCoord.Y() + TileOffset) / ChunkSize) - GridOffset
	return Coord(x, y)
}
