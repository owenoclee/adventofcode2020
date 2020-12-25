// Example usage: cat input.txt | go run . -p 1
package main

import (
	"flag"
	"os"
	"sort"

	"github.com/owenoclee/adventofcode2020/out"
	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()

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
	if *part != 2 {
		differences := make(map[int]int)
		for i := 1; i < len(adapters); i++ {
			differences[adapters[i]-adapters[i-1]]++
		}
		out.Fatalln(differences[1] * differences[3])
	}

	paths := make([]int64, len(adapters))
	paths[len(adapters)-1] = 1
	for i := len(adapters) - 2; i >= 0; i-- {
		var p int64
		for ii := i + 1; ii <= i+3 && ii < len(adapters); ii++ {
			if adapters[ii] > adapters[i]+3 {
				break
			}
			p += paths[ii]
		}
		paths[i] = p
	}
	out.Fatalln(paths[0])
}
