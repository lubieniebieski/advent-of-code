package utils

import (
	"io"
	"os"
	"testing"
	"time"
)

func TestDisplayResults(t *testing.T) {
	// Capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Define test cases
	tests := []struct {
		day   int
		part1 Result
		part2 Result
		want  string
	}{
		{
			day: 1,
			part1: Result{
				Value:    42,
				Duration: 2 * time.Second,
			},
			part2: Result{
				Value:    84,
				Duration: 3 * time.Second,
			},
			want: "=== Day 01 ===\nPart 1: 42 (took 2s)\nPart 2: 84 (took 3s)\nTotal time: 5s\n============\n",
		},
		{
			day: 2,
			part1: Result{
				Value:    "foo",
				Duration: 1 * time.Millisecond,
			},
			part2: Result{
				Value:    "bar",
				Duration: 2 * time.Millisecond,
			},
			want: "=== Day 02 ===\nPart 1: foo (took 1ms)\nPart 2: bar (took 2ms)\nTotal time: 3ms\n============\n",
		},
	}

	for _, tt := range tests {
		// Run the function
		DisplayResults(tt.day, tt.part1, tt.part2)

		// Close the writer and read the output
		w.Close()
		out, _ := io.ReadAll(r)
		os.Stdout = old

		// Check the output
		if got := string(out); got != tt.want {
			t.Errorf("DisplayResults() = %v, want %v", got, tt.want)
		}

		// Reset the pipe for the next test case
		r, w, _ = os.Pipe()
		os.Stdout = w
	}
}
