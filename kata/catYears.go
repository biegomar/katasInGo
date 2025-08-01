package kata

func CalculateYears(years int) [3]int {
	catYears := 0
	dogYears := 0

	for i := 0; i < years; i++ {
		switch {
		case i == 0:
			{
				catYears = 15
				dogYears = 15
			}
		case i == 1:
			{
				catYears += 9
				dogYears += 9
			}
		default:
			{
				catYears += 4
				dogYears += 5
			}
		}
	}
	return [3]int{years, catYears, dogYears}
}
