package main

import (
	"fmt"
	"sort"
)

// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//
// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
//
// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low := i + 1
			high := len(nums) - 1
			sum := 0 - nums[i]

			for low < high {
				if nums[low]+nums[high] == sum {
					res = append(res, []int{nums[i], nums[low], nums[high]})
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if nums[low]+nums[high] > sum {
					high--
				} else {
					low++
				}
			}
		}
	}
	return res
}

func main() {
	array := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(array))
}
