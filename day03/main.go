// Example usage: cat input.txt | go run . 1:1 3:1 5:1 7:1 1:2
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

const tree byte = '#'

type slope struct {
	right int
	down  int
}

func main() {
	slopeArgs := []string{"3:1"}
	if len(os.Args) > 1 {
		slopeArgs = os.Args[1:]
	}
	var slopes []slope
	for _, s := range slopeArgs {
		parts := strings.Split(s, ":")
		if len(parts) != 2 {
			panic(invalidSlopeError(s))
		}
		right, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(invalidSlopeError(s))
		}
		down, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(invalidSlopeError(s))
		}

		slopes = append(slopes, slope{
			right: right,
			down:  down,
		})
	}

	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	lineCount := len(lines)
	if lineCount == 0 {
		panic("no input lines given")
	}
	width := len(lines[0])

	treesEncountered := make([]int, len(slopes))
	for i, s := range slopes {
		x := s.right
		for y := s.down; y < lineCount; y += s.down {
			if lines[y][x%width] == tree {
				treesEncountered[i]++
			}
			x += s.right
		}
	}
	solution := 1
	for _, trees := range treesEncountered {
		solution *= trees
	}
	fmt.Println(solution)
}

func invalidSlopeError(s string) error {
	return fmt.Errorf("invalid slope: %s", s)
}
