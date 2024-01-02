package kata

import (
	"fmt"
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