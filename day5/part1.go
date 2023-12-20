package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type mapping struct {
	destination int
	source      int
	quantity    int
}

var order = map[string]int{
	"seed-to-soil":            0,
	"soil-to-fertilizer":      1,
	"fertilizer-to-water":     2,
	"water-to-light":          3,
	"light-to-temperature":    4,
	"temperature-to-humidity": 5,
	"humidity-to-location":    6,
}

func partOne() {
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

	for _, seed := range seedList {
		location := getSeedLocation(seed, toDestinationMap, keys)
		minLocation = min(minLocation, location)
	}

	fmt.Println(minLocation)

}

func getSeedLocation(seed int, toDestinationMap map[string][]mapping, sortedKeys []string) int {
	for _, k := range sortedKeys {
		for _, m := range toDestinationMap[k] {
			newSeed, converted := conversion(seed, m)
			seed = newSeed
			if converted {
				break
			}
		}
	}

	return seed
}

func conversion(seed int, mapTo mapping) (int, bool) {
	if seed >= mapTo.source && seed <= (mapTo.source+(mapTo.quantity-1)) {
		return mapTo.destination + (seed - mapTo.source), true
	}
	return seed, false
}
