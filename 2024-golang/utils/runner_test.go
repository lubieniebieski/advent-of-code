package utils

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestRun(t *testing.T) {
	// Create a temporary directory
	dir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(dir)

	// Create a temporary input file in the directory
	inputFilePath := filepath.Join(dir, "input.txt")
	err = os.WriteFile(inputFilePath, []byte("line1\nline2\nline3\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	// Change the working directory to the temporary directory
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(dir)

	// Define a mock solution
	mockSolution := Solution{
		Part1: func(input []string) int {
			return len(input)
		},
		Part2: func(input []string) int {
			return len(input) * 2
		},
	}

	// Capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	Run(1, mockSolution)

	// Close the writer and read the output
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	// Check the output
	expectedPattern := `=== Day 01 ===
Part 1: 3 \(took.+s\)
Part 2: 6 \(took.+s\)
Total time:.+s
============
`
	if matched, err := regexp.MatchString(expectedPattern, string(out)); err != nil || !matched {
		t.Errorf("Run() output doesn't match expected pattern.\nGot: %v\nExpected pattern: %v", string(out), expectedPattern)
	}
}
