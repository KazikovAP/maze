package pathfinding

import "github.com/KazikovAP/maze/internal/domain"

type Bfs struct{}

func (b *Bfs) FindPath(maze *domain.Maze, start, end domain.Point) (path []domain.Point, found bool) {
	queue := []domain.Point{start}
	visited := make(map[domain.Point]bool)
	visited[start] = true

	directions := []domain.Point{
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
	}

	prev := make(map[domain.Point]domain.Point)
	pathFound := false

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			pathFound = true
			break
		}

		for _, dir := range directions {
			next := domain.Point{X: current.X + dir.X, Y: current.Y + dir.Y}

			if maze.InBounds(next.X, next.Y) && maze.IsAccessibleAndUnvisited(next, visited) {
				queue = append(queue, next)
				visited[next] = true
				prev[next] = current
			}
		}
	}

	if pathFound {
		current := end
		for current != start {
			path = append([]domain.Point{current}, path...)
			current = prev[current]
		}

		path = append([]domain.Point{start}, path...)

		return path, true
	}

	return nil, false
}
