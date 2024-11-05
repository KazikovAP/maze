package application

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/KazikovAP/maze/config"
	"github.com/KazikovAP/maze/internal/domain"
	"github.com/KazikovAP/maze/internal/domain/generation"
	"github.com/KazikovAP/maze/internal/domain/pathfinding"
	"github.com/KazikovAP/maze/internal/infrastructure"
)

type MazeGenerator interface {
	Generate(maze *domain.Maze)
}

type Pathfinder interface {
	FindPath(maze *domain.Maze, start, end domain.Point) ([]domain.Point, bool)
}

type ioAdapter interface {
	Input() (string, error)
	Output(content string)
}

type App struct {
	maze       *domain.Maze
	io         ioAdapter
	cfg        *config.Config
	generator  MazeGenerator
	pathfinder Pathfinder
}

func NewApp(cfg *config.Config, io *infrastructure.IOAdapter) *App {
	return &App{
		cfg: cfg,
		io:  io,
	}
}

func (a *App) Start() error {
	a.greetUser()
	pathAlgo := a.initializeMaze()

	start, end := a.maze.SetStartAndEndPoints()
	path, found := a.pathfinder.FindPath(a.maze, start, end)
	a.outputPathResult(found, pathAlgo)

	if found {
		a.markPath(path)
	}

	a.io.Output(a.maze.Output())

	return nil
}

func (a *App) initializeMaze() string {
	width := a.validateDimension(a.InputWidth(a.cfg.MazeWidth), 60)
	height := a.validateDimension(a.InputHeight(a.cfg.MazeHeight), 30)
	genAlgo := a.InputGenerationAlgorithm(a.cfg.GenerationAlgorithm)
	pathAlgo := a.InputPathfindingAlgorithm(a.cfg.PathfindingAlgorithm)

	a.maze = domain.CreateMaze(width, height)
	a.selectGenerator(genAlgo)
	a.generator.Generate(a.maze)
	a.selectPathfinder(pathAlgo)

	a.summaryInformation(width, height, genAlgo, pathAlgo)

	return pathAlgo
}

func (a *App) InputWidth(defaultWidth int) int {
	a.io.Output(fmt.Sprintf("Enter maze width (default: %d): ", defaultWidth))

	input, err := a.io.Input()
	if err != nil || input == "" {
		return defaultWidth
	}

	width, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		a.io.Output("\"Invalid width input, using default = 15\"\n")
		return defaultWidth
	}

	return width
}

func (a *App) InputHeight(defaultHeight int) int {
	a.io.Output(fmt.Sprintf("Enter maze height (default: %d): ", defaultHeight))

	input, err := a.io.Input()
	if err != nil || input == "" {
		return defaultHeight
	}

	height, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		a.io.Output("\"Invalid height input, using default = 15\"\n")
		return defaultHeight
	}

	return height
}

func (a *App) InputGenerationAlgorithm(defaultAlgo string) string {
	a.io.Output(fmt.Sprintf("Choose maze generation algorithm (1 - Prim, 2 - Kruskal, default %s): ", defaultAlgo))

	input, err := a.io.Input()
	if err != nil || input == "" {
		return defaultAlgo
	}

	switch strings.TrimSpace(input) {
	case "1":
		return "Prim"
	case "2":
		return "Kruskal"
	default:
		a.io.Output("\"Invalid choice, using default Prim generation algorithm\"\n")
		return defaultAlgo
	}
}

func (a *App) InputPathfindingAlgorithm(defaultAlgo string) string {
	a.io.Output(fmt.Sprintf("Choose pathfinding algorithm (1 - BFS, 2 - A*, default %s): ", defaultAlgo))

	input, err := a.io.Input()
	if err != nil || input == "" {
		return defaultAlgo
	}

	switch strings.TrimSpace(input) {
	case "1":
		return "BFS"
	case "2":
		return "A*"
	default:
		a.io.Output("\"Invalid choice, using default BFS pathfinding algorithm\"\n")
		return defaultAlgo
	}
}

func (a *App) selectGenerator(genAlgo string) {
	switch genAlgo {
	case "Prim":
		a.generator = &generation.PrimGenerator{}
	case "Kruskal":
		a.generator = &generation.KruskalGenerator{}
	default:
		a.io.Output("Invalid generator choice, using default Prim generator.\n")
		a.generator = &generation.PrimGenerator{}
	}
}

func (a *App) selectPathfinder(pathAlgo string) {
	switch pathAlgo {
	case "BFS":
		a.pathfinder = &pathfinding.Bfs{}
	case "A*":
		a.pathfinder = &pathfinding.AStar{}
	default:
		a.io.Output("Invalid pathfinder choice, using default BFS.\n")
		a.pathfinder = &pathfinding.Bfs{}
	}
}

func (a *App) validateDimension(dimension, valMax int) int {
	if dimension < 3 {
		return 3
	}

	if dimension > valMax {
		return valMax
	}

	return dimension
}

func (a *App) greetUser() {
	a.io.Output("Welcome to the Maze Game!\n")
}

func (a *App) summaryInformation(width, height int, genAlgo, pathAlgo string) {
	a.io.Output(fmt.Sprintf("Generating a %dx%d maze using %s generation and %s pathfinding\n", width, height, genAlgo, pathAlgo))
}

func (a *App) outputPathResult(found bool, pathAlgo string) {
	if found {
		a.io.Output(fmt.Sprintf("Path found using %s!\n", pathAlgo))
	} else {
		a.io.Output(fmt.Sprintf("No path found using %s.\n", pathAlgo))
	}
}

func (a *App) markPath(path []domain.Point) {
	for _, p := range path {
		a.maze.SetPathSymbol(p)
	}
}
