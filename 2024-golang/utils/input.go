package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// ReadInput reads the input file and returns a slice of strings
func ReadInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
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

// ReadInputFile reads the input file for a specific day
func ReadInputFile(day int) ([]string, error) {
	// Start with the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %v", err)
	}

	// Try to find input.txt in the current directory first
	inputPath := filepath.Join(currentDir, "input.txt")
	if _, err := os.Stat(inputPath); err == nil {
		return ReadInput(inputPath)
	}

	// If not found and we're in the root directory, try the day directory
	dayDir := fmt.Sprintf("day%02d", day)
	inputPath = filepath.Join(currentDir, dayDir, "input.txt")

	return ReadInput(inputPath)
}
