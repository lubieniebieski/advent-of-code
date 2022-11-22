package tools

import (
	"reflect"
	"testing"
)

func TestExtractIntegersFromString(t *testing.T) {
	input := `1
  2
  3`
	want := []int{1, 2, 3}
	got := ExtractIntegersFromString(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestExtractStringsFromString(t *testing.T) {
	input := `abc
  cba
  bac
  two words`
	want := []string{"abc", "cba", "bac", "two words"}
	got := ExtractStringsFromString(input)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
