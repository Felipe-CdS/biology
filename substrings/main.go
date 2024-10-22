package main

import (
	"fmt"
	"log"
	"os"
)

func readEntryFromFile(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error reading file.")
	}

	return string(data)
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

func searchAllSubsequences(dna string, s string) []int {

	positions := []int{}

	searchNextInSubseq(dna, s, positions)
	fmt.Println(positions)
	return positions
}

func searchNextInSubseq(dna string, s string, p []int) {
	for i := 0; i < len(dna); i++ {
		if string(dna[i]) == string(s[0]) {
			p = append(p, i)
			searchNextInSubseq(dna[i:], s[1:], p)
		}
	}
}

func main() {
	entry := readEntryFromFile("entry.txt")
	substrings := []string{"TTGACA", "TATATT", "CGCGCGCGCG"}
	var substrPos []int

	subsequences := []string{"TTGACATATATT", "CGCGCGCGCG", "ATCCAGAATTCTCGGA"}
	var subseqPos [][]int

	/* Substrings Search */
	for _, s := range substrings {
		fmt.Printf(">Searching substring %s...\n", s)
		substrPos = searchAllSubstrings(entry, s)

		if len(substrPos) == 0 {
			fmt.Printf(">No \"%s\" found.\n\n", s)
			continue
		}

		for _, pos := range substrPos {
			fmt.Printf(">\"%s\" in position [%d].\n",
				entry[pos:(pos+len(s))],
				pos)
		}
		fmt.Println()
	}

	fmt.Print("#################\n\n")

	/* Subsequences Search */
	for _, s := range subsequences {
		fmt.Printf(">Searching subsequence %s...\n", s)

		subseqPos = append(subseqPos, searchAllSubsequences(entry, s))

		for _, s := range subseqPos {
			for _, p := range s {
				fmt.Println(p)
			}
		}
	}
}
