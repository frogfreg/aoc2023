package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxRed = 12
const maxGreen = 13
const maxBlue = 14

func partOne() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	countValid := 0

	for i, line := range lines {
		if line == "" {
			continue
		}
		gameId := i + 1

		setGroup := strings.Split(line, ":")[1]
		setGroup = strings.TrimSpace(setGroup)

		if isValidGame(setGroup) {
			countValid += gameId
		}

	}

	fmt.Println(countValid)

}

func isValidGame(setGroup string) bool {
	sets := strings.Split(setGroup, ";")

	for _, set := range sets {
		for _, cube := range strings.Split(set, ",") {
			numColor := strings.Fields(cube)
			num, _ := strconv.Atoi(numColor[0])
			color := numColor[1]

			switch color {
			case "red":
				if num > maxRed {
					return false
				}
			case "green":
				if num > maxGreen {
					return false
				}
			case "blue":
				if num > maxBlue {
					return false
				}
			}
		}
	}

	return true
}
