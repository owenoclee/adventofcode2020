package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

type state struct {
	index       int
	accumulator int
}

func main() {
	ops, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}

	stateStack := make(map[int]state)
	visitedIndexes := make(map[int]struct{})
	for {
		lastState := state{}
		if len(stateStack) != 0 {
			lastState = stateStack[len(stateStack)-1]
		}
		if _, visitedBefore := visitedIndexes[lastState.index]; visitedBefore {
			break
		}
		visitedIndexes[lastState.index] = struct{}{}
		opArg := strings.Split(ops[lastState.index], " ")
		if len(opArg) < 2 {
			log.Fatalf("error on line %d: line does not satisfy format `<op> <arg>`", lastState.index+1)
		}
		op, sarg := opArg[0], opArg[1]
		arg, err := strconv.Atoi(sarg)
		if err != nil {
			log.Fatalf("error on line %d: cannot convert arg to int", lastState.index+1)
		}
		stateStack[len(stateStack)] = nextState(stateStack[len(stateStack)-1], lastState.index, op, arg)
	}
	fmt.Println(stateStack[len(stateStack)-1].accumulator)
}

func nextState(lastState state, index int, op string, arg int) state {
	accumulator := lastState.accumulator
	switch op {
	case "acc":
		return state{
			accumulator: accumulator + arg,
			index:       index + 1,
		}
	case "jmp":
		return state{
			accumulator: accumulator,
			index:       index + arg,
		}
	case "nop":
		return state{
			accumulator: accumulator,
			index:       index + 1,
		}
	default:
		panic(fmt.Errorf("invalid operation '%s'", op))
	}
}
