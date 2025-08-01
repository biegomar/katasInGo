package kata

func QuarterOf(month int) int {
	quarterBase := month / 3
	quarterRest := month % 3

	if quarterRest == 0 {
		return quarterBase
	}

	return quarterBase + 1
}
