// Example usage: cat input.txt | go run . -p 1
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

type instruction struct {
	op  string
	arg int
}

type state struct {
	index       int
	accumulator int
}

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()

	sInstructions, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	var instructions []instruction
	for i, inst := range sInstructions {
		opArg := strings.Split(inst, " ")
		if len(opArg) < 2 {
			log.Fatalf("error on line %d: line does not satisfy format `<op> <arg>`", i+1)
		}
		op, sArg := opArg[0], opArg[1]
		arg, err := strconv.Atoi(sArg)
		if err != nil {
			log.Fatalf("error on line %d: cannot convert arg to int", i+1)
		}
		instructions = append(instructions, instruction{
			op:  op,
			arg: arg,
		})
	}

	stateStack := make(map[int]state)
	visitedIndexes := make(map[int]struct{})
	for {
		s := stateStack[len(stateStack)-1]
		if _, visitedBefore := visitedIndexes[s.index]; visitedBefore {
			break
		}
		visitedIndexes[s.index] = struct{}{}
		stateStack[len(stateStack)] = nextState(s, instructions[s.index])
	}
	if *part != 2 {
		fmt.Println(stateStack[len(stateStack)-1].accumulator)
		return
	}

	finalStateIdx := len(instructions)
	// work backwards from the last valid state, amending instructions until a fix is found
	for i := len(stateStack) - 1; i >= 0; i-- {
		s := stateStack[i]
		inst := instructions[s.index]
		switch inst.op {
		case "nop":
			inst.op = "jmp"
		case "jmp":
			inst.op = "nop"
		}
		visitedIndexes := make(map[int]struct{})
		for {
			if _, visitedBefore := visitedIndexes[s.index]; visitedBefore {
				break
			}
			visitedIndexes[s.index] = struct{}{}
			s = nextState(s, inst)
			if s.index == finalStateIdx {
				fmt.Println(s.accumulator)
				return
			}
			inst = instructions[s.index]
		}
	}
	fmt.Println("no solution")
}

func nextState(s state, inst instruction) state {
	switch inst.op {
	case "acc":
		return state{
			accumulator: s.accumulator + inst.arg,
			index:       s.index + 1,
		}
	case "jmp":
		return state{
			accumulator: s.accumulator,
			index:       s.index + inst.arg,
		}
	case "nop":
		return state{
			accumulator: s.accumulator,
			index:       s.index + 1,
		}
	default:
		panic(fmt.Errorf("invalid operation '%s'", inst.op))
	}
}
