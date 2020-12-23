// Example usage: cat input.txt | go run . -p 1
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()

	text, err := parse.TextFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	groups := strings.Split(text, "\n\n")

	total := 0
	for _, g := range groups {
		gg := strings.ReplaceAll(g, "\n", "")
		answerCountMap := make(map[rune]int)
		for _, l := range gg {
			answerCountMap[l] += 1
		}
		if *part != 2 {
			total += len(answerCountMap)
			continue
		}
		membersInGroup := strings.Count(g, "\n") + 1
		for _, count := range answerCountMap {
			if count == membersInGroup {
				total += 1
			}
		}
	}
	fmt.Println(total)
}
