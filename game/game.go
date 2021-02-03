package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements Game interface of ebitent.Game
type Game struct {
	Img *ebiten.Image
}

// Update frame every second
func (g *Game) Update() error {
	return nil
}

// Draw on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Img, nil)
}

// Layout defines width and height of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
