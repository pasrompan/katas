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
