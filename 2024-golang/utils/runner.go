package utils

import (
	"fmt"
	"os"
	"time"
)

type Solution struct {
	Part1 func([]string) int
	Part2 func([]string) int
}

func Run(day int, solution Solution) {
	input, err := ReadInputFile(day)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
		os.Exit(1)
	}

	// Measure Part 1
	start := time.Now()
	part1Value := solution.Part1(input)
	part1 := Result{Value: part1Value, Duration: time.Since(start)}

	// Measure Part 2
	start = time.Now()
	part2Value := solution.Part2(input)
	part2 := Result{Value: part2Value, Duration: time.Since(start)}

	DisplayResults(day, part1, part2)
}
