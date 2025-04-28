package app

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

type Game struct {
}

func NewGame() (*Game, error) {
	return nil, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
