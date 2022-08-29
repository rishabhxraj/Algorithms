package main

import (
	"fmt"
	"math"
)

func findMaxConsecutiveOnes(nums []int) int {
	count, max := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		max = int(math.Max(float64(count), float64(max)))
	}
	return max
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
