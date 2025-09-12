package happynumbers

import "testing"

func TestIsHappy(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{19, true},
		{2, false},
		{7, true},
	}

	for _, tc := range testCases {
		actual := IsHappy(tc.input)
		if actual != tc.expected {
			t.Fatalf("Expected %t, got %t", tc.expected, actual)
		}
	}
}
