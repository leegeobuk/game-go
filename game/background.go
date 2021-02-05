package game

import "github.com/hajimehoshi/ebiten/v2"

// Background of the game
type Background struct {
	Img            *ebiten.Image
	ScaleX, ScaleY float64
}
