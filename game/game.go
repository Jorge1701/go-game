package game

import (
	"fmt"
	"game/audio"
	"game/collision"
	"game/configuration"
	"game/graphics"
	"math/rand"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var gameBoundary = &collision.Rectangle{
	X:      -10,
	Y:      -10,
	Width:  configuration.Width + 10,
	Height: configuration.Height + 10,
}

type Game struct {
	player  *Player
	enemies []*Enemy
	bullets []*Bullet

	stage         int
	killedEnemies int
	maxEnemyCount int

	IsGameOver bool
}

func NewGame() *Game {
	g := &Game{
		enemies:       []*Enemy{},
		bullets:       []*Bullet{},
		stage:         1,
		killedEnemies: 0,
		maxEnemyCount: configuration.StartingEnemyCount,
		IsGameOver:    false,
	}

	g.player = NewPlayer(
		g,
		configuration.Width/2,
		configuration.Height/2,
	)

	return g
}

func (g *Game) Update() error {
	g.player.Update()

	for _, e := range g.enemies {
		e.Update()
	}

	for _, b := range g.bullets {
		if collision.CheckCollision(b, gameBoundary) {
			b.Update(g.enemies)
		} else {
			g.deleteBullet(b)
		}
	}

	for len(g.enemies) < g.maxEnemyCount {
		g.createEnemy()
	}

	fmt.Printf("stage %d - enemies %d - killed %d - next %d - bullets %d\n",
		g.stage,
		len(g.enemies),
		g.killedEnemies,
		g.nextStage(),
		len(g.bullets),
	)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	graphics.AllImages["background"].Draw(screen, 0, 0)

	graphics.Draw(g.player, screen)

	for _, e := range g.enemies {
		graphics.Draw(e, screen)
	}

	for _, b := range g.bullets {
		graphics.Draw(b, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configuration.Width, configuration.Height
}

func (g *Game) createBullet(x, y, dir float64) {
	g.bullets = append(g.bullets, NewBullet(g, x, y, dir))
}

func (g *Game) createEnemy() {
	var x, y float64

	side := rand.Intn(4)

	switch side {
	case 0: // top
		x = float64(rand.Intn(configuration.Width))
		y = -50
	case 1: // right
		x = configuration.Width + 50
		y = float64(rand.Intn(configuration.Height))
	case 2: // bottom
		x = float64(rand.Intn(configuration.Width))
		y = configuration.Height + 50
	case 3: // left
		x = -50
		y = float64(rand.Intn(configuration.Height))
	}

	e, err := NewEnemy(g, x, y)

	if err != nil {
		panic(fmt.Sprintf("Error generating enemy: %v", err))
	}

	g.enemies = append(g.enemies, e)
}

func (g *Game) nextStage() int {
	return configuration.StartingEnemyCount*g.stage + g.stage
}

func (g *Game) deleteEnemy(enemyToDelete *Enemy) {
	for i, e := range g.enemies {
		if e == enemyToDelete {
			g.enemies = slices.Delete(g.enemies, i, i+1)

			g.killedEnemies++
			if g.killedEnemies%g.nextStage() == 0 {
				g.stage++
				g.maxEnemyCount++
			}

			return
		}
	}
}

func (g *Game) deleteBullet(bulletToDelete *Bullet) {
	for i, b := range g.bullets {
		if b == bulletToDelete {
			g.bullets = slices.Delete(g.bullets, i, i+1)
			return
		}
	}
}

func (g *Game) GameOver() {
	g.IsGameOver = true
	audio.AllAudios["game_over"].Play()
}
