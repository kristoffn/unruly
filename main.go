package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/kristoffn/unruly/app"
)

func main() {
	game, err := app.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(app.ScreenWidth, app.ScreenHeight)
	ebiten.SetWindowTitle("Unruly")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
