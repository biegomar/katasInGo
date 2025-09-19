package ipvalidation

import "testing"

func TestIpValidation(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"127.0.0.1.1", false},
		{"1.2.3.4", true},
		{"1.2.3", false},
		{"123.456.78.90", false},
		{"123.045.067.089", false},
		{"123.45.67.89", true},
	}

	for _, tc := range testCases {
		actual := Is_valid_ip(tc.input)
		if actual != tc.expected {
			t.Fatalf("[%s] Expected %t, got %t", tc.input, tc.expected, actual)
		}
	}
}
