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

func searchSubsequence(dna string, s string, idxDna int, idxS int, idxs []int, solutions *[][]int) {

	if idxS >= len(s) {
		*solutions = append(*solutions, idxs)
		return
	}

	if idxDna >= len(dna) {
		return
	}

	charS := string(s[idxS])

	for i := idxDna; i < len(dna); i++ {

		charDna := string(dna[idxDna])

		if charS == charDna {
			idxs = append(idxs, i)
			searchSubsequence(dna, s, i+1, idxS+1, idxs, solutions)
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
		fmt.Printf(">Searching substring %s...\n", s)
		found := searchAllSubstrings(entry, s)

		if len(found) == 0 {
			fmt.Printf(">No \"%s\" found.\n\n", s)
			continue
		}

		for _, pos := range found {
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
		solutions := [][]int{}

		searchSubsequence(subsequencesEntry, s, 0, 0, []int{}, &solutions)

		fmt.Printf(">\"%s\" subsequence found in positions:\n", s)

		for _, x := range solutions {
			fmt.Println(x)
		}

		fmt.Println()
	}
}
