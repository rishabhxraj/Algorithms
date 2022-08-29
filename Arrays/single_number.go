package main

import "fmt"

// Example 1:
//
// Input: nums = [2,2,1]
// Output: 1
//
// Example 2:
//
// Input: nums = [4,1,2,1,2]
// Output: 4
//
// Example 3:
//
// Input: nums = [1]
// Output: 1

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	array := []int{2, 2, 1}
	fmt.Println(singleNumber(array))
}
