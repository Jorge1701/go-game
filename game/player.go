package game

import (
	"game/graphics"
	"game/utils"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var lastFire = int64(0)

type Player struct {
	x float64
	y float64

	speed    float64
	fireRate int64

	health int

	image *graphics.Image

	game *Game
}

func NewPlayer(game *Game, x, y float64) *Player {
	return &Player{
		x:        x,
		y:        y,
		speed:    1,
		fireRate: 500,
		health:   5,
		image:    graphics.AllImages["player"],
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
			dirToMouse := utils.Direction(&utils.Point{X: p.x, Y: p.y}, &utils.Point{X: float64(mouseX), Y: float64(mouseY)})
			p.game.createBullet(p.x, p.y, dirToMouse)
			// audio.AllAudios["shot"].Play()
		}
	}

	// Checkout keyboard input
	keys := inpututil.AppendPressedKeys([]ebiten.Key{})

	movePoint := &utils.Point{}

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

		p.x += math.Cos(moveDirection) * p.speed
		p.y += math.Sin(moveDirection) * p.speed
	}
}

func (p *Player) GetHit() {
	p.health--
	p.game.audioPlayer.PlayFromBytes("player_hit")
	// audio.AllAudios["player_hit"].Play()

	if p.health <= 0 {
		p.game.GameOver()
	}
}

func (p *Player) GetX() float64 {
	return p.x - float64(p.GetWidth()/2)
}

func (p *Player) GetY() float64 {
	return p.y - float64(p.GetHeight()/2)
}

func (p *Player) GetImage() *graphics.Image {
	return p.image
}

func (p *Player) GetWidth() int {
	return p.image.Image.Bounds().Dx()
}

func (p *Player) GetHeight() int {
	return p.image.Image.Bounds().Dy()
}
