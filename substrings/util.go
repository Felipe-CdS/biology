package main

import (
	"bufio"
	"flag"
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

func openOutputfile(filename string) (*os.File, *bufio.Writer) {

	f, err := os.Create(filename)

	if err != nil {
		log.Fatalln("Error creating output file.")
	}

	return f, bufio.NewWriter(f)
}

func writeOutputOnFile(w *bufio.Writer, s []int, dna string) {
	fmt.Fprint(w, fmt.Sprint(s, " "))

	if DEBUG_OUTPUT {
		for _, p := range s {
			fmt.Fprint(w, fmt.Sprint(string(dna[p])))
		}
	}

	fmt.Fprintln(w)
}

func flagsSetup() {
	flag.IntVar(&MAX_SUBSEQ_QUANT, "subseq", 1000, "Max number of subsequences found.")
	flag.BoolVar(&DEBUG_OUTPUT, "debug", false, "debug")
	ALL_SUBSEQ_ALLOW := flag.Bool("all", false, "Allow to search all subsequences.")

	flag.Parse()

	if *ALL_SUBSEQ_ALLOW {
		MAX_SUBSEQ_QUANT = -1
	}
}
