package wordcountreader

import (
	"bufio"
	"os"
	"strings"
)

type FileContentInfo struct {
	Words                     int
	Characters                int
	CharactersExcludingSpaces int
}

func (fileContentInfo *FileContentInfo) ReadFileContent(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileContentInfo.Words += len(strings.Fields(fileScanner.Text()))
		fileContentInfo.Characters += len(fileScanner.Text())
		fileContentInfo.CharactersExcludingSpaces += len(strings.Replace(fileScanner.Text(), " ", "", -1))
	}
}
