// Example usage: cat input.txt | go run . -p 1
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/owenoclee/adventofcode2020/parse"
)

func main() {
	part := flag.Int("p", 1, "Specify which part of the puzzle to solve")
	flag.Parse()
	useStrictValidation := *part == 2

	text, err := parse.TextFrom(os.Stdin)
	if err != nil {
		panic(err)
	}
	rawPassportLines := strings.Split(text, "\n\n")
	var passportLines []string
	for _, rpLine := range rawPassportLines {
		passportLines = append(passportLines, strings.ReplaceAll(rpLine, "\n", " "))
	}

	var validCount int
	for _, pLine := range passportLines {
		if isLineValid(pLine, useStrictValidation) {
			validCount++
		}
	}
	fmt.Println(validCount)
}

func isLineValid(line string, useStrictValidation bool) bool {
	fields := strings.Split(line, " ")
	switch len(fields) {
	case 8:
		if useStrictValidation {
			return areFieldsValid(fields)
		}
		return true
	case 7:
		isMissingIgnoredField := !strings.Contains(line, "cid:")
		if useStrictValidation && isMissingIgnoredField {
			return areFieldsValid(fields)
		}
		return isMissingIgnoredField
	}
	return false
}

var eyeColours = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
var validHcl = regexp.MustCompile(`^#[0-9a-f]{6}$`)

func isFieldValid(f string) bool {
	parts := strings.Split(f, ":")
	if len(parts) != 2 {
		return false
	}
	name := parts[0]
	value := parts[1]

	switch name {
	case "byr":
		return isYearValid(value, 1920, 2002)
	case "iyr":
		return isYearValid(value, 2010, 2020)
	case "eyr":
		return isYearValid(value, 2020, 2030)
	case "hgt":
		if len(value) < 4 {
			return false
		}
		heightStr := value[:len(value)-2]
		unit := value[len(value)-2:]
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			return false
		}
		switch unit {
		case "cm":
			return height >= 150 && height <= 193
		case "in":
			return height >= 59 && height <= 76
		}
		return false
	case "hcl":
		return validHcl.Match([]byte(value))
	case "ecl":
		for _, validEcl := range eyeColours {
			if value == validEcl {
				return true
			}
		}
		return false
	case "pid":
		if len(value) != 9 {
			return false
		}
		_, err := strconv.Atoi(value)
		return err == nil
	case "cid":
		return true
	}
	return false
}

func isYearValid(yearStr string, min int, max int) bool {
	if len(yearStr) != 4 {
		return false
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}
	return year >= min && year <= max
}

func areFieldsValid(fields []string) bool {
	for _, f := range fields {
		if !isFieldValid(f) {
			return false
		}
	}
	return true
}
