package unruly

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Unruly")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
