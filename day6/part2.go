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

	timeStr := strings.Join(strings.Fields(strings.TrimPrefix(lines[0], "Time:")), "")
	distanceStr := strings.Join(strings.Fields(strings.TrimPrefix(lines[1], "Distance:")), "")

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	fmt.Println(waysToWin(time, distance))
}
