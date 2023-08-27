package game

import (
	"game/collision"
	"game/utils"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x      float64
	y      float64
	width  int
	height int

	speed float64

	game *Game
}

func NewEnemy(game *Game, x, y float64) (*Enemy, error) {
	return &Enemy{
		x:      x,
		y:      y,
		width:  19,
		height: 17,
		speed:  0.5,
		game:   game,
	}, nil
}

func (e *Enemy) Update() {
	directionToPlayer := utils.Direction(e, e.game.player)

	if collision.CheckCollision(e, e.game.player) {
		e.game.player.GetHit()
		e.game.deleteEnemy(e)
	}

	e.x += math.Cos(directionToPlayer) * e.speed
	e.y += math.Sin(directionToPlayer) * e.speed
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.game.imageManager.Draw(screen, "enemy", e.GetX(), e.GetY())
}

func (e *Enemy) GetX() float64 {
	return e.x - float64(e.GetWidth()/2)
}

func (e *Enemy) GetY() float64 {
	return e.y - float64(e.GetHeight()/2)
}

func (e *Enemy) GetWidth() int {
	return e.width
}

func (e *Enemy) GetHeight() int {
	return e.height
}
