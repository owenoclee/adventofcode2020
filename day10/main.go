package main

import (
	"os"
	"sort"

	"github.com/owenoclee/adventofcode2020/out"
	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	adapters, err := parse.IntsFrom(os.Stdin)
	if err != nil {
		out.Fatalf("error parsing ints: %v", err)
	}
	var highest int
	for _, a := range adapters {
		if a > highest {
			highest = a
		}
	}
	adapters = append(adapters, 0, highest+3)

	sort.Ints(adapters)
	differences := make(map[int]int)
	for i := 1; i < len(adapters); i++ {
		differences[adapters[i]-adapters[i-1]]++
	}
	out.Fatalln(differences[1] * differences[3])
}
