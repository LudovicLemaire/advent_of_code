package main

import (
	"os"
	"strings"
)

func ex01() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	var total int = 0
	games := strings.Split(string(dat), "\n")
	for i, _ := range games {
		players := strings.Split(games[i], " ")
		g := Game{player1: players[0], player2: players[1]}
		g.getShapeValue()
		g.getGameValue()
		g.getTotalScore()
		total += g.totalScore
	}
	return total
}
