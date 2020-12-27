package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/owenoclee/adventofcode2020/day12/compass"
	"github.com/owenoclee/adventofcode2020/out"
	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()

	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		out.Fatalf("error parsing lines: %v", err)
	}

	waypoint := compass.East.Scale(10).Add(compass.North)
	if *part != 2 {
		waypoint = compass.East
	}
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
		var motion compass.Vector
		switch action {
		case 'N':
			motion = compass.North.Scale(arg)
		case 'E':
			motion = compass.East.Scale(arg)
		case 'S':
			motion = compass.South.Scale(arg)
		case 'W':
			motion = compass.West.Scale(arg)
		case 'F':
			location = location.Add(waypoint.Scale(arg))
		case 'R':
			waypoint = waypoint.Rotate(arg)
		case 'L':
			waypoint = waypoint.Rotate(-arg)
		}
		if *part != 2 {
			location = location.Add(motion)
			continue
		}
		waypoint = waypoint.Add(motion)
	}
	out.Fatalln(location.ManhattanDistance())
}
