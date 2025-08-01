package kata

import "testing"

func TestQuarterOf(t *testing.T) {
	// Arrange
	testCases := []struct {
		input    int
		expected int
	}{
		{1, 1},  // Januar   -> Q1
		{2, 1},  // Februar  -> Q1
		{3, 1},  // März     -> Q1
		{4, 2},  // April    -> Q2
		{5, 2},  // Mai      -> Q2
		{6, 2},  // Juni     -> Q2
		{7, 3},  // Juli     -> Q3
		{8, 3},  // August   -> Q3
		{9, 3},  // September-> Q3
		{10, 4}, // Oktober  -> Q4
		{11, 4}, // November -> Q4
		{12, 4}, // Dezember -> Q4
	}

	// Act & Assert
	for _, tc := range testCases {
		result := QuarterOf(tc.input)
		if result != tc.expected {
			t.Errorf("Für Monat %d: erwartete %d, bekam %d",
				tc.input, tc.expected, result)
		}
	}
}
