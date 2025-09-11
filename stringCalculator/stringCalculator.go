package stringCalculator

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if numbers == "" {
		return 0
	}

	values := strings.Split(numbers, ",")
	result := 0
	for i := 0; i < len(values); i++ {
		n, err := strconv.Atoi(values[i])
		if err == nil {
			result += n
		}
	}

	return result
}
