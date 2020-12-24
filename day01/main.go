// Example usage: cat input.txt | go run . -s 2
package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/owenoclee/adventofcode2020/parse"
)

const goal = 2020

func main() {
	summandsCount := flag.Int("s", 2, "Specify how many summands to use")
	flag.Parse()

	expenseReport, err := parse.IntsFrom(os.Stdin)
	if err != nil {
		panic(err)
	}

	solution, err := solve(expenseReport, *summandsCount)
	if err != nil {
		panic(err)
	}
	fmt.Println(solution)
}

func solve(expenseReport []int, summandsCount int) (int, error) {
	divisors := make([]int, summandsCount)
	for col := 0; col < summandsCount; col++ {
		divisors[col] = int(math.Pow(float64(len(expenseReport)), float64(col)))
	}

	combinations := int(math.Pow(float64(len(expenseReport)), float64(summandsCount)))
	for combo := 0; combo < combinations; {
	Start:
		combo++
		indexes := make(map[int]bool)
		for col := 0; col < summandsCount; col++ {
			idx := (combo / divisors[col]) % len(expenseReport)
			if indexes[idx] {
				goto Start
			}
			indexes[idx] = true
		}

		var total int
		for i := range indexes {
			total += expenseReport[i]
		}
		if total == goal {
			solution := 1
			for i := range indexes {
				solution *= expenseReport[i]
			}
			return solution, nil
		}
	}
	return 0, errors.New("no solution")
}
