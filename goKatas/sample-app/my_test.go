package kata

import (
	"fmt"
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

func TestRGB(t *testing.T) {
	tests := []struct {
		r        int
		g        int
		b        int
		expected string
	}{
		{255, 255, 255, "FFFFFF"},
		{0, 0, 0, "000000"},
		{255, 0, 0, "FF0000"},
		{0, 255, 0, "00FF00"},
		{0, 0, 255, "0000FF"},
		{123, 45, 67, "7B2D43"},
		{256, 0, 0, "FF0000"},
		{0, 256, 0, "00FF00"},
		{0, 0, 256, "0000FF"},
		{-1, 0, 0, "000000"},
		{0, -1, 0, "000000"},
		{0, 0, -1, "000000"},
		{100, 200, 300, "64C8FF"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("RGB(%d, %d, %d)", tt.r, tt.g, tt.b), func(t *testing.T) {
			result := RGB(tt.r, tt.g, tt.b)
			if result != tt.expected {
				t.Errorf("RGB(%d, %d, %d) = %s; expected %s", tt.r, tt.g, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBeeramid(t *testing.T) {
	tests := []struct {
		bonus    int
		price    float64
		expected int
	}{
		{9, 2.0, 1},
		{21, 1.5, 3},
		{-1, 4, 0},
	}

	for _, tt := range tests {
		result := Beeramid(tt.bonus, tt.price)
		if result != tt.expected {
			t.Errorf("Beeramid(%d, %f) = %d; expected %d", tt.bonus, tt.price, result, tt.expected)
		}
	}
}

func TestValidISBN10_ValidISBN(t *testing.T) {
	isbn := "0471958697"
	expected := true

	result := ValidISBN10(isbn)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestValidISBN10_InvalidISBN(t *testing.T) {
	isbn := "1112223339"
	expected := true

	result := ValidISBN10(isbn)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestValidISBN10_InvalidCharacter(t *testing.T) {
	isbn := "X123456788"
	expected := false

	result := ValidISBN10(isbn)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestValidISBN10_InvalidLength(t *testing.T) {
	isbn := "047195869"
	expected := false

	result := ValidISBN10(isbn)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestValidISBN10_ValidLowerCaseX(t *testing.T) {
	isbn := "048665088x"
	expected := true

	result := ValidISBN10(isbn)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestLCS(t *testing.T) {
	tests := []struct {
		x        string
		y        string
		expected string
	}{
		{"", "", ""},
		{"abc", "", ""},
		{"", "def", ""},
		{"abc", "def", ""},
		{"abc", "abc", "abc"},
		{"abc", "defg", ""},
		{"abcdef", "def", "def"},
		{"abcdef", "abcf", "abcf"},
		{"abcdef", "acdf", "acdf"},
		{"abcdef", "abcdef", "abcdef"},
		{"abcdef", "aceg", "ace"},
		{"abcdef", "acef", "acef"},
	}

	for _, tt := range tests {
		result := LCS(tt.x, tt.y)
		if result != tt.expected {
			t.Errorf("LCS(%s, %s) = %s; expected %s", tt.x, tt.y, result, tt.expected)
		}
	}
}
