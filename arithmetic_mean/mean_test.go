package arithmetic_mean

import "testing"

func TestMean(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{1, 1, 1}, 1.0},
		{[]int{1, 2, 3}, 2.0},
		{[]int{-3, 0, 3, 6}, 1.5},
		{[]int{2, 4, 6, 8}, 5.0},
		{[]int{1, 2, 3, 4, 6}, 3.2},
	}

	for _, tc := range testCases {
		actual := Mean(tc.numbers)
		if actual != tc.expected {
			t.Fatalf("Expected %.2f, got %.2f", tc.expected, actual)
		}
	}
}
