package main

import (
	"os"
	"strconv"

	"github.com/owenoclee/adventofcode2020/day12/compass"
	"github.com/owenoclee/adventofcode2020/out"
	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		out.Fatalf("error parsing lines: %v", err)
	}

	direction := compass.East
	location := compass.Zero
	for i, l := range lines {
		if len(l) < 2 {
			out.Fatalf("invalid line %d", i)
		}
		action := l[0]
		sarg := l[1:]
		arg, err := strconv.Atoi(sarg)
		if err != nil {
			out.Fatalf("argument is not a number on line %d", i)
		}
		switch action {
		case 'N':
			location = location.Add(compass.North.Multiply(arg))
		case 'E':
			location = location.Add(compass.East.Multiply(arg))
		case 'S':
			location = location.Add(compass.South.Multiply(arg))
		case 'W':
			location = location.Add(compass.West.Multiply(arg))
		case 'F':
			location = location.Add(direction.Multiply(arg))
		case 'R':
			direction = direction.Turn(arg)
		case 'L':
			direction = direction.Turn(-arg)
		}
	}
	out.Fatalln(location.ManhattanDistance())
}
