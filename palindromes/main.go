package main

import (
	"fmt"
	"log"
	"os"
)

type palindromes struct {
	pos    int
	length int
}

func readEntryFromFile(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error reading file.")
	}

	return string(data)
}

func printPalindrome(dna string, p palindromes) {
	fmt.Printf("[%03d] %s\n", p.pos, dna[p.pos:(p.pos+p.length)])
}

func checkPalindrome(dna string, pos int, s int) (*palindromes, error) {

	if pos < 0 || pos+s >= len(dna) {
		if s > 3 {
			if pos == -1 {
				return &palindromes{0, s - 2 + 1}, nil
			}
			return &palindromes{len(dna) - s, s}, nil
		}
		return &palindromes{}, fmt.Errorf("oob")
	}

	a := string(dna[pos])
	b := string(dna[pos+s])

	if (a == "C" && b == "G" || a == "G" && b == "C") ||
		(a == "T" && b == "A" || a == "A" && b == "T") {

		p, err := checkPalindrome(dna, pos-1, s+2)
		if err != nil {
			if s > 3 {
				return &palindromes{pos, s + 1}, nil
			}
			return &palindromes{}, fmt.Errorf("too small")
		}
		return p, nil
	}

	if s > 3 {
		return &palindromes{pos + 1, s - 1}, nil
	}
	return &palindromes{}, fmt.Errorf("not pair")
}

func main() {
	entry := readEntryFromFile("entry2.txt")

	var p []*palindromes

	for i, c := range entry {

		//Remove break lines to run code
		if string(c) == "\n" {
			break
		}

		n, err := checkPalindrome(entry, i, 1)

		if err == nil {
			p = append(p, n)
		}
	}

	for _, e := range p {
		printPalindrome(entry, *e)
	}
}
