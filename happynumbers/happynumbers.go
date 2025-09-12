package happynumbers

import (
	"strconv"
)

func IsHappy(n int) bool {
	seen := make(map[int]bool)
	iterationSum := n
	for !seen[iterationSum] && iterationSum != 1 {
		seen[iterationSum] = true
		iterationSum = oneCycle(iterationSum)
	}
	return iterationSum == 1
}

func oneCycle(n int) int {
	squareSum := 0
	digits := intToDigits(n)
	for _, digit := range digits {
		squareSum += digit * digit
	}

	return squareSum
}

func intToDigits(n int) []int {
	var result []int
	numberAsString := strconv.Itoa(n)
	for _, digit := range numberAsString {
		result = append(result, int(digit-'0'))
	}

	return result
}
