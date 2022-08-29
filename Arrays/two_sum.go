package main

import "fmt"

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
func twoSum(arr []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if (k != j+1) && (arr[k] == arr[k-1]) {
					continue
				}
				if arr[i]+arr[j]+arr[k] == target {
					ans := []int{arr[i], arr[j], arr[k]}
					res = append(res, ans)
				}
			}
		}
	}
	return res
}

func main() {
	array := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(twoSum(array, 0))
}
