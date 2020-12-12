package main

import (
	"fmt"
	"os"

	"github.com/owenoclee/adventofcode2020/parse"
)

const (
	right      = 3
	down       = 1
	tree  byte = '#'
)

func main() {
	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	lineCount := len(lines)
	if lineCount == 0 {
		panic("no input lines given")
	}
	width := len(lines[0])

	var treesEncountered int
	x := right
	for y := down; y < lineCount; y += down {
		if lines[y][x%width] == tree {
			treesEncountered++
		}
		x += right
	}
	fmt.Println(treesEncountered)
}
