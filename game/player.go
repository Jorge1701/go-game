package game

import (
	"game/graphics"
	"game/utils"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
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
	keys := sdl.GetKeyboardState()

	mouseX, mouseY, mouseState := sdl.GetMouseState()

	if mouseState == sdl.ButtonLMask {
		currT := time.Now().UnixMilli()
		if currT-lastFire > p.fireRate {
			lastFire = currT
			dirToMouse := utils.Direction(&utils.Point{X: p.x, Y: p.y}, &utils.Point{X: float64(mouseX), Y: float64(mouseY)})
			p.game.createBullet(p.x, p.y, dirToMouse)
			// audio.AllAudios["shot"].Play()
		}
	}

	movePoint := &utils.Point{}

	if keys[sdl.SCANCODE_A] == 1 {
		movePoint.X--
	}
	if keys[sdl.SCANCODE_D] == 1 {
		movePoint.X++
	}
	if keys[sdl.SCANCODE_W] == 1 {
		movePoint.Y--
	}
	if keys[sdl.SCANCODE_S] == 1 {
		movePoint.Y++
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
