package kata

import "testing"

func TestDNAStrand(t *testing.T) {
	var testCases = []struct {
		input    string
		expected string
	}{
		{"ATTGC", "TAACG"},
		{"AAAA", "TTTT"},
		{"GTAT", "CATA"},
	}

	for _, tc := range testCases {
		if DNAStrand(tc.input) != tc.expected {
			t.Fatalf("Expected %s, got %s", tc.expected, DNAStrand(tc.input))
		}
	}
}
