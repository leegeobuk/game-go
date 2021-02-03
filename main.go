package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leegeobuk/game-go/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
