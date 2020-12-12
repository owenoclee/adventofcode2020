// Example usage: cat input.txt | go run . -p 1
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

type entry struct {
	min      int
	max      int
	letter   byte
	password string
}

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()

	validator := partOneValidator
	if *part == 2 {
		validator = partTwoValidator
	}

	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	entries := parseLinesToEntries(lines)
	var validCount int
	for _, e := range entries {
		if validator(e) {
			validCount++
		}
	}
	fmt.Println(validCount)
}

func partOneValidator(e entry) bool {
	occurrences := strings.Count(e.password, string(e.letter))
	return occurrences >= e.min && occurrences <= e.max
}

func partTwoValidator(e entry) bool {
	return (e.password[e.min-1] == e.letter) != (e.password[e.max-1] == e.letter)
}

func parseLinesToEntries(lines []string) []entry {
	var entries []entry
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			panic(lineMalformedError(i + 1))
		}
		bounds := strings.Split(parts[0], "-")
		if len(bounds) != 2 {
			panic(lineMalformedError(i + 1))
		}
		min, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(lineMalformedError(i + 1))
		}
		max, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(lineMalformedError(i + 1))
		}
		letter := parts[1][0]
		password := parts[2]

		entries = append(entries, entry{
			min:      min,
			max:      max,
			letter:   letter,
			password: password,
		})
	}
	return entries
}

func lineMalformedError(lineNumber int) error {
	return fmt.Errorf("line %d is malformed", lineNumber)
}
