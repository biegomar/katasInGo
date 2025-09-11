package diamond

import "strings"

func RenderDiamond(item string) string {
	if item == "A" {
		return "A"
	}

	result := strings.Builder{}
	baseRune := rune(item[0])
	distance := baseRune - 'A'
	gapDistance := 1

	for nextItem := 'A'; nextItem <= baseRune; nextItem++ {
		startDistance := distance - (nextItem - 'A')
		result.WriteString(strings.Repeat(" ", int(startDistance)))
		result.WriteRune(nextItem)
		if nextItem != 'A' {
			result.WriteString(strings.Repeat(" ", gapDistance))
			gapDistance += 2
			result.WriteRune(nextItem)
		}
		result.WriteString("\n")
	}

	gapDistance -= 4
	for nextItem := baseRune - 1; nextItem >= 'A'; nextItem-- {
		startDistance := distance - (nextItem - 'A')
		result.WriteString(strings.Repeat(" ", int(startDistance)))
		result.WriteRune(nextItem)
		if nextItem != 'A' {
			result.WriteString(strings.Repeat(" ", gapDistance))
			gapDistance -= 2
			result.WriteRune(nextItem)
		}
		result.WriteString("\n")
	}

	return result.String()
}
