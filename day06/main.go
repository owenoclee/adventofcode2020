package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	text, err := parse.TextFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	groups := strings.Split(text, "\n\n")

	uniqueAnswersPerGroupTotal := 0
	for _, g := range groups {
		gg := strings.ReplaceAll(g, "\n", "")
		uniqueAnswers := make(map[rune]struct{})
		for _, l := range gg {
			uniqueAnswers[l] = struct{}{}
		}
		uniqueAnswersPerGroupTotal += len(uniqueAnswers)
	}
	fmt.Println(uniqueAnswersPerGroupTotal)
}
