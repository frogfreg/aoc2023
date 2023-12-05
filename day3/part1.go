package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type num struct {
	adjacent bool
	digits   string
}

var digitRegex = regexp.MustCompile(`[0-9]`)
var validNumStrings []string
var rows []string

func partOne() {
	inputData, err := os.ReadFile("./test-input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		rows = append(rows, line)
	}

	fmt.Printf("%#v\n", rows)

}

func searchPartNumbers(rowIndex, colIndex int) {

	if regexp.MatchString(digitRegex, string(rows[rowIndex][colIndex])) {

	}

	digits := ""

}
