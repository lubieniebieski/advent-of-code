package utils

import (
	"fmt"
	"os"
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

	DisplayResults(day, solution.Part1(input), solution.Part2(input))
}
