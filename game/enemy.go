package game

import (
	"game/collision"
	"game/images"
	"game/utils"
	"math"
)

type Enemy struct {
	x float64
	y float64

	speed float64

	image *images.Image

	game *Game
}

func NewEnemy(game *Game, x, y float64) (*Enemy, error) {
	return &Enemy{
		x:     x,
		y:     y,
		speed: 0.5,
		image: images.AllImages["enemy"],
		game:  game,
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

func (e *Enemy) GetX() float64 {
	return e.x - float64(e.GetWidth()/2)
}

func (e *Enemy) GetY() float64 {
	return e.y - float64(e.GetHeight()/2)
}

func (e *Enemy) GetImage() *images.Image {
	return e.image
}

func (e *Enemy) GetWidth() int {
	return e.image.Image.Bounds().Dx()
}

func (e *Enemy) GetHeight() int {
	return e.image.Image.Bounds().Dy()
}
