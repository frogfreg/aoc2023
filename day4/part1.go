package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func partOne() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")
	sum := 0
	for _, line := range lines {
		sum += getPointsFromCard(line)
	}
	fmt.Println(sum)
}

func getPointsFromCard(card string) int {
	raw := strings.Split(strings.Split(card, ":")[1], "|")
	winners := strings.Fields(raw[0])
	have := strings.Fields(raw[1])
	matches := 0
	for _, num := range winners {
		if slices.Contains(have, num) {
			matches++
		}
	}

	return calculatePoints(matches)

}

func calculatePoints(matches int) int {
	if matches <= 0 {
		return 0
	}
	return int(math.Pow(2, float64(matches-1)))
}
