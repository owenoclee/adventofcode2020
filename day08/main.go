package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	ops, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}

	visitedIndexes := make(map[int]struct{})
	accumulator := 0
	for i := 0; i < len(ops) && i >= 0; {
		if _, visitedBefore := visitedIndexes[i]; visitedBefore {
			break
		}
		visitedIndexes[i] = struct{}{}
		opArg := strings.Split(ops[i], " ")
		if len(opArg) < 2 {
			log.Fatalf("error on line %d: line does not satisfy format `<op> <arg>`", i+1)
		}
		op, sarg := opArg[0], opArg[1]
		arg, err := strconv.Atoi(sarg)
		if err != nil {
			log.Fatalf("error on line %d: cannot convert arg to int", i+1)
		}
		switch op {
		case "acc":
			accumulator += arg
			i++
		case "jmp":
			i += arg
		case "nop":
			i++
		}
	}
	fmt.Println(accumulator)
}
