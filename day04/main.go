package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	text, err := parse.TextFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	rawPassportLines := strings.Split(text, "\n\n")
	var passportLines []string
	for _, rpLine := range rawPassportLines {
		passportLines = append(passportLines, strings.ReplaceAll(rpLine, "\n", " "))
	}

	var validCount int
	for _, pLine := range passportLines {
		if isLineValid(pLine) {
			validCount++
		}
	}
	fmt.Println(validCount)
}

func isLineValid(line string) bool {
	parts := strings.Split(line, " ")
	switch len(parts) {
	case 8:
		return true
	case 7:
		return !strings.Contains(line, "cid:")
	}
	return false
}
