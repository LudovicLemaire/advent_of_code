package main

import (
	"os"
	"strings"
)

func ex02() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	total := 0
	rucksacks := strings.Split(string(dat), "\n")
	for i, rucksack := range rucksacks {
		if (i)%3 == 0 {
			var key rune
			for _, currRune := range rucksack {
				if strings.Contains(rucksacks[i+1], string(currRune)) && strings.Contains(rucksacks[i+2], string(currRune)) {
					key = currRune
				}
			}
			if key >= 97 {
				key -= 96
			} else {
				key -= 38
			}
			total += int(key)
		}
	}

	return total
}

/*
var key rune
		for _, rune1 := range part1 {

		}
		if key >= 97 {
			key -= 96
		} else {
			key -= 38
		}
		total += int(key)
*/
