package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type cardInfo struct {
	generates int
	quantity  int
}

func partTwo() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	cards := map[int]cardInfo{}

	lines := strings.Split(string(inputData), "\n")
	for _, line := range lines {
		id, card := countWinners(line)
		cards[id] = card
	}

	for cardId := 1; cardId <= len(cards); cardId++ {
		for i := 1; i <= cards[cardId].generates; i++ {
			if childCard, exists := cards[cardId+i]; exists {
				childCard.quantity += cards[cardId].quantity
				cards[cardId+i] = childCard
			}
		}
	}

	sum := 0
	for _, card := range cards {
		sum += card.quantity
	}

	fmt.Println(sum)
}

func countWinners(card string) (int, cardInfo) {

	cardParts := strings.Split(card, ":")
	id, _ := strconv.Atoi(strings.Fields(cardParts[0])[1])
	raw := strings.Split(cardParts[1], "|")
	winners := strings.Fields(raw[0])
	have := strings.Fields(raw[1])
	matches := 0

	for _, num := range winners {
		if slices.Contains(have, num) {
			matches++
		}
	}

	return id, cardInfo{matches, 1}
}
