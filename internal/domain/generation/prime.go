package generation

import (
	"github.com/KazikovAP/maze/internal/domain"
	"github.com/KazikovAP/maze/pkg/srandom"
)

type PrimGenerator struct{}

func (g *PrimGenerator) Generate(maze *domain.Maze) {
	GenerateMazePrime(maze)
}

func GenerateMazePrime(maze *domain.Maze) {
	start := domain.Point{X: 1, Y: 1}
	maze.Field[start.Y][start.X] = ' '

	walls := []domain.Point{
		{X: start.X + 1, Y: start.Y},
		{X: start.X - 1, Y: start.Y},
		{X: start.X, Y: start.Y + 1},
		{X: start.X, Y: start.Y - 1},
	}

	for len(walls) > 0 {
		index := srandom.Intn(len(walls))
		wall := walls[index]
		walls = append(walls[:index], walls[index+1:]...)

		if !maze.InBounds(wall.X, wall.Y) {
			continue
		}

		adjCount := 0
		if maze.Field[wall.Y-1][wall.X] == ' ' {
			adjCount++
		}

		if maze.Field[wall.Y+1][wall.X] == ' ' {
			adjCount++
		}

		if maze.Field[wall.Y][wall.X-1] == ' ' {
			adjCount++
		}

		if maze.Field[wall.Y][wall.X+1] == ' ' {
			adjCount++
		}

		if adjCount == 1 {
			maze.Field[wall.Y][wall.X] = ' '

			walls = append(walls,
				domain.Point{X: wall.X + 1, Y: wall.Y},
				domain.Point{X: wall.X - 1, Y: wall.Y},
				domain.Point{X: wall.X, Y: wall.Y + 1},
				domain.Point{X: wall.X, Y: wall.Y - 1},
			)
		}
	}
}
