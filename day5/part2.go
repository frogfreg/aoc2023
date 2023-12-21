package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type seedPair struct {
	start  int
	sRange int
}

func partTwo() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(inputData), "\n")

	seedStrList := strings.Fields(strings.TrimPrefix(lines[0], "seeds:"))
	seedList := []int{}

	for _, seedStr := range seedStrList {
		num, _ := strconv.Atoi(seedStr)
		seedList = append(seedList, num)
	}

	seedPairs := []seedPair{}

	currPair := []int{}
	for _, s := range seedList {
		if len(currPair) == 2 {
			seedPairs = append(seedPairs, seedPair{currPair[0], currPair[1]})
			currPair = []int{}
		}

		currPair = append(currPair, s)
	}
	seedPairs = append(seedPairs, seedPair{currPair[0], currPair[1]})

	toDestinationMap := map[string][]mapping{}

	currKey := ""

	for _, l := range lines[1:] {
		if strings.TrimSpace(l) == "" {
			continue
		}
		if strings.Contains(l, "map:") {
			currKey = strings.Fields(l)[0]
			continue
		}

		mInts := []int{}

		for _, num := range strings.Fields(l) {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			mInts = append(mInts, n)
		}

		m := mapping{mInts[0], mInts[1], mInts[2]}

		toDestinationMap[currKey] = append(toDestinationMap[currKey], m)

	}

	keys := []string{}

	for k := range toDestinationMap {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a, b string) int {
		return order[a] - order[b]
	})

	minLocation := math.MaxInt

	for _, sp := range seedPairs {
		for i := sp.start; i < sp.start+sp.sRange; i++ {
			location := getSeedLocation(i, toDestinationMap, keys)
			minLocation = min(minLocation, location)
		}
	}

	fmt.Println(minLocation)
}
