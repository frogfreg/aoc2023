package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	var times []int
	var distances []int

	for _, l := range lines {
		if strings.Contains(l, "Time:") {
			for _, digit := range strings.Fields(strings.TrimPrefix(l, "Time:")) {
				num, _ := strconv.Atoi(digit)
				times = append(times, num)
			}
		}
		if strings.Contains(l, "Distance:") {
			for _, digit := range strings.Fields(strings.TrimPrefix(l, "Distance:")) {
				num, _ := strconv.Atoi(digit)
				distances = append(distances, num)
			}
		}
	}

	product := 1

	for i := 0; i < len(times); i++ {
		product *= waysToWin(times[i], distances[i])
	}

	fmt.Println(product)
}

func waysToWin(time int, record int) int {
	var winnerQuantity int
	for i := 0; i <= time; i++ {
		if i*(time-i) > record {
			winnerQuantity++
		}
	}
	return winnerQuantity
}
