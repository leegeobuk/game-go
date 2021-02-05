package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements Game interface of ebitent.Game
type Game struct {
	Background *Background
	Player     *Player
}

// Update logic every tick
func (g *Game) Update() error {
	return nil
}

// Draw background and player on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	bg := g.Background
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(bg.ScaleX, bg.ScaleY)
	screen.DrawImage(bg.Img, op)

	character := g.Player.Character
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(character.ScaleX, character.ScaleY)
	playerOp.GeoM.Translate(g.Player.PosX, g.Player.PosY)
	screen.DrawImage(character.Image, playerOp)
}

// Layout defines width and height of the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
