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

func TestTwiceAsOld(t *testing.T) {
	testCases := []struct {
		dadYearsOld int
		sonYearsOld int
		expected    int
	}{
		{36, 7, 22},
		{55, 30, 5},
		{42, 21, 0},
		{22, 1, 20},
		{29, 0, 29},
	}

	// Act & Assert
	for _, tc := range testCases {
		result := TwiceAsOld(tc.dadYearsOld, tc.sonYearsOld)
		if result != tc.expected {
			t.Errorf("[Alter Vater: %d, Alter Sohn %d] Erwartet wurde der zeitliche Abstand von %d Jahren und nicht %d Jahre.", tc.dadYearsOld, tc.sonYearsOld, tc.expected, result)
		}
	}
}

func TestTwiceAsOldClever(t *testing.T) {
	testCases := []struct {
		dadYearsOld int
		sonYearsOld int
		expected    int
	}{
		{36, 7, 22},
		{55, 30, 5},
		{42, 21, 0},
		{22, 1, 20},
		{29, 0, 29},
	}

	// Act & Assert
	for _, tc := range testCases {
		result := TwiceAsOldClever(tc.dadYearsOld, tc.sonYearsOld)
		if result != tc.expected {
			t.Errorf("[Alter Vater: %d, Alter Sohn %d] Erwartet wurde der zeitliche Abstand von %d Jahren und nicht %d Jahre.", tc.dadYearsOld, tc.sonYearsOld, tc.expected, result)
		}
	}
}
