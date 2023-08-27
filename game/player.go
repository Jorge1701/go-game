package game

import (
	"game/engine"
	"game/graphics"
	"game/utils"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var lastFire = int64(0)

type Player struct {
	drawable *graphics.Drawable

	speed    float64
	fireRate int64

	health int

	game *Game
}

func NewPlayer(game *Game, x, y float64) *Player {
	return &Player{
		drawable: &graphics.Drawable{
			Rect: &engine.Rectangle{
				Position: &engine.Point{
					X: x,
					Y: y,
				},
				Width:  19,
				Height: 40,
			},
			ImageAlias: "player",
		},
		speed:    1.3,
		fireRate: 500,
		health:   5,
		game:     game,
	}
}

func (p *Player) Update() {
	// Check mouse input
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		mouseX, mouseY := ebiten.CursorPosition()
		currT := time.Now().UnixMilli()

		if currT-lastFire > p.fireRate {
			lastFire = currT

			dirToMouse := utils.Direction(
				p.drawable.Rect.Position,
				&engine.Point{X: float64(mouseX), Y: float64(mouseY)},
			)

			p.game.createBullet(p.drawable.Rect.Position, dirToMouse)
			p.game.audioPlayer.PlayFromBytes("shot")
		}
	}

	// Checkout keyboard input
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})

	movePoint := &engine.Point{}

	for _, key := range keys {
		if key == ebiten.KeyA {
			movePoint.X--
		}
		if key == ebiten.KeyD {
			movePoint.X++
		}
		if key == ebiten.KeyW {
			movePoint.Y--
		}
		if key == ebiten.KeyS {
			movePoint.Y++
		}
	}

	if movePoint.X != 0 || movePoint.Y != 0 {
		moveDirection := utils.DirectionTo(movePoint)

		p.drawable.Rect.Position.X += math.Cos(moveDirection) * p.speed
		p.drawable.Rect.Position.Y += math.Sin(moveDirection) * p.speed
	}
}

func (p *Player) GetHit() {
	p.health--
	p.game.audioPlayer.PlayFromBytes("player_hit")

	if p.health <= 0 {
		p.game.GameOver()
	}
}
