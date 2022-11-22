package main

import (
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022/tools"
)

func TestPartOne(t *testing.T) {
	input := `199
  200
  208
  210
  200
  207
  240
  269
  260
  263
  `
	want := 7
	got, _ := PartOne(tools.ExtractIntegersFromString(input))
	compareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	input := `199
  200
  208
  210
  200
  207
  240
  269
  260
  263
  `
	want := 5
	got, _ := PartTwo(tools.ExtractIntegersFromString(input))
	compareWithExample(t, want, got)
}

func compareWithExample(t *testing.T, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("According to the example, we should have %d, but we've got %d", want, got)
	}
}
