package kata

import "strings"

func Points(games []string) int {
	points := 0
	for _, game := range games {
		gameResult := strings.Split(game, ":")
		if gameResult[0] == gameResult[1] {
			points += 1
		}
		if gameResult[0] > gameResult[1] {
			points += 3
		}
	}

	return points
}
