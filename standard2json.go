package main

import (
	"log"
	"os"
)

func main() {
	const firstNAGLine = 1398
	const lastNAGLine = 1537

	filename := "doc/pgnstd.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %q: %s", filename, err)
	}

	defer file.Close()

}
