package main

import (
	"os"
	"strings"
)

func ex01() int {
	dat, err := os.ReadFile("./input.txt")
	check(err)

	var totals []int
	elfs := strings.Split(string(dat), "\n\n")
	for i, _ := range elfs {
		cals := strings.Split(elfs[i], "\n")
		totals = append(totals, sum(cals))
	}
	return getMax(totals)
}
