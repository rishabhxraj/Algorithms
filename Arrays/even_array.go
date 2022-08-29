package main

import (
	"fmt"
	"math"
)

func findNumbers(nums []int) int {
	count := 0
	for _, n := range nums {
		if isEvenCount(n) {
			count++
		}
	}
	return count
}

func isEvenCount(n int) bool {
	count := int(math.Log10(float64(n)) + 1)
	if count%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
}
