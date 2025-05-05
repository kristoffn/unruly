package game

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	arcadeFaceSource *text.GoTextFaceSource
)

const (
	ScreenDimension = 640
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	arcadeFaceSource = s
}

type Game struct {
	grid *Grid
}

func NewGame() (*Game, error) {
	grid, err := NewGrid(20)
	if err != nil {
		return nil, err
	}
	g := &Game{grid: grid}
	return g, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.grid.Draw(screen)
	x, y := ebiten.CursorPosition()
	op := &text.DrawOptions{}
	text.Draw(screen, fmt.Sprintf("x: %v, y: %v", x, y),
		&text.GoTextFace{
			Source: arcadeFaceSource,
			Size:   20,
		}, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenDimension, ScreenDimension
}
