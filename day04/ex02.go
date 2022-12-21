package main

import (
	"os"
	"strings"
)

func ex02() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	duplicate := 0
	pairs := strings.Split(string(dat), "\n")
	for _, pair := range pairs {
		startSectionElf1 := 0
		endSectionElf1 := 0
		startSectionElf2 := 0
		endSectionElf2 := 0
		parts := strings.Split(string(pair), ",")
		for i, part := range parts {
			sections := strings.Split(string(part), "-")
			for y, _ := range sections {
				switch y {
				case 0:
					if i == 0 {
						startSectionElf1 = reducedAtoi(sections[0])
					} else {
						startSectionElf2 = reducedAtoi(sections[0])
					}
				case 1:
					if i == 0 {
						endSectionElf1 = reducedAtoi(sections[1])
					} else {
						endSectionElf2 = reducedAtoi(sections[1])
					}
				}
			}

		}
		if startSectionElf2 >= startSectionElf1 && startSectionElf2 <= endSectionElf1 {
			duplicate++
		} else if startSectionElf1 >= startSectionElf2 && startSectionElf1 <= endSectionElf2 {
			duplicate++
		}
	}
	return duplicate
}
