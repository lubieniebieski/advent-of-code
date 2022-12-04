package main

import (
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestPartOne(t *testing.T) {
	want := 24000
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := 45000
	got := PartTwo(testData())
	toolstest.CompareWithExample(t, want, got)
}

func testData() []string {
	return tools.ExtractStringsFromString(testInput())
}

func testInput() string {
	return `1000
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
}
