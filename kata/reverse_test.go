package kata

import "testing"

func TestReverseWords(t *testing.T) {
	var TestCases = []struct {
		input    string
		expected string
	}{
		{"The quick brown fox jumps over the lazy dog.", "ehT kciuq nworb xof spmuj revo eht yzal .god"},
		{"apple", "elppa"},
		{"a b c d", "a b c d"},
		{"double  spaced  words", "elbuod  decaps  sdrow"},
		{"stressed desserts", "desserts stressed"},
		{"hello hello", "olleh olleh"},
	}

	for _, tc := range TestCases {
		actual := ReverseWords(tc.input)
		if actual != tc.expected {
			t.Fatalf("Expected %s, got %s", tc.expected, actual)
		}
	}
}
