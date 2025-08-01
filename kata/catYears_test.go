package kata

import "testing"

func TestCatYears(t *testing.T) {
	testCases := []struct {
		input    int
		expected [3]int
	}{
		{1, [3]int{1, 15, 15}},
		{2, [3]int{2, 24, 24}},
		{10, [3]int{10, 56, 64}},
	}

	for _, tc := range testCases {
		actual := CalculateYears(tc.input)
		if actual != tc.expected {
			t.Fatalf("Expected %d, got %d", tc.expected, actual)
		}
	}
}
