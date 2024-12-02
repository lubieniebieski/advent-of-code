package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadInput(t *testing.T) {
	// Create a temporary file with test data
	file, err := os.CreateTemp("", "test_input.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString("line1\nline2\nline3\n")
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	file.Close()

	// Test ReadInput function
	lines, err := ReadInput(file.Name())
	if err != nil {
		t.Fatalf("ReadInput() error = %v, want nil", err)
	}
	want := []string{"line1", "line2", "line3"}
	if len(lines) != len(want) {
		t.Fatalf("ReadInput() = %v, want %v", lines, want)
	}
	for i, line := range lines {
		if line != want[i] {
			t.Errorf("ReadInput()[%d] = %v, want %v", i, line, want[i])
		}
	}
}

func TestReadInputFile(t *testing.T) {
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

	// Test ReadInputFile function
	lines, err := ReadInputFile(1)
	if err != nil {
		t.Fatalf("ReadInputFile() error = %v, want nil", err)
	}
	want := []string{"line1", "line2", "line3"}
	if len(lines) != len(want) {
		t.Fatalf("ReadInputFile() = %v, want %v", lines, want)
	}
	for i, line := range lines {
		if line != want[i] {
			t.Errorf("ReadInputFile()[%d] = %v, want %v", i, line, want[i])
		}
	}
}
