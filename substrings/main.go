package main

import (
	"fmt"
)

func searchAllSubstrings(dna string, s string) []int {

	positions := []int{}

	for i := 0; i < len(dna); i++ {

		if string(dna[i]) == string(s[0]) {
			flag := true
			for j := 0; j < len(s); j++ {
				if dna[(j+i)] != s[j] {
					flag = false
					break
				}
			}

			if flag {
				positions = append(positions, i)
			}
		}
	}
	return positions
}

func searchSubsequence(dna string, needle string, dnaIdx int, needleIdx int, idxs []int, solutions *[][]int) {

	if needleIdx == len(needle) {
		*solutions = append(*solutions, idxs)
		return
	}

	if dnaIdx == len(dna) {
		return
	}

	charNeedle := string(needle[needleIdx])

	for i := dnaIdx; i < len(dna); i++ {

		charDna := string(dna[i])

		if charNeedle == charDna {
			idxs[needleIdx] = i
			searchSubsequence(dna, needle, i+1, needleIdx+1, idxs, solutions)
		}
	}
}

func main() {
	entry := readEntryFromFile("entry.txt")
	subsequencesEntry := readEntryFromFile("entry2.txt")

	substrings := []string{"TTGACA", "TATATT", "CGCGCGCGCG"}
	subsequences := []string{"TTGACATATATT", "CGCGCGCGCG", "ATCCAGAATTCTCGGA"}

	/* Substrings Search */
	for _, s := range substrings {
		fmt.Printf(">Procurando substring %s...\n", s)
		found := searchAllSubstrings(entry, s)

		if len(found) == 0 {
			fmt.Printf(">Nenhuma instância de \"%s\" encontrada.\n\n", s)
			continue
		}

		for _, pos := range found {
			fmt.Printf(">\"%s\" encontrada na posição [%d].\n",
				entry[pos:(pos+len(s))],
				pos)
		}
		fmt.Println()
	}

	fmt.Print("#################\n\n")

	/* Subsequences Search */
	for _, s := range subsequences {
		fmt.Printf(">Procurando subsequência %s...\n", s)
		solutions := [][]int{}

		searchSubsequence(subsequencesEntry, s, 0, 0, make([]int, len(s)), &solutions)
		fmt.Printf(">A subsequência \"%s\" foi encontrada em %d posições distintas.\n", s, len(solutions))
		fmt.Println(solutions[0])
		fmt.Println(solutions[10])
		fmt.Printf(">A listagem de todas as posições pode ser encontrada no arquivo output_%s.txt.\n", s)
		writeOutputOnFile(fmt.Sprintf("output_%s.txt", s), solutions)

		fmt.Println()
	}
}
