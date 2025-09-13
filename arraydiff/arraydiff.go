package arraydiff

func ArrayDiff(a, b []int) []int {
	var result []int

	for _, valueOfA := range a {
		found := false
		for _, valueOfB := range b {
			if valueOfA == valueOfB {
				found = true
			}
		}
		if !found {
			result = append(result, valueOfA)
		}
	}
	return result
}
