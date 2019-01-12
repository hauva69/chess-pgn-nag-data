package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// NAG (Numeric Annotation Glyphs or NAGs are used to annotate chess games),
// providing an assessment of a chess move or a chess position.
type NAG struct {
	Value       int    `json:"value"`
	Description string `json:"description"`
}

func (n NAG) String() string {
	return fmt.Sprintf("%d\t%s", n.Value, n.Description)
}

func main() {
	const firstNAGLine = 1398
	const lastNAGLine = 1537

	filename := "doc/pgnstd.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %q: %q", filename, err)
	}

	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatalf("Unable to read %q: %s", line, err)
		}

		if lineCount >= lastNAGLine {
			break
		} else if lineCount >= firstNAGLine {
			fields := strings.Fields(line)
			value, err := strconv.Atoi(fields[0])
			if err != nil {
				log.Fatalf("Unable to parse NAG: %s", err)
			}
			description := strings.Join(fields[1:], " ")
			nag := NAG{value, description}
			fmt.Println(nag)

		}

		lineCount++
	}
}
