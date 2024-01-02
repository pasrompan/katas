package kata

import (
	"testing"
)

func TestIsSolved(t *testing.T) {
	var board [3][3]int
	board = [3][3]int{
		{0, 0, 1},
		{0, 1, 2},
		{2, 1, 0},
	}

	result := IsSolved(board)

	if result != -1 { // replace 0 with the expected result
		t.Errorf("Expected -1, but got %d", result)
	}
}

func TestRowin1(t *testing.T) {
	board := [3][3]int{{1, 1, 1}, {0, 1, 2}, {2, 1, 0}}

	result := IsSolved(board)

	if result != 1 { // replace 0 with the expected result
		t.Errorf("Expected 1, but got %d", result)
	}
}

func Test2Rowin1(t *testing.T) {
	board := [3][3]int{{1, 1, 2}, {1, 1, 1}, {2, 1, 0}}

	result := IsSolved(board)

	if result != 1 { // replace 0 with the expected result
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestColumnin1(t *testing.T) {
	board := [3][3]int{{1, 0, 1}, {1, 1, 2}, {1, 2, 0}}

	result := IsSolved(board)

	if result != 1 { // replace 0 with the expected result
		t.Errorf("Expected 1, but got %d", result)
	}
}

func Test2Columnin1(t *testing.T) {
	board := [3][3]int{{1, 1, 1}, {2, 1, 2}, {1, 1, 0}}

	result := IsSolved(board)

	if result != 1 { // replace 0 with the expected result
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestDiagonal(t *testing.T) {
	board := [3][3]int{{1, 2, 1}, {2, 1, 2}, {2, 0, 1}}

	result := IsSolved(board)

	if result != 1 { // replace 0 with the expected result
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestInvDiagonal(t *testing.T) {
	board := [3][3]int{{1, 2, 1}, {2, 1, 2}, {1, 0, 2}}

	result := IsSolved(board)

	if result != 1 { // replace 0 with the expected result
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestMiddleColumn(t *testing.T) {
	board := [3][3]int{{2, 1, 2}, {2, 1, 1}, {1, 1, 2}}

	result := IsSolved(board)

	if result != 1 { // replace 0 with the expected result
		t.Errorf("Expected 1, but got %d", result)
	}
}

func TestRandom(t *testing.T) {
	board := [3][3]int{{0, 2, 1}, {1, 0, 2}, {1, 2, 0}}

	result := IsSolved(board)

	if result != -1 { // replace 0 with the expected result
		t.Errorf("Expected -1, but got %d", result)
	}
}

func TestHumanReadable(t *testing.T) {
	result := HumanReadableTime(0)

	if result != "00:00:00" { // replace 0 with the expected result
		t.Errorf("Expected 00:00:00, but got %s", result)
	}
}

func TestHumanReadableTime(t *testing.T) {
	tests := []struct {
		name     string
		seconds  int
		expected string
	}{
		{"Zero seconds", 0, "00:00:00"},
		{"One minute", 60, "00:01:00"},
		{"One hour", 3600, "01:00:00"},
		{"One day", 86400, "24:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HumanReadableTime(tt.seconds)
			if got != tt.expected {
				t.Errorf("HumanReadableTime(%d) = %s; want %s", tt.seconds, got, tt.expected)
			}
		})
	}
}
