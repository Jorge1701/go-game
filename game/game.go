package game

import (
	"fmt"
	"game/configuration"
	"game/render"
	"math/rand"
	"slices"
)

type Game struct {
	renderer *render.Renderer

	player  *Player
	enemies []*Enemy
	bullets []*Bullet

	stage         int
	killedEnemies int
	maxEnemyCount int

	IsGameOver bool
}

func NewGame(renderer *render.Renderer) (g *Game, err error) {
	g = &Game{
		renderer:      renderer,
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

	return g, err
}

func (g *Game) Update() {
	g.player.Update()

	for _, e := range g.enemies {
		e.Update()
	}

	for _, b := range g.bullets {
		b.Update(g.enemies)
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
}

func (g *Game) Render() {
	g.renderer.RenderDrawable(g.player)

	for _, e := range g.enemies {
		g.renderer.RenderDrawable(e)
	}

	for _, b := range g.bullets {
		g.renderer.RenderDrawable(b)
	}
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
}
