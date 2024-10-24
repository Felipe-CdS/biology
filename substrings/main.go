package main

import (
	"bufio"
	"fmt"
)

var MAX_SUBSEQ_QUANT int = 1000
var DEBUG_OUTPUT bool = false

func main() {

	flagsSetup()

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
			fmt.Printf(">\"%s\" encontrada na posição [%d].\n\n",
				entry[pos:(pos+len(s))],
				pos)
		}
	}

	fmt.Print("#################\n\n")

	/* Subsequences Search */
	for _, s := range subsequences {
		fmt.Printf(">Procurando subsequência %s...\n", s)
		fmt.Printf(">Esse passo pode demorar dependendo da quantidade de subsequências escolhida...\n")

		filename := fmt.Sprintf("output_%s.txt", s)
		f, w := openOutputfile(filename)

		quant := 0

		searchSubsequence(
			subsequencesEntry,
			s,
			0,
			0,
			&quant,
			make([]int, len(s)),
			w)

		fmt.Printf(">A subsequência \"%s\" foi encontrada em %d posições distintas.\n",
			s,
			quant,
		)

		fmt.Printf(">A listagem das posições pode ser encontrada no arquivo output_%s.txt.\n", s)
		fmt.Printf(">Dependendo da quantidade de subsequencias encontradas esse arquivo pode ser extremamente grande. Cuidado ao abri-lo.\n\n")

		w.Flush()
		f.Close()
	}
}

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

func searchSubsequence(dna string,
	needle string,
	dnaIdx int,
	needleIdx int,
	quant *int,
	idxs []int,
	w *bufio.Writer) {

	if *quant >= MAX_SUBSEQ_QUANT && MAX_SUBSEQ_QUANT != -1 {
		return
	}

	if needleIdx == len(needle) {
		*quant++
		writeOutputOnFile(w, idxs, dna)
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
			searchSubsequence(
				dna,
				needle,
				i+1,
				needleIdx+1,
				quant,
				idxs,
				w)
		}
	}
}
