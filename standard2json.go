package main

import (
	"bufio"
	"encoding/json"
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
	Extension   string `json:"extension,omitempty"`
}

// NAGList is a slice of NAGs
type NAGList struct {
	NAGs []NAG `json:"nags"`
}

func (n NAG) String() string {
	return fmt.Sprintf("%d\t%s", n.Value, n.Description)
}

func getStandardNAGList() NAGList {
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
	nags := NAGList{}

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
			nag := NAG{value, description, ""}
			nags.NAGs = append(nags.NAGs, nag)
		}

		lineCount++
	}

	return nags
}

func getChessPadExtensions() []NAG {
	var nags []NAG

	nags = append(nags, NAG{140, "With the idea...", "ChessPad"})
	nags = append(nags, NAG{141, "Aimed against...", "ChessPad"})
	nags = append(nags, NAG{142, "Better is...", "ChessPad"})
	nags = append(nags, NAG{143, "Worse is...", "ChessPad"})
	nags = append(nags, NAG{144, "Equivalent is...", "ChessPad"})
	nags = append(nags, NAG{145, "Editorial comment", "ChessPad"})
	nags = append(nags, NAG{146, "Novelty", "ChessPad"})
	nags = append(nags, NAG{220, "Diagram", "ChessPad"})
	nags = append(nags, NAG{221, "Diagram (from Black)", "ChessPad"})
	nags = append(nags, NAG{238, "Space advantage", "ChessPad"})
	nags = append(nags, NAG{239, "File", "ChessPad"})
	nags = append(nags, NAG{240, "Diagonal", "ChessPad"})
	nags = append(nags, NAG{241, "Centre", "ChessPad"})
	nags = append(nags, NAG{242, "Kingside", "ChessPad"})
	nags = append(nags, NAG{243, "Queenside", "ChessPad"})
	nags = append(nags, NAG{244, "Weak point", "ChessPad"})
	nags = append(nags, NAG{245, "Ending", "ChessPad"})
	nags = append(nags, NAG{246, "Bishop pair", "ChessPad"})
	nags = append(nags, NAG{247, "Opposite bishops", "ChessPad"})
	nags = append(nags, NAG{248, "Same bishops", "ChessPad"})
	nags = append(nags, NAG{249, "Connected pawns", "ChessPad"})
	nags = append(nags, NAG{250, "Isolated pawns", "ChessPad"})
	nags = append(nags, NAG{251, "Doubled pawns", "ChessPad"})
	nags = append(nags, NAG{252, "Passed pawn", "ChessPad"})
	nags = append(nags, NAG{253, "Pawn majority", "ChessPad"})
	nags = append(nags, NAG{254, "With", "ChessPad"})
	nags = append(nags, NAG{255, "Without", "ChessPad"})

	return nags
}

func main() {
	nags := getStandardNAGList()
	nags.NAGs = append(nags.NAGs, getChessPadExtensions()...)
	js, err := json.Marshal(nags)
	if err != nil {
		log.Fatalf("Unable to marshal %q: %q", nags, err)
	}

	fmt.Println(string(js))
}
