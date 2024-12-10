package day10 // Change this for each new day

import (
	"slices"
	"strconv"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type Point struct {
	i, j  int
	value string
}

type Path struct {
	points []Point
}

// Add this new method
func (p Path) WithPoint(point Point) Path {
	newPoints := make([]Point, len(p.points), len(p.points)+1)
	copy(newPoints, p.points)
	return Path{points: append(newPoints, point)}
}

func Explore(grid [][]string, p Point, path Path) []Path {
	currentValue := grid[p.i][p.j]
	val, _ := strconv.Atoi(currentValue)

	if val == 9 {
		return []Path{path}
	}

	var completePaths []Path

	// Only check orthogonal moves (up, down, left, right)
	moves := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, move := range moves {
		newI, newJ := p.i+move[0], p.j+move[1]
		if isValidPoint(grid, newI, newJ) {
			newPoint := Point{newI, newJ, grid[newI][newJ]}
			nextVal, _ := strconv.Atoi(newPoint.value)
			if nextVal == val+1 && !slices.Contains(path.points, newPoint) {
				newPath := path.WithPoint(newPoint)
				paths := Explore(grid, newPoint, newPath)
				completePaths = append(completePaths, paths...)
			}
		}
	}

	return completePaths
}

func UniqueEnds(paths []Path) []Point {
	uniqueEnds := []Point{}
	for _, path := range paths {
		if !slices.Contains(uniqueEnds, path.points[len(path.points)-1]) {
			uniqueEnds = append(uniqueEnds, path.points[len(path.points)-1])
		}
	}
	return uniqueEnds
}

func isValidPoint(grid [][]string, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func Solve1(input []string) int {
	trailsheads := []Point{}
	grid := utils.StringsTo2DArray(input)
	for i, row := range grid {
		for j, cell := range row {
			if cell == "0" {
				point := Point{i, j, cell}
				trailsheads = append(trailsheads, point)
			}
		}
	}
	result := 0
	for _, point := range trailsheads {
		paths := Explore(grid, point, Path{points: []Point{point}})
		result += len(UniqueEnds(paths))
	}

	return result
}

func Solve2(input []string) int {
	trailsheads := []Point{}
	grid := utils.StringsTo2DArray(input)
	for i, row := range grid {
		for j, cell := range row {
			if cell == "0" {
				point := Point{i, j, cell}
				trailsheads = append(trailsheads, point)
			}
		}
	}
	result := 0
	for _, point := range trailsheads {
		paths := Explore(grid, point, Path{points: []Point{point}})
		result += len(paths)
	}

	return result
}
