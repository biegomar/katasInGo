package kata

import "math"

func QuarterOf(month int) int {
	quarterBase := month / 3
	quarterRest := month % 3

	if quarterRest == 0 {
		return quarterBase
	}

	return quarterBase + 1
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	twiceAsOldYears := 0

	if sonYearsOld == 0 {
		return dadYearsOld
	}

	if dadYearsOld == sonYearsOld*2 {
		return 0
	}

	if dadYearsOld > sonYearsOld*2 {
		for sonYearsOld > 0 && dadYearsOld != sonYearsOld*2 {
			twiceAsOldYears++
			dadYearsOld--
		}
		return twiceAsOldYears
	}

	if dadYearsOld < sonYearsOld*2 {
		for dadYearsOld != sonYearsOld*2 {
			twiceAsOldYears++
			dadYearsOld++
		}
		return twiceAsOldYears
	}

	return twiceAsOldYears
}

func TwiceAsOldClever(dadYearsOld, sonYearsOld int) int {
	return int(math.Abs(float64(dadYearsOld - 2*sonYearsOld)))
}
