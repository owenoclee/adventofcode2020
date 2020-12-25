package main

import (
	"flag"
	"math"
	"os"
	"strings"

	"github.com/owenoclee/adventofcode2020/out"
	"github.com/owenoclee/adventofcode2020/parse"
)

const (
	empty    = 'L'
	occupied = '#'
	floor    = '.'
)

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()
	lookDistance := math.MaxInt32
	if *part != 2 {
		lookDistance = 1
	}

	curLayout, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		out.Fatalf("error parsing lines: %v", err)
	}
	if len(curLayout) == 0 {
		out.Fatalln("no lines in input")
	}
	height := len(curLayout)
	width := len(curLayout[0])

	lastLayout := make([]string, len(curLayout))
	for {
		copy(lastLayout, curLayout)
		curLayout = make([]string, len(curLayout))
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				if lastLayout[row][col] == floor {
					curLayout[row] += string(floor)
					continue
				}
				occupiedCount := nearbyOccupiedCount(lastLayout, row, col, lookDistance)
				switch {
				case occupiedCount == 0:
					curLayout[row] += string(occupied)
				case occupiedCount >= 5 || (*part != 2 && occupiedCount >= 4):
					curLayout[row] += string(empty)
				default:
					curLayout[row] += string(lastLayout[row][col])
				}
			}
		}
		out.Println(layoutToString(lastLayout))
		out.Println()
		if layoutToString(curLayout) == layoutToString(lastLayout) {
			break
		}
	}
	out.Fatalln(strings.Count(layoutToString(curLayout), string(occupied)))
}

func layoutToString(l []string) string {
	return strings.Join(l, "\n")
}

func nearbyOccupiedCount(layout []string, row, col, maxDistance int) int {
	if maxDistance == 0 || !isInBounds(layout, row, col) {
		return 0
	}

	var unitVectors [8][2]int
	i := 0
	for rRow := -1; rRow <= 1; rRow++ {
		for rCol := -1; rCol <= 1; rCol++ {
			if rRow == 0 && rCol == 0 {
				continue
			}
			unitVectors[i] = [2]int{rRow, rCol}
			i++
		}
	}

	occupiedCount := 0
	for _, v := range unitVectors {
		for d := 1; d <= maxDistance; d++ {
			aRow := row + (v[0] * d)
			aCol := col + (v[1] * d)
			if !isInBounds(layout, aRow, aCol) {
				break
			}
			if curElem := layout[aRow][aCol]; curElem != floor {
				if curElem == occupied {
					occupiedCount++
				}
				break
			}
		}
	}
	return occupiedCount
}

func isInBounds(layout []string, row, col int) bool {
	return row >= 0 && col >= 0 && row < len(layout) && col < len(layout[row])
}
