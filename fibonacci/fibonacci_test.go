package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		sequenceNumber int
		expected       int
	}{
		{12, 144},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
	}

	for _, tc := range testCases {
		actual := Fibonacci(tc.sequenceNumber)
		if actual != tc.expected {
			t.Fatalf("[%d] Expected %d, got %d", tc.sequenceNumber, tc.expected, actual)
		}
	}
}
