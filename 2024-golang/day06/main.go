package day06

import (
	"runtime"
	"sync"
	"time"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type Position struct {
	row, col int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var movements = map[Direction]Position{
	Up:    {-1, 0},
	Right: {0, 1},
	Down:  {1, 0},
	Left:  {0, -1},
}

func rotateRight(d Direction) Direction {
	return Direction((int(d) + 1) % 4)
}

type PathState struct {
	pos Position
	dir Direction
}

func followPath(grid [][]string, start Position, blocked Position) (bool, map[Position]bool) {
	visited := make(map[Position]bool)
	states := make(map[PathState]bool)
	current := PathState{start, Up}

	rows, cols := len(grid), len(grid[0])

	for {
		// Early return if we've seen this state
		if states[current] {
			return true, visited
		}

		visited[current.pos] = true
		states[current] = true

		// Calculate next position
		move := movements[current.dir]
		nextRow := current.pos.row + move.row
		nextCol := current.pos.col + move.col

		// Fast boundary check
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
			return false, visited
		}

		// Fast blocked check
		nextPos := Position{nextRow, nextCol}
		if grid[nextRow][nextCol] == "#" || (nextPos.row == blocked.row && nextPos.col == blocked.col) {
			current.dir = rotateRight(current.dir)
		} else {
			current.pos = nextPos
		}
	}
}

func findStart(grid [][]string) Position {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "^" {
				return Position{i, j}
			}
		}
	}
	return Position{} // or handle error
}

func Solve1(input []string) int {
	grid := utils.StringsTo2DArray(input)
	start := findStart(grid)

	_, path := followPath(grid, start, Position{-1, -1})
	return len(path)
}

func Solve2(input []string) int {
	grid := utils.StringsTo2DArray(input)
	startPos := findStart(grid)

	// Get original path without blocks
	_, origPath := followPath(grid, startPos, Position{-1, -1})

	var positions []Position
	for pos := range origPath {
		if pos == startPos {
			continue
		}
		positions = append(positions, pos)
	}

	// Create result channel and wait group
	results := make(chan bool, len(positions))
	var wg sync.WaitGroup

	// Calculate number of workers
	workers := runtime.NumCPU()
	if workers > len(positions) {
		workers = len(positions)
	}

	// Distribute work more evenly
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// Each worker processes positions at index workerID, workerID+workers, workerID+2*workers, etc.
			for j := workerID; j < len(positions); j += workers {
				isLoop, _ := followPath(grid, startPos, positions[j])
				results <- isLoop
			}
		}(i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(results)
		close(done)
	}()

	loops := 0
	select {
	case <-done:
		for isLoop := range results {
			if isLoop {
				loops++
			}
		}
	case <-time.After(5 * time.Second):
		return 0
	}

	return loops
}
