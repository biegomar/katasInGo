package kata

import "strings"

var dnaComplement = map[rune]string{
	'A': "T",
	'T': "A",
	'C': "G",
	'G': "C",
}

func DNAStrand(dna string) string {
	var result = strings.Builder{}
	result.Grow(len(dna))

	for _, c := range dna {
		result.WriteString(dnaComplement[c])
	}

	return result.String()
}
