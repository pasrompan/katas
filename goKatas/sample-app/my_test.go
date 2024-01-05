package kata

import (
	"reflect"
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

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"All words have less than five letters", "This is a test", "This is a test"},
		{"One word has five or more letters", "This is another test", "This is rehtona test"},
		{"All words have five or more letters", "Hello world", "olleH dlrow"},
		{"Empty string", "", ""},
		{"should test that the solution returns the correct value for single word inputs", "Welcome", "emocleW"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.input); got != tt.expected {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.expected)
			}
		})
	}

}

func TestMaxSumSubarray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"All positives", []int{1, 2, 3, 4, 5}, 15},
		{"All negatives", []int{-1, -2, -3, -4, -5}, 0},
		{"Mixed numbers", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"Empty array", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaximumSubarraySum(tt.arr)
			if got != tt.want {
				t.Errorf("MaxSumSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTribonacci(t *testing.T) {
	tests := []struct {
		signature [3]float64
		n         int
		want      []float64
	}{
		{[3]float64{1, 1, 1}, 10, []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105}},
		{[3]float64{0, 0, 1}, 10, []float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}},
		{[3]float64{0, 1, 1}, 0, []float64{}},
	}

	for _, tt := range tests {
		if got := Tribonacci(tt.signature, tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Tribonacci(%v, %v) = %v, want %v", tt.signature, tt.n, got, tt.want)
		}
	}
}

func TestFindNb(t *testing.T) {
	testCases := []struct {
		m        int
		expected int
	}{
		{1071225, 45},
		{91716553919377, -1},
		{4183059834009, 2022},
		{1, 1},
		{9, 2},
	}

	for _, tc := range testCases {
		result := findNb(tc.m)
		if result != tc.expected {
			t.Errorf("findNb(%d) = %d; expected %d", tc.m, result, tc.expected)
		}
	}
}
