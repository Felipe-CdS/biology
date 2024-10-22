package main

import (
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
