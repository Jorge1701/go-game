package game

import (
	"game/collision"
	"game/render"
	"game/utils"
	"math"
)

type Enemy struct {
	x float64
	y float64

	speed float64

	texture *render.Texture

	game *Game
}

func NewEnemy(game *Game, x, y float64, texture *render.Texture) (*Enemy, error) {
	return &Enemy{
		x:       x,
		y:       y,
		speed:   1,
		texture: texture,
		game:    game,
	}, nil
}

func (e *Enemy) Update() {
	directionToPlayer := utils.Direction(e, e.game.player)

	if collision.CheckCollision(e, e.game.player) {
		e.game.player.GetHit()
		e.game.DeleteEnemy(e)
	}

	e.x += math.Cos(directionToPlayer) * e.speed
	e.y += math.Sin(directionToPlayer) * e.speed
}

func (e *Enemy) GetX() float64 {
	return e.x
}

func (e *Enemy) GetY() float64 {
	return e.y
}

func (e *Enemy) GetTexture() *render.Texture {
	return e.texture
}

func (e *Enemy) GetWidth() float64 {
	return e.texture.Width
}

func (e *Enemy) GetHeight() float64 {
	return e.texture.Height
}
