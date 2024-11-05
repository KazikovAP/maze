package generation

import (
	"github.com/KazikovAP/maze/internal/domain"
	"github.com/KazikovAP/maze/pkg/srandom"
)

type Edge struct {
	X1, Y1, X2, Y2 int
}

type DisjointSet struct {
	parent, rank []int
}

type KruskalGenerator struct{}

func (g *KruskalGenerator) Generate(maze *domain.Maze) {
	width, height := maze.Width, maze.Height

	g.initializeWalls(maze)

	ds := g.newDisjointSet(width * height)
	edges := g.generateEdges(width, height)
	g.shuffleEdges(edges)

	g.initializeEmptyCells(maze)
	g.connectCells(maze, ds, edges)
}

func (g *KruskalGenerator) initializeWalls(maze *domain.Maze) {
	for i := 0; i < maze.Height; i++ {
		for j := 0; j < maze.Width; j++ {
			maze.Field[i][j] = '▓'
		}
	}
}

func (g *KruskalGenerator) newDisjointSet(size int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := 0; i < size; i++ {
		ds.parent[i] = i
	}

	return ds
}

func (g *KruskalGenerator) generateEdges(width, height int) []Edge {
	var edges []Edge

	for y := 1; y < height-1; y += 2 {
		for x := 1; x < width-1; x += 2 {
			if x+2 < width {
				edges = append(edges, Edge{x, y, x + 2, y})
			}

			if y+2 < height {
				edges = append(edges, Edge{x, y, x, y + 2})
			}
		}
	}

	return edges
}

func (g *KruskalGenerator) shuffleEdges(edges []Edge) {
	for i := len(edges) - 1; i > 0; i-- {
		j := srandom.Intn(i + 1)
		edges[i], edges[j] = edges[j], edges[i]
	}
}

func (g *KruskalGenerator) initializeEmptyCells(maze *domain.Maze) {
	for y := 1; y < maze.Height; y += 2 {
		for x := 1; x < maze.Width; x += 2 {
			maze.Field[y][x] = ' '
		}
	}
}

func (g *KruskalGenerator) connectCells(maze *domain.Maze, ds *DisjointSet, edges []Edge) {
	for _, edge := range edges {
		cell1 := edge.Y1*maze.Width + edge.X1
		cell2 := edge.Y2*maze.Width + edge.X2

		if ds.find(cell1) != ds.find(cell2) {
			ds.union(cell1, cell2)

			maze.Field[(edge.Y1+edge.Y2)/2][(edge.X1+edge.X2)/2] = ' '
		}
	}
}

func (ds *DisjointSet) find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.find(ds.parent[x])
	}

	return ds.parent[x]
}

func (ds *DisjointSet) union(x, y int) {
	rootX := ds.find(x)
	rootY := ds.find(y)

	if rootX != rootY {
		switch {
		case ds.rank[rootX] > ds.rank[rootY]:
			ds.parent[rootY] = rootX
		case ds.rank[rootX] < ds.rank[rootY]:
			ds.parent[rootX] = rootY
		default:
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}
