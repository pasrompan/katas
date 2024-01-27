package kata

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
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

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	years := seconds / 31536000
	seconds %= 31536000
	days := seconds / 86400
	seconds %= 86400
	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	seconds %= 60

	parts := []string{}
	if years > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", years, pluralize("year", years)))
	}
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", days, pluralize("day", days)))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", hours, pluralize("hour", hours)))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", minutes, pluralize("minute", minutes)))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%d %s", seconds, pluralize("second", seconds)))
	}

	return humanize(parts)
}

func pluralize(word string, count int64) string {
	if count > 1 {
		return word + "s"
	}
	return word
}

var fibMatrix = [2][2]*big.Int{
	{big.NewInt(1), big.NewInt(1)},
	{big.NewInt(1), big.NewInt(0)},
}

func multiply(a, b [2][2]*big.Int) [2][2]*big.Int {
	x := new(big.Int).Mul(a[0][0], b[0][0])
	y := new(big.Int).Mul(a[0][1], b[1][0])
	x.Add(x, y)

	y = new(big.Int).Mul(a[0][0], b[0][1])
	z := new(big.Int).Mul(a[0][1], b[1][1])
	y.Add(y, z)

	z = new(big.Int).Mul(a[1][0], b[0][0])
	w := new(big.Int).Mul(a[1][1], b[1][0])
	z.Add(z, w)

	w = new(big.Int).Mul(a[1][0], b[0][1])
	u := new(big.Int).Mul(a[1][1], b[1][1])
	w.Add(w, u)

	return [2][2]*big.Int{
		{x, y},
		{z, w},
	}
}

func power(mat [2][2]*big.Int, n int64) [2][2]*big.Int {
	if n == 0 || n == 1 {
		return mat
	}

	mat = power(mat, n/2)
	mat = multiply(mat, mat)

	if n%2 != 0 {
		mat = multiply(mat, fibMatrix)
	}

	return mat
}

func fib(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}

	if n < 0 {
		n = -n
		if n%2 == 0 {
			return new(big.Int).Neg(power(fibMatrix, n)[1][0])
		}
	}

	return power(fibMatrix, n)[1][0]
}
func humanize(parts []string) string {
	if len(parts) == 1 {
		return parts[0]
	}
	return strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
}

func PermutationalPrimes(nMax, kPerms int) [3]int {
	primes := sieve(nMax)
	primePerms := make(map[int][]int)
	primeChecked := make(map[int]bool)

	for _, prime := range primes {
		if primeChecked[prime] {
			continue
		}
		perms := primePermutations(prime, primes, kPerms)
		for _, perm := range perms {
			primeChecked[perm] = true
		}
		if len(perms) == kPerms {
			primePerms[prime] = perms
		}
	}

	if len(primePerms) == 0 {
		return [3]int{0, 0, 0}
	}

	keys := make([]int, 0, len(primePerms))
	for k := range primePerms {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return [3]int{len(keys), keys[0], keys[len(keys)-1]}
}

func sieve(n int) []int {
	sieve := make([]bool, n+1)
	for x := 1; x*x <= n; x++ {
		for y := 1; y*y <= n; y++ {
			z := 4*x*x + y*y
			if z <= n && (z%12 == 1 || z%12 == 5) {
				sieve[z] = !sieve[z]
			}
			z = 3*x*x + y*y
			if z <= n && z%12 == 7 {
				sieve[z] = !sieve[z]
			}
			if x > y {
				z = 3*x*x - y*y
				if z <= n && z%12 == 11 {
					sieve[z] = !sieve[z]
				}
			}
		}
	}
	for r := 5; r*r <= n; r++ {
		if sieve[r] {
			for i := r * r; i <= n; i += r * r {
				sieve[i] = false
			}
		}
	}
	primes := []int{2, 3}
	for x := 5; x <= n; x += 2 {
		if sieve[x] {
			primes = append(primes, x)
		}
	}
	return primes
}

func primePermutations(n int, primes []int, kPerms int) []int {
	perms := make([]int, 0)
	str := strconv.Itoa(n)
	for _, prime := range primes {
		if prime != n && len(str) == len(strconv.Itoa(prime)) && isPermutation(str, strconv.Itoa(prime)) {
			perms = append(perms, prime)
		}
		if len(perms) > kPerms {
			return perms
		}
	}
	return perms
}

func isPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	counts := [10]int{}
	for _, r := range a {
		counts[r-'0']++
	}

	for _, r := range b {
		counts[r-'0']--
		if counts[r-'0'] < 0 {
			return false
		}
	}

	return true
}

func ListPosition(word string) *big.Int {
	// Count occurrences of each letter.
	counts := make(map[rune]int64)
	for _, letter := range word {
		counts[letter]++
	}

	rank := big.NewInt(1)
	length := int64(len(word))

	// Calculate the rank.
	for i := int64(0); i < length; i++ {
		smaller := big.NewInt(0)
		currentLetter := rune(word[i])

		// Calculate the number of words that would start with a letter smaller than the current one.
		for letter, _ := range counts {
			if letter < currentLetter {
				temp := factorialBig(length - i - 1)
				for k, v := range counts {
					if k != letter {
						temp.Div(temp, factorialBig(v))
					} else {
						temp.Div(temp, factorialBig(v-1))
					}
				}
				smaller.Add(smaller, temp)
			}
		}

		// Add the number of words starting with smaller letters to the rank.
		rank.Add(rank, smaller)
		counts[currentLetter]--

		if counts[currentLetter] == 0 {
			delete(counts, currentLetter)
		}
	}

	return rank
}

// factorial computes the factorial of n.
func factorialBig(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func UniquePermutationsCount(s string) int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	totalChars := len(s)
	result := factorial(totalChars)

	for _, count := range charCount {
		result /= factorial(count)
	}

	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func toJadenCase(str string) string {
	return strings.Title(str)
}
