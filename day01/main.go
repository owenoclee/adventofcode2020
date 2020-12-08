package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const desiredNumber = 2020

func main() {
	availableNumbers := make(map[int]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		availableNumbers[num] = true
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for currentNumber := range availableNumbers {
		neededNumber := desiredNumber - currentNumber
		if neededNumber == currentNumber {
			continue
		}
		if availableNumbers[neededNumber] {
			fmt.Println(currentNumber * neededNumber)
			return
		}
	}
	fmt.Println("no solution")
}
