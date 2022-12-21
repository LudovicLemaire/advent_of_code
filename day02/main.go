package main

import (
	"fmt"
)

type Game struct {
	player1    string
	player2    string
	totalScore int
	shapeValue int
	gameValue  int
}

func (g *Game) getShapeValue() {
	r := []rune(g.player2)
	g.shapeValue = int(r[0]) - 87
}

func (g *Game) getGameValue() {
	if (g.player2 == "X" && g.player1 == "C") || (g.player2 == "Y" && g.player1 == "A") || (g.player2 == "Z" && g.player1 == "B") {
		g.gameValue = 6
	} else if (g.player2 == "X" && g.player1 == "A") || (g.player2 == "Y" && g.player1 == "B") || (g.player2 == "Z" && g.player1 == "C") {
		g.gameValue = 3
	} else {
		g.gameValue = 0
	}
}

func (g *Game) getTotalScore() {
	g.totalScore = g.shapeValue + g.gameValue
}

func (g *Game) getCorrectAnswer() {
	switch g.player2 {
	case "X":
		switch g.player1 {
		case "A":
			g.player2 = "Z"
		case "B":
			g.player2 = "X"
		case "C":
			g.player2 = "Y"
		}
	case "Y":
		switch g.player1 {
		case "A":
			g.player2 = "X"
		case "B":
			g.player2 = "Y"
		case "C":
			g.player2 = "Z"
		}
	case "Z":
		switch g.player1 {
		case "A":
			g.player2 = "Y"
		case "B":
			g.player2 = "Z"
		case "C":
			g.player2 = "X"
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(ex01())
	fmt.Println(ex02())
}
