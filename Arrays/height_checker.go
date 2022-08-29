package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	sortedheight := make([]int, len(heights))
	at := copy(sortedheight, heights)
	sort.Ints(sortedheight)

	count := 0
	for i := 0; i < at; i++ {
		if heights[i] != sortedheight[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))
}
