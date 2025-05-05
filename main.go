package main

import (
	"log"
	"unruly/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame()
	if err != nil {

	}

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Unruly Puzzle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
