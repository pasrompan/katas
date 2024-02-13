package kata

import (
	"fmt"
	"math/big"
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

func TestDeterminant(t *testing.T) {
	tests := []struct {
		name string
		m    [][]int
		want int
	}{
		{
			name: "2x2 matrix",
			m:    [][]int{{1}},
			want: 1,
		},
		{
			name: "2x2 matrix",
			m:    [][]int{{1, 2}, {3, 4}},
			want: -2,
		},
		{
			name: "3x3 matrix",
			m:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: 0,
		},
		{
			name: "3x3 matrix",
			m:    [][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}},
			want: -20,
		},
		//Expect(Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}})).To(Equal(-20)) Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Determinant(tt.m); got != tt.want {
				t.Errorf("Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		x        int64
		expected string
	}{
		{0, "now"},
		{30, "30 seconds"},
		{150, "2 minutes and 30 seconds"},
		{3720, "1 hour and 2 minutes"},
		{90000, "1 day and 1 hour"},
		{90001, "1 day, 1 hour and 1 second"},
	}

	for _, tt := range tests {
		result := FormatDuration(tt.x)
		if result != tt.expected {
			t.Errorf("FormatDuration(%d) = %s; expected %s", tt.x, result, tt.expected)
		}
	}
}
func TestFib(t *testing.T) {
	tests := []struct {
		n        int64
		expected *big.Int
	}{
		{0, big.NewInt(0)},
		{1, big.NewInt(1)},
		{2, big.NewInt(1)},
		{3, big.NewInt(2)},
		{4, big.NewInt(3)},
		{5, big.NewInt(5)},
		{6, big.NewInt(8)},
		{7, big.NewInt(13)},
		{8, big.NewInt(21)},
		{9, big.NewInt(34)},
		{10, big.NewInt(55)},
		{11, big.NewInt(89)},
		{12, big.NewInt(144)},
		{13, big.NewInt(233)},
		{14, big.NewInt(377)},
		{15, big.NewInt(610)},
	}

	for _, tt := range tests {
		t.Run("Fib("+string(rune(tt.n))+")", func(t *testing.T) {
			result := fib(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Fib(%d) = %v; expected %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestPermutationalPrimes(t *testing.T) {
	tests := []struct {
		name   string
		nMax   int
		kPerms int
		want   [3]int
	}{
		{
			name:   "Test Case 1",
			nMax:   1000,
			kPerms: 1,
			want:   [3]int{34, 13, 797},
		},
		{
			name:   "Test Case 2",
			nMax:   1000,
			kPerms: 2,
			want:   [3]int{9, 113, 389},
		},
		{
			name:   "Test Case 3",
			nMax:   1000,
			kPerms: 3,
			want:   [3]int{3, 149, 379},
		},
		// add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermutationalPrimes(tt.nMax, tt.kPerms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermutationalPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListPosition(t *testing.T) {
	tests := []struct {
		word     string
		expected *big.Int
	}{
		{"ABAB", big.NewInt(2)},
		{"AAAB", big.NewInt(1)},
		{"BAAA", big.NewInt(4)},
		{"QUESTION", big.NewInt(24572)},
		{"BOOKKEEPER", big.NewInt(10743)},
	}

	for _, test := range tests {
		if got := ListPosition(test.word); got.Cmp(test.expected) != 0 {
			t.Errorf("ListPosition(%s) = %s, want %s", test.word, got.String(), test.expected.String())
		}
	}
}

func TestUniquePermutationsCount(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"abc", 6},
		{"aab", 3},
		{"aaa", 1},
		{"abab", 6},
		{"", 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %s", test.s), func(t *testing.T) {
			result := UniquePermutationsCount(test.s)
			if result != test.expected {
				t.Errorf("UniquePermutationsCount(%s) = %d, expected %d", test.s, result, test.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, false},
		{17, true},
		{18, false},
		{19, true},
		{20, false},
		{21, false},
		{22, false},
		{23, true},
		{24, false},
		{25, false},
		{26, false},
		{27, false},
		{28, false},
		{29, true},
		{30, false},
		{31, true},
		{32, false},
		{33, false},
		{34, false},
		{35, false},
		{36, false},
		{37, true},
		{38, false},
		{39, false},
		{40, false},
		{41, true},
		{42, false},
		{43, true},
		{44, false},
		{45, false},
		{46, false},
		{47, true},
		{48, false},
		{49, false},
		{50, false},
		{51, false},
		{52, false},
		{53, true},
		{54, false},
		{55, false},
		{56, false},
		{57, false},
		{58, false},
		{59, true},
		{60, false},
		{61, true},
		{62, false},
		{63, false},
		{64, false},
		{65, false},
		{66, false},
		{67, true},
		{68, false},
		{69, false},
		{70, false},
		{71, true},
		{72, false},
		{73, true},
		{74, false},
		{75, false},
		{76, false},
		{77, false},
		{78, false},
		{79, true},
		{80, false},
		{81, false},
		{82, false},
		{83, true},
		{84, false},
		{85, false},
		{86, false},
		{87, false},
		{88, false},
		{89, true},
		{90, false},
		{91, false},
		{92, false},
		{93, false},
		{94, false},
		{95, false},
		{96, false},
		{97, true},
		{98, false},
		{99, false},
		{100, false},
	}

	for _, tt := range tests {
		t.Run("isPrime("+string(rune(tt.n))+")", func(t *testing.T) {
			result := isPrime(tt.n)
			if result != tt.expected {
				t.Errorf("isPrime(%d) = %v; expected %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestToJadenCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"this is a test", "This Is A Test"},
		{"jaden case", "Jaden Case"},
		{"", ""},
		{"all lowercase", "All Lowercase"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := toJadenCase(tt.input)
			if result != tt.expected {
				t.Errorf("toJadenCase(%s) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDigitalRoot(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{9, 9},
		{10, 1},
		{38, 2},
		{123, 6},
		{9876, 3},
		{123456789, 9},
		{999999999, 9},
	}

	for _, tt := range tests {
		t.Run("DigitalRoot("+string(rune(tt.n))+")", func(t *testing.T) {
			result := DigitalRoot(tt.n)
			if result != tt.expected {
				t.Errorf("DigitalRoot(%d) = %d; expected %d", tt.n, result, tt.expected)
			}
		})
	}
}

func TestSpinWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello world", "olleH dlrow"},
		{"Welcome", "emocleW"},
		{"This is a test", "This is a test"},
		{"Spin words longer than five", "Spin sdrow regnol than five"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := SpinWords(tt.input)
			if result != tt.expected {
				t.Errorf("SpinWords(%s) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountBits(t *testing.T) {
	tests := []struct {
		n        uint
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{10, 2},
		{15, 4},
		{16, 1},
		{100, 3},
	}

	for _, tt := range tests {
		result := CountBits(tt.n)
		if result != tt.expected {
			t.Errorf("CountBits(%d) = %d; expected %d", tt.n, result, tt.expected)
		}
	}
}

func TestFindOdd(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{
			name:     "Test Case 1",
			arr:      []int{1, 1, 2, 2, 3},
			expected: 3,
		},
		{
			name:     "Test Case 2",
			arr:      []int{1, 1, 2, 2, 3, 3, 3},
			expected: 3,
		},
		{
			name:     "Test Case 3",
			arr:      []int{1, 2, 2, 3, 3, 3},
			expected: 1,
		},
		{
			name:     "Test Case 4",
			arr:      []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindOdd(tt.arr)
			if result != tt.expected {
				t.Errorf("FindOdd() = %d; expected %d", result, tt.expected)
			}
		})
	}
}

func TestSolution(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"abcdef", []string{"ab", "cd", "ef"}},
		{"abcde", []string{"ab", "cd", "e_"}},
		{"a", []string{"a_"}},
	}

	for _, tt := range tests {
		t.Run("Solution("+tt.input+")", func(t *testing.T) {
			result := Solution(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Solution(%s) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
