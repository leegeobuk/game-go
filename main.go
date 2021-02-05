package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/leegeobuk/game-go/game"
	"github.com/leegeobuk/game-go/util"
)

const (
	winWidth  = 640
	winHeight = 480
	chWidth   = 32
	chHeight  = 32
)

var (
	background         *game.Background
	bgScaleX, bgScaleY float64
	character          *game.Character
	player             *game.Player
)

func init() {
	bg, _, err := ebitenutil.NewImageFromFile("./image/space.jpg")
	if err != nil {
		log.Fatal(err)
	}

	bgWidth, bgHeight := bg.Size()
	bgScaleX, bgScaleY = util.DivideInt(winWidth, bgWidth), util.DivideInt(winHeight, bgHeight)
	background = &game.Background{Img: bg, ScaleX: bgScaleX, ScaleY: bgScaleY}

	spaceship, _, err := ebitenutil.NewImageFromFile("./image/spaceship.png")
	if err != nil {
		log.Fatal(err)
	}

	ssWidth, ssHeight := spaceship.Size()
	ssScaleX, ssScaleY := util.DivideInt(chWidth, ssWidth), util.DivideInt(chHeight, ssHeight)
	character = &game.Character{Image: spaceship, ScaleX: ssScaleX, ScaleY: ssScaleY}

	player = &game.Player{
		Character: character,
		PosX:      float64(winWidth / 2),
		PosY:      float64(winHeight - chHeight),
		Speed:     4,
	}
}

func main() {
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowTitle("Space runner")

	if err := ebiten.RunGame(&game.Game{Background: background, Player: player}); err != nil {
		log.Fatal(err)
	}
}
