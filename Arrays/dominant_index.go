package main

import (
	"fmt"
	"math"
)

// Example 1:

// Input: nums = [3,6,1,0]
// Output: 1
// Explanation: 6 is the largest integer.
// For every other number in the array x, 6 is at least twice as big as x.
// The index of value 6 is 1, so we return 1.
// Example 2:

// Input: nums = [1,2,3,4]
// Output: -1
// Explanation: 4 is less than twice the value of 3, so we return -1.

func dominantIndex(nums []int) int {
	maxIndex := max(nums)
	return maxIndex
}

func max(array []int) int {

	max := math.MinInt
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}
