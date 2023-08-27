package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize = 32
)

var tiles = map[int]string{
	0: "grass_tile",
	1: "dirt_tile",
	2: "sand_tile",
	3: "water_tile",
}

var world = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 2, 3},
	{0, 0, 0, 0, 1, 1, 0, 0, 2, 3},
	{0, 0, 1, 1, 1, 1, 0, 0, 2, 3},
	{0, 0, 0, 1, 1, 0, 0, 0, 2, 3},
	{0, 0, 0, 1, 1, 0, 0, 0, 2, 3},
	{0, 0, 0, 0, 0, 0, 0, 0, 2, 3},
	{0, 0, 0, 0, 0, 0, 0, 0, 2, 3},
	{0, 0, 0, 0, 0, 0, 0, 0, 2, 3},
	{0, 0, 0, 0, 0, 0, 2, 2, 2, 3},
	{2, 2, 2, 2, 2, 2, 2, 3, 3, 3},
	{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
}

type Terrain struct {
	game *Game
}

func NewTerrain(game *Game) *Terrain {
	return &Terrain{
		game: game,
	}
}

func (t *Terrain) Draw(screen *ebiten.Image) {
	for x, row := range world {
		for y, tileId := range row {
			t.game.imageManager.DrawImage(
				screen,
				tiles[tileId],
				float64(x*tileSize),
				float64(y*tileSize),
			)
		}
	}
}
