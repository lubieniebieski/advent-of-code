package day08 // Change this for each new day

import (
	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type Point struct {
	i int
	j int
}

func findPoints(char string, grid [][]string) []Point {
	points := []Point{}
	for i, row := range grid {
		for j, cell := range row {
			if cell == char {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func antinodesForPair(pair []Point, grid [][]string) []Point {
	antinodes := []Point{}
	rows, cols := len(grid), len(grid[0])

	di := pair[1].i - pair[0].i
	dj := pair[1].j - pair[0].j

	beforePoint := Point{pair[0].i - di, pair[0].j - dj}
	afterPoint := Point{pair[1].i + di, pair[1].j + dj}

	if beforePoint.i >= 0 && beforePoint.i < rows &&
		beforePoint.j >= 0 && beforePoint.j < cols {
		antinodes = append(antinodes, beforePoint)
	}

	if afterPoint.i >= 0 && afterPoint.i < rows &&
		afterPoint.j >= 0 && afterPoint.j < cols {
		antinodes = append(antinodes, afterPoint)
	}

	return antinodes
}

func antinodesForPairV2(pair []Point, grid [][]string) []Point {
	antinodes := []Point{}
	rows, cols := len(grid), len(grid[0])

	// Calculate step size
	di := pair[1].i - pair[0].i
	dj := pair[1].j - pair[0].j

	// Check each integer multiple of the step
	multiplier := 1
	for {
		// Try next point after the pair
		nextI := pair[1].i + (di * multiplier)
		nextJ := pair[1].j + (dj * multiplier)

		// Break if out of bounds
		if nextI < 0 || nextI >= rows || nextJ < 0 || nextJ >= cols {
			break
		}

		antinodes = append(antinodes, Point{nextI, nextJ})
		multiplier++
	}

	// Check each integer multiple in the opposite direction
	multiplier = 1
	for {
		// Try previous point before the pair
		prevI := pair[0].i - (di * multiplier)
		prevJ := pair[0].j - (dj * multiplier)

		// Break if out of bounds
		if prevI < 0 || prevI >= rows || prevJ < 0 || prevJ >= cols {
			break
		}

		antinodes = append(antinodes, Point{prevI, prevJ})
		multiplier++
	}

	return antinodes
}

func createPairs(char string, grid [][]string) [][]Point {
	points := findPoints(char, grid)
	pairs := [][]Point{}
	for i, point := range points {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, []Point{point, points[j]})
		}
	}
	return pairs
}

func Solve1(input []string) int {
	grid := utils.StringsTo2DArray(input)

	uniqueChars := make(map[string]bool)
	for _, row := range grid {
		for _, char := range row {
			if char != "." {
				uniqueChars[char] = true
			}
		}
	}
	antinodes := make(map[Point]bool)
	for key := range uniqueChars {
		pairs := createPairs(key, grid)
		for _, pair := range pairs {
			for _, antinode := range antinodesForPair(pair, grid) {
				antinodes[antinode] = true
			}
		}
	}

	return len(antinodes)

}

func Solve2(input []string) int {
	grid := utils.StringsTo2DArray(input)

	antinodes := make(map[Point]bool)
	uniqueChars := make(map[string]bool)
	for i, row := range grid {
		for j, char := range row {
			if char != "." {
				uniqueChars[char] = true
				antinodes[Point{i, j}] = true
			}
		}
	}
	for key := range uniqueChars {
		pairs := createPairs(key, grid)
		for _, pair := range pairs {
			for _, antinode := range antinodesForPairV2(pair, grid) {
				antinodes[antinode] = true
			}
		}
	}
	for antinode := range antinodes {
		grid[antinode.i][antinode.j] = "#"
	}

	return len(antinodes)
}
