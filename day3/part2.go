package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type gear struct {
	row int
	col int
}

type digitLocation struct {
	row, col int
}

func partTwo() {
	var grid [][]string

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

	gears := []gear{}

	for rowIndex, row := range grid {
		for colIndex, letter := range row {
			if letter == "*" {
				g := gear{rowIndex, colIndex}
				gears = append(gears, g)
			}
		}
	}

	sum := 0
	for _, g := range gears {
		sum += getGearRatio(grid, g)
	}

	fmt.Println(sum)
}

func getGearRatio(grid [][]string, g gear) int {
	var digitRegex = regexp.MustCompile(`[0-9]`)
	adjacentCount := 0
	dLocations := []digitLocation{}

	//up
	if g.row > 0 {
		upAdjacent := false
		dl := digitLocation{}

		var rightCorner *digitLocation
		var leftCorner *digitLocation
		var up *digitLocation

		if digitRegex.MatchString(grid[g.row-1][g.col]) {
			upAdjacent = true
			dl = digitLocation{g.row - 1, g.col}
			up = &digitLocation{g.row - 1, g.col}
		}
		if g.col > 0 && digitRegex.MatchString(grid[g.row-1][g.col-1]) {
			upAdjacent = true
			dl = digitLocation{g.row - 1, g.col - 1}
			leftCorner = &digitLocation{g.row - 1, g.col - 1}
		}
		if g.col+1 < len(grid[g.row]) && digitRegex.MatchString(grid[g.row-1][g.col+1]) {
			upAdjacent = true
			dl = digitLocation{g.row - 1, g.col + 1}
			rightCorner = &digitLocation{g.row - 1, g.col + 1}
		}
		if upAdjacent {
			if leftCorner != nil && rightCorner != nil && up == nil {
				dLocations = append(dLocations, *leftCorner, *rightCorner)
				adjacentCount += 2
			} else {
				dLocations = append(dLocations, dl)
				adjacentCount++
			}
		}
	}
	// left
	if g.col > 0 && digitRegex.MatchString(grid[g.row][g.col-1]) {
		adjacentCount++
		dLocations = append(dLocations, digitLocation{g.row, g.col - 1})
	}
	// right
	if g.col+1 > len(grid[g.row][g.col+1]) && digitRegex.MatchString(grid[g.row][g.col+1]) {
		adjacentCount++
		dLocations = append(dLocations, digitLocation{g.row, g.col + 1})
	}
	// down
	if g.row+1 < len(grid) {
		downAdjacent := false
		dl := digitLocation{}

		var rightCorner *digitLocation
		var leftCorner *digitLocation
		var down *digitLocation

		if digitRegex.MatchString(grid[g.row+1][g.col]) {
			downAdjacent = true
			dl = digitLocation{g.row + 1, g.col}
			down = &digitLocation{g.row + 1, g.col}
		}
		if g.col > 0 && digitRegex.MatchString(grid[g.row+1][g.col-1]) {
			downAdjacent = true
			dl = digitLocation{g.row + 1, g.col - 1}
			leftCorner = &digitLocation{g.row + 1, g.col - 1}
		}
		if g.col+1 < len(grid[g.row]) && digitRegex.MatchString(grid[g.row+1][g.col+1]) {
			downAdjacent = true
			dl = digitLocation{g.row + 1, g.col + 1}
			rightCorner = &digitLocation{g.row + 1, g.col + 1}
		}
		if downAdjacent {
			if leftCorner != nil && rightCorner != nil && down == nil {
				dLocations = append(dLocations, *leftCorner, *rightCorner)
				adjacentCount += 2
			} else {
				dLocations = append(dLocations, dl)
				adjacentCount++
			}
		}
	}

	if adjacentCount != 2 {
		return 0
	}

	gr := findAndMultiplyDigits(grid, dLocations)
	return gr
}

func findAndMultiplyDigits(grid [][]string, dLocations []digitLocation) int {
	product := 1

	var digitRegex = regexp.MustCompile(`[0-9]`)
	for _, dl := range dLocations {
		digitString := grid[dl.row][dl.col]
		currentCol := dl.col

		for currentCol-1 >= 0 && digitRegex.MatchString(grid[dl.row][currentCol-1]) {
			digitString = grid[dl.row][currentCol-1] + digitString
			currentCol--
		}
		currentCol = dl.col
		for currentCol+1 < len(grid[dl.row]) && digitRegex.MatchString(grid[dl.row][currentCol+1]) {
			digitString += grid[dl.row][currentCol+1]
			currentCol++
		}

		num, _ := strconv.Atoi(digitString)
		product *= num
	}

	return product
}
