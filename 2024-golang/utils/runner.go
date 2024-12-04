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

func RunAndStore(day int, solution Solution, input []string) DayResults {
	start := time.Now()
	part1Value := solution.Part1(input)
	part1 := Result{Value: part1Value, Duration: time.Since(start)}

	start = time.Now()
	part2Value := solution.Part2(input)
	part2 := Result{Value: part2Value, Duration: time.Since(start)}

	return DayResults{
		Day:       day,
		Part1:     part1,
		Part2:     part2,
		TotalTime: (part1.Duration + part2.Duration).String(),
	}
}

// Map of available solutions
var solutions = make(map[int]Solution)

func RegisterSolution(day int, solution Solution) {
	solutions[day] = solution
}

func GetSolution(day int) (Solution, bool) {
	sol, ok := solutions[day]
	return sol, ok
}
