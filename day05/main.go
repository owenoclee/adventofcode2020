// Example usage: cat input.txt | go run . -p 1
package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()

	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}

	seatIDs := make(map[int]struct{})
	for i, l := range lines {
		row, col, err := decodeSeat(l)
		if err != nil {
			panic(invalidSeatError(i+1, l))
		}
		seatIDs[seatID(row, col)] = struct{}{}
	}

	lowestSeatID := math.MaxInt32
	highestSeatID := 0
	for id := range seatIDs {
		if id > highestSeatID {
			highestSeatID = id
		} else if id < lowestSeatID {
			lowestSeatID = id
		}
	}
	if *part != 2 {
		fmt.Println(highestSeatID)
		return
	}

	for i := lowestSeatID; i < highestSeatID; i++ {
		if _, ok := seatIDs[i]; !ok {
			fmt.Println(i)
			return
		}
	}
}

var seatRegexp = regexp.MustCompile(`[FB]{7}[RL]{3}`)
var binaryReplacer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

func decodeSeat(seat string) (int, int, error) {
	if len(seat) != 10 || !seatRegexp.MatchString(seat) {
		return 0, 0, errors.New("invalid seat specification")
	}
	rowBinary := binaryReplacer.Replace(seat[:7])
	colBinary := binaryReplacer.Replace(seat[7:])
	row, err := strconv.ParseInt(rowBinary, 2, 32)
	if err != nil {
		return 0, 0, err
	}
	col, err := strconv.ParseInt(colBinary, 2, 32)
	if err != nil {
		return 0, 0, err
	}
	return int(row), int(col), nil
}

func seatID(row, column int) int {
	return row*8 + column
}

func invalidSeatError(line int, s string) error {
	return fmt.Errorf("invalid seat specification on line %d: %s", line, s)
}
