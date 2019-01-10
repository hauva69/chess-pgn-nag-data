package main

import (
	"bufio"
	"fmt"
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

	lineCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatalf("Unable to read %q: %s", line, err)
		}

		if lineCount > lastNAGLine {
			break
		} else if lineCount >= firstNAGLine {
			fmt.Println(line)
		}

		lineCount++
	}
}
