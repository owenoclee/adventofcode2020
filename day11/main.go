package main

import (
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
	curLayout, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		out.Fatalf("error parsing lines: %v", err)
	}
	if len(curLayout) == 0 {
		out.Fatalln("no lines in input")
	}
	width := len(curLayout[0])
	height := len(curLayout)

	var lastLayout []string
	for {
		lastLayout = make([]string, len(curLayout))
		for i, l := range curLayout {
			lastLayout[i] = l
		}
		curLayout = make([]string, len(curLayout))
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				if lastLayout[row][col] == floor {
					curLayout[row] += string(floor)
					continue
				}
				occupiedCount := 0
				// look at adjacent positions
				for aRow := row - 1; aRow <= row+1; aRow++ {
					for aCol := col - 1; aCol <= col+1; aCol++ {
						isWithinBounds := aRow >= 0 && aRow < height && aCol >= 0 && aCol < width
						if !(aRow == row && aCol == col) && isWithinBounds {
							if lastLayout[aRow][aCol] == occupied {
								occupiedCount++
							}
						}
					}
				}
				switch {
				case occupiedCount == 0:
					curLayout[row] += string(occupied)
				case occupiedCount >= 4:
					curLayout[row] += string(empty)
				default:
					curLayout[row] += string(lastLayout[row][col])
				}
			}
		}
		if layoutToString(curLayout) == layoutToString(lastLayout) {
			break
		}
	}
	out.Fatalln(strings.Count(layoutToString(curLayout), string(occupied)))
}

func layoutToString(l []string) string {
	return strings.Join(l, "\n")
}
