package main

import (
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(arr []string) int {
	var total int = 0
	for _, s := range arr {
		value, err := strconv.Atoi(s)
		check(err)
		total += value
	}
	return total
}

func getMax(arr []int) int {
	max := arr[0]
	for _, el := range arr {
		if el > max {
			max = el
		}
	}
	return max
}

func main() {
	fmt.Println(ex01())
	fmt.Println(ex02())
}
