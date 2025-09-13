package wordcountreader

import "testing"

func TestFileContentInfo_ReadFileContent(t *testing.T) {
	testcases := []struct {
		filename string
		expected FileContentInfo
	}{
		{"example.txt", FileContentInfo{Words: 18, Characters: 86, CharactersExcludingSpaces: 70}},
	}

	for _, tc := range testcases {
		var fileContentInfo FileContentInfo
		fileContentInfo.ReadFileContent(tc.filename)
		if fileContentInfo != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, fileContentInfo)
		}
	}
}
