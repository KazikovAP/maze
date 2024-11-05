package config

import (
	"log"
	"os"
	"strconv"
)

const defaultValue int = 15

type Config struct {
	MazeWidth            int
	MazeHeight           int
	GenerationAlgorithm  string
	PathfindingAlgorithm string
}

func NewConfig() *Config {
	return &Config{
		MazeWidth:            defaultValue,
		MazeHeight:           defaultValue,
		GenerationAlgorithm:  "Prim",
		PathfindingAlgorithm: "BFS",
	}
}

func (c *Config) Init() {
	if width, exists := os.LookupEnv("MAZE_WIDTH"); exists {
		if w, err := strconv.Atoi(width); err == nil {
			c.MazeWidth = w
		} else {
			log.Println("MAZE_WIDTH is not specified, using default:", c.MazeWidth)
		}
	}

	if height, exists := os.LookupEnv("MAZE_HEIGHT"); exists {
		if h, err := strconv.Atoi(height); err == nil {
			c.MazeHeight = h
		} else {
			log.Println("MAZE_HEIGHT is not specified, using default:", c.MazeHeight)
		}
	}

	if genAlgo, exists := os.LookupEnv("GENERATION_ALGO"); exists {
		c.GenerationAlgorithm = genAlgo
	}

	if pathAlgo, exists := os.LookupEnv("PATHFINDING_ALGO"); exists {
		c.PathfindingAlgorithm = pathAlgo
	}
}
