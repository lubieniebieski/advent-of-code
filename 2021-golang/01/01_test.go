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
	got, _ := PartOne(tools.ParseIntegersFromString(input))

	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
