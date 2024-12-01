package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve1(input []string) int {
	// TODO: Implement solution for part 1
	return 0
}

func solve2(input []string) int {
	// TODO: Implement solution for part 2
	return 0
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
		os.Exit(1)
	}

	result1 := solve1(input)
	fmt.Printf("Part 1: %d\n", result1)

	result2 := solve2(input)
	fmt.Printf("Part 2: %d\n", result2)
}
