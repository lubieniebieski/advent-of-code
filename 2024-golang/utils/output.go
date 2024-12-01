package utils

import "fmt"

// DisplayResults prints the results for both parts of the puzzle
func DisplayResults(day int, part1, part2 interface{}) {
	fmt.Printf("=== Day %02d ===\n", day)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
	fmt.Println("============")
}
