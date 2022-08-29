package main

import "fmt"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	// Two pointers
	left := 0
	right := len(nums) - 1
	i := len(nums) - 1
	for left < right {
		if nums[left] < 0 && nums[left]+nums[right] < 0 {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
		i--
	}
	res[0] = nums[left] * nums[left]
	return res
}
func main() {
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
