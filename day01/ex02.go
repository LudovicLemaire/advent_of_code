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
	elfs := strings.Split(string(dat), "\n\n")
	for i, _ := range elfs {
		cals := strings.Split(elfs[i], "\n")
		totals = append(totals, sum(cals))
	}

	sort.Ints(sort.IntSlice(totals))
	l := len(totals)
	total := totals[l-1] + totals[l-2] + totals[l-3]
	return total
}
