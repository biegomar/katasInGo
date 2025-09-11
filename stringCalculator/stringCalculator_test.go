package stringCalculator

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		stringItems string
		expected    int
	}{
		{"", 0},
		{"1,2", 3},
		{"4", 4},
		{"1,2,3", 6},
		{"127,23", 150},
	}

	for _, tc := range testCases {
		actual := Add(tc.stringItems)
		if actual != tc.expected {
			t.Fatalf("[%s] Expected %d, got %d", tc.stringItems, tc.expected, actual)
		}
	}
}
