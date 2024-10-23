package main

import (
	"bufio"
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

func writeOutputOnFile(filename string, solutions [][]int) {
	f, err := os.Create(filename)

	if err != nil {
		log.Fatalln("Error creating output file.")
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	for i, s := range solutions {
		if i > 100 {
			break
		}
		fmt.Fprintln(w, fmt.Sprint(s))
	}
	w.Flush()
}
