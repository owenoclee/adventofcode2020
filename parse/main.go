package parse

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func LinesFrom(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func TextFrom(r io.Reader) (string, error) {
	lines, err := LinesFrom(r)
	if err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

func LinesToInts(lines []string) ([]int, error) {
	var ints []int
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			return ints, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
