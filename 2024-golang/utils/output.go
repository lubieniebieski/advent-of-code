package utils

import (
	"fmt"
	"time"
)

// Result represents the solution result with its execution time
type Result struct {
	Value    interface{}
	Duration time.Duration
}

// DisplayResults prints the results for both parts of the puzzle with timing information
func DisplayResults(day int, part1, part2 Result) {
	totalTime := part1.Duration + part2.Duration
	fmt.Printf("=== Day %02d ===\n", day)
	fmt.Printf("Part 1: %v (took %v)\n", part1.Value, part1.Duration)
	fmt.Printf("Part 2: %v (took %v)\n", part2.Value, part2.Duration)
	fmt.Printf("Total time: %v\n", totalTime)
	fmt.Println("============")
}

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
