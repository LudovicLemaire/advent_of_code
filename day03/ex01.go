package main

import (
	"os"
	"strings"
)

func ex01() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	total := 0
	rucksacks := strings.Split(string(dat), "\n")
	for _, rucksack := range rucksacks {
		part1 := rucksack[:len(rucksack)/2]
		part2 := rucksack[len(rucksack)/2:]
		didBreak := false
		var key rune
		for _, rune1 := range part1 {
			if didBreak {
				break
			}
			for _, rune2 := range part2 {
				if rune1 == rune2 {
					key = rune1
					didBreak = true
					break
				}
			}
		}
		if key >= 97 {
			key -= 96
		} else {
			key -= 38
		}
		total += int(key)
	}

	return total
}
