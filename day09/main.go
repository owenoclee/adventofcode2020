package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	lines, err := parse.LinesFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	var numbers []int
	for i, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalf("error on line %d: not a number", i+1)
		}
		numbers = append(numbers, n)
	}

	for i := 25; i < len(numbers); i++ {
		if !numbersCanSumTo(numbers[i-25:i], numbers[i]) {
			fmt.Println(numbers[i])
			return
		}
	}
	fmt.Println("no solution")
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
