package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

type bag struct {
	name     string
	contents map[*bag]int // bag -> quantity
}

type bagSet map[string]*bag

var ruleRegexp = regexp.MustCompile(`^(\w+ \w+) bags contain (.*)\.$`)
var contentRegexp = regexp.MustCompile(`(\d) (\w+ \w+) bags?`)

func main() {
	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	var bags bagSet = make(map[string]*bag)
	for i, l := range lines {
		matches := ruleRegexp.FindStringSubmatch(l)
		if len(matches) != 3 {
			panic(fmt.Errorf("error on line %d", i+1))
		}

		contents := make(map[*bag]int)
		if matches[2] != "no other bags" {
			contentStrings := strings.Split(matches[2], ",")
			for _, c := range contentStrings {
				matches := contentRegexp.FindStringSubmatch(c)
				if len(matches) < 1 {
					fmt.Println(contentStrings)
				}
				cQuantity, err := strconv.Atoi(matches[1])
				if err != nil {
					panic(fmt.Errorf("error on line %d", i+1))
				}
				cBagName := matches[2]
				cBag := bags.get(cBagName)
				contents[cBag] = cQuantity
			}
		}

		mBagName := matches[1]
		mBag := bags.get(mBagName)
		mBag.contents = contents
	}

	total := 0
	shinyGold := bags.get("shiny gold")
	for _, b := range bags {
		if b.canContain(shinyGold) {
			total++
		}
	}
	fmt.Println(total)
}

func (s bagSet) get(bagName string) *bag {
	if b := s[bagName]; b != nil {
		return b
	}
	b := &bag{name: bagName}
	s[bagName] = b
	return b
}

func (b *bag) canContain(desiredBag *bag) bool {
	for cb := range b.contents {
		if cb == desiredBag {
			return true
		}
		if cb.canContain(desiredBag) {
			return true
		}
	}
	return false
}
