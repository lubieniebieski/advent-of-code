package main

import (
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

func TestPartOne(t *testing.T) {
	input := `1000
  2000
  3000

  4000

  5000
  6000

  7000
  8000
  9000

  10000
  `
	want := 24000
	got, _ := PartOne(tools.ExtractStringsFromString(input))
	compareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	input := `1000
  2000
  3000

  4000

  5000
  6000

  7000
  8000
  9000

  10000
  `
	want := 45000
	got, _ := PartTwo(tools.ExtractStringsFromString(input))
	compareWithExample(t, want, got)
}

func compareWithExample(t *testing.T, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("According to the example, we should have %d, but we've got %d", want, got)
	}
}
