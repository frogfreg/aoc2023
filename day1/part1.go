package main

import (
	"fmt"
	"os"
)

func partOne() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	_ = inputData

	fmt.Println("preparing for aoc2023")
}
