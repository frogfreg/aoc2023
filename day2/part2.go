package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	count := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		setGroup := strings.Split(line, ":")[1]
		setGroup = strings.TrimSpace(setGroup)

		count += getPower(setGroup)
	}

	fmt.Println(count)

}

func getPower(setGroup string) int {
	sets := strings.Split(setGroup, ";")

	red := 0
	green := 0
	blue := 0

	for _, set := range sets {
		for _, cube := range strings.Split(set, ",") {
			numColor := strings.Fields(cube)
			num, _ := strconv.Atoi(numColor[0])
			color := numColor[1]

			switch color {
			case "red":
				red = max(red, num)
			case "green":
				green = max(green, num)
			case "blue":
				blue = max(blue, num)

			}

		}
	}

	return red * green * blue
}
