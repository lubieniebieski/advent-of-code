package main

import (
	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func getAllStringsFromPoint(grid [][]string, i, j, n int) []string {
	strings := []string{}

	// Horizontal left
	if j-n >= 0 {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i][j-k]
		}
		strings = append(strings, str)
	}

	// Horizontal right
	if j+n < len(grid[0]) {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i][j+k]
		}
		strings = append(strings, str)
	}

	// Vertical up
	if i-n >= 0 {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i-k][j]
		}
		strings = append(strings, str)
	}

	// Vertical down
	if i+n < len(grid) {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i+k][j]
		}
		strings = append(strings, str)
	}

	// Diagonal up-left
	if i-n >= 0 && j-n >= 0 {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i-k][j-k]
		}
		strings = append(strings, str)
	}

	// Diagonal up-right
	if i-n >= 0 && j+n < len(grid[0]) {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i-k][j+k]
		}
		strings = append(strings, str)
	}

	// Diagonal down-left
	if i+n < len(grid) && j-n >= 0 {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i+k][j-k]
		}
		strings = append(strings, str)
	}

	// Diagonal down-right
	if i+n < len(grid) && j+n < len(grid[0]) {
		str := ""
		for k := 0; k <= n; k++ {
			str += grid[i+k][j+k]
		}
		strings = append(strings, str)
	}

	return strings
}

func isIsOkSam(grid [][]string, i, j int) bool {
	if i-1 < 0 || i+1 >= len(grid) || j-1 < 0 || j+1 >= len(grid[0]) {
		return false
	}
	return ((grid[i-1][j-1] == "S" && grid[i+1][j+1] == "M") ||
		(grid[i-1][j-1] == "M" && grid[i+1][j+1] == "S")) &&
		((grid[i+1][j-1] == "S" && grid[i-1][j+1] == "M") ||
			(grid[i+1][j-1] == "M" && grid[i-1][j+1] == "S"))

}

func solve1(input []string) int {
	sum := 0

	grid := utils.StringsTo2DArray(input)
	maxI := len(grid)
	maxJ := len(grid[0])

	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {
			if grid[i][j] == "X" {
				strings := getAllStringsFromPoint(grid, i, j, 3)
				for _, s := range strings {
					if s == "XMAS" {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func solve2(input []string) int {
	sum := 0

	grid := utils.StringsTo2DArray(input)
	maxI := len(grid)
	maxJ := len(grid[0])

	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {
			if grid[i][j] == "A" {
				if isIsOkSam(grid, i, j) {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	utils.Run(4, utils.Solution{Part1: solve1, Part2: solve2}) // Just update the day number
}
