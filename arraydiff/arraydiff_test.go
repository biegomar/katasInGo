package arraydiff

import (
	"slices"
	"testing"
)

func TestArrayDiff(t *testing.T) {
	testCases := []struct {
		inputA   []int
		inputB   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2}, []int{1, 3}},
		{[]int{1, 2, 3}, []int{1, 2}, []int{3}},
		{[]int{1, 2}, []int{2}, []int{1}},
	}

	for _, tc := range testCases {
		actual := ArrayDiff(tc.inputA, tc.inputB)
		if !slices.Equal(tc.expected, actual) {
			t.Fatalf("[%d, %d] Expected %d, got %d", tc.inputA, tc.inputB, tc.expected, actual)
		}

	}
}
