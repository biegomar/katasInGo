package kata

import "strings"

func ReverseWords(str string) string {
	splitted := strings.Split(str, " ")

	for i := 0; i < len(splitted); i++ {
		splitted[i] = reverseWords(splitted[i])
	}

	return strings.Join(splitted, " ")
}

func reverseWords(str string) string {
	var result = ""

	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}

	return result
}
