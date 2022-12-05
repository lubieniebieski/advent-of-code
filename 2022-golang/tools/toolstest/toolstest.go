package toolstest

import "testing"

func CompareWithExample(t *testing.T, want, got interface{}) {
	t.Helper()

	if got != want {
		t.Errorf("According to the example, we should have %v, but we've got %v", want, got)
	}
}
