package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dict = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func partTwo() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	count := 0

	for _, line := range lines {

		first := getFirstOcurring(line)
		last := getLastOcurring(line)

		digit := first + last

		num, err := strconv.Atoi(digit)

		if err != nil {
			panic(err)
		}

		count += num

	}

	fmt.Println(count)

}

func getFirstOcurring(s string) string {
	writtenDigit := ""
	minOcurringIndex := 100000

	for k := range dict {
		index := strings.Index(s, k)
		if index != -1 && index < minOcurringIndex {
			minOcurringIndex = index
			writtenDigit = k
		}
	}

	value, ok := dict[writtenDigit]
	if !ok {
		panic("whaaat")
	}

	return value
}
func getLastOcurring(s string) string {
	writtenDigit := ""
	maxOcurringIndex := -10000000

	for k := range dict {
		index := strings.LastIndex(s, k)
		if index != -1 && index > maxOcurringIndex {
			maxOcurringIndex = index
			writtenDigit = k
		}
	}

	value, ok := dict[writtenDigit]
	if !ok {
		panic("whaaat")
	}

	return value
}
