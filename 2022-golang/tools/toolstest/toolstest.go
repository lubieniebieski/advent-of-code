package toolstest

import "testing"

func CompareWithExample(t *testing.T, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("According to the example, we should have %d, but we've got %d", want, got)
	}
}
