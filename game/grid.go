package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	Trivial = iota
	Easy
	Normal
	Hard
)

const (
	Black = iota
	White
)

type Grid struct {
	size       int
	difficulty int
	cellMatrix [][]Cell
}

type Cell struct {
	color int
	fixed bool
}

func NewGrid(gridsize int) (*Grid, error) {
	if gridsize < 1 {
		return nil, fmt.Errorf("gridsize must be bigger than 0")
	}
	dim := gridsize * 2
	matrix := make([][]Cell, dim)
	for i := 0; i < dim; i++ {
		matrix[i] = make([]Cell, dim)
	}
	g := &Grid{
		size:       gridsize,
		cellMatrix: matrix,
	}
	return g, nil
}

func (g *Grid) Update() error {
	return nil
}

func (g *Grid) Draw(screen *ebiten.Image) {
	cellSize := ScreenDimension / g.size / 2
	dim := g.size * 2
	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			vector.DrawFilledRect(screen, float32(x*cellSize), float32(y*cellSize), float32(cellSize-1), float32(cellSize-1), color.RGBA{0x80, 0x00, 0x80, 0x80}, true)
		}
	}
}

func (g *Grid) SetDifficulty(diff int) error {
	if diff < 0 || diff > 3 {
		return fmt.Errorf("invalid difficulty: %v", diff)
	}
	g.difficulty = diff
	return nil
}
