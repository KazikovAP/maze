package pathfinding

import (
	"container/heap"
	"math"

	"github.com/KazikovAP/maze/internal/domain"
)

type PriorityPoint struct {
	Point    domain.Point
	Priority float64
}

type PointHeap []PriorityPoint

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(PriorityPoint))
}
func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]

	return item
}

func heuristic(a, b domain.Point) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

type AStar struct{}

func (a *AStar) FindPath(maze *domain.Maze, start, end domain.Point) (path []domain.Point, found bool) {
	cameFrom := make(map[domain.Point]domain.Point)
	costSoFar := make(map[domain.Point]float64)

	startNode := PriorityPoint{
		Point:    start,
		Priority: 0,
	}
	cameFrom[start] = start
	costSoFar[start] = 0

	var openSet PointHeap

	heap.Init(&openSet)
	heap.Push(&openSet, startNode)

	directions := []domain.Point{
		{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}, {X: -1, Y: 0},
	}

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(PriorityPoint).Point

		if current == end {
			path := []domain.Point{}
			for current != start {
				path = append(path, current)
				current = cameFrom[current]
			}

			path = append(path, start)

			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}

			return path, true
		}

		for _, dir := range directions {
			neighbor := domain.Point{X: current.X + dir.X, Y: current.Y + dir.Y}
			if !maze.InBounds(neighbor.X, neighbor.Y) || !maze.IsPassable(neighbor) {
				continue
			}

			newCost := costSoFar[current] + 1
			if _, ok := costSoFar[neighbor]; !ok || newCost < costSoFar[neighbor] {
				costSoFar[neighbor] = newCost
				priority := newCost + heuristic(neighbor, end)
				heap.Push(&openSet, PriorityPoint{
					Point:    neighbor,
					Priority: priority,
				})

				cameFrom[neighbor] = current
			}
		}
	}

	return nil, false
}
