package main

import (
	"os"
	"strings"
)

func ex01() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	var totals []int
	elf := strings.Split(string(dat), "\n\n")
	for i, _ := range elf {
		cal := strings.Split(elf[i], "\n")
		totals = append(totals, sum(cal))
	}
	return getMax(totals)
}
