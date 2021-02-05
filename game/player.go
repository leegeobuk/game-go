package game

// Player of the game
type Player struct {
	Character      *Character
	PosX, PosY     float64
	ScaleX, ScaleY float64
	Speed          int
}
