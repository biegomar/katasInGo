package fizzbuzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int) string {
	var result = strings.Builder{}
	for i := 1; i <= n; i++ {
		if i > 1 {
			result.WriteString("\n")
		}
		result.WriteString(numberToFizzBuzz(i))
	}

	return result.String()
}

func numberToFizzBuzz(number int) string {
	var result = ""

	if number%3 == 0 {
		result = "Fizz"
	}
	if number%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
