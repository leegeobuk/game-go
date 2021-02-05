package game

import "github.com/hajimehoshi/ebiten/v2"

// Character of the game
type Character struct {
	Image          *ebiten.Image
	ScaleX, ScaleY float64
}
