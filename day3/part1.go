package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type partNum struct {
	adjacent bool
	digits   string
}

var digitRegex = regexp.MustCompile(`[0-9]`)
var symbolRegex = regexp.MustCompile(`[^\w\.\d]`)
var validNumStrings []string
var grid [][]string

func partOne() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		grid = append(grid, strings.Split(line, ""))
	}

	for rowIndex, row := range grid {
		digits := ""
		adjacent := false

		for colIndex, letter := range row {
			if digitRegex.MatchString(letter) {
				digits += letter
				if adjacentSymbol(rowIndex, colIndex) {
					adjacent = true
				}
			} else {
				if adjacent {
					validNumStrings = append(validNumStrings, digits)
					adjacent = false
				}
				digits = ""
			}

		}

		if adjacent {
			validNumStrings = append(validNumStrings, digits)
		}
	}

	sum := 0
	for _, digits := range validNumStrings {
		num, _ := strconv.Atoi(digits)
		sum += num
	}

	fmt.Println(sum)

}

func adjacentSymbol(rowIndex, colIndex int) bool {

	if rowIndex-1 >= 0 && symbolRegex.MatchString(grid[rowIndex-1][colIndex]) {
		return true
	}

	if rowIndex+1 < len(grid) && symbolRegex.MatchString(grid[rowIndex+1][colIndex]) {
		return true
	}

	if colIndex-1 >= 0 && symbolRegex.MatchString(grid[rowIndex][colIndex-1]) {
		return true
	}

	if colIndex+1 < len(grid[rowIndex]) && symbolRegex.MatchString(grid[rowIndex][colIndex+1]) {
		return true
	}

	if rowIndex-1 >= 0 && colIndex-1 >= 0 && symbolRegex.MatchString(grid[rowIndex-1][colIndex-1]) {
		return true
	}

	if rowIndex-1 >= 0 && colIndex+1 < len(grid[rowIndex]) && symbolRegex.MatchString(grid[rowIndex-1][colIndex+1]) {
		return true
	}

	if rowIndex+1 < len(grid) && colIndex-1 >= 0 && symbolRegex.MatchString(grid[rowIndex+1][colIndex-1]) {
		return true
	}

	if rowIndex+1 < len(grid) && colIndex+1 < len(grid[rowIndex]) && symbolRegex.MatchString(grid[rowIndex+1][colIndex+1]) {
		return true
	}
	return false
}
