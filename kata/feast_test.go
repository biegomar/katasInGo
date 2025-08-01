package kata

import "testing"

func TestFeast(t *testing.T) {
	testCases := []struct {
		beast    string
		dish     string
		expected bool
	}{
		{"great blue heron", "garlic naan", true},
		{"chickadee", "chocolate cake", true},
		{"brown bear", "bear claw", false},
	}

	for _, tc := range testCases {
		actual := Feast(tc.beast, tc.dish)
		if actual != tc.expected {
			t.Fatalf("[%s, %s] Expected %t, got %t", tc.beast, tc.dish, tc.expected, actual)
		}
	}
}
