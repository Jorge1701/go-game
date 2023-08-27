package game

import (
	"game/engine"
	"game/graphics"
	"game/utils"
	"math"
)

type Enemy struct {
	drawable *graphics.Drawable

	speed float64

	game *Game
}

func NewEnemy(game *Game, x, y float64) (*Enemy, error) {
	return &Enemy{
		drawable: &graphics.Drawable{
			Rect: &engine.Rectangle{
				Position: &engine.Point{
					X: x,
					Y: y,
				},
				Width:  19,
				Height: 17,
			},
			ImageAlias: "enemy",
		},
		speed: 0.2,
		game:  game,
	}, nil
}

func (e *Enemy) Update() {
	directionToPlayer := utils.Direction(
		e.drawable.Rect.Position,
		e.game.player.drawable.Rect.Position,
	)

	if engine.CheckCollision(e.drawable.Rect, e.game.player.drawable.Rect) {
		e.game.player.GetHit()
		e.game.deleteEnemy(e)
	}

	e.drawable.Rect.Position.X += math.Cos(directionToPlayer) * e.speed
	e.drawable.Rect.Position.Y += math.Sin(directionToPlayer) * e.speed
}
