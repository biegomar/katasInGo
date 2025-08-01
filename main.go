package main

import (
	"fmt"
	"katasInGo/kata"
	"strconv"
)

func main() {
	fmt.Println(countSheep(4))

	fmt.Println(kata.QuarterOf(1))
	fmt.Println(kata.QuarterOf(4))
	fmt.Println(kata.QuarterOf(9))
	fmt.Println(kata.QuarterOf(11))
}

func countSheep(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		result += strconv.Itoa(i) + " sheep..."
	}

	return result
}
