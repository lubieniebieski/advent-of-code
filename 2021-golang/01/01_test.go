package main

import (
	"testing"
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
	got, _ := PartOne(ParseIntegersFromString(input))

	if got != want {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}
