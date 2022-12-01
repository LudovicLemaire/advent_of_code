package main

import (
	"os"
	"sort"
	"strings"
)

func ex02() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	var totals []int
	elf := strings.Split(string(dat), "\n\n")
	for i, _ := range elf {
		cal := strings.Split(elf[i], "\n")
		totals = append(totals, sum(cal))
	}

	sort.Ints(sort.IntSlice(totals))
	l := len(totals)
	total := totals[l-1] + totals[l-2] + totals[l-3]
	return total
}
