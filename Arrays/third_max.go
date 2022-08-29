package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	first := nums[0]
	second := math.MinInt
	third := math.MinInt

	for i := 1; i < len(nums); i++ {
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if second > nums[i] {
			third = second
			second = nums[i]
		} else if third > nums[i] {
			third = nums[i]
		}
	}
	return third
}

func main() {
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}
