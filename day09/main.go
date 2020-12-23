// Example usage: cat input.txt | go run . -p 1
package main

import (
	"flag"
	"math"
	"os"
	"strconv"

	"github.com/owenoclee/adventofcode2020/out"
	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()

	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	var numbers []int
	for i, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			out.Fatalf("error on line %d: not a number", i+1)
		}
		numbers = append(numbers, n)
	}

	var invalidNumber int
	for i := 25; i < len(numbers); i++ {
		if !numbersCanSumTo(numbers[i-25:i], numbers[i]) {
			invalidNumber = numbers[i]
			break
		}
	}
	if *part != 2 {
		out.Fatalln(invalidNumber)
	}

	for i := 0; i < len(numbers); i++ {
		total := 0
		smallest := math.MaxInt32
		largest := math.MinInt32
		for ni := i; ni < len(numbers); ni++ {
			curNum := numbers[ni]
			if curNum < smallest {
				smallest = curNum
			}
			if curNum > largest {
				largest = curNum
			}
			total += curNum
			if total > invalidNumber {
				break
			}
			if total == invalidNumber {
				out.Fatalln(smallest + largest)
			}
		}
	}
	out.Fatalln("no solution")
}

func numbersCanSumTo(numbers []int, desired int) bool {
	availableNumbers := make(map[int]int)
	for _, n := range numbers {
		availableNumbers[n]++
	}
	for n, q := range availableNumbers {
		needed := desired - n
		if needed != n && availableNumbers[needed] > 0 || needed == n && q > 1 {
			return true
		}
	}
	return false
}
