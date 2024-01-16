package kata

import (
	"fmt"
	"strings"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
}

func IsSolved(board [3][3]int) int {
	hasEmpty := false
	for i := 0; i < len(board); i++ {
		found := false
		continueSearch := true
		firstValue := board[i][0]
		for j := 0; j < len(board[i]); j++ {
			currentValue := board[i][j]
			if currentValue == 0 {
				hasEmpty = true
			}
			if currentValue == firstValue && currentValue != 0 && continueSearch {
				found = true
			} else {
				found = false
				continueSearch = false
			}
		}
		if found {
			return firstValue
		}
	}

	for i := 0; i < len(board); i++ {
		found := false
		continueSearch := true
		firstValue := board[0][i]
		for j := 0; j < len(board[i]); j++ {
			currentValue := board[j][i]
			if currentValue == 0 {
				hasEmpty = true
			}
			if currentValue == firstValue && currentValue != 0 && continueSearch {
				found = true
			} else {
				found = false
				continueSearch = false
			}
		}
		if found {
			return firstValue
		}
	}

	found := false
	continueSearch := true
	firstValue := board[0][0]
	for i := 0; i < len(board); i++ {
		currentValue := board[i][i]
		if currentValue == firstValue && currentValue != 0 && continueSearch {
			found = true
		} else {
			found = false
			continueSearch = false
		}
	}
	if found {
		return firstValue
	}

	firstValue = board[0][2]
	continueSearch = true
	for i := 0; i < len(board); i++ {
		currentValue := board[i][len(board)-i-1]
		if currentValue == firstValue && currentValue != 0 && continueSearch {
			found = true
		} else {
			found = false
			continueSearch = false
		}
	}
	if found {
		return firstValue
	}
	if hasEmpty {
		return -1
	}
	return 0
}

func HumanReadableTime(seconds int) string {
	hours := seconds / 3600
	seconds = seconds % 3600
	minutes := seconds / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

}

func ReverseWords(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverse(word)
		}
	}
	return strings.Join(words, " ")
}

func reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func MaximumSubarraySum(arr []int) int {
	maxSum, currentSum := 0, 0
	for _, num := range arr {
		currentSum = max(0, currentSum+num)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Tribonacci(signature [3]float64, n int) []float64 {

	if n == 0 {
		return []float64{}
	}

	tribonacci := make([]float64, n)
	copy(tribonacci, signature[:])

	for i := 3; i < n; i++ {
		tribonacci[i] = tribonacci[i-1] + tribonacci[i-2] + tribonacci[i-3]
	}

	return tribonacci[:n]
}

func findNb(m int) int {
	total := 0
	n := 0
	for total < m {
		n++
		total += n * n * n
	}
	if total == m {
		return n
	}
	return -1
}

func RGB(r, g, b int) string {
	r = clamp(r, 0, 255)
	g = clamp(g, 0, 255)
	b = clamp(b, 0, 255)
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func Beeramid(bonus int, price float64) int {
	totalCans := int(float64(bonus) / price) // Total number of cans that can be bought with the bonus
	level := 0
	cansUsed := 0

	for cansUsed <= totalCans {
		level++
		levelCans := level * level // Number of cans needed for the current level
		if cansUsed+levelCans > totalCans {
			return level - 1 // Return the last completed level
		}
		cansUsed += levelCans
	}

	return level
}

func ValidISBN10(isbn string) bool {
	var isbnInts []int
	for position, char := range isbn {
		if char >= '0' && char <= '9' {
			isbnInts = append(isbnInts, int(char-'0'))
		} else if (char == 'X' || char == 'x') && (position == len(isbn)-1) {
			isbnInts = append(isbnInts, 10)
		} else {
			fmt.Errorf("Invalid character %c in ISBN", char)
			return false
		}
	}
	sum := 0
	for i := 0; i < len(isbnInts); i++ {
		sum = sum + isbnInts[i]*(i+1)
	}
	return sum%11 == 0 && len(isbnInts) == 10
}

func LCS(x, y string) string {
	m := len(x)
	n := len(y)

	// Create a 2D slice to store the lengths of LCS
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	// Build the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Reconstruct the LCS
	lcsLength := dp[m][n]
	lcs := make([]byte, lcsLength)
	i, j := m, n
	for i > 0 && j > 0 {
		if x[i-1] == y[j-1] {
			lcs[lcsLength-1] = x[i-1]
			i--
			j--
			lcsLength--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(lcs)
}

func Determinant(m [][]int) int {
	if len(m) == 1 {
		return m[0][0]
	}
	total := 0
	s := 1

	// Store indexes in a list for row referencing
	index := make([]int, len(m))
	for i := range index {
		index[i] = i
	}

	// when at 2x2 matrices
	if len(m) == 2 && len(m[0]) == 2 {
		// the determinant of a 2x2 matrix.
		return m[0][0]*m[1][1] - m[1][0]*m[0][1]
	} else {
		// define submatrix for focus column
		mSub := make([][]int, len(m)-1)
		for i := range mSub {
			mSub[i] = make([]int, len(m)-1)
		}

		// for each focus column
		for c := 0; c < len(m); c++ {
			// populate submatrix with all rows excluding the first
			for r := 1; r < len(m); r++ {
				// for each new subcolumn excluding the focus
				i := 0
				for cSub := 0; cSub < len(m); cSub++ {
					if cSub == c {
						continue
					}
					mSub[r-1][i] = m[r][cSub]
					i++
				}
			}
			total += s * m[0][c] * Determinant(mSub)
			// switch sign
			s = -s
		}
	}

	return total
}
