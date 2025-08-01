package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countSheep(3))
}

func countSheep(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		result += strconv.Itoa(i) + " sheep..."
	}

	return result
}
