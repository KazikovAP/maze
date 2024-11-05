package domain

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type Point struct {
	X, Y int
}

type Maze struct {
	Field  [][]rune
	Width  int
	Height int
}

func (m *Maze) IsPassable(point Point) bool {
	return m.Field[point.Y][point.X] != '▓'
}

func (m *Maze) SetPathSymbol(point Point) {
	if m.Field[point.Y][point.X] != 'S' && m.Field[point.Y][point.X] != 'E' {
		m.Field[point.Y][point.X] = '*'
	}
}

func (m *Maze) IsAccessibleAndUnvisited(point Point, visited map[Point]bool) bool {
	return (m.Field[point.Y][point.X] == ' ' || m.Field[point.Y][point.X] == 'E') && !visited[point]
}

func CreateMaze(width, height int) *Maze {
	field := make([][]rune, height)
	for i := range field {
		field[i] = make([]rune, width)
		for j := range field[i] {
			field[i][j] = '▓'
		}
	}

	return &Maze{Field: field, Width: width, Height: height}
}

func (m *Maze) InBounds(x, y int) bool {
	return x > 0 && x < m.Width-1 && y > 0 && y < m.Height-1
}

func (m *Maze) SetStartAndEndPoints() (start, end Point) {
	start = Point{1, randomInt(1, m.Height-2)}
	end = Point{m.Width - 2, randomInt(1, m.Height-2)}

	m.Field[start.Y][start.X] = 'S'
	m.Field[end.Y][end.X] = 'E'

	fmt.Printf("Start: %v, End: %v\n", start, end)

	return start, end
}

func (m *Maze) Output() string {
	var builder strings.Builder

	builder.WriteString("\n")

	for _, row := range m.Field {
		for _, cell := range row {
			switch cell {
			case '*':
				builder.WriteString("\033[42m \033[0m")
			default:
				builder.WriteRune(cell)
			}
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

func randomInt(minVal, maxVal int) int {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(maxVal-minVal+1)))
	return minVal + int(nBig.Int64())
}
