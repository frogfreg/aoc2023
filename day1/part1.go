package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func partOne() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	count := 0

	for _, line := range lines {
		firstIndex := strings.IndexFunc(line, func(letter rune) bool {
			match, err := regexp.MatchString(`[0-9]`, string(letter))
			if err != nil {
				panic(err)
			}

			return match
		})

		lastIndex := strings.LastIndexFunc(line, func(letter rune) bool {
			match, err := regexp.MatchString(`[0-9]`, string(letter))
			if err != nil {
				panic(err)
			}

			return match
		})

		if firstIndex < 0 || lastIndex < 0 {
			fmt.Println(line)
			panic("wtf")
		}

		d := string(line[firstIndex]) + string(line[lastIndex])

		digit, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}

		count += digit

	}

	fmt.Println(count)

}
