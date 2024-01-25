package kata

func DigitalRoot(n int) int {
	digitalRoot := parseDigits(n)
	for digitalRoot > 9 {
		digitalRoot = parseDigits(digitalRoot)
	}
	return digitalRoot
}

func parseDigits(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += digit
	}
	return sum
}
