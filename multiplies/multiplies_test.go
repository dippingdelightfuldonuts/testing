package multiplies

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	x := 67
	y := 89

	strategies := []string{
		"egyptian",
		"normal",
		"russian",
		"long",
		"addition",
		"lattice",
	}

	for _, sgy := range strategies {
		result := Multiply(x, y, sgy)

		if result != 5963 {
			t.Errorf("Multiply was incorrect for %s, got: %d, want: %d", sgy, result, 5963)
		}
	}
}

type mockStrategy struct{}

func (mock *mockStrategy) Multiply(x, y int) int {
	return 87
}

func TestMultiplyGivenStrategy(t *testing.T) {
	mock := &mockStrategy{}

	result := MultiplyGivenStrategy(5, 5, mock)

	if result != 87 {
		t.Error("Multiply did not call strategy")
	}
}
