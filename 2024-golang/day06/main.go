package day06

import "github.com/lubieniebieski/advent-of-code/2024-golang/utils"

func Solve1(input []string) int {
	grid := utils.StringsTo2DArray(input)
	var startI, startJ int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "^" {
				startI, startJ = i, j
				grid[i][j] = "1" // Mark start position
			}
		}
	}
	LookUp(grid, startI, startJ)

	// Count all "1"s in the grid
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "1" {
				sum++
			}
		}
	}
	return sum
}

func LookUp(grid [][]string, i, j int) {
	if i-1 < 0 {
		return
	}
	next := grid[i-1][j]
	if next == "#" {
		LookRight(grid, i, j)
		return
	}
	grid[i-1][j] = "1"
	LookUp(grid, i-1, j)
}

func LookRight(grid [][]string, i, j int) {
	if j+1 >= len(grid[i]) {
		return
	}
	next := grid[i][j+1]
	if next == "#" {
		LookDown(grid, i, j)
		return
	}
	grid[i][j+1] = "1"
	LookRight(grid, i, j+1)
}

func LookDown(grid [][]string, i, j int) {
	if i+1 >= len(grid) {
		return
	}
	next := grid[i+1][j]
	if next == "#" {
		LookLeft(grid, i, j)
		return
	}
	grid[i+1][j] = "1"
	LookDown(grid, i+1, j)
}

func LookLeft(grid [][]string, i, j int) {
	if j-1 < 0 {
		return
	}
	next := grid[i][j-1]
	if next == "#" {
		LookUp(grid, i, j)
		return
	}
	grid[i][j-1] = "1"
	LookLeft(grid, i, j-1)
}
