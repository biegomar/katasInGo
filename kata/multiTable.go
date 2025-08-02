package kata

import (
	"strconv"
)

func MultiTable(number int) string {
	var result = ""

	for i := 1; i <= 10; i++ {
		result = result + strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(i*number)
		if i < 10 {
			result = result + "\n"
		}
	}

	return result
}
