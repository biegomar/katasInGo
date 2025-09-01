package arithmetic_mean

func Mean(numbers []int) float64 {
	mean := float64(sumForMean(numbers)) / float64(len(numbers))

	return mean
}

func sumForMean(numbers []int) int {
	var sum = 0
	for num := 0; num < len(numbers); num++ {
		sum += numbers[num]
	}

	return sum
}
